package ws

import (
	"fastchat/manager"
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
	defer session.Close()

	//2.Create Session Manager
	sessionManager := manager.NewSessionManager()

	//3.异步启动监听器
	go sessionManager.StartListen()

	//4.注册Session
	sessionManager.Registrar <- session

	//5.异步读取信息
	go session.Read(func(result Result) error {

		err := session.WriteMsg(result.MsgType, result.Data)
		if err != nil {
			return err
		}
		return nil
	})

}
