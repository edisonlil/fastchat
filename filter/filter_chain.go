package filter

import (
	"context"
	"net/http"
)

func NewFilterChain() *FilterChain {

	return &FilterChain{
		Filters: make([]func(chain *FilterChain) error, 10),
	}
}

type HttpContext struct {
	Ctx context.Context

	Response http.ResponseWriter

	Request *http.Request
}

func NewHttpContext(w http.ResponseWriter, r *http.Request) *HttpContext {

	return &HttpContext{
		Ctx:      context.TODO(),
		Request:  r,
		Response: w,
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
