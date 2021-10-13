package filter

func NewFilterChain() *FilterChain {

	return &FilterChain{
		Filters: make([]func(chain *FilterChain) error, 10),
	}
}

//FilterChain 过滤链

type FilterChain struct {
	Filters []func(chain *FilterChain) error //过滤链
}

func (p *FilterChain) DoFilter() error {

	for _, filter := range p.Filters {

		if filter != nil {
			err := filter(p)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func (p *FilterChain) AddFilter(filter func(chain *FilterChain) error) {
	p.Filters = append(p.Filters, filter)
}
