package contract

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

//
//
//
type Contract struct {
	name      string
	resources []R
}

func NewContract(data *ContractData) (*Contract, error) {

	res := []R{}
	for _, r := range data.Resources {

		cases := []*Pair{}
		for _, c := range r.Cases {

			//create pair from data
			p, err := NewPairFromData(c)
			if err != nil {
				return nil, err
			}

			cases = append(cases, p)
		}

		res = append(res, NewResource(r.Pattern, cases...))
	}

	return &Contract{name: data.Name, resources: res}, nil
}

func (c *Contract) Name() string {
	return c.name
}

func (c *Contract) Resources() ([]R, error) {
	return c.resources, nil
}

// walk resources, actions and pairs to map all necessary states to
// create a router that mocks the contract
func (c *Contract) Mock() (*web.Mux, error) {
	mux := web.New()

	res, err := c.Resources()
	if err != nil {
		return mux, err
	}

	//look at each resource
	for _, r := range res {

		acs, err := r.Actions()
		if err != nil {
			return mux, err
		}

		//create middleware that routes to the correct example
		fn := func(ctx web.C, w http.ResponseWriter, r *http.Request) {

			//match the request method tot the correct action
			for _, a := range acs {

				if a.Method() == r.Method {

					//ask the action for a handle
					h, err := a.Handler(r)
					if err != nil {
						http.Error(w, fmt.Sprintf("%s", err.Error()), http.StatusInternalServerError)
						return
					}

					//serve mock and return
					h.ServeHTTP(w, r)
					return
				}
			}
		}

		//route everything that matches the resource pattern to the dynamic middleware
		mux.Handle(r.Pattern(), fn)
	}

	return mux, nil
}

// walk resources, actions and pairs to map all necessary states
// to test against the contract
func (c *Contract) States() (map[string][]string, error) {
	states := map[string][]string{}

	res, err := c.Resources()
	if err != nil {
		return states, err
	}

	for _, r := range res {
		as, err := r.Actions()
		if err != nil {
			return states, err
		}

		//loop over all pairs to map states
		for _, a := range as {
			for _, p := range a.Pairs() {
				for pname, g := range p.Given {

					//add state
					if _, ok := states[pname]; !ok {
						states[pname] = []string{}
					}

					//@todo prevend duplicate snames?
					states[pname] = append(states[pname], g.Name)

				}
			}
		}
	}

	return states, nil
}

// walk resources, actions and pairs to map all necessary dependencies
// to be mocked for isolation
func (c *Contract) Dependencies() (map[string][]string, error) {
	deps := map[string][]string{}

	res, err := c.Resources()
	if err != nil {
		return deps, err
	}

	for _, r := range res {
		as, err := r.Actions()
		if err != nil {
			return deps, err
		}

		//loop over all pairs to map deps
		for _, a := range as {
			for _, p := range a.Pairs() {
				for _, w := range p.While {
					if _, ok := deps[w.ID]; !ok {
						deps[w.ID] = []string{}

						//@todo also add casenames?
					}
				}
			}
		}
	}

	return deps, nil
}
