package contract

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

		res = append(res, NewResource(r.Name, r.Pattern, cases...))
	}

	return &Contract{name: data.Name, resources: res}, nil
}

func (c *Contract) Name() string {
	return c.name
}

func (c *Contract) Resources() ([]R, error) {
	return c.resources, nil
}
