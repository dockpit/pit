package runner

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

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
			for _, tt := range a.Tests() {

				//execute the actual test
				err := tt(subject.String(), docker.String(), http.DefaultClient, sm, conf)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// //start default states for each provider
// for _, pc := range conf.ProviderConfigs() {
// 	_, err = sm.Start(pc.Name(), pc.DefaultState())
// 	if err != nil {
// 		return nil, nil, err
// 	}
// }

// //get command from configuration
// run := conf.RunConfig().Cmd

// //create command that can has timedout start& stop
// tcmd := exec.Command(run[0], run[1:]...)
// tcmd.Stdout = os.Stdout
// tcmd.Stderr = os.Stderr

// err = tcmd.StartWithTimeout(conf.RunConfig().ReadyTimeout, conf.RunConfig().ReadyExp)
// if err != nil {
// 	return nil, nil, err
// }

// //@todo make this configurable
// defer tcmd.StopWithTimeout(time.Second)

// //run all tests, for all resources
// res, err := contract.Resources()
// if err != nil {
// 	return nil, nil, err
// }

// errs := []error{}
// for _, r := range res {
// 	acs, err := r.Actions()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	for _, a := range acs {
// 		for _, tt := range a.Tests() {

// 			err := tt(host, dhost, http.DefaultClient, sm, conf)
// 			if err != nil {
// 				//@todo handle failed tests better

// 				//gather testing errors
// 				errs = append(errs, err)

// 				fmt.Fprintf(c.out, "%s\n", err.Error())
// 			}
// 		}
// 	}
// }
