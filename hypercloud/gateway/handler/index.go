package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Handle(name string) func(*gin.Context) {
	log.Println("do " + name)
	switch name {
	case "GetUser":
		return GetUser
	case "CreateUser":
		return CreateUser
	}
	return nil
}
