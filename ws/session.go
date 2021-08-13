package ws

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Session struct {
	id string //会话ID

	UserId string // 用户Id，用户登录以后才有

	conn *websocket.Conn // 用户连接

	Healthy uint64 // 用户上次心跳时间

	LoginTime uint64 // 登录时间 登录以后才有

}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

type Result struct {
	Data []byte

	MsgType int
}

func (p Session) Read(reader func(result Result) error) {

	for {

		mt, message, err := p.conn.ReadMessage()
		if err != nil {
			panic(err.Error())
		}

		err = reader(Result{
			Data:    message,
			MsgType: mt,
		})

		if err != nil {
			panic(err.Error())
		}
	}

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

func (p Session) WriteMsg(mt int, message []byte) error {
	err := p.conn.WriteMessage(mt, message)
	if err != nil {
		log.Error("write:", err)
		return err
	}
	return err
}

func (p Session) Close() {
	p.conn.Close()
}
