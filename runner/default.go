package runner

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dockpit/exec"
	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/contract"
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

func (d *Default) RunOne(conf config.C, p *contract.Pair, sm *state.Manager, subject *url.URL, docker *url.URL) (reerr error) {
	t := p.GenerateTest()

	//start states
	for pname, g := range p.Given {
		_, err := sm.Start(pname, g.Name)
		if err != nil {
			return err
		}
	}

	//stop states
	defer func() {
		for pname, g := range p.Given {
			err := sm.Stop(pname, g.Name)
			if err != nil {
				reerr = err
				return
			}
		}
	}()

	//start subject
	tcmd := exec.Command(conf.RunConfig().Cmd[0], conf.RunConfig().Cmd[1:]...)
	tcmd.Stdout = d.out
	tcmd.Stderr = d.out

	err := tcmd.StartWithTimeout(conf.RunConfig().ReadyTimeout, conf.RunConfig().ReadyExp)
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

func (d *Default) Run(conf config.C, c contract.C, sm *state.Manager, subject *url.URL, docker *url.URL) error {
	res, err := c.Resources()
	if err != nil {
		return err
	}

	//loop over each resource in the contract
	for _, r := range res {
		fmt.Fprintf(d.out, "Resource '%s'", r.Pattern())

		acs, err := r.Actions()
		if err != nil {
			return err
		}

		//loop over each action
		for _, a := range acs {

			//for each test
			for _, p := range a.Pairs() {
				err = d.RunOne(conf, p, sm, subject, docker)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
