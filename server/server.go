package server

import (
	"bytes"
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
	view     *View
	client   *client.Docker
	model    *model.Model
	bind     string
	done     chan interface{}
	listener net.Listener

	*http.Server
}

func New(v, baddr string, m *model.Model, client *client.Docker) (*Server, error) {
	mux := web.New()
	dbmeta, err := m.GetDBMetaData()
	if err != nil {
		return nil, errwrap.Wrapf("Failed to get DB Meta data: {{err}}", err)
	}

	s := &Server{
		view:     NewView(dbmeta, m.DBPath, v),
		client:   client,
		model:    m,
		bind:     baddr,
		done:     make(chan interface{}),
		listener: bind.Socket(baddr),

		Server: &http.Server{Handler: mux},
	}

	//api endpoints
	mux.Get("/api/isolations", s.ListIsolations)
	mux.Post("/api/isolations", s.CreateIsolation)
	mux.Delete("/api/isolations/:id", s.RemoveIsolation)
	mux.Put("/api/isolations/:id", s.UpdateIsolation)
	mux.Get("/api/deps", s.ListDeps)
	mux.Post("/api/deps", s.CreateDep)
	mux.Delete("/api/deps/:name", s.RemoveDep)

	mux.Put("/api/deps/:name/states/:state_name", s.UpdateDepState)
	mux.Get("/api/deps/:name/states/:state_name", s.OneDepState)
	mux.Delete("/api/deps/:name/states/:state_name", s.RemoveDepState)

	//page endpoints
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

	// @TODO DEPRECATE

	//add state to dep
	mux.Get("/deps/:id/add-state", func(c web.C, w http.ResponseWriter, r *http.Request) {
		ID := c.URLParams["id"]
		dep, err := s.model.FindDepByName(ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find dep with ID '%s': %s", ID, err), http.StatusBadRequest)
			return
		}

		s.view.RenderAddState(w, dep)
	})

	//add state to dep
	mux.Post("/deps/:name/add-state", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		dep, err := s.model.FindDepByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form data: %s", err), http.StatusBadRequest)
			return
		}

		state, err := model.NewState(r.Form.Get("name"), r.Form.Get("type"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create state: %s", err), http.StatusBadRequest)
			return
		}

		dep.AddState(state)
		err = s.model.UpdateDep(dep)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed update dep with name '%s': %s", name, err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//edit a state
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

	//edit a state
	mux.Post("/deps/:name/states/:state_name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form data: %s", err), http.StatusBadRequest)
			return
		}

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

		state.Files["Dockerfile"] = r.Form.Get("dockerfile")
		if r.Form.Get("save") != "" {
			state.ImageName = r.Form.Get("image_name")
			if state.ImageName == "" {
				http.Error(w, fmt.Sprintf("Image name cannot be empty, please build it first"), http.StatusBadRequest)
				return
			}

			err = s.model.UpdateDep(dep)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed update dep state with name '%s'::'%s': %s", name, sname, err), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {

			buf := bytes.NewBuffer(nil)
			iname, err := s.client.Build(dep, state, buf)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to build dep::state '%s'::'%s': %s (%s)", name, sname, err, buf.String()), http.StatusInternalServerError)
				return
			}

			state.ImageName = iname
			s.view.RenderEditStateBuilt(w, dep, state, buf.String())
		}
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
