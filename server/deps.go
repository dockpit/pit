package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

	"github.com/zenazn/goji/web"
)

func (s *Server) ListDeps(c web.C, w http.ResponseWriter, r *http.Request) {
	deps, err := s.model.GetAllDeps()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get all dependencies: %s", err), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(deps)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) ListTemplates(c web.C, w http.ResponseWriter, r *http.Request) {
	tmpls := map[string]*model.Template{}

	//@todo in the future retrieve  this list of templates solely from GitHub
	tmplCustom := &model.Template{
		ID:          "custom",
		Name:        "Custom",
		Category:    "",
		DefaultName: "",
		TemplateFiles: map[string]*model.TemplateFile{
			"Dockerfile": &model.TemplateFile{"FROM scratch"},
		},
	}

	tmplMySQL := &model.Template{
		ID:          "mysql",
		Name:        "MySQL",
		DefaultName: "mysql",
		Category:    "Data Stores",
		TemplateFiles: map[string]*model.TemplateFile{
			"Dockerfile": &model.TemplateFile{"FROM scratch"},
		},
	}

	tmplPostgres := &model.Template{
		ID:          "postgres",
		Name:        "PostgreSQL",
		Category:    "Data Stores",
		DefaultName: "postgres",
		TemplateFiles: map[string]*model.TemplateFile{
			"Dockerfile": &model.TemplateFile{"FROM scratch"},
		},
	}

	tmplEtcd := &model.Template{
		ID:          "etcd",
		Name:        "etcd",
		Category:    "Service Discovery",
		DefaultName: "etcd",
		TemplateFiles: map[string]*model.TemplateFile{
			"Dockerfile": &model.TemplateFile{"FROM scratch"},
		},
	}

	tmpls[tmplCustom.ID] = tmplCustom
	tmpls[tmplMySQL.ID] = tmplMySQL
	tmpls[tmplPostgres.ID] = tmplPostgres
	tmpls[tmplEtcd.ID] = tmplEtcd

	enc := json.NewEncoder(w)
	err := enc.Encode(tmpls)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) CreateDep(c web.C, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newdep *model.Dep
	err := dec.Decode(&newdep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode bew dep: %s", err), http.StatusBadRequest)
		return
	}

	dep, err := model.NewDep(newdep.Name, newdep.Template)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create dep: %s", err), http.StatusBadRequest)
		return
	}

	err = s.model.InsertDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert dep: %s", err), http.StatusBadRequest)
		return
	}
}

func (s *Server) RemoveDep(c web.C, w http.ResponseWriter, r *http.Request) {
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
}
