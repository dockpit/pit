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

// server_ui_list_isolations_html reads file data from disk. It returns an error on failure.
func server_ui_list_isolations_html() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/list_isolations.html"
	name := "server/ui/list_isolations.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
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

// server_ui_src_js_components_hello_jsx reads file data from disk. It returns an error on failure.
func server_ui_src_js_components_hello_jsx() (*asset, error) {
	path := "/Users/advanderveer/Documents/Projects/go/src/github.com/dockpit/pit/server/ui/src/js/components/Hello.jsx"
	name := "server/ui/src/js/components/Hello.jsx"
	bytes, err := bindata_read(path, name)
	if err != nil {
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
	"server/ui/.gitignore": server_ui_gitignore,
	"server/ui/.jshintrc": server_ui_jshintrc,
	"server/ui/add_dep.html": server_ui_add_dep_html,
	"server/ui/add_state.html": server_ui_add_state_html,
	"server/ui/dist/app.js": server_ui_dist_app_js,
	"server/ui/edit_state.html": server_ui_edit_state_html,
	"server/ui/foot.html": server_ui_foot_html,
	"server/ui/gulpfile.js": server_ui_gulpfile_js,
	"server/ui/head.html": server_ui_head_html,
	"server/ui/list_isolations.html": server_ui_list_isolations_html,
	"server/ui/package.json": server_ui_package_json,
	"server/ui/src/js/app.js": server_ui_src_js_app_js,
	"server/ui/src/js/components/Hello.jsx": server_ui_src_js_components_hello_jsx,
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
			".gitignore": &_bintree_t{server_ui_gitignore, map[string]*_bintree_t{
			}},
			".jshintrc": &_bintree_t{server_ui_jshintrc, map[string]*_bintree_t{
			}},
			"add_dep.html": &_bintree_t{server_ui_add_dep_html, map[string]*_bintree_t{
			}},
			"add_state.html": &_bintree_t{server_ui_add_state_html, map[string]*_bintree_t{
			}},
			"dist": &_bintree_t{nil, map[string]*_bintree_t{
				"app.js": &_bintree_t{server_ui_dist_app_js, map[string]*_bintree_t{
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
			"list_isolations.html": &_bintree_t{server_ui_list_isolations_html, map[string]*_bintree_t{
			}},
			"package.json": &_bintree_t{server_ui_package_json, map[string]*_bintree_t{
			}},
			"src": &_bintree_t{nil, map[string]*_bintree_t{
				"js": &_bintree_t{nil, map[string]*_bintree_t{
					"app.js": &_bintree_t{server_ui_src_js_app_js, map[string]*_bintree_t{
					}},
					"components": &_bintree_t{nil, map[string]*_bintree_t{
						"Hello.jsx": &_bintree_t{server_ui_src_js_components_hello_jsx, map[string]*_bintree_t{
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

