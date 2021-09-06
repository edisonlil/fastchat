package ws

import (
	"context"
	"fastchat/domain"
	"net/http"
)

const CurrentUser = "CurrentUser"

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

func (p *HttpContext) GetCurrentUser() *domain.User {
	return p.Ctx.Value(CurrentUser).(*domain.User)
}

func (p *HttpContext) SetCurrentUser(user *domain.User) {
	context.WithValue(p.Ctx, CurrentUser, user)
}
