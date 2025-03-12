package main

import (
	"hyper-gateway/hypercloud/gateway/handler"
	"hyper-gateway/hypercloud/gateway/middleware"
	"hyper-gateway/lht"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liuhongtian/hypergo/misc"
)

// HyperCloud API Gateway
func main() {
	lht.Help()
	log.Println("Server is UP!")
	defer log.Println("Server will Down!")

	misc.Sample()
	misc.SampleWithParam("liuhongtian")

	router := gin.Default()

	// add middlewares
	router.Use(middleware.Logger())

	// add routes
	router.GET("/users/:loginName", handler.GetUser)
	router.POST("/users", handler.CreateUser)

	router.Run(":9080")
}
