package server

import (
	"encoding/json"
	"fmt"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/web"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"
	"github.com/dockpit/pit/server/ui/bin"
)

type Server struct {
	view      *View
	client    *client.Docker
	model     *model.Model
	bind      string
	done      chan interface{}
	listener  net.Listener
	templates map[string]*model.Template

	*http.Server
}

func New(v, baddr string, m *model.Model, client *client.Docker) (*Server, error) {
	mux := web.New()
	dbmeta, err := m.GetDBMetaData()
	if err != nil {
		return nil, errwrap.Wrapf("Failed to get DB Meta data: {{err}}", err)
	}

	s := &Server{
		view:      NewView(dbmeta, m.DBPath, v),
		client:    client,
		model:     m,
		bind:      baddr,
		done:      make(chan interface{}),
		listener:  bind.Socket(baddr),
		templates: map[string]*model.Template{},

		Server: &http.Server{Handler: mux},
	}

	//load templates from json assets
	tmpldata, err := uibin.Asset(filepath.Join(TemplateDir, "templates.json"))
	if err != nil {
		return nil, errwrap.Wrapf("Failed to load Dep templates: {{err}}", err)
	}

	err = json.Unmarshal(tmpldata, &s.templates)
	if err != nil {
		return nil, errwrap.Wrapf("Failed to decode Dep templates: {{err}}", err)
	}

	//api endpoints
	mux.Get("/api/isolations", s.ListIsolations)
	mux.Post("/api/isolations", s.CreateIsolation)
	mux.Delete("/api/isolations/:id", s.RemoveIsolation)
	mux.Put("/api/isolations/:id", s.UpdateIsolation)
	mux.Get("/api/templates", s.ListTemplates)
	mux.Get("/api/deps", s.ListDeps)
	mux.Post("/api/deps", s.CreateDep)
	mux.Delete("/api/deps/:name", s.RemoveDep)

	mux.Post("/api/deps/:name/states", s.CreateState)
	mux.Put("/api/deps/:name/states/:state_name", s.UpdateDepState)
	mux.Get("/api/deps/:name/states/:state_name", s.OneDepState)
	mux.Delete("/api/deps/:name/states/:state_name", s.RemoveDepState)
	mux.Post("/api/deps/:name/states/:state_name/builds", s.BuildDepState)
	mux.Post("/api/deps/:name/states/:state_name/runs", s.RunDepState)

	mux.Get("/api/builds/:id", s.OneBuild)

	//dashboard
	mux.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		isos, err := s.model.GetAllIsolations()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get all isolations: %s", err), http.StatusInternalServerError)
			return
		}

		deps, err := s.model.GetAllDeps()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get all dependencies: %s", err), http.StatusInternalServerError)
			return
		}

		s.view.RenderDashboard(w, isos, deps)
	})

	//editor
	mux.Get("/deps/:name/states/:state_name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		dep, err := s.model.FindDepByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		if dep == nil {
			http.NotFound(w, r)
			return
		}

		sname := c.URLParams["state_name"]
		state := dep.GetState(sname)
		if state == nil {
			http.NotFound(w, r)
			return
		}

		s.view.RenderEditor(w, dep, state)
	})

	//static files
	mux.Get("/static/*", func(c web.C, w http.ResponseWriter, r *http.Request) {
		path := filepath.Clean(r.URL.Path)
		path, _ = filepath.Rel("/static", path)
		path = filepath.Join(TemplateDir, path)

		data, err := uibin.Asset(path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ct := mime.TypeByExtension(filepath.Ext(path))
		w.Header().Set("Content-Type", ct)
		w.Write(data)
	})

	return s, nil
}

func (s *Server) Start() error {
	err := s.Server.Serve(s.listener)
	select {
	case <-s.done:
		if strings.HasSuffix(err.Error(), "use of closed network connection") {
			return nil
		}
	}
	return err
}

func (s *Server) URL() string {
	return fmt.Sprintf("http://%s", s.listener.Addr().String())
}

func (s *Server) Stop() error {
	close(s.done)
	err := s.listener.Close()
	return err
}
