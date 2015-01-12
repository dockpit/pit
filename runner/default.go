package runner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/dockpit/exec"
	"github.com/dockpit/lang/manifest"
	"github.com/dockpit/pit/config"
	"github.com/dockpit/state"
)

func NewDefault(out io.Writer) *Default {
	return &Default{out}
}

// Default runner starts and shutsdown the test subject
// for every new test case, this is slow but guarantees
// that internal states are clean and doesn't require
// the developer to implement logic for handling dropping
// state connections
type Default struct {
	out io.Writer
}

func (d *Default) Name() string { return "default" }

func (d *Default) TemplatedCommandArgs(conf config.C, args ...string) ([]string, error) {

	res := []string{}
	for i, arg := range args {

		//parse each arg as a template
		t := template.New(fmt.Sprintf("arg_%d", i))
		t, err := t.Parse(arg)
		if err != nil {
			return args, err
		}

		//execute with conf for each arg and write
		//to a buffer
		buff := bytes.NewBuffer(nil)
		err = t.Execute(buff, conf.Data())
		if err != nil {
			return args, err
		}

		res = append(res, buff.String())
	}

	return res, nil
}

func (d *Default) RunOne(conf config.C, p *manifest.Pair, sm *state.Manager, subject *url.URL, docker *url.URL) (reerr error) {
	t := p.GenerateTest()

	//map states in manifest
	states := map[string]string{}
	for pname, g := range p.Given {
		states[pname] = g.Name

	}

	//add default states if given doesn't give one
	for _, pc := range conf.ProviderConfigs() {
		if _, ok := states[pc.Name()]; !ok {
			states[pc.Name()] = pc.DefaultState()
		}
	}

	//actually start the states
	for pname, sname := range states {
		_, err := sm.Start(pname, sname)
		if err != nil {
			return err
		}
	}

	//stop states
	defer func() {
		for pname, sname := range states {
			err := sm.Stop(pname, sname)
			if err != nil {
				reerr = err
				return
			}
		}
	}()

	//parse args as template
	args, err := d.TemplatedCommandArgs(conf, conf.RunConfig().Cmd...)
	if err != nil {
		return err
	}

	//start subject
	tcmd := exec.Command(args[0], args[1:]...)
	tcmd.Stdout = d.out
	tcmd.Stderr = d.out

	err = tcmd.StartWithTimeout(conf.RunConfig().ReadyTimeout, conf.RunConfig().ReadyExp)
	if err != nil {
		return err
	}

	//stop subject, @todo this to be configurable
	defer tcmd.StopWithTimeout(time.Second * 2)

	//execute the actual test
	err = t(subject.String(), docker.String(), http.DefaultClient, conf)
	if err != nil {
		return err
	}

	return nil
}

func (d *Default) Run(conf config.C, m manifest.M, sel Selector, sm *state.Manager, subject *url.URL, docker *url.URL) error {
	res, err := m.Resources()
	if err != nil {
		return err
	}

	//loop over each resource in the manifest
	for _, r := range res {
		fmt.Fprintf(d.out, "%s\n", r.Pattern())

		acs, err := r.Actions()
		if err != nil {
			return err
		}

		//loop over each action
		for _, a := range acs {
			fmt.Fprintf(d.out, "%s %s\n", a.Method(), r.Pattern())

			//for each test
			for _, p := range a.Pairs() {
				if sel.ShouldRun(p) {
					err = d.RunOne(conf, p, sm, subject, docker)
					if err != nil {
						return err
					}
				} else {
					fmt.Fprintf(d.out, "skipping\n")
				}
			}
		}
	}

	return nil
}
