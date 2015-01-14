package reporter

import (
	"bytes"
	"io"
	"log"
	"strings"
)

type Terminal struct {
	part P
	path []string
	*bytes.Buffer
	*log.Logger
}

func NewTerminal(w io.Writer) *Terminal {
	buff := bytes.NewBuffer(nil)
	mw := io.MultiWriter(w, buff)

	return &Terminal{
		Buffer: buff,
		Logger: log.New(mw, "", 0),
	}
}

func (t *Terminal) prefix() string {
	fix := ""
	for _, _ = range t.path {
		fix += "\t"
	}
	return fix
}

func (t *Terminal) Enter(p P) {
	t.part = p
	t.path = append(t.path, t.part.ID())
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Exit() {
	t.path = t.path[0 : len(t.path)-1]
	t.Logger.SetPrefix(t.prefix())
}

func (t *Terminal) Path() string {
	return strings.Join(t.path, ".")
}
