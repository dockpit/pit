package server

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/web"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"
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

func New(baddr string, m *model.Model, client *client.Docker) (*Server, error) {
	mux := web.New()
	dbmeta, err := m.GetDBMetaData()
	if err != nil {
		return nil, errwrap.Wrapf("Failed to get DB Meta data: {{err}}", err)
	}

	s := &Server{
		view:     NewView(dbmeta),
		client:   client,
		model:    m,
		bind:     baddr,
		done:     make(chan interface{}),
		listener: bind.Socket(baddr),

		Server: &http.Server{Handler: mux},
	}

	//list isolations
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

		s.view.RenderIsolationList(w, isos, deps)
	})

	//create new isolation
	mux.Post("/isolations", func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form data: %s", err), http.StatusBadRequest)
			return
		}

		i, err := model.NewIsolation(r.Form.Get("name"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create isolation: %s", err), http.StatusBadRequest)
			return
		}

		err = s.model.InsertIsolation(i)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert isolation: %s", err), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//create new dep
	mux.Post("/deps", func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form data: %s", err), http.StatusBadRequest)
			return
		}

		dep, err := model.NewDep(r.Form.Get("name"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create dep: %s", err), http.StatusBadRequest)
			return
		}

		err = s.model.InsertDep(dep)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert dep: %s", err), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//remove a dep
	mux.Post("/deps/:name/delete", func(c web.C, w http.ResponseWriter, r *http.Request) {
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

		err = s.model.RemoveDep(dep)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed remove dep with name '%s': %s", name, err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//remove a dep state
	mux.Post("/deps/:name/states/:state_name/delete", func(c web.C, w http.ResponseWriter, r *http.Request) {
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
		err = s.model.RemoveDepStateByName(dep, sname)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to remove state from dep: %s", err), http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//add a new dep to isolation
	mux.Get("/isolations/:name/add-dep", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		iso, err := s.model.FindIsolationByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find isolation with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		if iso == nil {
			http.NotFound(w, r)
			return
		}

		deps, err := s.model.GetAllDeps()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get all dependencies: %s", err), http.StatusInternalServerError)
			return
		}

		s.view.RenderAddDep(w, iso, deps)
	})

	//add state to dep
	mux.Get("/deps/:name/add-state", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		dep, err := s.model.FindDepByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		s.view.RenderAddState(w, dep)
	})

	//create state
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

	//add dep to state
	mux.Post("/isolations/:name/add-dep", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		iso, err := s.model.FindIsolationByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find isolation with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		if iso == nil {
			http.NotFound(w, r)
			return
		}

		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse form data: %s", err), http.StatusBadRequest)
			return
		}

		selection := strings.SplitN(r.Form.Get("dep"), "::", 2)
		if len(selection) != 2 {
			http.Error(w, fmt.Sprintf("Unexpected value for dependency: '%s'", r.Form.Get("dep")), http.StatusBadRequest)
		}

		dep, err := s.model.FindDepByName(selection[0])
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		if dep == nil {
			http.NotFound(w, r)
			return
		}

		iso.AddDep(dep, selection[1])
		err = s.model.UpdateIsolation(iso)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed update isolation with name '%s': %s", name, err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	//remove an isolation
	mux.Post("/isolations/:name/delete", func(c web.C, w http.ResponseWriter, r *http.Request) {
		name := c.URLParams["name"]
		iso, err := s.model.FindIsolationByName(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to find isolation with name '%s': %s", name, err), http.StatusBadRequest)
			return
		}

		if iso == nil {
			http.NotFound(w, r)
			return
		}

		err = s.model.RemoveIsolation(iso)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed remove isolation with name '%s': %s", name, err), http.StatusInternalServerError)
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

		s.view.RenderEditState(w, dep, state)
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

		state.Dockerfile = r.Form.Get("dockerfile")
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
