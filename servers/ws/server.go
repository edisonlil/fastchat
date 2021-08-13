package ws

import (
	"net/http"
)

const wsPattern = "/ws"

func StartWebSocket(addr string) {

	http.HandleFunc(wsPattern, run)

	http.ListenAndServe(addr, nil)

}

func run(w http.ResponseWriter, r *http.Request) {

	//1.获取Session
	session := NewSession(w, r)

	//2.Create Session Manager
	sessionManager := NewSessionManager()

	//3.异步启动监听器
	go sessionManager.StartListen()

	//4.注册Session
	sessionManager.Registrar <- session
	session.Manager = sessionManager

	//5.异步读取信息
	go session.Read(func(msg Message) error {

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
