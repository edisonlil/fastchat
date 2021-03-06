package ws

import (
	"fastchat/auth"
	"fastchat/config"
	"fastchat/domain"
	"fastchat/err"
	"fastchat/filter"
	"fastchat/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const wsPattern = "/ws"

var prop = config.GetWebsocketProp()

func StartWebSocket() {

	http.HandleFunc(prop.Pattern, run)

	http.ListenAndServe(prop.Addr, nil)
}

func InitFilter(chain *filter.FilterChain, ctx *HttpContext) {

	//Jwt鉴权
	chain.AddFilter(func(chain *filter.FilterChain) error {

		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.Response.Write([]byte("未携带Token..."))
			return &err.UnauthorizedError{
				Msg: "Token Unauthorized",
			}
		}

		claims, err := auth.ParseJwtToken(token)

		if err != nil {
			log.Error(err.Error())
			ctx.Response.Write([]byte(err.Error()))
			ctx.Response.WriteHeader(500)
			//错误则请求返回失败,不执行下一个过滤器
			return err
		}

		//设置当前用户
		ctx.SetCurrentUser(service.GetUserById(claims.UserId))

		return nil
	})

}

func run(w http.ResponseWriter, r *http.Request) {

	//过滤器
	wsFilterChain := filter.NewFilterChain()

	httpContext := NewHttpContext(w, r)
	//初始化过滤器
	InitFilter(wsFilterChain, httpContext)

	//执行过滤器
	err := wsFilterChain.DoFilter()

	if err != nil {
		return
	}

	servlet(w, r, httpContext)

}

func servlet(w http.ResponseWriter, r *http.Request, ctx *HttpContext) {

	//1.获取Session && 升级为ws
	session := NewSession(w, r, ctx)

	//2.Create Session Manager
	sessionManager := InitSessionManager()

	//3.异步启动监听器
	go sessionManager.StartListen()

	//4.注册Session
	sessionManager.Registrar <- session
	session.Manager = sessionManager

	//5.异步读取信息
	go session.Read(func(msg *domain.Message) error {

		var err error
		if msg.MsgType == CloseConn {
			return ReadExitError{}
		}

		if msg.MsgType == Healthy {
			session.Healthy = time.Now().Unix()
		}

		err = session.WriteMsg(msg)

		if err != nil {
			return err
		}

		return nil
	})

}
