package spec

import (
	"io"

	"github.com/zenazn/goji/web"
)

//Specfication Mock
type M interface {
	Mux() *web.Mux
}

//Specification Endpoint Case interface
type C interface {
	Study() (*Study, error)
	Name() string
}

//Specification Endpoint interface
type EP interface {
	Name() string
	Cases() []C
}

//Specification interface
type S interface {
	Endpoints() []EP
	Mock() (M, error)
}

//Loader interface
type L interface {
	Load(loc string) (io.ReadCloser, error)
}

//Factory Interface
type F interface {
	Create(loc string) (S, error)
}
