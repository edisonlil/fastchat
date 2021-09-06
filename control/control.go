package control

import (
	"fastchat/domain"
	"fastchat/servers/ws"
	"fmt"
	"github.com/gin-gonic/gin"
)

//SendMsg 发送消息
func SendMsg(c *gin.Context) {

	msg := &domain.Message{

		Id:   "1",
		Data: []byte(c.GetString("Data")),

		Namespace:    c.GetString("Namespace"),
		SenderOpenId: c.GetString("SenderOpenId"),
		SendTime:     c.GetUint64("SendTime"),

		AcceptOpenId: c.GetString("AcceptOpenId"),
	}

	//TODO...
	ws.Manager.SendMsg(msg)

	fmt.Println(msg)

}

//SendMsgToNamespace 发送信息到指定命名空间
func SendMsgToNamespace() {

}
