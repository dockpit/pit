package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dockpit/pit/model"
	"github.com/dockpit/pit/server/ui/bin"
)

var TemplateDir = filepath.Join("server", "ui")

type View struct {
	dbMeta  *model.Meta
	dbPath  string
	version string
}

func NewView(dbmeta *model.Meta, dbpath, version string) *View {
	return &View{
		dbMeta:  dbmeta,
		dbPath:  dbpath,
		version: version,
	}
}

func (v *View) addSystemTemplates(t *template.Template) error {
	head, err := uibin.Asset(filepath.Join(TemplateDir, "head.html"))
	if err != nil {
		return err
	}

	t.Funcs(template.FuncMap{
		"DATABASE_ID": func() string {
			return v.dbMeta.ID
		},
	})

	foot, err := uibin.Asset(filepath.Join(TemplateDir, "foot.html"))
	if err != nil {
		return err
	}

	_, err = t.New("head").Parse(string(head))
	if err != nil {
		return err
	}

	_, err = t.New("foot").Parse(string(foot))
	if err != nil {
		return err
	}

	return nil
}

func (v *View) Render(w http.ResponseWriter, name string, data map[string]interface{}) {
	b, err := uibin.Asset(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed load template '%s': %s", name, err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("name").Parse(string(b))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed parse template '%s': %s", name, err), http.StatusInternalServerError)
		return
	}

	err = v.addSystemTemplates(tmpl)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to system templates: %s", err), http.StatusInternalServerError)
		return
	}

	//expose some system variables on eacht render
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get CWD while rendering: %s", err), http.StatusInternalServerError)
		return
	}

	data["Meta"] = v.dbMeta
	data["Version"] = v.version
	data["DBPath"] = filepath.Join(cwd, v.dbPath)

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute template '%s': %s", name, err), http.StatusInternalServerError)
	}
}

func (v *View) RenderDashboard(w http.ResponseWriter, isos []*model.Isolation, deps []*model.Dep) {
	v.Render(w, filepath.Join(TemplateDir, "dashboard.html"), map[string]interface{}{"Isolations": isos, "Deps": deps})
}

// @TODO DEPRECATE

func (v *View) RenderAddDep(w http.ResponseWriter, iso *model.Isolation, deps []*model.Dep) {
	v.Render(w, filepath.Join(TemplateDir, "add_dep.html"), map[string]interface{}{"Isolation": iso, "Deps": deps})
}

func (v *View) RenderAddState(w http.ResponseWriter, dep *model.Dep) {
	v.Render(w, filepath.Join(TemplateDir, "add_state.html"), map[string]interface{}{"Dep": dep})
}

func (v *View) RenderEditState(w http.ResponseWriter, dep *model.Dep, s *model.State) {
	v.Render(w, filepath.Join(TemplateDir, "edit_state.html"), map[string]interface{}{"Dep": dep, "State": s})
}

func (v *View) RenderEditStateBuilt(w http.ResponseWriter, dep *model.Dep, s *model.State, output string) {
	v.Render(w, filepath.Join(TemplateDir, "edit_state.html"), map[string]interface{}{"Dep": dep, "State": s, "Output": output})
}

func (v *View) RenderOneIsolation(w http.ResponseWriter, iso *model.Isolation) {
	v.Render(w, filepath.Join(TemplateDir, "one_isolation.html"), map[string]interface{}{"Isolation": iso})
}
