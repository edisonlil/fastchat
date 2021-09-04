package ws

import (
	"context"
	"fastchat/auth"
	"fastchat/filter"
	"fastchat/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const wsPattern = "/ws"

func StartWebSocket(addr string) {

	http.HandleFunc(wsPattern, run)

	http.ListenAndServe(addr, nil)
}

func InitFilter(chain *filter.FilterChain, ctx *filter.HttpContext) {

	//Jwt鉴权
	chain.AddFilter(func(chain *filter.FilterChain) error {

		token := ctx.Request.Header.Get("Authorization")

		claims, err := auth.ParseJwtToken(token)

		if err != nil {
			log.Error(err.Error())
			ctx.Response.Write([]byte(err.Error()))
			ctx.Response.WriteHeader(500)

			//TODO 错误则请求返回失败,不执行下一个过滤器
			return err
		}

		context.WithValue(ctx.Ctx, "UserDetail", service.GetUserById(claims.UserId))

		return nil
	})

}

func run(w http.ResponseWriter, r *http.Request) {

	//过滤器
	wsFilterChain := filter.NewFilterChain()

	//初始化过滤器
	InitFilter(wsFilterChain, filter.NewHttpContext(w, r))

	//执行过滤器
	err := wsFilterChain.DoFilter()

	if err != nil {
		return
	}

	//执行
	servlet(w, r)

}

func servlet(w http.ResponseWriter, r *http.Request) {

	//1.获取Session && 升级为ws
	session := NewSession(w, r)

	//2.Create Session Manager
	sessionManager := InitSessionManager()

	//3.异步启动监听器
	go sessionManager.StartListen()

	//4.注册Session
	sessionManager.Registrar <- session
	session.Manager = sessionManager

	//5.异步读取信息
	go session.Read(func(msg WsMessage) error {

		var err error
		if msg.MsgType == CloseConn {
			return ReadExitError{}
		}

		err = session.WriteMsg(msg)

		if err != nil {
			return err
		}

		return nil
	})

}
