package main

import (
	//"10.1.1.42/research/samplego/misc"
	"github.com/gin-gonic/gin"
	"github.com/liuhongtian/hypergo/misc"
	"hyper-gateway/hypercloud/gateway/handler"
	"hyper-gateway/lht"
	"log"
)

//HyperCloud API Gateway
func main() {
	lht.Help()
	log.Println("Server is UP!")
	defer log.Println("Server will Down!")

	misc.Sample()
	misc.SampleWithParam("liuhongtian")

	router := gin.Default()
	router.GET("/users/:loginName", handler.GetUser)
	router.POST("/users", handler.CreateUser)

	router.Run(":9080")
}
