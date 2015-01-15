package reporter

import (
	"io"
)

type Result struct {
	Succeeded int
	Failed    int
	Skipped   int
}

type StepFunc func(...interface{}) (int, string)

type P interface {
	ID() string
}

type R interface {
	Bytes() []byte
	String() string
	Enter(P, StepFunc, ...interface{})
	Exit()
	Path() string
	Printf(string, ...interface{})
	Pipe() io.Writer

	SetStatusCode(int)
	StatusCode() int

	Report(StepFunc, ...interface{})
	Success(StepFunc, ...interface{})
	Warning(StepFunc, ...interface{})
	Error(StepFunc, ...interface{})
}
