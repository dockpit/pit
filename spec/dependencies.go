package spec

import (
	"strings"
)

type Dependencies struct {
	// map[<service_name>]map[<endpoint_name>]map[<case_names>]bool
	Map map[string]map[string]map[string]bool
}

func NewDependencies(s S) (*Dependencies, error) {
	m := make(map[string]map[string]map[string]bool)

	//analyse endpoints
	for _, ep := range s.Endpoints() {
		for _, c := range ep.Cases() {
			for _, ws := range c.Whiles() {
				for sname, paths := range ws {
					for _, p := range paths {

						//lazily service map
						var ss map[string]map[string]bool
						var ok bool

						if ss, ok = m[sname]; !ok {
							ss = make(map[string]map[string]bool)
							m[sname] = ss
						}

						//split path names, first is endpoint, second case
						parts := strings.SplitN(p, "/", 2)

						//lazily create cases list per endpoint
						var cases map[string]bool
						if cases, ok = ss[parts[0]]; !ok {
							cases = make(map[string]bool)
							ss[parts[0]] = cases
						}

						//add case to cases if its not in there yet
						if _, ok = cases[parts[1]]; !ok {
							cases[parts[1]] = true
						}
					}
				}
			}
		}
	}

	return &Dependencies{
		Map: m,
	}, nil
}
