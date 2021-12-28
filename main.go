package main

import (
	"github.com/gin-gonic/gin"
	"hyper-gateway/hypercloud/gateway/handler"
	"log"
)

//HyperCloud API Gateway
func main() {
	log.Println("Server is UP!")
	defer log.Println("Server will Down!")

	router := gin.Default()
	router.GET("/users/:loginName", handler.GetUser)
	router.POST("/users", handler.CreateUser)

	router.Run(":9080")
}
