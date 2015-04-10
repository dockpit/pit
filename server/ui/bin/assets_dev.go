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

// server_ui_editor_html reads file data from disk. It returns an error on failure.
func server_ui_editor_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/editor.html"
	name := "server/ui/editor.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_img_icon_build_image_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_build_image_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_build_image.png"
	name := "server/ui/src/img/icon_build_image.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_build_image_1x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_build_image_1x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_build_image@1x.png"
	name := "server/ui/src/img/icon_build_image@1x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_built_image_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_built_image_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_built_image.png"
	name := "server/ui/src/img/icon_built_image.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_dep_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_dep_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_dep.png"
	name := "server/ui/src/img/icon_dep.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_dep_1x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_dep_1x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_dep@1x.png"
	name := "server/ui/src/img/icon_dep@1x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_fixture_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_fixture_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_fixture.png"
	name := "server/ui/src/img/icon_fixture.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_fixture_1x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_fixture_1x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_fixture@1x.png"
	name := "server/ui/src/img/icon_fixture@1x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_run_image_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_run_image_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_run_image.png"
	name := "server/ui/src/img/icon_run_image.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_run_image_1x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_run_image_1x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_run_image@1x.png"
	name := "server/ui/src/img/icon_run_image@1x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_api_blueprint_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_api_blueprint_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_api-blueprint.png"
	name := "server/ui/src/img/icon_t_api-blueprint.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_api_blueprint_0_7x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_api_blueprint_0_7x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_api-blueprint@0,7x.png"
	name := "server/ui/src/img/icon_t_api-blueprint@0,7x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_custom_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_custom_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_custom.png"
	name := "server/ui/src/img/icon_t_custom.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_custom_0_7x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_custom_0_7x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_custom@0,7x.png"
	name := "server/ui/src/img/icon_t_custom@0,7x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_etcd_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_etcd_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_etcd.png"
	name := "server/ui/src/img/icon_t_etcd.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_etcd_0_7x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_etcd_0_7x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_etcd@0,7x.png"
	name := "server/ui/src/img/icon_t_etcd@0,7x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_postgres_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_postgres_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_postgres.png"
	name := "server/ui/src/img/icon_t_postgres.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_t_postgres_0_7x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_t_postgres_0_7x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_t_postgres@0,7x.png"
	name := "server/ui/src/img/icon_t_postgres@0,7x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_icon_unbuilt_image_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_icon_unbuilt_image_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/icon_unbuilt_image.png"
	name := "server/ui/src/img/icon_unbuilt_image.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_logo_brand_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_logo_brand_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/logo_brand.png"
	name := "server/ui/src/img/logo_brand.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_img_logo_brand_2x_png reads file data from disk. It returns an error on failure.
func server_ui_src_img_logo_brand_2x_png() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/img/logo_brand@2x.png"
	name := "server/ui/src/img/logo_brand@2x.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_actions_editoractions_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_actions_editoractions_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/actions/EditorActions.js"
	name := "server/ui/src/js/actions/EditorActions.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_actions_statactions_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_actions_statactions_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/actions/StatActions.js"
	name := "server/ui/src/js/actions/StatActions.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_components_achievementpopup_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_achievementpopup_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/AchievementPopup.jsx"
	name := "server/ui/src/js/components/AchievementPopup.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_components_editorfilepanel_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_editorfilepanel_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/EditorFilePanel.jsx"
	name := "server/ui/src/js/components/EditorFilePanel.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_editorrunpanel_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_editorrunpanel_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/EditorRunPanel.jsx"
	name := "server/ui/src/js/components/EditorRunPanel.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_src_js_components_editorworkspace_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_editorworkspace_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/EditorWorkspace.jsx"
	name := "server/ui/src/js/components/EditorWorkspace.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_stores_editorstore_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_stores_editorstore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/stores/EditorStore.js"
	name := "server/ui/src/js/stores/EditorStore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_stores_statstore_js reads file data from disk. It returns an error on failure.
func server_ui_src_js_stores_statstore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/stores/StatStore.js"
	name := "server/ui/src/js/stores/StatStore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_templates_json reads file data from disk. It returns an error on failure.
func server_ui_templates_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/templates.json"
	name := "server/ui/templates.json"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_vendor_ace_1_1_8_changelog_txt reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_changelog_txt() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/ChangeLog.txt"
	name := "server/ui/vendor/ace-1.1.8/ChangeLog.txt"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_license reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_license() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/LICENSE"
	name := "server/ui/vendor/ace-1.1.8/LICENSE"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_readme_md reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_readme_md() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/README.md"
	name := "server/ui/vendor/ace-1.1.8/README.md"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_package_json reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_package_json() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/package.json"
	name := "server/ui/vendor/ace-1.1.8/package.json"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ace_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ace_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ace.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ace.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_beautify_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_beautify_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-beautify.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-beautify.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_chromevox_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_chromevox_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-chromevox.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-chromevox.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_elastic_tabstops_lite_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_elastic_tabstops_lite_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-elastic_tabstops_lite.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-elastic_tabstops_lite.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_emmet_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_emmet_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-emmet.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-emmet.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_error_marker_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_error_marker_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-error_marker.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-error_marker.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_keybinding_menu_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_keybinding_menu_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-keybinding_menu.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-keybinding_menu.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_language_tools_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_language_tools_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-language_tools.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-language_tools.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_linking_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_linking_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-linking.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-linking.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_modelist_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_modelist_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-modelist.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-modelist.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_old_ie_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_old_ie_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-old_ie.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-old_ie.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_searchbox_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_searchbox_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-searchbox.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-searchbox.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_settings_menu_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_settings_menu_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-settings_menu.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-settings_menu.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_spellcheck_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_spellcheck_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-spellcheck.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-spellcheck.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_split_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_split_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-split.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-split.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_static_highlight_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_static_highlight_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-static_highlight.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-static_highlight.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_statusbar_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_statusbar_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-statusbar.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-statusbar.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_textarea_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_textarea_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-textarea.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-textarea.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_themelist_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_themelist_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-themelist.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-themelist.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_whitespace_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_whitespace_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-whitespace.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-whitespace.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_emacs_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_emacs_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-emacs.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-emacs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_vim_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_vim_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-vim.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-vim.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_abap_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_abap_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-abap.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-abap.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_actionscript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_actionscript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-actionscript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-actionscript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ada_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ada_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ada.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ada.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_apache_conf_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_apache_conf_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-apache_conf.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-apache_conf.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_applescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_applescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-applescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-applescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_asciidoc_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_asciidoc_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-asciidoc.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-asciidoc.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_assembly_x86_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_assembly_x86_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-assembly_x86.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-assembly_x86.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_autohotkey_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_autohotkey_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-autohotkey.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-autohotkey.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_batchfile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_batchfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-batchfile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-batchfile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c9search_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c9search_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c9search.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c9search.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c_cpp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c_cpp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c_cpp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c_cpp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cirru_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cirru_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cirru.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cirru.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_clojure_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_clojure_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-clojure.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-clojure.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cobol_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cobol_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cobol.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cobol.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coffee_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coffee_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coffee.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coffee.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coldfusion_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coldfusion_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coldfusion.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coldfusion.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_csharp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_csharp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-csharp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-csharp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_css_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_css_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-css.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-css.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_curly_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_curly_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-curly.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-curly.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_d_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_d_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-d.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-d.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dart_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dart_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dart.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dart.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_diff_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_diff_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-diff.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-diff.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_django_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_django_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-django.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-django.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dockerfile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dockerfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dockerfile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dockerfile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dot_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dot_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dot.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dot.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_eiffel_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_eiffel_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-eiffel.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-eiffel.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ejs_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ejs_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ejs.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ejs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elixir_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elixir_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elixir.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elixir.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elm_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elm_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elm.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elm.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_erlang_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_erlang_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-erlang.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-erlang.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_forth_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_forth_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-forth.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-forth.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ftl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ftl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ftl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ftl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gcode_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gcode_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gcode.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gcode.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gherkin_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gherkin_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gherkin.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gherkin.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gitignore_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gitignore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gitignore.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gitignore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_glsl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_glsl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-glsl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-glsl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_golang_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_golang_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-golang.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-golang.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_groovy_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_groovy_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-groovy.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-groovy.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_handlebars_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_handlebars_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-handlebars.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-handlebars.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haskell_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haskell_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haskell.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haskell.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haxe_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haxe_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haxe.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haxe.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_ruby_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_ruby_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html_ruby.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html_ruby.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ini_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ini_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ini.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ini.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_io_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_io_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-io.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-io.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jack_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jack_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jack.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jack.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jade_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jade_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jade.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jade.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_java_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_java_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-java.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-java.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_javascript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_javascript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-javascript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-javascript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_json_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_json_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-json.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-json.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsoniq_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsoniq_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsoniq.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsoniq.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsx_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsx_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsx.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsx.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_julia_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_julia_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-julia.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-julia.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_latex_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_latex_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-latex.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-latex.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_less_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_less_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-less.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-less.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_liquid_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_liquid_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-liquid.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-liquid.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lisp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lisp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lisp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lisp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_livescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_livescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-livescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-livescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_logiql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_logiql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-logiql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-logiql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lsl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lsl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lsl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lsl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lua_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lua_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lua.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lua.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_luapage_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_luapage_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-luapage.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-luapage.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lucene_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lucene_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lucene.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lucene.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_makefile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_makefile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-makefile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-makefile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_markdown_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_markdown_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-markdown.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-markdown.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_matlab_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_matlab_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-matlab.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-matlab.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mel_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mel_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mel.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mel.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mushcode_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mushcode_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mushcode.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mushcode.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mysql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mysql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mysql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mysql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_nix_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_nix_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-nix.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-nix.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_objectivec_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_objectivec_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-objectivec.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-objectivec.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ocaml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ocaml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ocaml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ocaml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pascal_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pascal_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pascal.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pascal.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_perl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_perl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-perl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-perl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pgsql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pgsql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pgsql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pgsql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_php_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_php_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-php.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-php.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_plain_text_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_plain_text_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-plain_text.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-plain_text.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_powershell_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_powershell_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-powershell.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-powershell.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_praat_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_praat_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-praat.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-praat.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_prolog_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_prolog_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-prolog.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-prolog.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_properties_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_properties_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-properties.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-properties.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_protobuf_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_protobuf_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-protobuf.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-protobuf.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_python_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_python_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-python.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-python.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_r_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_r_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-r.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-r.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rdoc_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rdoc_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rdoc.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rdoc.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rhtml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rhtml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rhtml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rhtml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ruby_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ruby_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ruby.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ruby.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rust_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rust_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rust.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rust.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sass_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sass_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sass.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sass.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scad_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scad_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scad.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scad.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scala_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scala_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scala.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scala.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scheme_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scheme_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scheme.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scheme.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scss_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scss_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scss.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scss.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sh_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sh_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sh.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sh.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sjs_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sjs_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sjs.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sjs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_smarty_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_smarty_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-smarty.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-smarty.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_snippets_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_snippets_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-snippets.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-snippets.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_soy_template_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_soy_template_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-soy_template.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-soy_template.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_space_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_space_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-space.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-space.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_stylus_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_stylus_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-stylus.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-stylus.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_svg_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_svg_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-svg.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-svg.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tcl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tcl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tcl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tcl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tex_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tex_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tex.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tex.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_text_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_text_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-text.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-text.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_textile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_textile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-textile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-textile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_toml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_toml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-toml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-toml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_twig_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_twig_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-twig.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-twig.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_typescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_typescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-typescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-typescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vala_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vala_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vala.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vala.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vbscript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vbscript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vbscript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vbscript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_velocity_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_velocity_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-velocity.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-velocity.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_verilog_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_verilog_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-verilog.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-verilog.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vhdl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vhdl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vhdl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vhdl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xquery_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xquery_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xquery.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xquery.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_yaml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_yaml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-yaml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-yaml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_abap_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_abap_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/abap.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/abap.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_actionscript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_actionscript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/actionscript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/actionscript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ada_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ada_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ada.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ada.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_apache_conf_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_apache_conf_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/apache_conf.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/apache_conf.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_applescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_applescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/applescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/applescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_asciidoc_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_asciidoc_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/asciidoc.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/asciidoc.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_assembly_x86_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_assembly_x86_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/assembly_x86.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/assembly_x86.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_autohotkey_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_autohotkey_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/autohotkey.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/autohotkey.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_batchfile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_batchfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/batchfile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/batchfile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c9search_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c9search_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c9search.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c9search.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c_cpp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c_cpp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c_cpp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c_cpp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cirru_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cirru_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cirru.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cirru.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_clojure_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_clojure_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/clojure.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/clojure.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cobol_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cobol_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cobol.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cobol.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coffee_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coffee_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coffee.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coffee.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coldfusion_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coldfusion_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coldfusion.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coldfusion.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_csharp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_csharp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/csharp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/csharp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_css_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_css_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/css.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/css.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_curly_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_curly_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/curly.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/curly.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_d_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_d_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/d.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/d.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dart_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dart_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dart.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dart.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_diff_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_diff_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/diff.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/diff.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_django_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_django_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/django.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/django.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dockerfile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dockerfile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dockerfile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dockerfile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dot_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dot_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dot.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dot.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_eiffel_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_eiffel_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/eiffel.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/eiffel.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ejs_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ejs_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ejs.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ejs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elixir_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elixir_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elixir.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elixir.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elm_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elm_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elm.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elm.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_erlang_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_erlang_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/erlang.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/erlang.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_forth_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_forth_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/forth.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/forth.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ftl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ftl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ftl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ftl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gcode_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gcode_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gcode.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gcode.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gherkin_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gherkin_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gherkin.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gherkin.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gitignore_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gitignore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gitignore.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gitignore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_glsl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_glsl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/glsl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/glsl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_golang_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_golang_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/golang.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/golang.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_groovy_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_groovy_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/groovy.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/groovy.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_handlebars_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_handlebars_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/handlebars.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/handlebars.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haskell_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haskell_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haskell.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haskell.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haxe_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haxe_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haxe.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haxe.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_ruby_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_ruby_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html_ruby.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html_ruby.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ini_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ini_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ini.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ini.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_io_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_io_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/io.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/io.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jack_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jack_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jack.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jack.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jade_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jade_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jade.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jade.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_java_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_java_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/java.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/java.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_javascript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_javascript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/javascript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/javascript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_json_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_json_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/json.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/json.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsoniq_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsoniq_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsoniq.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsoniq.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsx_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsx_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsx.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsx.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_julia_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_julia_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/julia.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/julia.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_latex_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_latex_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/latex.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/latex.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_less_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_less_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/less.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/less.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_liquid_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_liquid_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/liquid.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/liquid.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lisp_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lisp_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lisp.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lisp.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_livescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_livescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/livescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/livescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_logiql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_logiql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/logiql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/logiql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lsl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lsl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lsl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lsl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lua_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lua_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lua.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lua.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_luapage_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_luapage_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/luapage.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/luapage.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lucene_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lucene_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lucene.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lucene.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_makefile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_makefile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/makefile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/makefile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_markdown_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_markdown_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/markdown.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/markdown.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_matlab_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_matlab_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/matlab.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/matlab.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mel_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mel_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mel.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mel.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mushcode_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mushcode_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mushcode.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mushcode.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mysql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mysql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mysql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mysql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_nix_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_nix_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/nix.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/nix.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_objectivec_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_objectivec_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/objectivec.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/objectivec.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ocaml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ocaml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ocaml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ocaml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pascal_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pascal_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pascal.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pascal.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_perl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_perl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/perl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/perl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pgsql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pgsql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pgsql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pgsql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_php_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_php_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/php.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/php.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_plain_text_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_plain_text_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/plain_text.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/plain_text.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_powershell_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_powershell_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/powershell.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/powershell.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_praat_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_praat_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/praat.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/praat.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_prolog_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_prolog_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/prolog.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/prolog.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_properties_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_properties_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/properties.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/properties.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_protobuf_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_protobuf_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/protobuf.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/protobuf.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_python_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_python_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/python.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/python.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_r_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_r_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/r.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/r.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rdoc_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rdoc_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rdoc.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rdoc.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rhtml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rhtml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rhtml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rhtml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ruby_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ruby_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ruby.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ruby.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rust_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rust_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rust.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rust.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sass_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sass_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sass.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sass.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scad_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scad_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scad.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scad.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scala_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scala_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scala.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scala.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scheme_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scheme_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scheme.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scheme.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scss_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scss_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scss.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scss.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sh_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sh_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sh.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sh.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sjs_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sjs_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sjs.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sjs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_smarty_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_smarty_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/smarty.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/smarty.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_snippets_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_snippets_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/snippets.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/snippets.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_soy_template_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_soy_template_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/soy_template.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/soy_template.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_space_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_space_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/space.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/space.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sql_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sql_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sql.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sql.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_stylus_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_stylus_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/stylus.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/stylus.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_svg_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_svg_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/svg.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/svg.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tcl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tcl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tcl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tcl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tex_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tex_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tex.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tex.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_text_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_text_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/text.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/text.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_textile_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_textile_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/textile.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/textile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_toml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_toml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/toml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/toml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_twig_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_twig_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/twig.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/twig.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_typescript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_typescript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/typescript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/typescript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vala_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vala_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vala.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vala.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vbscript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vbscript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vbscript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vbscript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_velocity_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_velocity_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/velocity.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/velocity.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_verilog_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_verilog_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/verilog.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/verilog.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vhdl_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vhdl_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vhdl.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vhdl.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xquery_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xquery_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xquery.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xquery.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_yaml_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_yaml_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/yaml.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/yaml.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_ambiance_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_ambiance_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-ambiance.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-ambiance.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chaos_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chaos_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chaos.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chaos.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chrome_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chrome_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chrome.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chrome.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_midnight_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_midnight_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds_midnight.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds_midnight.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_cobalt_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_cobalt_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-cobalt.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-cobalt.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_crimson_editor_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_crimson_editor_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-crimson_editor.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-crimson_editor.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dawn_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dawn_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dawn.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dawn.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dreamweaver_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dreamweaver_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dreamweaver.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dreamweaver.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_eclipse_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_eclipse_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-eclipse.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-eclipse.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_github_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_github_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-github.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-github.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_idle_fingers_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_idle_fingers_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-idle_fingers.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-idle_fingers.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_katzenmilch_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_katzenmilch_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-katzenmilch.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-katzenmilch.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kr_theme_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kr_theme_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kr_theme.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kr_theme.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kuroir_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kuroir_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kuroir.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kuroir.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_soft_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_soft_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore_soft.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore_soft.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_mono_industrial_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_mono_industrial_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-mono_industrial.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-mono_industrial.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_monokai_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_monokai_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-monokai.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-monokai.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_pastel_on_dark_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_pastel_on_dark_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-pastel_on_dark.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-pastel_on_dark.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_dark_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_dark_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_dark.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_dark.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_light_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_light_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_light.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_light.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_terminal_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_terminal_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-terminal.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-terminal.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_textmate_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_textmate_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-textmate.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-textmate.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_blue_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_blue_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_blue.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_blue.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_bright_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_bright_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_bright.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_bright.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_eighties_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_eighties_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_eighties.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_eighties.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_twilight_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_twilight_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-twilight.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-twilight.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_vibrant_ink_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_vibrant_ink_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-vibrant_ink.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-vibrant_ink.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_xcode_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_xcode_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-xcode.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-xcode.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_coffee_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_coffee_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-coffee.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-coffee.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_css_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_css_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-css.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-css.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_html_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_html_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-html.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-html.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_javascript_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_javascript_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-javascript.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-javascript.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_json_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_json_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-json.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-json.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_lua_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_lua_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-lua.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-lua.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_php_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_php_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-php.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-php.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_xquery_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_xquery_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-xquery.js"
	name := "server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-xquery.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// server_ui_vendor_jquery_2_1_3_jquery_min_js reads file data from disk. It returns an error on failure.
func server_ui_vendor_jquery_2_1_3_jquery_min_js() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/vendor/jquery-2.1.3/jquery.min.js"
	name := "server/ui/vendor/jquery-2.1.3/jquery.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
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
	"server/ui/dashboard.html": server_ui_dashboard_html,
	"server/ui/dist/app.js": server_ui_dist_app_js,
	"server/ui/dist/main.css": server_ui_dist_main_css,
	"server/ui/editor.html": server_ui_editor_html,
	"server/ui/foot.html": server_ui_foot_html,
	"server/ui/gulpfile.js": server_ui_gulpfile_js,
	"server/ui/head.html": server_ui_head_html,
	"server/ui/package.json": server_ui_package_json,
	"server/ui/src/.DS_Store": server_ui_src_ds_store,
	"server/ui/src/img/icon_build_image.png": server_ui_src_img_icon_build_image_png,
	"server/ui/src/img/icon_build_image@1x.png": server_ui_src_img_icon_build_image_1x_png,
	"server/ui/src/img/icon_built_image.png": server_ui_src_img_icon_built_image_png,
	"server/ui/src/img/icon_dep.png": server_ui_src_img_icon_dep_png,
	"server/ui/src/img/icon_dep@1x.png": server_ui_src_img_icon_dep_1x_png,
	"server/ui/src/img/icon_fixture.png": server_ui_src_img_icon_fixture_png,
	"server/ui/src/img/icon_fixture@1x.png": server_ui_src_img_icon_fixture_1x_png,
	"server/ui/src/img/icon_run_image.png": server_ui_src_img_icon_run_image_png,
	"server/ui/src/img/icon_run_image@1x.png": server_ui_src_img_icon_run_image_1x_png,
	"server/ui/src/img/icon_t_api-blueprint.png": server_ui_src_img_icon_t_api_blueprint_png,
	"server/ui/src/img/icon_t_api-blueprint@0,7x.png": server_ui_src_img_icon_t_api_blueprint_0_7x_png,
	"server/ui/src/img/icon_t_custom.png": server_ui_src_img_icon_t_custom_png,
	"server/ui/src/img/icon_t_custom@0,7x.png": server_ui_src_img_icon_t_custom_0_7x_png,
	"server/ui/src/img/icon_t_etcd.png": server_ui_src_img_icon_t_etcd_png,
	"server/ui/src/img/icon_t_etcd@0,7x.png": server_ui_src_img_icon_t_etcd_0_7x_png,
	"server/ui/src/img/icon_t_postgres.png": server_ui_src_img_icon_t_postgres_png,
	"server/ui/src/img/icon_t_postgres@0,7x.png": server_ui_src_img_icon_t_postgres_0_7x_png,
	"server/ui/src/img/icon_unbuilt_image.png": server_ui_src_img_icon_unbuilt_image_png,
	"server/ui/src/img/logo_brand.png": server_ui_src_img_logo_brand_png,
	"server/ui/src/img/logo_brand@2x.png": server_ui_src_img_logo_brand_2x_png,
	"server/ui/src/js/actions/DepActions.js": server_ui_src_js_actions_depactions_js,
	"server/ui/src/js/actions/EditorActions.js": server_ui_src_js_actions_editoractions_js,
	"server/ui/src/js/actions/IsolationActions.js": server_ui_src_js_actions_isolationactions_js,
	"server/ui/src/js/actions/StatActions.js": server_ui_src_js_actions_statactions_js,
	"server/ui/src/js/app.js": server_ui_src_js_app_js,
	"server/ui/src/js/components/AchievementPopup.jsx": server_ui_src_js_components_achievementpopup_jsx,
	"server/ui/src/js/components/DepForm.jsx": server_ui_src_js_components_depform_jsx,
	"server/ui/src/js/components/DepItem.jsx": server_ui_src_js_components_depitem_jsx,
	"server/ui/src/js/components/DepPanel.jsx": server_ui_src_js_components_deppanel_jsx,
	"server/ui/src/js/components/EditorFilePanel.jsx": server_ui_src_js_components_editorfilepanel_jsx,
	"server/ui/src/js/components/EditorRunPanel.jsx": server_ui_src_js_components_editorrunpanel_jsx,
	"server/ui/src/js/components/EditorWorkspace.jsx": server_ui_src_js_components_editorworkspace_jsx,
	"server/ui/src/js/components/IsolationForm.jsx": server_ui_src_js_components_isolationform_jsx,
	"server/ui/src/js/components/IsolationItem.jsx": server_ui_src_js_components_isolationitem_jsx,
	"server/ui/src/js/components/IsolationList.jsx": server_ui_src_js_components_isolationlist_jsx,
	"server/ui/src/js/components/IsolationPanel.jsx": server_ui_src_js_components_isolationpanel_jsx,
	"server/ui/src/js/dispatcher/AppDispatcher.js": server_ui_src_js_dispatcher_appdispatcher_js,
	"server/ui/src/js/stores/DepStore.js": server_ui_src_js_stores_depstore_js,
	"server/ui/src/js/stores/EditorStore.js": server_ui_src_js_stores_editorstore_js,
	"server/ui/src/js/stores/IsolationStore.js": server_ui_src_js_stores_isolationstore_js,
	"server/ui/src/js/stores/StatStore.js": server_ui_src_js_stores_statstore_js,
	"server/ui/src/less/main.less": server_ui_src_less_main_less,
	"server/ui/templates.json": server_ui_templates_json,
	"server/ui/vendor/.DS_Store": server_ui_vendor_ds_store,
	"server/ui/vendor/ace-1.1.8/ChangeLog.txt": server_ui_vendor_ace_1_1_8_changelog_txt,
	"server/ui/vendor/ace-1.1.8/LICENSE": server_ui_vendor_ace_1_1_8_license,
	"server/ui/vendor/ace-1.1.8/README.md": server_ui_vendor_ace_1_1_8_readme_md,
	"server/ui/vendor/ace-1.1.8/package.json": server_ui_vendor_ace_1_1_8_package_json,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ace.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ace_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-beautify.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_beautify_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-chromevox.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_chromevox_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-elastic_tabstops_lite.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_elastic_tabstops_lite_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-emmet.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_emmet_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-error_marker.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_error_marker_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-keybinding_menu.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_keybinding_menu_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-language_tools.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_language_tools_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-linking.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_linking_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-modelist.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_modelist_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-old_ie.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_old_ie_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-searchbox.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_searchbox_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-settings_menu.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_settings_menu_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-spellcheck.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_spellcheck_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-split.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_split_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-static_highlight.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_static_highlight_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-statusbar.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_statusbar_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-textarea.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_textarea_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-themelist.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_themelist_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/ext-whitespace.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_whitespace_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-emacs.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_emacs_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/keybinding-vim.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_vim_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-abap.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_abap_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-actionscript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_actionscript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ada.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ada_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-apache_conf.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_apache_conf_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-applescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_applescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-asciidoc.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_asciidoc_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-assembly_x86.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_assembly_x86_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-autohotkey.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_autohotkey_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-batchfile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_batchfile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c9search.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c9search_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-c_cpp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c_cpp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cirru.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cirru_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-clojure.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_clojure_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-cobol.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cobol_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coffee.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coffee_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-coldfusion.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coldfusion_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-csharp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_csharp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-css.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_css_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-curly.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_curly_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-d.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_d_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dart.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dart_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-diff.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_diff_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-django.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_django_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dockerfile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dockerfile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-dot.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dot_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-eiffel.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_eiffel_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ejs.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ejs_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elixir.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elixir_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-elm.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elm_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-erlang.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_erlang_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-forth.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_forth_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ftl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ftl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gcode.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gcode_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gherkin.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gherkin_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-gitignore.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gitignore_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-glsl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_glsl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-golang.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_golang_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-groovy.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_groovy_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-handlebars.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_handlebars_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haskell.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haskell_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-haxe.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haxe_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-html_ruby.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_ruby_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ini.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ini_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-io.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_io_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jack.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jack_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jade.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jade_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-java.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_java_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-javascript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_javascript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-json.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_json_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsoniq.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsoniq_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-jsx.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsx_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-julia.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_julia_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-latex.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_latex_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-less.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_less_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-liquid.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_liquid_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lisp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lisp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-livescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_livescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-logiql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_logiql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lsl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lsl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lua.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lua_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-luapage.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_luapage_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-lucene.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lucene_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-makefile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_makefile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-markdown.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_markdown_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-matlab.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_matlab_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mel.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mel_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mushcode.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mushcode_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-mysql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mysql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-nix.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_nix_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-objectivec.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_objectivec_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ocaml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ocaml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pascal.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pascal_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-perl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_perl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-pgsql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pgsql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-php.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_php_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-plain_text.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_plain_text_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-powershell.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_powershell_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-praat.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_praat_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-prolog.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_prolog_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-properties.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_properties_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-protobuf.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_protobuf_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-python.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_python_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-r.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_r_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rdoc.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rdoc_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rhtml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rhtml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-ruby.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ruby_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-rust.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rust_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sass.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sass_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scad.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scad_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scala.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scala_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scheme.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scheme_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-scss.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scss_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sh.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sh_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sjs.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sjs_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-smarty.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_smarty_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-snippets.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_snippets_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-soy_template.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_soy_template_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-space.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_space_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-sql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-stylus.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_stylus_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-svg.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_svg_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tcl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tcl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-tex.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tex_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-text.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_text_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-textile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_textile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-toml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_toml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-twig.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_twig_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-typescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_typescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vala.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vala_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vbscript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vbscript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-velocity.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_velocity_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-verilog.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_verilog_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-vhdl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vhdl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-xquery.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xquery_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/mode-yaml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_yaml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/abap.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_abap_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/actionscript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_actionscript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ada.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ada_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/apache_conf.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_apache_conf_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/applescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_applescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/asciidoc.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_asciidoc_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/assembly_x86.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_assembly_x86_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/autohotkey.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_autohotkey_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/batchfile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_batchfile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c9search.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c9search_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/c_cpp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c_cpp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cirru.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cirru_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/clojure.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_clojure_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/cobol.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cobol_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coffee.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coffee_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/coldfusion.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coldfusion_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/csharp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_csharp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/css.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_css_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/curly.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_curly_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/d.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_d_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dart.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dart_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/diff.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_diff_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/django.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_django_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dockerfile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dockerfile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/dot.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dot_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/eiffel.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_eiffel_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ejs.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ejs_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elixir.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elixir_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/elm.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elm_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/erlang.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_erlang_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/forth.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_forth_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ftl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ftl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gcode.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gcode_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gherkin.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gherkin_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/gitignore.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gitignore_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/glsl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_glsl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/golang.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_golang_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/groovy.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_groovy_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/handlebars.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_handlebars_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haskell.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haskell_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/haxe.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haxe_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/html_ruby.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_ruby_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ini.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ini_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/io.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_io_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jack.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jack_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jade.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jade_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/java.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_java_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/javascript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_javascript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/json.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_json_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsoniq.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsoniq_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/jsx.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsx_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/julia.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_julia_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/latex.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_latex_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/less.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_less_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/liquid.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_liquid_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lisp.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lisp_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/livescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_livescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/logiql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_logiql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lsl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lsl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lua.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lua_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/luapage.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_luapage_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/lucene.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lucene_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/makefile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_makefile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/markdown.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_markdown_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/matlab.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_matlab_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mel.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mel_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mushcode.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mushcode_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/mysql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mysql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/nix.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_nix_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/objectivec.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_objectivec_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ocaml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ocaml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pascal.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pascal_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/perl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_perl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/pgsql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pgsql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/php.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_php_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/plain_text.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_plain_text_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/powershell.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_powershell_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/praat.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_praat_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/prolog.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_prolog_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/properties.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_properties_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/protobuf.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_protobuf_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/python.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_python_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/r.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_r_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rdoc.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rdoc_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rhtml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rhtml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/ruby.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ruby_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/rust.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rust_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sass.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sass_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scad.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scad_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scala.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scala_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scheme.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scheme_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/scss.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scss_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sh.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sh_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sjs.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sjs_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/smarty.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_smarty_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/snippets.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_snippets_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/soy_template.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_soy_template_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/space.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_space_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/sql.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sql_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/stylus.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_stylus_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/svg.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_svg_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tcl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tcl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/tex.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tex_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/text.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_text_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/textile.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_textile_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/toml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_toml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/twig.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_twig_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/typescript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_typescript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vala.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vala_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vbscript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vbscript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/velocity.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_velocity_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/verilog.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_verilog_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/vhdl.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vhdl_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/xquery.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xquery_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/snippets/yaml.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_yaml_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-ambiance.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_ambiance_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chaos.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chaos_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-chrome.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chrome_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-clouds_midnight.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_midnight_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-cobalt.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_cobalt_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-crimson_editor.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_crimson_editor_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dawn.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dawn_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-dreamweaver.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dreamweaver_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-eclipse.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_eclipse_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-github.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_github_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-idle_fingers.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_idle_fingers_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-katzenmilch.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_katzenmilch_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kr_theme.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kr_theme_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-kuroir.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kuroir_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-merbivore_soft.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_soft_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-mono_industrial.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_mono_industrial_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-monokai.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_monokai_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-pastel_on_dark.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_pastel_on_dark_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_dark.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_dark_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-solarized_light.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_light_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-terminal.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_terminal_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-textmate.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_textmate_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_blue.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_blue_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_bright.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_bright_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-tomorrow_night_eighties.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_eighties_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-twilight.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_twilight_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-vibrant_ink.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_vibrant_ink_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/theme-xcode.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_xcode_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-coffee.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_coffee_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-css.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_css_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-html.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_html_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-javascript.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_javascript_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-json.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_json_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-lua.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_lua_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-php.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_php_js,
	"server/ui/vendor/ace-1.1.8/src-min-noconflict/worker-xquery.js": server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_xquery_js,
	"server/ui/vendor/jquery-2.1.3/jquery.min.js": server_ui_vendor_jquery_2_1_3_jquery_min_js,
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
			"dashboard.html": &_bintree_t{server_ui_dashboard_html, map[string]*_bintree_t{
			}},
			"dist": &_bintree_t{nil, map[string]*_bintree_t{
				"app.js": &_bintree_t{server_ui_dist_app_js, map[string]*_bintree_t{
				}},
				"main.css": &_bintree_t{server_ui_dist_main_css, map[string]*_bintree_t{
				}},
			}},
			"editor.html": &_bintree_t{server_ui_editor_html, map[string]*_bintree_t{
			}},
			"foot.html": &_bintree_t{server_ui_foot_html, map[string]*_bintree_t{
			}},
			"gulpfile.js": &_bintree_t{server_ui_gulpfile_js, map[string]*_bintree_t{
			}},
			"head.html": &_bintree_t{server_ui_head_html, map[string]*_bintree_t{
			}},
			"package.json": &_bintree_t{server_ui_package_json, map[string]*_bintree_t{
			}},
			"src": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{server_ui_src_ds_store, map[string]*_bintree_t{
				}},
				"img": &_bintree_t{nil, map[string]*_bintree_t{
					"icon_build_image.png": &_bintree_t{server_ui_src_img_icon_build_image_png, map[string]*_bintree_t{
					}},
					"icon_build_image@1x.png": &_bintree_t{server_ui_src_img_icon_build_image_1x_png, map[string]*_bintree_t{
					}},
					"icon_built_image.png": &_bintree_t{server_ui_src_img_icon_built_image_png, map[string]*_bintree_t{
					}},
					"icon_dep.png": &_bintree_t{server_ui_src_img_icon_dep_png, map[string]*_bintree_t{
					}},
					"icon_dep@1x.png": &_bintree_t{server_ui_src_img_icon_dep_1x_png, map[string]*_bintree_t{
					}},
					"icon_fixture.png": &_bintree_t{server_ui_src_img_icon_fixture_png, map[string]*_bintree_t{
					}},
					"icon_fixture@1x.png": &_bintree_t{server_ui_src_img_icon_fixture_1x_png, map[string]*_bintree_t{
					}},
					"icon_run_image.png": &_bintree_t{server_ui_src_img_icon_run_image_png, map[string]*_bintree_t{
					}},
					"icon_run_image@1x.png": &_bintree_t{server_ui_src_img_icon_run_image_1x_png, map[string]*_bintree_t{
					}},
					"icon_t_api-blueprint.png": &_bintree_t{server_ui_src_img_icon_t_api_blueprint_png, map[string]*_bintree_t{
					}},
					"icon_t_api-blueprint@0,7x.png": &_bintree_t{server_ui_src_img_icon_t_api_blueprint_0_7x_png, map[string]*_bintree_t{
					}},
					"icon_t_custom.png": &_bintree_t{server_ui_src_img_icon_t_custom_png, map[string]*_bintree_t{
					}},
					"icon_t_custom@0,7x.png": &_bintree_t{server_ui_src_img_icon_t_custom_0_7x_png, map[string]*_bintree_t{
					}},
					"icon_t_etcd.png": &_bintree_t{server_ui_src_img_icon_t_etcd_png, map[string]*_bintree_t{
					}},
					"icon_t_etcd@0,7x.png": &_bintree_t{server_ui_src_img_icon_t_etcd_0_7x_png, map[string]*_bintree_t{
					}},
					"icon_t_postgres.png": &_bintree_t{server_ui_src_img_icon_t_postgres_png, map[string]*_bintree_t{
					}},
					"icon_t_postgres@0,7x.png": &_bintree_t{server_ui_src_img_icon_t_postgres_0_7x_png, map[string]*_bintree_t{
					}},
					"icon_unbuilt_image.png": &_bintree_t{server_ui_src_img_icon_unbuilt_image_png, map[string]*_bintree_t{
					}},
					"logo_brand.png": &_bintree_t{server_ui_src_img_logo_brand_png, map[string]*_bintree_t{
					}},
					"logo_brand@2x.png": &_bintree_t{server_ui_src_img_logo_brand_2x_png, map[string]*_bintree_t{
					}},
				}},
				"js": &_bintree_t{nil, map[string]*_bintree_t{
					"actions": &_bintree_t{nil, map[string]*_bintree_t{
						"DepActions.js": &_bintree_t{server_ui_src_js_actions_depactions_js, map[string]*_bintree_t{
						}},
						"EditorActions.js": &_bintree_t{server_ui_src_js_actions_editoractions_js, map[string]*_bintree_t{
						}},
						"IsolationActions.js": &_bintree_t{server_ui_src_js_actions_isolationactions_js, map[string]*_bintree_t{
						}},
						"StatActions.js": &_bintree_t{server_ui_src_js_actions_statactions_js, map[string]*_bintree_t{
						}},
					}},
					"app.js": &_bintree_t{server_ui_src_js_app_js, map[string]*_bintree_t{
					}},
					"components": &_bintree_t{nil, map[string]*_bintree_t{
						"AchievementPopup.jsx": &_bintree_t{server_ui_src_js_components_achievementpopup_jsx, map[string]*_bintree_t{
						}},
						"DepForm.jsx": &_bintree_t{server_ui_src_js_components_depform_jsx, map[string]*_bintree_t{
						}},
						"DepItem.jsx": &_bintree_t{server_ui_src_js_components_depitem_jsx, map[string]*_bintree_t{
						}},
						"DepPanel.jsx": &_bintree_t{server_ui_src_js_components_deppanel_jsx, map[string]*_bintree_t{
						}},
						"EditorFilePanel.jsx": &_bintree_t{server_ui_src_js_components_editorfilepanel_jsx, map[string]*_bintree_t{
						}},
						"EditorRunPanel.jsx": &_bintree_t{server_ui_src_js_components_editorrunpanel_jsx, map[string]*_bintree_t{
						}},
						"EditorWorkspace.jsx": &_bintree_t{server_ui_src_js_components_editorworkspace_jsx, map[string]*_bintree_t{
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
						"EditorStore.js": &_bintree_t{server_ui_src_js_stores_editorstore_js, map[string]*_bintree_t{
						}},
						"IsolationStore.js": &_bintree_t{server_ui_src_js_stores_isolationstore_js, map[string]*_bintree_t{
						}},
						"StatStore.js": &_bintree_t{server_ui_src_js_stores_statstore_js, map[string]*_bintree_t{
						}},
					}},
				}},
				"less": &_bintree_t{nil, map[string]*_bintree_t{
					"main.less": &_bintree_t{server_ui_src_less_main_less, map[string]*_bintree_t{
					}},
				}},
			}},
			"templates.json": &_bintree_t{server_ui_templates_json, map[string]*_bintree_t{
			}},
			"vendor": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{server_ui_vendor_ds_store, map[string]*_bintree_t{
				}},
				"ace-1.1.8": &_bintree_t{nil, map[string]*_bintree_t{
					"ChangeLog.txt": &_bintree_t{server_ui_vendor_ace_1_1_8_changelog_txt, map[string]*_bintree_t{
					}},
					"LICENSE": &_bintree_t{server_ui_vendor_ace_1_1_8_license, map[string]*_bintree_t{
					}},
					"README.md": &_bintree_t{server_ui_vendor_ace_1_1_8_readme_md, map[string]*_bintree_t{
					}},
					"package.json": &_bintree_t{server_ui_vendor_ace_1_1_8_package_json, map[string]*_bintree_t{
					}},
					"src-min-noconflict": &_bintree_t{nil, map[string]*_bintree_t{
						"ace.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ace_js, map[string]*_bintree_t{
						}},
						"ext-beautify.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_beautify_js, map[string]*_bintree_t{
						}},
						"ext-chromevox.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_chromevox_js, map[string]*_bintree_t{
						}},
						"ext-elastic_tabstops_lite.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_elastic_tabstops_lite_js, map[string]*_bintree_t{
						}},
						"ext-emmet.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_emmet_js, map[string]*_bintree_t{
						}},
						"ext-error_marker.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_error_marker_js, map[string]*_bintree_t{
						}},
						"ext-keybinding_menu.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_keybinding_menu_js, map[string]*_bintree_t{
						}},
						"ext-language_tools.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_language_tools_js, map[string]*_bintree_t{
						}},
						"ext-linking.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_linking_js, map[string]*_bintree_t{
						}},
						"ext-modelist.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_modelist_js, map[string]*_bintree_t{
						}},
						"ext-old_ie.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_old_ie_js, map[string]*_bintree_t{
						}},
						"ext-searchbox.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_searchbox_js, map[string]*_bintree_t{
						}},
						"ext-settings_menu.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_settings_menu_js, map[string]*_bintree_t{
						}},
						"ext-spellcheck.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_spellcheck_js, map[string]*_bintree_t{
						}},
						"ext-split.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_split_js, map[string]*_bintree_t{
						}},
						"ext-static_highlight.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_static_highlight_js, map[string]*_bintree_t{
						}},
						"ext-statusbar.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_statusbar_js, map[string]*_bintree_t{
						}},
						"ext-textarea.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_textarea_js, map[string]*_bintree_t{
						}},
						"ext-themelist.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_themelist_js, map[string]*_bintree_t{
						}},
						"ext-whitespace.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_ext_whitespace_js, map[string]*_bintree_t{
						}},
						"keybinding-emacs.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_emacs_js, map[string]*_bintree_t{
						}},
						"keybinding-vim.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_keybinding_vim_js, map[string]*_bintree_t{
						}},
						"mode-abap.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_abap_js, map[string]*_bintree_t{
						}},
						"mode-actionscript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_actionscript_js, map[string]*_bintree_t{
						}},
						"mode-ada.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ada_js, map[string]*_bintree_t{
						}},
						"mode-apache_conf.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_apache_conf_js, map[string]*_bintree_t{
						}},
						"mode-applescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_applescript_js, map[string]*_bintree_t{
						}},
						"mode-asciidoc.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_asciidoc_js, map[string]*_bintree_t{
						}},
						"mode-assembly_x86.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_assembly_x86_js, map[string]*_bintree_t{
						}},
						"mode-autohotkey.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_autohotkey_js, map[string]*_bintree_t{
						}},
						"mode-batchfile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_batchfile_js, map[string]*_bintree_t{
						}},
						"mode-c9search.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c9search_js, map[string]*_bintree_t{
						}},
						"mode-c_cpp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_c_cpp_js, map[string]*_bintree_t{
						}},
						"mode-cirru.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cirru_js, map[string]*_bintree_t{
						}},
						"mode-clojure.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_clojure_js, map[string]*_bintree_t{
						}},
						"mode-cobol.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_cobol_js, map[string]*_bintree_t{
						}},
						"mode-coffee.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coffee_js, map[string]*_bintree_t{
						}},
						"mode-coldfusion.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_coldfusion_js, map[string]*_bintree_t{
						}},
						"mode-csharp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_csharp_js, map[string]*_bintree_t{
						}},
						"mode-css.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_css_js, map[string]*_bintree_t{
						}},
						"mode-curly.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_curly_js, map[string]*_bintree_t{
						}},
						"mode-d.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_d_js, map[string]*_bintree_t{
						}},
						"mode-dart.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dart_js, map[string]*_bintree_t{
						}},
						"mode-diff.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_diff_js, map[string]*_bintree_t{
						}},
						"mode-django.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_django_js, map[string]*_bintree_t{
						}},
						"mode-dockerfile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dockerfile_js, map[string]*_bintree_t{
						}},
						"mode-dot.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_dot_js, map[string]*_bintree_t{
						}},
						"mode-eiffel.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_eiffel_js, map[string]*_bintree_t{
						}},
						"mode-ejs.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ejs_js, map[string]*_bintree_t{
						}},
						"mode-elixir.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elixir_js, map[string]*_bintree_t{
						}},
						"mode-elm.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_elm_js, map[string]*_bintree_t{
						}},
						"mode-erlang.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_erlang_js, map[string]*_bintree_t{
						}},
						"mode-forth.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_forth_js, map[string]*_bintree_t{
						}},
						"mode-ftl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ftl_js, map[string]*_bintree_t{
						}},
						"mode-gcode.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gcode_js, map[string]*_bintree_t{
						}},
						"mode-gherkin.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gherkin_js, map[string]*_bintree_t{
						}},
						"mode-gitignore.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_gitignore_js, map[string]*_bintree_t{
						}},
						"mode-glsl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_glsl_js, map[string]*_bintree_t{
						}},
						"mode-golang.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_golang_js, map[string]*_bintree_t{
						}},
						"mode-groovy.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_groovy_js, map[string]*_bintree_t{
						}},
						"mode-haml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haml_js, map[string]*_bintree_t{
						}},
						"mode-handlebars.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_handlebars_js, map[string]*_bintree_t{
						}},
						"mode-haskell.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haskell_js, map[string]*_bintree_t{
						}},
						"mode-haxe.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_haxe_js, map[string]*_bintree_t{
						}},
						"mode-html.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_js, map[string]*_bintree_t{
						}},
						"mode-html_ruby.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_html_ruby_js, map[string]*_bintree_t{
						}},
						"mode-ini.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ini_js, map[string]*_bintree_t{
						}},
						"mode-io.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_io_js, map[string]*_bintree_t{
						}},
						"mode-jack.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jack_js, map[string]*_bintree_t{
						}},
						"mode-jade.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jade_js, map[string]*_bintree_t{
						}},
						"mode-java.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_java_js, map[string]*_bintree_t{
						}},
						"mode-javascript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_javascript_js, map[string]*_bintree_t{
						}},
						"mode-json.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_json_js, map[string]*_bintree_t{
						}},
						"mode-jsoniq.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsoniq_js, map[string]*_bintree_t{
						}},
						"mode-jsp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsp_js, map[string]*_bintree_t{
						}},
						"mode-jsx.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_jsx_js, map[string]*_bintree_t{
						}},
						"mode-julia.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_julia_js, map[string]*_bintree_t{
						}},
						"mode-latex.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_latex_js, map[string]*_bintree_t{
						}},
						"mode-less.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_less_js, map[string]*_bintree_t{
						}},
						"mode-liquid.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_liquid_js, map[string]*_bintree_t{
						}},
						"mode-lisp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lisp_js, map[string]*_bintree_t{
						}},
						"mode-livescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_livescript_js, map[string]*_bintree_t{
						}},
						"mode-logiql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_logiql_js, map[string]*_bintree_t{
						}},
						"mode-lsl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lsl_js, map[string]*_bintree_t{
						}},
						"mode-lua.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lua_js, map[string]*_bintree_t{
						}},
						"mode-luapage.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_luapage_js, map[string]*_bintree_t{
						}},
						"mode-lucene.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_lucene_js, map[string]*_bintree_t{
						}},
						"mode-makefile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_makefile_js, map[string]*_bintree_t{
						}},
						"mode-markdown.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_markdown_js, map[string]*_bintree_t{
						}},
						"mode-matlab.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_matlab_js, map[string]*_bintree_t{
						}},
						"mode-mel.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mel_js, map[string]*_bintree_t{
						}},
						"mode-mushcode.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mushcode_js, map[string]*_bintree_t{
						}},
						"mode-mysql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_mysql_js, map[string]*_bintree_t{
						}},
						"mode-nix.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_nix_js, map[string]*_bintree_t{
						}},
						"mode-objectivec.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_objectivec_js, map[string]*_bintree_t{
						}},
						"mode-ocaml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ocaml_js, map[string]*_bintree_t{
						}},
						"mode-pascal.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pascal_js, map[string]*_bintree_t{
						}},
						"mode-perl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_perl_js, map[string]*_bintree_t{
						}},
						"mode-pgsql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_pgsql_js, map[string]*_bintree_t{
						}},
						"mode-php.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_php_js, map[string]*_bintree_t{
						}},
						"mode-plain_text.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_plain_text_js, map[string]*_bintree_t{
						}},
						"mode-powershell.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_powershell_js, map[string]*_bintree_t{
						}},
						"mode-praat.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_praat_js, map[string]*_bintree_t{
						}},
						"mode-prolog.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_prolog_js, map[string]*_bintree_t{
						}},
						"mode-properties.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_properties_js, map[string]*_bintree_t{
						}},
						"mode-protobuf.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_protobuf_js, map[string]*_bintree_t{
						}},
						"mode-python.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_python_js, map[string]*_bintree_t{
						}},
						"mode-r.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_r_js, map[string]*_bintree_t{
						}},
						"mode-rdoc.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rdoc_js, map[string]*_bintree_t{
						}},
						"mode-rhtml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rhtml_js, map[string]*_bintree_t{
						}},
						"mode-ruby.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_ruby_js, map[string]*_bintree_t{
						}},
						"mode-rust.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_rust_js, map[string]*_bintree_t{
						}},
						"mode-sass.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sass_js, map[string]*_bintree_t{
						}},
						"mode-scad.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scad_js, map[string]*_bintree_t{
						}},
						"mode-scala.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scala_js, map[string]*_bintree_t{
						}},
						"mode-scheme.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scheme_js, map[string]*_bintree_t{
						}},
						"mode-scss.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_scss_js, map[string]*_bintree_t{
						}},
						"mode-sh.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sh_js, map[string]*_bintree_t{
						}},
						"mode-sjs.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sjs_js, map[string]*_bintree_t{
						}},
						"mode-smarty.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_smarty_js, map[string]*_bintree_t{
						}},
						"mode-snippets.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_snippets_js, map[string]*_bintree_t{
						}},
						"mode-soy_template.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_soy_template_js, map[string]*_bintree_t{
						}},
						"mode-space.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_space_js, map[string]*_bintree_t{
						}},
						"mode-sql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_sql_js, map[string]*_bintree_t{
						}},
						"mode-stylus.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_stylus_js, map[string]*_bintree_t{
						}},
						"mode-svg.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_svg_js, map[string]*_bintree_t{
						}},
						"mode-tcl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tcl_js, map[string]*_bintree_t{
						}},
						"mode-tex.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_tex_js, map[string]*_bintree_t{
						}},
						"mode-text.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_text_js, map[string]*_bintree_t{
						}},
						"mode-textile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_textile_js, map[string]*_bintree_t{
						}},
						"mode-toml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_toml_js, map[string]*_bintree_t{
						}},
						"mode-twig.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_twig_js, map[string]*_bintree_t{
						}},
						"mode-typescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_typescript_js, map[string]*_bintree_t{
						}},
						"mode-vala.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vala_js, map[string]*_bintree_t{
						}},
						"mode-vbscript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vbscript_js, map[string]*_bintree_t{
						}},
						"mode-velocity.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_velocity_js, map[string]*_bintree_t{
						}},
						"mode-verilog.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_verilog_js, map[string]*_bintree_t{
						}},
						"mode-vhdl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_vhdl_js, map[string]*_bintree_t{
						}},
						"mode-xml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xml_js, map[string]*_bintree_t{
						}},
						"mode-xquery.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_xquery_js, map[string]*_bintree_t{
						}},
						"mode-yaml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_mode_yaml_js, map[string]*_bintree_t{
						}},
						"snippets": &_bintree_t{nil, map[string]*_bintree_t{
							"abap.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_abap_js, map[string]*_bintree_t{
							}},
							"actionscript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_actionscript_js, map[string]*_bintree_t{
							}},
							"ada.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ada_js, map[string]*_bintree_t{
							}},
							"apache_conf.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_apache_conf_js, map[string]*_bintree_t{
							}},
							"applescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_applescript_js, map[string]*_bintree_t{
							}},
							"asciidoc.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_asciidoc_js, map[string]*_bintree_t{
							}},
							"assembly_x86.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_assembly_x86_js, map[string]*_bintree_t{
							}},
							"autohotkey.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_autohotkey_js, map[string]*_bintree_t{
							}},
							"batchfile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_batchfile_js, map[string]*_bintree_t{
							}},
							"c9search.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c9search_js, map[string]*_bintree_t{
							}},
							"c_cpp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_c_cpp_js, map[string]*_bintree_t{
							}},
							"cirru.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cirru_js, map[string]*_bintree_t{
							}},
							"clojure.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_clojure_js, map[string]*_bintree_t{
							}},
							"cobol.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_cobol_js, map[string]*_bintree_t{
							}},
							"coffee.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coffee_js, map[string]*_bintree_t{
							}},
							"coldfusion.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_coldfusion_js, map[string]*_bintree_t{
							}},
							"csharp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_csharp_js, map[string]*_bintree_t{
							}},
							"css.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_css_js, map[string]*_bintree_t{
							}},
							"curly.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_curly_js, map[string]*_bintree_t{
							}},
							"d.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_d_js, map[string]*_bintree_t{
							}},
							"dart.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dart_js, map[string]*_bintree_t{
							}},
							"diff.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_diff_js, map[string]*_bintree_t{
							}},
							"django.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_django_js, map[string]*_bintree_t{
							}},
							"dockerfile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dockerfile_js, map[string]*_bintree_t{
							}},
							"dot.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_dot_js, map[string]*_bintree_t{
							}},
							"eiffel.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_eiffel_js, map[string]*_bintree_t{
							}},
							"ejs.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ejs_js, map[string]*_bintree_t{
							}},
							"elixir.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elixir_js, map[string]*_bintree_t{
							}},
							"elm.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_elm_js, map[string]*_bintree_t{
							}},
							"erlang.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_erlang_js, map[string]*_bintree_t{
							}},
							"forth.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_forth_js, map[string]*_bintree_t{
							}},
							"ftl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ftl_js, map[string]*_bintree_t{
							}},
							"gcode.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gcode_js, map[string]*_bintree_t{
							}},
							"gherkin.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gherkin_js, map[string]*_bintree_t{
							}},
							"gitignore.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_gitignore_js, map[string]*_bintree_t{
							}},
							"glsl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_glsl_js, map[string]*_bintree_t{
							}},
							"golang.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_golang_js, map[string]*_bintree_t{
							}},
							"groovy.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_groovy_js, map[string]*_bintree_t{
							}},
							"haml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haml_js, map[string]*_bintree_t{
							}},
							"handlebars.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_handlebars_js, map[string]*_bintree_t{
							}},
							"haskell.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haskell_js, map[string]*_bintree_t{
							}},
							"haxe.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_haxe_js, map[string]*_bintree_t{
							}},
							"html.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_js, map[string]*_bintree_t{
							}},
							"html_ruby.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_html_ruby_js, map[string]*_bintree_t{
							}},
							"ini.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ini_js, map[string]*_bintree_t{
							}},
							"io.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_io_js, map[string]*_bintree_t{
							}},
							"jack.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jack_js, map[string]*_bintree_t{
							}},
							"jade.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jade_js, map[string]*_bintree_t{
							}},
							"java.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_java_js, map[string]*_bintree_t{
							}},
							"javascript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_javascript_js, map[string]*_bintree_t{
							}},
							"json.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_json_js, map[string]*_bintree_t{
							}},
							"jsoniq.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsoniq_js, map[string]*_bintree_t{
							}},
							"jsp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsp_js, map[string]*_bintree_t{
							}},
							"jsx.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_jsx_js, map[string]*_bintree_t{
							}},
							"julia.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_julia_js, map[string]*_bintree_t{
							}},
							"latex.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_latex_js, map[string]*_bintree_t{
							}},
							"less.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_less_js, map[string]*_bintree_t{
							}},
							"liquid.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_liquid_js, map[string]*_bintree_t{
							}},
							"lisp.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lisp_js, map[string]*_bintree_t{
							}},
							"livescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_livescript_js, map[string]*_bintree_t{
							}},
							"logiql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_logiql_js, map[string]*_bintree_t{
							}},
							"lsl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lsl_js, map[string]*_bintree_t{
							}},
							"lua.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lua_js, map[string]*_bintree_t{
							}},
							"luapage.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_luapage_js, map[string]*_bintree_t{
							}},
							"lucene.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_lucene_js, map[string]*_bintree_t{
							}},
							"makefile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_makefile_js, map[string]*_bintree_t{
							}},
							"markdown.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_markdown_js, map[string]*_bintree_t{
							}},
							"matlab.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_matlab_js, map[string]*_bintree_t{
							}},
							"mel.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mel_js, map[string]*_bintree_t{
							}},
							"mushcode.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mushcode_js, map[string]*_bintree_t{
							}},
							"mysql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_mysql_js, map[string]*_bintree_t{
							}},
							"nix.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_nix_js, map[string]*_bintree_t{
							}},
							"objectivec.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_objectivec_js, map[string]*_bintree_t{
							}},
							"ocaml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ocaml_js, map[string]*_bintree_t{
							}},
							"pascal.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pascal_js, map[string]*_bintree_t{
							}},
							"perl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_perl_js, map[string]*_bintree_t{
							}},
							"pgsql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_pgsql_js, map[string]*_bintree_t{
							}},
							"php.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_php_js, map[string]*_bintree_t{
							}},
							"plain_text.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_plain_text_js, map[string]*_bintree_t{
							}},
							"powershell.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_powershell_js, map[string]*_bintree_t{
							}},
							"praat.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_praat_js, map[string]*_bintree_t{
							}},
							"prolog.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_prolog_js, map[string]*_bintree_t{
							}},
							"properties.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_properties_js, map[string]*_bintree_t{
							}},
							"protobuf.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_protobuf_js, map[string]*_bintree_t{
							}},
							"python.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_python_js, map[string]*_bintree_t{
							}},
							"r.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_r_js, map[string]*_bintree_t{
							}},
							"rdoc.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rdoc_js, map[string]*_bintree_t{
							}},
							"rhtml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rhtml_js, map[string]*_bintree_t{
							}},
							"ruby.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_ruby_js, map[string]*_bintree_t{
							}},
							"rust.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_rust_js, map[string]*_bintree_t{
							}},
							"sass.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sass_js, map[string]*_bintree_t{
							}},
							"scad.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scad_js, map[string]*_bintree_t{
							}},
							"scala.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scala_js, map[string]*_bintree_t{
							}},
							"scheme.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scheme_js, map[string]*_bintree_t{
							}},
							"scss.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_scss_js, map[string]*_bintree_t{
							}},
							"sh.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sh_js, map[string]*_bintree_t{
							}},
							"sjs.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sjs_js, map[string]*_bintree_t{
							}},
							"smarty.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_smarty_js, map[string]*_bintree_t{
							}},
							"snippets.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_snippets_js, map[string]*_bintree_t{
							}},
							"soy_template.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_soy_template_js, map[string]*_bintree_t{
							}},
							"space.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_space_js, map[string]*_bintree_t{
							}},
							"sql.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_sql_js, map[string]*_bintree_t{
							}},
							"stylus.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_stylus_js, map[string]*_bintree_t{
							}},
							"svg.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_svg_js, map[string]*_bintree_t{
							}},
							"tcl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tcl_js, map[string]*_bintree_t{
							}},
							"tex.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_tex_js, map[string]*_bintree_t{
							}},
							"text.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_text_js, map[string]*_bintree_t{
							}},
							"textile.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_textile_js, map[string]*_bintree_t{
							}},
							"toml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_toml_js, map[string]*_bintree_t{
							}},
							"twig.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_twig_js, map[string]*_bintree_t{
							}},
							"typescript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_typescript_js, map[string]*_bintree_t{
							}},
							"vala.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vala_js, map[string]*_bintree_t{
							}},
							"vbscript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vbscript_js, map[string]*_bintree_t{
							}},
							"velocity.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_velocity_js, map[string]*_bintree_t{
							}},
							"verilog.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_verilog_js, map[string]*_bintree_t{
							}},
							"vhdl.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_vhdl_js, map[string]*_bintree_t{
							}},
							"xml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xml_js, map[string]*_bintree_t{
							}},
							"xquery.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_xquery_js, map[string]*_bintree_t{
							}},
							"yaml.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_snippets_yaml_js, map[string]*_bintree_t{
							}},
						}},
						"theme-ambiance.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_ambiance_js, map[string]*_bintree_t{
						}},
						"theme-chaos.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chaos_js, map[string]*_bintree_t{
						}},
						"theme-chrome.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_chrome_js, map[string]*_bintree_t{
						}},
						"theme-clouds.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_js, map[string]*_bintree_t{
						}},
						"theme-clouds_midnight.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_clouds_midnight_js, map[string]*_bintree_t{
						}},
						"theme-cobalt.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_cobalt_js, map[string]*_bintree_t{
						}},
						"theme-crimson_editor.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_crimson_editor_js, map[string]*_bintree_t{
						}},
						"theme-dawn.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dawn_js, map[string]*_bintree_t{
						}},
						"theme-dreamweaver.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_dreamweaver_js, map[string]*_bintree_t{
						}},
						"theme-eclipse.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_eclipse_js, map[string]*_bintree_t{
						}},
						"theme-github.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_github_js, map[string]*_bintree_t{
						}},
						"theme-idle_fingers.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_idle_fingers_js, map[string]*_bintree_t{
						}},
						"theme-katzenmilch.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_katzenmilch_js, map[string]*_bintree_t{
						}},
						"theme-kr_theme.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kr_theme_js, map[string]*_bintree_t{
						}},
						"theme-kuroir.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_kuroir_js, map[string]*_bintree_t{
						}},
						"theme-merbivore.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_js, map[string]*_bintree_t{
						}},
						"theme-merbivore_soft.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_merbivore_soft_js, map[string]*_bintree_t{
						}},
						"theme-mono_industrial.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_mono_industrial_js, map[string]*_bintree_t{
						}},
						"theme-monokai.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_monokai_js, map[string]*_bintree_t{
						}},
						"theme-pastel_on_dark.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_pastel_on_dark_js, map[string]*_bintree_t{
						}},
						"theme-solarized_dark.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_dark_js, map[string]*_bintree_t{
						}},
						"theme-solarized_light.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_solarized_light_js, map[string]*_bintree_t{
						}},
						"theme-terminal.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_terminal_js, map[string]*_bintree_t{
						}},
						"theme-textmate.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_textmate_js, map[string]*_bintree_t{
						}},
						"theme-tomorrow.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_js, map[string]*_bintree_t{
						}},
						"theme-tomorrow_night.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_js, map[string]*_bintree_t{
						}},
						"theme-tomorrow_night_blue.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_blue_js, map[string]*_bintree_t{
						}},
						"theme-tomorrow_night_bright.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_bright_js, map[string]*_bintree_t{
						}},
						"theme-tomorrow_night_eighties.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_tomorrow_night_eighties_js, map[string]*_bintree_t{
						}},
						"theme-twilight.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_twilight_js, map[string]*_bintree_t{
						}},
						"theme-vibrant_ink.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_vibrant_ink_js, map[string]*_bintree_t{
						}},
						"theme-xcode.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_theme_xcode_js, map[string]*_bintree_t{
						}},
						"worker-coffee.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_coffee_js, map[string]*_bintree_t{
						}},
						"worker-css.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_css_js, map[string]*_bintree_t{
						}},
						"worker-html.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_html_js, map[string]*_bintree_t{
						}},
						"worker-javascript.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_javascript_js, map[string]*_bintree_t{
						}},
						"worker-json.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_json_js, map[string]*_bintree_t{
						}},
						"worker-lua.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_lua_js, map[string]*_bintree_t{
						}},
						"worker-php.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_php_js, map[string]*_bintree_t{
						}},
						"worker-xquery.js": &_bintree_t{server_ui_vendor_ace_1_1_8_src_min_noconflict_worker_xquery_js, map[string]*_bintree_t{
						}},
					}},
				}},
				"jquery-2.1.3": &_bintree_t{nil, map[string]*_bintree_t{
					"jquery.min.js": &_bintree_t{server_ui_vendor_jquery_2_1_3_jquery_min_js, map[string]*_bintree_t{
					}},
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

