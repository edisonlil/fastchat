package filter

import (
	"context"
	"net/http"
)

func NewFilterChain(w http.ResponseWriter, r *http.Request) *FilterChain {

	return &FilterChain{
		Filters:  make([]func(chain *FilterChain) error, 10),
		Response: w,
		Request:  r,
	}
}

//FilterChain 过滤链
type FilterChain struct {
	Ctx context.Context

	Response http.ResponseWriter

	Request *http.Request

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
