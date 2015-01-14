package reporter

import (
	"bufio"
	"bytes"
	"io"
	"log"

	"github.com/simonwaldherr/golibs/cli"
)

type Terminal struct {
	path []P
	mw   io.Writer
	*bytes.Buffer
	*log.Logger
}

func NewTerminal(w io.Writer) *Terminal {
	buff := bytes.NewBuffer(nil)
	mw := io.MultiWriter(w, buff)

	return &Terminal{
		mw:     mw,
		Buffer: buff,
		Logger: log.New(mw, "", 0),
	}
}

func (t *Terminal) prefix() string {
	fix := ""
	for _, _ = range t.path {
		fix += "   "
	}
	return fix
}

func (t *Terminal) Pipe() io.Writer {
	r, w := io.Pipe()
	s := bufio.NewScanner(r)
	go func() {
		for s.Scan() {
			t.Print(s.Text())
		}
	}()

	return w
}

func (t *Terminal) Success(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(cli.Color(str, cli.Green))
}

func (t *Terminal) Report(stepFn StepFunc, args ...interface{}) {
	_, str := stepFn(args...)
	t.Print(str)
}

func (t *Terminal) Enter(p P, stepFn StepFunc, args ...interface{}) {
	if stepFn != nil {
		_, str := stepFn(args...)
		t.Print(cli.Underline(str))
	}

	t.path = append(t.path, p)
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Exit() {
	t.path = t.path[0 : len(t.path)-1]
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Path() string {
	path := ""
	for _, p := range t.path {
		path += "." + p.ID()
	}
	return path
}
