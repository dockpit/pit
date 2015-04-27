package host

type Local struct {
	hostURL string
	certDir string
}

func NewLocal(host, cert string) *Local {
	return &Local{host, cert}
}

func (l *Local) CertDir() string { return l.certDir }
func (l *Local) HostURL() string { return l.hostURL }

func (l *Local) IsUsable() bool {
	//@todo check if local docker host is usable
	return false
}
