package control

import (
	"fastchat/domain"
	"fastchat/servers/ws"
	"fmt"
	"github.com/gin-gonic/gin"
)

//SendMsg 发送消息
func SendMsg(c *gin.Context) {

	msg := &domain.Message{}
	c.BindJSON(msg)
	ws.Manager.SendMsg(msg)
	fmt.Println(msg)

}

//SendMsgToNamespace 发送信息到指定命名空间
func SendMsgToNamespace(c *gin.Context) {

	msg := &domain.Message{}
	c.BindJSON(msg)
	ws.Manager.SendMsgToNamespace(msg)

	fmt.Println(msg)

}
