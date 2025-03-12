package main

import (
	"fmt"
	"hyper-gateway/config"
	"hyper-gateway/hypercloud/gateway/handler"
	"hyper-gateway/hypercloud/gateway/middleware"
	"hyper-gateway/lht"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/liuhongtian/hypergo/misc"
)

// HyperCloud API Gateway
func main() {
	// 加载配置文件
	if err := config.LoadConfig("config.yml"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	lht.Help()
	log.Println("Server is UP!")
	defer log.Println("Server will Down!")

	misc.Sample()
	misc.SampleWithParam("liuhongtian")

	r := gin.Default()

	// 添加中间件
	r.Use(middleware.Logger())

	// 动态添加路由
	for _, route := range config.AppConfig.Routes {
		r.Handle(route.Method, route.Path, handler.Handle(route.Handler))
	}

	// 从配置文件读取端口
	r.Run(fmt.Sprintf(":%d", config.AppConfig.Server.Port))
}
