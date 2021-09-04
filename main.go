package main

import (
	"fastchat/store"
	"github.com/gin-gonic/gin"
)

func main() {

	//MongoInit()
	//
	//StartWeb()
	//
	//ws.StartWebSocket(":8000")
}

func MongoInit() {
	store.InitMongoClient()
}

func StartWeb() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
