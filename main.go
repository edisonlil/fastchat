package main

import (
	"fastchat/filter"
	"fastchat/store"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	wsFilterChain := filter.WsFilterChain

	wsFilterChain.AddFilter(func(chain *filter.FilterChain) {
		fmt.Println("1")
	})

	wsFilterChain.AddFilter(func(chain *filter.FilterChain) {
		fmt.Println("2")
	})

	wsFilterChain.DoFilter()
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
