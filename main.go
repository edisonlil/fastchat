package main

import (
	"fastchat/auth"
	"fastchat/store"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	token, _ := auth.CreateJwtToken(map[string]interface{}{

		"id": 123,
	})

	fmt.Println(token)

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
