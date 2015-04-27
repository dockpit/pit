package client

type Host interface {
	HostURL() string
	CertDir() string
}
