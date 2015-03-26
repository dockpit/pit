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

// server_ui_vendor_ds_store reads file data from disk. It returns an error on failure.
func server_ui_vendor_ds_store() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/.DS_Store"
	name := "server/ui/vendor/.DS_Store"
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

// server_ui_vendor_semantic_ui_1_11_5_csscomb_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_csscomb_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/.csscomb.json"
	name := "server/ui/vendor/semantic-ui-1.11.5/.csscomb.json"
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

// server_ui_vendor_semantic_ui_1_11_5_csslintrc reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_csslintrc() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/.csslintrc"
	name := "server/ui/vendor/semantic-ui-1.11.5/.csslintrc"
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

// server_ui_vendor_semantic_ui_1_11_5_gitignore reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_gitignore() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/.gitignore"
	name := "server/ui/vendor/semantic-ui-1.11.5/.gitignore"
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

// server_ui_vendor_semantic_ui_1_11_5_jshintrc reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_jshintrc() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/.jshintrc"
	name := "server/ui/vendor/semantic-ui-1.11.5/.jshintrc"
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

// server_ui_vendor_semantic_ui_1_11_5_contributing_md reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_contributing_md() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/CONTRIBUTING.md"
	name := "server/ui/vendor/semantic-ui-1.11.5/CONTRIBUTING.md"
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

// server_ui_vendor_semantic_ui_1_11_5_license_md reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_license_md() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/LICENSE.md"
	name := "server/ui/vendor/semantic-ui-1.11.5/LICENSE.md"
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

// server_ui_vendor_semantic_ui_1_11_5_readme_md reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_readme_md() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/README.md"
	name := "server/ui/vendor/semantic-ui-1.11.5/README.md"
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

// server_ui_vendor_semantic_ui_1_11_5_release_notes_md reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_release_notes_md() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/RELEASE-NOTES.md"
	name := "server/ui/vendor/semantic-ui-1.11.5/RELEASE-NOTES.md"
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

// server_ui_vendor_semantic_ui_1_11_5_bower_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_bower_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/bower.json"
	name := "server/ui/vendor/semantic-ui-1.11.5/bower.json"
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

// server_ui_vendor_semantic_ui_1_11_5_composer_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_composer_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/composer.json"
	name := "server/ui/vendor/semantic-ui-1.11.5/composer.json"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_api_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_api_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/api.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/api.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_api_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_api_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/api.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/api.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_button_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_button_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/button.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/button.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_button_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_button_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/button.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/button.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_card_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_card_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/card.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/card.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_card_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_card_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/card.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/card.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_form_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_form_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/form.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/form.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_form_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_form_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/form.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/form.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_header_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_header_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/header.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/header.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_header_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_header_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/header.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/header.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_image_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_image_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/image.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/image.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_image_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_image_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/image.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/image.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_input_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_input_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/input.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/input.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_input_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_input_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/input.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/input.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_item_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_item_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/item.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/item.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_item_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_item_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/item.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/item.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_label_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_label_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/label.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/label.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_label_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_label_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/label.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/label.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_list_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_list_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/list.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/list.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_list_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_list_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/list.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/list.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_message_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_message_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/message.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/message.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_message_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_message_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/message.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/message.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_search_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_search_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/search.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/search.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_search_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_search_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/search.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/search.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_site_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_site_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/site.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/site.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_site_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_site_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/site.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/site.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_state_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_state_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/state.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/state.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_state_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_state_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/state.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/state.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_step_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_step_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/step.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/step.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_step_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_step_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/step.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/step.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_table_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_table_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/table.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/table.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_table_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_table_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/table.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/table.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_video_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_video_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/video.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/video.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_video_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_video_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/video.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/video.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_semantic_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_semantic_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/semantic.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/semantic.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_semantic_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_semantic_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/semantic.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/semantic.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_css reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_css() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.css"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.css"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.js"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_eot reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_eot() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.eot"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.eot"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_svg reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_svg() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.svg"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.svg"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_ttf reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_ttf() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.ttf"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.ttf"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_woff reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_woff() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.woff"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.woff"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_eot reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_eot() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.eot"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.eot"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_otf reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_otf() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.otf"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.otf"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_svg reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_svg() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.svg"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.svg"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_ttf reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_ttf() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.ttf"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.ttf"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff2 reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff2() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff2"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff2"
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

// server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_images_flags_png reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_images_flags_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/images/flags.png"
	name := "server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/images/flags.png"
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

// server_ui_vendor_semantic_ui_1_11_5_gulpfile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_gulpfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/gulpfile.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/gulpfile.js"
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

// server_ui_vendor_semantic_ui_1_11_5_karma_conf_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_karma_conf_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/karma.conf.js"
	name := "server/ui/vendor/semantic-ui-1.11.5/karma.conf.js"
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

// server_ui_vendor_semantic_ui_1_11_5_logo_png reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_logo_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/logo.png"
	name := "server/ui/vendor/semantic-ui-1.11.5/logo.png"
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

// server_ui_vendor_semantic_ui_1_11_5_package_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_package_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/package.json"
	name := "server/ui/vendor/semantic-ui-1.11.5/package.json"
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

// server_ui_vendor_semantic_ui_1_11_5_semantic_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_semantic_ui_1_11_5_semantic_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/semantic-ui-1.11.5/semantic.json"
	name := "server/ui/vendor/semantic-ui-1.11.5/semantic.json"
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
	"server/ui/vendor/.DS_Store": server_ui_vendor_ds_store,
	"server/ui/vendor/semantic-ui-1.11.5/.csscomb.json": server_ui_vendor_semantic_ui_1_11_5_csscomb_json,
	"server/ui/vendor/semantic-ui-1.11.5/.csslintrc": server_ui_vendor_semantic_ui_1_11_5_csslintrc,
	"server/ui/vendor/semantic-ui-1.11.5/.gitignore": server_ui_vendor_semantic_ui_1_11_5_gitignore,
	"server/ui/vendor/semantic-ui-1.11.5/.jshintrc": server_ui_vendor_semantic_ui_1_11_5_jshintrc,
	"server/ui/vendor/semantic-ui-1.11.5/CONTRIBUTING.md": server_ui_vendor_semantic_ui_1_11_5_contributing_md,
	"server/ui/vendor/semantic-ui-1.11.5/LICENSE.md": server_ui_vendor_semantic_ui_1_11_5_license_md,
	"server/ui/vendor/semantic-ui-1.11.5/README.md": server_ui_vendor_semantic_ui_1_11_5_readme_md,
	"server/ui/vendor/semantic-ui-1.11.5/RELEASE-NOTES.md": server_ui_vendor_semantic_ui_1_11_5_release_notes_md,
	"server/ui/vendor/semantic-ui-1.11.5/bower.json": server_ui_vendor_semantic_ui_1_11_5_bower_json,
	"server/ui/vendor/semantic-ui-1.11.5/composer.json": server_ui_vendor_semantic_ui_1_11_5_composer_json,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/accordion.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/ad.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/api.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_api_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/api.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_api_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/breadcrumb.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/button.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_button_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/button.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_button_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/card.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_card_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/card.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_card_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/checkbox.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/comment.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dimmer.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/divider.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/dropdown.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/feed.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/flag.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/form.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_form_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/form.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_form_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/form.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/grid.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/header.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_header_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/header.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_header_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/icon.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/image.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_image_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/image.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_image_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/input.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_input_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/input.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_input_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/item.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_item_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/item.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_item_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/label.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_label_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/label.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_label_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/list.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_list_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/list.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_list_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/loader.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/menu.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/message.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_message_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/message.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_message_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/modal.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/nag.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/popup.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/progress.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rail.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/rating.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/reset.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/reveal.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/search.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_search_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/search.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_search_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/search.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/segment.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/shape.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sidebar.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/site.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_site_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/site.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_site_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/site.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/state.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_state_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/state.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_state_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/statistic.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/step.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_step_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/step.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_step_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/sticky.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/tab.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/table.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_table_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/table.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_table_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/transition.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/video.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_video_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/video.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_video_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/video.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/components/visibility.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/semantic.css": server_ui_vendor_semantic_ui_1_11_5_dist_semantic_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/semantic.js": server_ui_vendor_semantic_ui_1_11_5_dist_semantic_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.css": server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_css,
	"server/ui/vendor/semantic-ui-1.11.5/dist/semantic.min.js": server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_js,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.eot": server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_eot,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.svg": server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_svg,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.ttf": server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_ttf,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/basic/assets/fonts/icons.woff": server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_woff,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.eot": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_eot,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.otf": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_otf,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.svg": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_svg,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.ttf": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_ttf,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/fonts/icons.woff2": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff2,
	"server/ui/vendor/semantic-ui-1.11.5/dist/themes/default/assets/images/flags.png": server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_images_flags_png,
	"server/ui/vendor/semantic-ui-1.11.5/gulpfile.js": server_ui_vendor_semantic_ui_1_11_5_gulpfile_js,
	"server/ui/vendor/semantic-ui-1.11.5/karma.conf.js": server_ui_vendor_semantic_ui_1_11_5_karma_conf_js,
	"server/ui/vendor/semantic-ui-1.11.5/logo.png": server_ui_vendor_semantic_ui_1_11_5_logo_png,
	"server/ui/vendor/semantic-ui-1.11.5/package.json": server_ui_vendor_semantic_ui_1_11_5_package_json,
	"server/ui/vendor/semantic-ui-1.11.5/semantic.json": server_ui_vendor_semantic_ui_1_11_5_semantic_json,
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
			}},
			"vendor": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{server_ui_vendor_ds_store, map[string]*_bintree_t{
				}},
				"semantic-ui-1.11.5": &_bintree_t{nil, map[string]*_bintree_t{
					".csscomb.json": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_csscomb_json, map[string]*_bintree_t{
					}},
					".csslintrc": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_csslintrc, map[string]*_bintree_t{
					}},
					".gitignore": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_gitignore, map[string]*_bintree_t{
					}},
					".jshintrc": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_jshintrc, map[string]*_bintree_t{
					}},
					"CONTRIBUTING.md": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_contributing_md, map[string]*_bintree_t{
					}},
					"LICENSE.md": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_license_md, map[string]*_bintree_t{
					}},
					"README.md": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_readme_md, map[string]*_bintree_t{
					}},
					"RELEASE-NOTES.md": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_release_notes_md, map[string]*_bintree_t{
					}},
					"bower.json": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_bower_json, map[string]*_bintree_t{
					}},
					"composer.json": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_composer_json, map[string]*_bintree_t{
					}},
					"dist": &_bintree_t{nil, map[string]*_bintree_t{
						"components": &_bintree_t{nil, map[string]*_bintree_t{
							"accordion.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_css, map[string]*_bintree_t{
							}},
							"accordion.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_js, map[string]*_bintree_t{
							}},
							"accordion.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_css, map[string]*_bintree_t{
							}},
							"accordion.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_accordion_min_js, map[string]*_bintree_t{
							}},
							"ad.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_css, map[string]*_bintree_t{
							}},
							"ad.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_ad_min_css, map[string]*_bintree_t{
							}},
							"api.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_api_js, map[string]*_bintree_t{
							}},
							"api.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_api_min_js, map[string]*_bintree_t{
							}},
							"breadcrumb.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_css, map[string]*_bintree_t{
							}},
							"breadcrumb.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_breadcrumb_min_css, map[string]*_bintree_t{
							}},
							"button.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_button_css, map[string]*_bintree_t{
							}},
							"button.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_button_min_css, map[string]*_bintree_t{
							}},
							"card.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_card_css, map[string]*_bintree_t{
							}},
							"card.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_card_min_css, map[string]*_bintree_t{
							}},
							"checkbox.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_css, map[string]*_bintree_t{
							}},
							"checkbox.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_js, map[string]*_bintree_t{
							}},
							"checkbox.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_css, map[string]*_bintree_t{
							}},
							"checkbox.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_checkbox_min_js, map[string]*_bintree_t{
							}},
							"comment.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_css, map[string]*_bintree_t{
							}},
							"comment.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_comment_min_css, map[string]*_bintree_t{
							}},
							"dimmer.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_css, map[string]*_bintree_t{
							}},
							"dimmer.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_js, map[string]*_bintree_t{
							}},
							"dimmer.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_css, map[string]*_bintree_t{
							}},
							"dimmer.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dimmer_min_js, map[string]*_bintree_t{
							}},
							"divider.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_css, map[string]*_bintree_t{
							}},
							"divider.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_divider_min_css, map[string]*_bintree_t{
							}},
							"dropdown.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_css, map[string]*_bintree_t{
							}},
							"dropdown.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_js, map[string]*_bintree_t{
							}},
							"dropdown.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_css, map[string]*_bintree_t{
							}},
							"dropdown.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_dropdown_min_js, map[string]*_bintree_t{
							}},
							"feed.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_css, map[string]*_bintree_t{
							}},
							"feed.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_feed_min_css, map[string]*_bintree_t{
							}},
							"flag.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_css, map[string]*_bintree_t{
							}},
							"flag.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_flag_min_css, map[string]*_bintree_t{
							}},
							"form.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_form_css, map[string]*_bintree_t{
							}},
							"form.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_form_js, map[string]*_bintree_t{
							}},
							"form.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_css, map[string]*_bintree_t{
							}},
							"form.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_form_min_js, map[string]*_bintree_t{
							}},
							"grid.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_css, map[string]*_bintree_t{
							}},
							"grid.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_grid_min_css, map[string]*_bintree_t{
							}},
							"header.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_header_css, map[string]*_bintree_t{
							}},
							"header.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_header_min_css, map[string]*_bintree_t{
							}},
							"icon.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_css, map[string]*_bintree_t{
							}},
							"icon.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_icon_min_css, map[string]*_bintree_t{
							}},
							"image.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_image_css, map[string]*_bintree_t{
							}},
							"image.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_image_min_css, map[string]*_bintree_t{
							}},
							"input.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_input_css, map[string]*_bintree_t{
							}},
							"input.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_input_min_css, map[string]*_bintree_t{
							}},
							"item.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_item_css, map[string]*_bintree_t{
							}},
							"item.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_item_min_css, map[string]*_bintree_t{
							}},
							"label.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_label_css, map[string]*_bintree_t{
							}},
							"label.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_label_min_css, map[string]*_bintree_t{
							}},
							"list.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_list_css, map[string]*_bintree_t{
							}},
							"list.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_list_min_css, map[string]*_bintree_t{
							}},
							"loader.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_css, map[string]*_bintree_t{
							}},
							"loader.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_loader_min_css, map[string]*_bintree_t{
							}},
							"menu.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_css, map[string]*_bintree_t{
							}},
							"menu.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_menu_min_css, map[string]*_bintree_t{
							}},
							"message.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_message_css, map[string]*_bintree_t{
							}},
							"message.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_message_min_css, map[string]*_bintree_t{
							}},
							"modal.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_css, map[string]*_bintree_t{
							}},
							"modal.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_js, map[string]*_bintree_t{
							}},
							"modal.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_css, map[string]*_bintree_t{
							}},
							"modal.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_modal_min_js, map[string]*_bintree_t{
							}},
							"nag.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_css, map[string]*_bintree_t{
							}},
							"nag.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_js, map[string]*_bintree_t{
							}},
							"nag.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_css, map[string]*_bintree_t{
							}},
							"nag.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_nag_min_js, map[string]*_bintree_t{
							}},
							"popup.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_css, map[string]*_bintree_t{
							}},
							"popup.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_js, map[string]*_bintree_t{
							}},
							"popup.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_css, map[string]*_bintree_t{
							}},
							"popup.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_popup_min_js, map[string]*_bintree_t{
							}},
							"progress.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_css, map[string]*_bintree_t{
							}},
							"progress.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_js, map[string]*_bintree_t{
							}},
							"progress.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_css, map[string]*_bintree_t{
							}},
							"progress.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_progress_min_js, map[string]*_bintree_t{
							}},
							"rail.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_css, map[string]*_bintree_t{
							}},
							"rail.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rail_min_css, map[string]*_bintree_t{
							}},
							"rating.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_css, map[string]*_bintree_t{
							}},
							"rating.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_js, map[string]*_bintree_t{
							}},
							"rating.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_css, map[string]*_bintree_t{
							}},
							"rating.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_rating_min_js, map[string]*_bintree_t{
							}},
							"reset.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_css, map[string]*_bintree_t{
							}},
							"reset.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_reset_min_css, map[string]*_bintree_t{
							}},
							"reveal.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_css, map[string]*_bintree_t{
							}},
							"reveal.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_reveal_min_css, map[string]*_bintree_t{
							}},
							"search.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_search_css, map[string]*_bintree_t{
							}},
							"search.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_search_js, map[string]*_bintree_t{
							}},
							"search.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_css, map[string]*_bintree_t{
							}},
							"search.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_search_min_js, map[string]*_bintree_t{
							}},
							"segment.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_css, map[string]*_bintree_t{
							}},
							"segment.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_segment_min_css, map[string]*_bintree_t{
							}},
							"shape.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_css, map[string]*_bintree_t{
							}},
							"shape.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_js, map[string]*_bintree_t{
							}},
							"shape.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_css, map[string]*_bintree_t{
							}},
							"shape.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_shape_min_js, map[string]*_bintree_t{
							}},
							"sidebar.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_css, map[string]*_bintree_t{
							}},
							"sidebar.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_js, map[string]*_bintree_t{
							}},
							"sidebar.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_css, map[string]*_bintree_t{
							}},
							"sidebar.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sidebar_min_js, map[string]*_bintree_t{
							}},
							"site.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_site_css, map[string]*_bintree_t{
							}},
							"site.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_site_js, map[string]*_bintree_t{
							}},
							"site.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_css, map[string]*_bintree_t{
							}},
							"site.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_site_min_js, map[string]*_bintree_t{
							}},
							"state.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_state_js, map[string]*_bintree_t{
							}},
							"state.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_state_min_js, map[string]*_bintree_t{
							}},
							"statistic.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_css, map[string]*_bintree_t{
							}},
							"statistic.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_statistic_min_css, map[string]*_bintree_t{
							}},
							"step.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_step_css, map[string]*_bintree_t{
							}},
							"step.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_step_min_css, map[string]*_bintree_t{
							}},
							"sticky.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_css, map[string]*_bintree_t{
							}},
							"sticky.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_js, map[string]*_bintree_t{
							}},
							"sticky.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_css, map[string]*_bintree_t{
							}},
							"sticky.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_sticky_min_js, map[string]*_bintree_t{
							}},
							"tab.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_css, map[string]*_bintree_t{
							}},
							"tab.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_js, map[string]*_bintree_t{
							}},
							"tab.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_css, map[string]*_bintree_t{
							}},
							"tab.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_tab_min_js, map[string]*_bintree_t{
							}},
							"table.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_table_css, map[string]*_bintree_t{
							}},
							"table.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_table_min_css, map[string]*_bintree_t{
							}},
							"transition.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_css, map[string]*_bintree_t{
							}},
							"transition.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_js, map[string]*_bintree_t{
							}},
							"transition.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_css, map[string]*_bintree_t{
							}},
							"transition.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_transition_min_js, map[string]*_bintree_t{
							}},
							"video.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_video_css, map[string]*_bintree_t{
							}},
							"video.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_video_js, map[string]*_bintree_t{
							}},
							"video.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_css, map[string]*_bintree_t{
							}},
							"video.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_video_min_js, map[string]*_bintree_t{
							}},
							"visibility.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_js, map[string]*_bintree_t{
							}},
							"visibility.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_components_visibility_min_js, map[string]*_bintree_t{
							}},
						}},
						"semantic.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_semantic_css, map[string]*_bintree_t{
						}},
						"semantic.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_semantic_js, map[string]*_bintree_t{
						}},
						"semantic.min.css": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_css, map[string]*_bintree_t{
						}},
						"semantic.min.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_semantic_min_js, map[string]*_bintree_t{
						}},
						"themes": &_bintree_t{nil, map[string]*_bintree_t{
							"basic": &_bintree_t{nil, map[string]*_bintree_t{
								"assets": &_bintree_t{nil, map[string]*_bintree_t{
									"fonts": &_bintree_t{nil, map[string]*_bintree_t{
										"icons.eot": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_eot, map[string]*_bintree_t{
										}},
										"icons.svg": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_svg, map[string]*_bintree_t{
										}},
										"icons.ttf": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_ttf, map[string]*_bintree_t{
										}},
										"icons.woff": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_basic_assets_fonts_icons_woff, map[string]*_bintree_t{
										}},
									}},
								}},
							}},
							"default": &_bintree_t{nil, map[string]*_bintree_t{
								"assets": &_bintree_t{nil, map[string]*_bintree_t{
									"fonts": &_bintree_t{nil, map[string]*_bintree_t{
										"icons.eot": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_eot, map[string]*_bintree_t{
										}},
										"icons.otf": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_otf, map[string]*_bintree_t{
										}},
										"icons.svg": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_svg, map[string]*_bintree_t{
										}},
										"icons.ttf": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_ttf, map[string]*_bintree_t{
										}},
										"icons.woff": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff, map[string]*_bintree_t{
										}},
										"icons.woff2": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_fonts_icons_woff2, map[string]*_bintree_t{
										}},
									}},
									"images": &_bintree_t{nil, map[string]*_bintree_t{
										"flags.png": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_dist_themes_default_assets_images_flags_png, map[string]*_bintree_t{
										}},
									}},
								}},
							}},
						}},
					}},
					"gulpfile.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_gulpfile_js, map[string]*_bintree_t{
					}},
					"karma.conf.js": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_karma_conf_js, map[string]*_bintree_t{
					}},
					"logo.png": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_logo_png, map[string]*_bintree_t{
					}},
					"package.json": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_package_json, map[string]*_bintree_t{
					}},
					"semantic.json": &_bintree_t{server_ui_vendor_semantic_ui_1_11_5_semantic_json, map[string]*_bintree_t{
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

