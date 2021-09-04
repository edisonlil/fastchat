package filter

var WsFilterChain = NewFilterChain()

func NewFilterChain() *FilterChain {

	return &FilterChain{
		Filters: make([]func(chain *FilterChain), 10),
	}
}

//FilterChain 过滤链
type FilterChain struct {
	Filters []func(chain *FilterChain) //过滤链

}

func (p *FilterChain) DoFilter() {

	for _, filter := range p.Filters {

		if filter != nil {
			filter(p)
		}
	}

}

func (p *FilterChain) AddFilter(filter func(chain *FilterChain)) {
	p.Filters = append(p.Filters, filter)
}
