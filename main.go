package main

import (
	"fastchat/control"
	"fastchat/docs"
	"fastchat/servers/ws"
	"fastchat/store"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Golang FastChat API
// @version 1.0

// @contact.name API Support
// @contact.email edisonlil@163.com
//@host 127.0.0.1:8080
func main() {

	MongoInit()

	StartWeb()

	ws.StartWebSocket(":8000")

}

func MongoInit() {
	store.InitMongoClient()
}

func StartWeb() {

	router := gin.Default()

	apiDocInit(router)

	ginRouter(router)

	router.Run(":8080")
}

func apiDocInit(router *gin.Engine) {
	docs.SwagInit()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func ginRouter(router *gin.Engine) {
	router.POST("/user/login", control.UserLogin)
	router.POST("/send-msg", control.SendMsg)
	router.POST("/namespace/send-msg", control.SendMsgToNamespace)
}
