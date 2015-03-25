// +build debug

package uibin

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// server_ui_ds_store reads file data from disk. It returns an error on failure.
func server_ui_ds_store() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/.DS_Store"
	name := "server/ui/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_gitignore reads file data from disk. It returns an error on failure.
func server_ui_gitignore() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/.gitignore"
	name := "server/ui/.gitignore"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_jshintrc reads file data from disk. It returns an error on failure.
func server_ui_jshintrc() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/.jshintrc"
	name := "server/ui/.jshintrc"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_add_dep_html reads file data from disk. It returns an error on failure.
func server_ui_add_dep_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/add_dep.html"
	name := "server/ui/add_dep.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_add_state_html reads file data from disk. It returns an error on failure.
func server_ui_add_state_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/add_state.html"
	name := "server/ui/add_state.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_dashboard_html reads file data from disk. It returns an error on failure.
func server_ui_dashboard_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/dashboard.html"
	name := "server/ui/dashboard.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_dist_app_js reads file data from disk. It returns an error on failure.
func server_ui_dist_app_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/dist/app.js"
	name := "server/ui/dist/app.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_dist_main_css reads file data from disk. It returns an error on failure.
func server_ui_dist_main_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/dist/main.css"
	name := "server/ui/dist/main.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_edit_state_html reads file data from disk. It returns an error on failure.
func server_ui_edit_state_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/edit_state.html"
	name := "server/ui/edit_state.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_foot_html reads file data from disk. It returns an error on failure.
func server_ui_foot_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/foot.html"
	name := "server/ui/foot.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_gulpfile_js reads file data from disk. It returns an error on failure.
func server_ui_gulpfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/gulpfile.js"
	name := "server/ui/gulpfile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_head_html reads file data from disk. It returns an error on failure.
func server_ui_head_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/head.html"
	name := "server/ui/head.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_one_isolation_html reads file data from disk. It returns an error on failure.
func server_ui_one_isolation_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/one_isolation.html"
	name := "server/ui/one_isolation.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_package_json reads file data from disk. It returns an error on failure.
func server_ui_package_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/package.json"
	name := "server/ui/package.json"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_ds_store reads file data from disk. It returns an error on failure.
func server_ui_src_ds_store() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/.DS_Store"
	name := "server/ui/src/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_actions_depactions_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_actions_depactions_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/actions/DepActions.js"
	name := "server/ui/src/js/actions/DepActions.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_actions_isolationactions_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_actions_isolationactions_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/actions/IsolationActions.js"
	name := "server/ui/src/js/actions/IsolationActions.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_app_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_app_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/app.js"
	name := "server/ui/src/js/app.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_depform_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_depform_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/DepForm.jsx"
	name := "server/ui/src/js/components/DepForm.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_depitem_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_depitem_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/DepItem.jsx"
	name := "server/ui/src/js/components/DepItem.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_deplist_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_deplist_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/DepList.jsx"
	name := "server/ui/src/js/components/DepList.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_deppanel_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_deppanel_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/DepPanel.jsx"
	name := "server/ui/src/js/components/DepPanel.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_isolationform_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_isolationform_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/IsolationForm.jsx"
	name := "server/ui/src/js/components/IsolationForm.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_isolationitem_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_isolationitem_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/IsolationItem.jsx"
	name := "server/ui/src/js/components/IsolationItem.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_isolationlist_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_isolationlist_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/IsolationList.jsx"
	name := "server/ui/src/js/components/IsolationList.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_isolationpanel_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_isolationpanel_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/IsolationPanel.jsx"
	name := "server/ui/src/js/components/IsolationPanel.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_dispatcher_appdispatcher_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_dispatcher_appdispatcher_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/dispatcher/AppDispatcher.js"
	name := "server/ui/src/js/dispatcher/AppDispatcher.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_stores_depstore_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_stores_depstore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/stores/DepStore.js"
	name := "server/ui/src/js/stores/DepStore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_stores_isolationstore_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_stores_isolationstore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/stores/IsolationStore.js"
	name := "server/ui/src/js/stores/IsolationStore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_less_main_less reads file data from disk. It returns an error on failure.
func server_ui_src_less_main_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/less/main.less"
	name := "server/ui/src/less/main.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_ds_store reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_ds_store() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/.DS_Store"
	name := "server/ui/src/vendor/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_alerts_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_alerts_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/alerts.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/alerts.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_badges_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_badges_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/badges.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/badges.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_bootstrap_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_bootstrap_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/bootstrap.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/bootstrap.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_breadcrumbs_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_breadcrumbs_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/breadcrumbs.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/breadcrumbs.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_button_groups_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_button_groups_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/button-groups.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/button-groups.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_buttons_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_buttons_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/buttons.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/buttons.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_carousel_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_carousel_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/carousel.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/carousel.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_close_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_close_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/close.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/close.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_code_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_code_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/code.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/code.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_component_animations_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_component_animations_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/component-animations.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/component-animations.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_dropdowns_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_dropdowns_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/dropdowns.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/dropdowns.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_forms_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_forms_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/forms.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/forms.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_glyphicons_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_glyphicons_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/glyphicons.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/glyphicons.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_grid_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_grid_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/grid.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/grid.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_input_groups_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_input_groups_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/input-groups.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/input-groups.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_jumbotron_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_jumbotron_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/jumbotron.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/jumbotron.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_labels_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_labels_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/labels.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/labels.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_list_group_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_list_group_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/list-group.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/list-group.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_media_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_media_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/media.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/media.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_alerts_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_alerts_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/alerts.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/alerts.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_background_variant_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_background_variant_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/background-variant.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/background-variant.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_border_radius_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_border_radius_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/border-radius.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/border-radius.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_buttons_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_buttons_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/buttons.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/buttons.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_center_block_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_center_block_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/center-block.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/center-block.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_clearfix_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_clearfix_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/clearfix.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/clearfix.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_forms_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_forms_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/forms.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/forms.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_gradients_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_gradients_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/gradients.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/gradients.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_framework_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_framework_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/grid-framework.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/grid-framework.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/grid.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/grid.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_hide_text_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_hide_text_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/hide-text.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/hide-text.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_image_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_image_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/image.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/image.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_labels_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_labels_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/labels.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/labels.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_list_group_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_list_group_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/list-group.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/list-group.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_divider_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_divider_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-divider.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-divider.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_vertical_align_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_vertical_align_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-vertical-align.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-vertical-align.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_opacity_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_opacity_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/opacity.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/opacity.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_pagination_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_pagination_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/pagination.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/pagination.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_panels_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_panels_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/panels.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/panels.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_progress_bar_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_progress_bar_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/progress-bar.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/progress-bar.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_reset_filter_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_reset_filter_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/reset-filter.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/reset-filter.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_resize_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_resize_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/resize.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/resize.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_responsive_visibility_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_responsive_visibility_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/responsive-visibility.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/responsive-visibility.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_size_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_size_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/size.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/size.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_tab_focus_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_tab_focus_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/tab-focus.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/tab-focus.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_table_row_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_table_row_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/table-row.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/table-row.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_text_emphasis_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_text_emphasis_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/text-emphasis.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/text-emphasis.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_text_overflow_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_text_overflow_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/text-overflow.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/text-overflow.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_vendor_prefixes_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_vendor_prefixes_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins/vendor-prefixes.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins/vendor-prefixes.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_mixins_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_mixins_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/mixins.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/mixins.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_modals_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_modals_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/modals.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/modals.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_navbar_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_navbar_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/navbar.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/navbar.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_navs_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_navs_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/navs.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/navs.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_normalize_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_normalize_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/normalize.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/normalize.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_pager_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_pager_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/pager.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/pager.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_pagination_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_pagination_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/pagination.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/pagination.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_panels_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_panels_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/panels.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/panels.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_popovers_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_popovers_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/popovers.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/popovers.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_print_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_print_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/print.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/print.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_progress_bars_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_progress_bars_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/progress-bars.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/progress-bars.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_responsive_embed_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_responsive_embed_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/responsive-embed.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/responsive-embed.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_responsive_utilities_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_responsive_utilities_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/responsive-utilities.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/responsive-utilities.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_scaffolding_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_scaffolding_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/scaffolding.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/scaffolding.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_tables_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_tables_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/tables.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/tables.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_theme_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_theme_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/theme.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/theme.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_thumbnails_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_thumbnails_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/thumbnails.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/thumbnails.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_tooltip_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_tooltip_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/tooltip.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/tooltip.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_type_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_type_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/type.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/type.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_utilities_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_utilities_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/utilities.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/utilities.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_variables_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_variables_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/variables.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/variables.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_vendor_bootstrap_3_3_4_wells_less reads file data from disk. It returns an error on failure.
func server_ui_src_vendor_bootstrap_3_3_4_wells_less() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/vendor/bootstrap-3.3.4/wells.less"
	name := "server/ui/src/vendor/bootstrap-3.3.4/wells.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"server/ui/.DS_Store": server_ui_ds_store,
	"server/ui/.gitignore": server_ui_gitignore,
	"server/ui/.jshintrc": server_ui_jshintrc,
	"server/ui/add_dep.html": server_ui_add_dep_html,
	"server/ui/add_state.html": server_ui_add_state_html,
	"server/ui/dashboard.html": server_ui_dashboard_html,
	"server/ui/dist/app.js": server_ui_dist_app_js,
	"server/ui/dist/main.css": server_ui_dist_main_css,
	"server/ui/edit_state.html": server_ui_edit_state_html,
	"server/ui/foot.html": server_ui_foot_html,
	"server/ui/gulpfile.js": server_ui_gulpfile_js,
	"server/ui/head.html": server_ui_head_html,
	"server/ui/one_isolation.html": server_ui_one_isolation_html,
	"server/ui/package.json": server_ui_package_json,
	"server/ui/src/.DS_Store": server_ui_src_ds_store,
	"server/ui/src/js/actions/DepActions.js": server_ui_src_js_actions_depactions_js,
	"server/ui/src/js/actions/IsolationActions.js": server_ui_src_js_actions_isolationactions_js,
	"server/ui/src/js/app.js": server_ui_src_js_app_js,
	"server/ui/src/js/components/DepForm.jsx": server_ui_src_js_components_depform_jsx,
	"server/ui/src/js/components/DepItem.jsx": server_ui_src_js_components_depitem_jsx,
	"server/ui/src/js/components/DepList.jsx": server_ui_src_js_components_deplist_jsx,
	"server/ui/src/js/components/DepPanel.jsx": server_ui_src_js_components_deppanel_jsx,
	"server/ui/src/js/components/IsolationForm.jsx": server_ui_src_js_components_isolationform_jsx,
	"server/ui/src/js/components/IsolationItem.jsx": server_ui_src_js_components_isolationitem_jsx,
	"server/ui/src/js/components/IsolationList.jsx": server_ui_src_js_components_isolationlist_jsx,
	"server/ui/src/js/components/IsolationPanel.jsx": server_ui_src_js_components_isolationpanel_jsx,
	"server/ui/src/js/dispatcher/AppDispatcher.js": server_ui_src_js_dispatcher_appdispatcher_js,
	"server/ui/src/js/stores/DepStore.js": server_ui_src_js_stores_depstore_js,
	"server/ui/src/js/stores/IsolationStore.js": server_ui_src_js_stores_isolationstore_js,
	"server/ui/src/less/main.less": server_ui_src_less_main_less,
	"server/ui/src/vendor/.DS_Store": server_ui_src_vendor_ds_store,
	"server/ui/src/vendor/bootstrap-3.3.4/alerts.less": server_ui_src_vendor_bootstrap_3_3_4_alerts_less,
	"server/ui/src/vendor/bootstrap-3.3.4/badges.less": server_ui_src_vendor_bootstrap_3_3_4_badges_less,
	"server/ui/src/vendor/bootstrap-3.3.4/bootstrap.less": server_ui_src_vendor_bootstrap_3_3_4_bootstrap_less,
	"server/ui/src/vendor/bootstrap-3.3.4/breadcrumbs.less": server_ui_src_vendor_bootstrap_3_3_4_breadcrumbs_less,
	"server/ui/src/vendor/bootstrap-3.3.4/button-groups.less": server_ui_src_vendor_bootstrap_3_3_4_button_groups_less,
	"server/ui/src/vendor/bootstrap-3.3.4/buttons.less": server_ui_src_vendor_bootstrap_3_3_4_buttons_less,
	"server/ui/src/vendor/bootstrap-3.3.4/carousel.less": server_ui_src_vendor_bootstrap_3_3_4_carousel_less,
	"server/ui/src/vendor/bootstrap-3.3.4/close.less": server_ui_src_vendor_bootstrap_3_3_4_close_less,
	"server/ui/src/vendor/bootstrap-3.3.4/code.less": server_ui_src_vendor_bootstrap_3_3_4_code_less,
	"server/ui/src/vendor/bootstrap-3.3.4/component-animations.less": server_ui_src_vendor_bootstrap_3_3_4_component_animations_less,
	"server/ui/src/vendor/bootstrap-3.3.4/dropdowns.less": server_ui_src_vendor_bootstrap_3_3_4_dropdowns_less,
	"server/ui/src/vendor/bootstrap-3.3.4/forms.less": server_ui_src_vendor_bootstrap_3_3_4_forms_less,
	"server/ui/src/vendor/bootstrap-3.3.4/glyphicons.less": server_ui_src_vendor_bootstrap_3_3_4_glyphicons_less,
	"server/ui/src/vendor/bootstrap-3.3.4/grid.less": server_ui_src_vendor_bootstrap_3_3_4_grid_less,
	"server/ui/src/vendor/bootstrap-3.3.4/input-groups.less": server_ui_src_vendor_bootstrap_3_3_4_input_groups_less,
	"server/ui/src/vendor/bootstrap-3.3.4/jumbotron.less": server_ui_src_vendor_bootstrap_3_3_4_jumbotron_less,
	"server/ui/src/vendor/bootstrap-3.3.4/labels.less": server_ui_src_vendor_bootstrap_3_3_4_labels_less,
	"server/ui/src/vendor/bootstrap-3.3.4/list-group.less": server_ui_src_vendor_bootstrap_3_3_4_list_group_less,
	"server/ui/src/vendor/bootstrap-3.3.4/media.less": server_ui_src_vendor_bootstrap_3_3_4_media_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/alerts.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_alerts_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/background-variant.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_background_variant_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/border-radius.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_border_radius_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/buttons.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_buttons_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/center-block.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_center_block_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/clearfix.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_clearfix_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/forms.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_forms_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/gradients.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_gradients_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/grid-framework.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_framework_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/grid.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/hide-text.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_hide_text_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/image.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_image_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/labels.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_labels_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/list-group.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_list_group_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-divider.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_divider_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/nav-vertical-align.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_vertical_align_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/opacity.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_opacity_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/pagination.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_pagination_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/panels.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_panels_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/progress-bar.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_progress_bar_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/reset-filter.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_reset_filter_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/resize.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_resize_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/responsive-visibility.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_responsive_visibility_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/size.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_size_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/tab-focus.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_tab_focus_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/table-row.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_table_row_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/text-emphasis.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_text_emphasis_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/text-overflow.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_text_overflow_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins/vendor-prefixes.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_vendor_prefixes_less,
	"server/ui/src/vendor/bootstrap-3.3.4/mixins.less": server_ui_src_vendor_bootstrap_3_3_4_mixins_less,
	"server/ui/src/vendor/bootstrap-3.3.4/modals.less": server_ui_src_vendor_bootstrap_3_3_4_modals_less,
	"server/ui/src/vendor/bootstrap-3.3.4/navbar.less": server_ui_src_vendor_bootstrap_3_3_4_navbar_less,
	"server/ui/src/vendor/bootstrap-3.3.4/navs.less": server_ui_src_vendor_bootstrap_3_3_4_navs_less,
	"server/ui/src/vendor/bootstrap-3.3.4/normalize.less": server_ui_src_vendor_bootstrap_3_3_4_normalize_less,
	"server/ui/src/vendor/bootstrap-3.3.4/pager.less": server_ui_src_vendor_bootstrap_3_3_4_pager_less,
	"server/ui/src/vendor/bootstrap-3.3.4/pagination.less": server_ui_src_vendor_bootstrap_3_3_4_pagination_less,
	"server/ui/src/vendor/bootstrap-3.3.4/panels.less": server_ui_src_vendor_bootstrap_3_3_4_panels_less,
	"server/ui/src/vendor/bootstrap-3.3.4/popovers.less": server_ui_src_vendor_bootstrap_3_3_4_popovers_less,
	"server/ui/src/vendor/bootstrap-3.3.4/print.less": server_ui_src_vendor_bootstrap_3_3_4_print_less,
	"server/ui/src/vendor/bootstrap-3.3.4/progress-bars.less": server_ui_src_vendor_bootstrap_3_3_4_progress_bars_less,
	"server/ui/src/vendor/bootstrap-3.3.4/responsive-embed.less": server_ui_src_vendor_bootstrap_3_3_4_responsive_embed_less,
	"server/ui/src/vendor/bootstrap-3.3.4/responsive-utilities.less": server_ui_src_vendor_bootstrap_3_3_4_responsive_utilities_less,
	"server/ui/src/vendor/bootstrap-3.3.4/scaffolding.less": server_ui_src_vendor_bootstrap_3_3_4_scaffolding_less,
	"server/ui/src/vendor/bootstrap-3.3.4/tables.less": server_ui_src_vendor_bootstrap_3_3_4_tables_less,
	"server/ui/src/vendor/bootstrap-3.3.4/theme.less": server_ui_src_vendor_bootstrap_3_3_4_theme_less,
	"server/ui/src/vendor/bootstrap-3.3.4/thumbnails.less": server_ui_src_vendor_bootstrap_3_3_4_thumbnails_less,
	"server/ui/src/vendor/bootstrap-3.3.4/tooltip.less": server_ui_src_vendor_bootstrap_3_3_4_tooltip_less,
	"server/ui/src/vendor/bootstrap-3.3.4/type.less": server_ui_src_vendor_bootstrap_3_3_4_type_less,
	"server/ui/src/vendor/bootstrap-3.3.4/utilities.less": server_ui_src_vendor_bootstrap_3_3_4_utilities_less,
	"server/ui/src/vendor/bootstrap-3.3.4/variables.less": server_ui_src_vendor_bootstrap_3_3_4_variables_less,
	"server/ui/src/vendor/bootstrap-3.3.4/wells.less": server_ui_src_vendor_bootstrap_3_3_4_wells_less,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"server": &_bintree_t{nil, map[string]*_bintree_t{
		"ui": &_bintree_t{nil, map[string]*_bintree_t{
			".DS_Store": &_bintree_t{server_ui_ds_store, map[string]*_bintree_t{
			}},
			".gitignore": &_bintree_t{server_ui_gitignore, map[string]*_bintree_t{
			}},
			".jshintrc": &_bintree_t{server_ui_jshintrc, map[string]*_bintree_t{
			}},
			"add_dep.html": &_bintree_t{server_ui_add_dep_html, map[string]*_bintree_t{
			}},
			"add_state.html": &_bintree_t{server_ui_add_state_html, map[string]*_bintree_t{
			}},
			"dashboard.html": &_bintree_t{server_ui_dashboard_html, map[string]*_bintree_t{
			}},
			"dist": &_bintree_t{nil, map[string]*_bintree_t{
				"app.js": &_bintree_t{server_ui_dist_app_js, map[string]*_bintree_t{
				}},
				"main.css": &_bintree_t{server_ui_dist_main_css, map[string]*_bintree_t{
				}},
			}},
			"edit_state.html": &_bintree_t{server_ui_edit_state_html, map[string]*_bintree_t{
			}},
			"foot.html": &_bintree_t{server_ui_foot_html, map[string]*_bintree_t{
			}},
			"gulpfile.js": &_bintree_t{server_ui_gulpfile_js, map[string]*_bintree_t{
			}},
			"head.html": &_bintree_t{server_ui_head_html, map[string]*_bintree_t{
			}},
			"one_isolation.html": &_bintree_t{server_ui_one_isolation_html, map[string]*_bintree_t{
			}},
			"package.json": &_bintree_t{server_ui_package_json, map[string]*_bintree_t{
			}},
			"src": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{server_ui_src_ds_store, map[string]*_bintree_t{
				}},
				"js": &_bintree_t{nil, map[string]*_bintree_t{
					"actions": &_bintree_t{nil, map[string]*_bintree_t{
						"DepActions.js": &_bintree_t{server_ui_src_js_actions_depactions_js, map[string]*_bintree_t{
						}},
						"IsolationActions.js": &_bintree_t{server_ui_src_js_actions_isolationactions_js, map[string]*_bintree_t{
						}},
					}},
					"app.js": &_bintree_t{server_ui_src_js_app_js, map[string]*_bintree_t{
					}},
					"components": &_bintree_t{nil, map[string]*_bintree_t{
						"DepForm.jsx": &_bintree_t{server_ui_src_js_components_depform_jsx, map[string]*_bintree_t{
						}},
						"DepItem.jsx": &_bintree_t{server_ui_src_js_components_depitem_jsx, map[string]*_bintree_t{
						}},
						"DepList.jsx": &_bintree_t{server_ui_src_js_components_deplist_jsx, map[string]*_bintree_t{
						}},
						"DepPanel.jsx": &_bintree_t{server_ui_src_js_components_deppanel_jsx, map[string]*_bintree_t{
						}},
						"IsolationForm.jsx": &_bintree_t{server_ui_src_js_components_isolationform_jsx, map[string]*_bintree_t{
						}},
						"IsolationItem.jsx": &_bintree_t{server_ui_src_js_components_isolationitem_jsx, map[string]*_bintree_t{
						}},
						"IsolationList.jsx": &_bintree_t{server_ui_src_js_components_isolationlist_jsx, map[string]*_bintree_t{
						}},
						"IsolationPanel.jsx": &_bintree_t{server_ui_src_js_components_isolationpanel_jsx, map[string]*_bintree_t{
						}},
					}},
					"dispatcher": &_bintree_t{nil, map[string]*_bintree_t{
						"AppDispatcher.js": &_bintree_t{server_ui_src_js_dispatcher_appdispatcher_js, map[string]*_bintree_t{
						}},
					}},
					"stores": &_bintree_t{nil, map[string]*_bintree_t{
						"DepStore.js": &_bintree_t{server_ui_src_js_stores_depstore_js, map[string]*_bintree_t{
						}},
						"IsolationStore.js": &_bintree_t{server_ui_src_js_stores_isolationstore_js, map[string]*_bintree_t{
						}},
					}},
				}},
				"less": &_bintree_t{nil, map[string]*_bintree_t{
					"main.less": &_bintree_t{server_ui_src_less_main_less, map[string]*_bintree_t{
					}},
				}},
				"vendor": &_bintree_t{nil, map[string]*_bintree_t{
					".DS_Store": &_bintree_t{server_ui_src_vendor_ds_store, map[string]*_bintree_t{
					}},
					"bootstrap-3.3.4": &_bintree_t{nil, map[string]*_bintree_t{
						"alerts.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_alerts_less, map[string]*_bintree_t{
						}},
						"badges.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_badges_less, map[string]*_bintree_t{
						}},
						"bootstrap.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_bootstrap_less, map[string]*_bintree_t{
						}},
						"breadcrumbs.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_breadcrumbs_less, map[string]*_bintree_t{
						}},
						"button-groups.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_button_groups_less, map[string]*_bintree_t{
						}},
						"buttons.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_buttons_less, map[string]*_bintree_t{
						}},
						"carousel.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_carousel_less, map[string]*_bintree_t{
						}},
						"close.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_close_less, map[string]*_bintree_t{
						}},
						"code.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_code_less, map[string]*_bintree_t{
						}},
						"component-animations.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_component_animations_less, map[string]*_bintree_t{
						}},
						"dropdowns.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_dropdowns_less, map[string]*_bintree_t{
						}},
						"forms.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_forms_less, map[string]*_bintree_t{
						}},
						"glyphicons.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_glyphicons_less, map[string]*_bintree_t{
						}},
						"grid.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_grid_less, map[string]*_bintree_t{
						}},
						"input-groups.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_input_groups_less, map[string]*_bintree_t{
						}},
						"jumbotron.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_jumbotron_less, map[string]*_bintree_t{
						}},
						"labels.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_labels_less, map[string]*_bintree_t{
						}},
						"list-group.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_list_group_less, map[string]*_bintree_t{
						}},
						"media.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_media_less, map[string]*_bintree_t{
						}},
						"mixins": &_bintree_t{nil, map[string]*_bintree_t{
							"alerts.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_alerts_less, map[string]*_bintree_t{
							}},
							"background-variant.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_background_variant_less, map[string]*_bintree_t{
							}},
							"border-radius.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_border_radius_less, map[string]*_bintree_t{
							}},
							"buttons.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_buttons_less, map[string]*_bintree_t{
							}},
							"center-block.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_center_block_less, map[string]*_bintree_t{
							}},
							"clearfix.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_clearfix_less, map[string]*_bintree_t{
							}},
							"forms.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_forms_less, map[string]*_bintree_t{
							}},
							"gradients.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_gradients_less, map[string]*_bintree_t{
							}},
							"grid-framework.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_framework_less, map[string]*_bintree_t{
							}},
							"grid.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_grid_less, map[string]*_bintree_t{
							}},
							"hide-text.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_hide_text_less, map[string]*_bintree_t{
							}},
							"image.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_image_less, map[string]*_bintree_t{
							}},
							"labels.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_labels_less, map[string]*_bintree_t{
							}},
							"list-group.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_list_group_less, map[string]*_bintree_t{
							}},
							"nav-divider.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_divider_less, map[string]*_bintree_t{
							}},
							"nav-vertical-align.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_nav_vertical_align_less, map[string]*_bintree_t{
							}},
							"opacity.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_opacity_less, map[string]*_bintree_t{
							}},
							"pagination.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_pagination_less, map[string]*_bintree_t{
							}},
							"panels.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_panels_less, map[string]*_bintree_t{
							}},
							"progress-bar.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_progress_bar_less, map[string]*_bintree_t{
							}},
							"reset-filter.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_reset_filter_less, map[string]*_bintree_t{
							}},
							"resize.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_resize_less, map[string]*_bintree_t{
							}},
							"responsive-visibility.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_responsive_visibility_less, map[string]*_bintree_t{
							}},
							"size.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_size_less, map[string]*_bintree_t{
							}},
							"tab-focus.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_tab_focus_less, map[string]*_bintree_t{
							}},
							"table-row.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_table_row_less, map[string]*_bintree_t{
							}},
							"text-emphasis.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_text_emphasis_less, map[string]*_bintree_t{
							}},
							"text-overflow.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_text_overflow_less, map[string]*_bintree_t{
							}},
							"vendor-prefixes.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_vendor_prefixes_less, map[string]*_bintree_t{
							}},
						}},
						"mixins.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_mixins_less, map[string]*_bintree_t{
						}},
						"modals.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_modals_less, map[string]*_bintree_t{
						}},
						"navbar.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_navbar_less, map[string]*_bintree_t{
						}},
						"navs.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_navs_less, map[string]*_bintree_t{
						}},
						"normalize.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_normalize_less, map[string]*_bintree_t{
						}},
						"pager.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_pager_less, map[string]*_bintree_t{
						}},
						"pagination.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_pagination_less, map[string]*_bintree_t{
						}},
						"panels.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_panels_less, map[string]*_bintree_t{
						}},
						"popovers.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_popovers_less, map[string]*_bintree_t{
						}},
						"print.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_print_less, map[string]*_bintree_t{
						}},
						"progress-bars.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_progress_bars_less, map[string]*_bintree_t{
						}},
						"responsive-embed.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_responsive_embed_less, map[string]*_bintree_t{
						}},
						"responsive-utilities.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_responsive_utilities_less, map[string]*_bintree_t{
						}},
						"scaffolding.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_scaffolding_less, map[string]*_bintree_t{
						}},
						"tables.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_tables_less, map[string]*_bintree_t{
						}},
						"theme.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_theme_less, map[string]*_bintree_t{
						}},
						"thumbnails.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_thumbnails_less, map[string]*_bintree_t{
						}},
						"tooltip.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_tooltip_less, map[string]*_bintree_t{
						}},
						"type.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_type_less, map[string]*_bintree_t{
						}},
						"utilities.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_utilities_less, map[string]*_bintree_t{
						}},
						"variables.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_variables_less, map[string]*_bintree_t{
						}},
						"wells.less": &_bintree_t{server_ui_src_vendor_bootstrap_3_3_4_wells_less, map[string]*_bintree_t{
						}},
					}},
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

