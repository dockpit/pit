package tool

type Manager struct {
	Dir string
}

// Manages microservice dependencies, uses go tool approach
// meaning that it retrieves dependencies by using go-like import paths
func NewManager(dir string) *Manager {
	return &Manager{
		Dir: dir,
	}
}

// install the dependency but if it already exists, do nothing
func (m *Manager) Install(ipath string) error {

	//create package
	p := NewPackage(ipath)

	//download it
	return DownloadPackage(p, m.Dir, false)
}

// update the dependency, if its not installed install
func (m *Manager) Upsert(ipath string) error {

	//create package
	p := NewPackage(ipath)

	//download it with update flag enabled
	return DownloadPackage(p, m.Dir, true)
}
