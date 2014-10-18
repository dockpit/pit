package spec

import (
	"io"
)

//Specification Endpoint Case interface
type C interface {
}

//Specification Endpoint interface
type EP interface {
	Name() string
	Cases() []C
}

//Specification interface
type S interface {
	Endpoints() []EP
}

//Loader interface
type L interface {
	Load(loc string) (io.ReadCloser, error)
}

//Factory Interface
type F interface {
	Create(loc string) (S, error)
}
