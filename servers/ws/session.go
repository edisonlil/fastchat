package ws

import (
	"encoding/json"
	"fastchat/domain"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"time"
)

type Session struct {
	Id string //会话ID

	UserId string // 用户Id，用户登录以后才有

	Healthy int64 // 用户上次心跳时间

	LoginTime int64 // 登录时间 登录以后才有

	conn *websocket.Conn // 用户连接

	Ctx *HttpContext //上下文

	Manager *SessionManager //会话管理器
}

type ReadExitError struct {
}

func (p ReadExitError) Error() string {
	return "Read Exit!"
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	CloseConn = 11
)

func NewSession(w http.ResponseWriter, r *http.Request, ctx *HttpContext) *Session {

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	return &Session{
		UserId:    ctx.GetCurrentUser().Id,
		Healthy:   time.Now().Unix(),
		LoginTime: time.Now().Unix(),
		conn:      conn,
		Ctx:       ctx,
	}
}

func (p Session) Read(reader func(msg *domain.Message) error) {

	for {

		mt, message, err := p.conn.ReadMessage()

		if err != nil {
			panic(err.Error())
		}

		var msg = &domain.Message{}

		err = json.Unmarshal(message, &msg)

		if err != nil {
			msg.Data = message
			msg.MsgType = mt
		}

		err = reader(msg)

		if reflect.TypeOf(err) == reflect.TypeOf(ReadExitError{}) {
			break
		}

		if err != nil {
			panic(err.Error())
		}
	}

	log.Infof("用户ID为 %s 已下线", p.UserId)
	p.Manager.Logout <- &p

}

func (p Session) ReadMsg() string {
	mt, message, err := p.conn.ReadMessage()

	if err != nil {
		panic(err.Error())
	}

	if mt != websocket.TextMessage {
		panic("the message not text.")
	}

	return string(message)
}

func (p Session) WriteMsg(msg *domain.Message) error {

	err := p.conn.WriteMessage(websocket.TextMessage, msg.Data)
	if err != nil {
		return err
	}
	return err
}

func (p Session) Close() {
	p.conn.Close()
	p.conn.UnderlyingConn().Close()
}
