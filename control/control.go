package control

import (
	"fastchat/base"
	"fastchat/servers/ws"
	"fmt"
	"github.com/gin-gonic/gin"
)

//SendMsg 发送消息
func SendMsg(c *gin.Context) {

	msg := base.Message{

		Id:   "1",
		Data: []byte(c.GetString("Data")),

		SenderAppId:  c.GetString("SenderAppId"),
		SenderOpenId: c.GetString("SenderOpenId"),
		SendTime:     c.GetUint64("SendTime"),

		AcceptAppId:  c.GetString("AcceptAppId"),
		AcceptOpenId: c.GetString("AcceptOpenId"),
	}

	//TODO...
	ws.Manager.Users[msg.AcceptAppId].WriteMsg(ws.WsMessage{
		Id:   "",
		Data: []byte(""),
	})

	fmt.Println(msg)

}
