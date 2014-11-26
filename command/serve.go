package command

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/zenazn/goji/bind"

	"github.com/dockpit/pit/contract"
)

var tmpl_serve = `Served successful!`

type Serve struct {
	*cmd
}

func NewServe(out io.Writer) *Serve {
	return &Serve{
		cmd: newCmd(out),
	}
}

func (c *Serve) Name() string {
	return "serve"
}

func (c *Serve) Description() string {
	return fmt.Sprintf("...")
}

func (c *Serve) Usage() string {
	return "Serve a folder of examples"
}

func (c *Serve) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, cli.StringFlag{Name: "bind, b", Value: ":8000", Usage: fmt.Sprintf("Specify on which address the server will listen.")})

	return fs
}

func (c *Serve) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Serve) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//catch relevant signals for terminating and reloading
	hup := make(chan os.Signal)
	stop := make(chan os.Signal)
	signal.Notify(hup, syscall.SIGHUP)
	signal.Notify(stop, os.Kill, os.Interrupt)

	//the listener can be reused
	l := bind.Socket(strings.TrimSpace(ctx.String("bind")))

	//reload http.DefaultServerMux
	loadMux := func() error {

		//get contract
		ctr, err := c.ParseExamples(ctx)
		if err != nil {
			return err
		}

		//create mux from contract
		mock := contract.NewMock(ctr, c.ExamplesPath(ctx))
		mux, err := mock.Mux()
		if err != nil {
			return err
		}

		//create and replace default mux
		h := http.NewServeMux()
		h.Handle("/", mux)
		http.DefaultServeMux = h

		return nil
	}

	//handle signalsk
	//@todo this probably causes run conditions within the http server
	go func() {
		for {
			select {

			//incase of a KILL/INT signal: stop listening
			case <-stop:

				fmt.Fprintf(c.out, "Stopping... \n")

				//close the current listener which
				//unblocks the http.Serve()
				l.Close()

			//in case of a HUP signal: reload the mux
			case <-hup:
				fmt.Fprintf(c.out, "Reloading... \n")

				//reload the default mux
				err := loadMux()
				if err != nil {
					fmt.Fprintln(c.out, err)
				}
			}
		}
	}()

	//load and serve with default server mux
	fmt.Fprintf(c.out, "Serving (%s)...\n", ctx.String("bind"))
	err := loadMux()
	if err != nil {
		return nil, nil, err
	}

	http.Serve(l, nil)

	//we're done serving, render results
	return template.Must(template.New("serve.success").Parse(tmpl_serve)), nil, nil
}
