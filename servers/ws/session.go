package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

type Session struct {
	Id string //会话ID

	Manager *SessionManager

	UserId string // 用户Id，用户登录以后才有

	conn *websocket.Conn // 用户连接

	Healthy uint64 // 用户上次心跳时间

	LoginTime uint64 // 登录时间 登录以后才有

}

type Message struct {
	Id string //消息ID

	SourceId string //发送者ID

	TargetId string //目标ID

	Data []byte //数据

	MsgType int //消息类型

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

func NewSession(w http.ResponseWriter, r *http.Request) *Session {

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	return &Session{
		conn: conn,
	}
}

func (p Session) Read(reader func(msg Message) error) {

	for {

		mt, message, err := p.conn.ReadMessage()

		if err != nil {
			panic(err.Error())
		}

		var msg = Message{}

		err = json.Unmarshal(message, &msg)

		if err != nil {

			msg = Message{
				MsgType: mt,
				Data:    message,
			}

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

func (p Session) WriteMsg(msg Message) error {

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
