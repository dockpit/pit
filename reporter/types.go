package reporter

type P interface {
	ID() string
}

type R interface {
	Printf(string, ...interface{})
	Bytes() []byte
	String() string
	Enter(P)
	Exit()
	Path() string
}
