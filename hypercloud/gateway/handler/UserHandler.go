package handler

import (
	"hyper-gateway/hypercloud/gateway/data"
	"hyper-gateway/hypercloud/gateway/store"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := store.Users[c.Param("loginName")]
	if user != nil {
		log.Println("found user: " + user.ToString())
	} else {
		log.Println("user not found: " + c.Param("loginName"))
	}
	c.IndentedJSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user data.User
	c.ShouldBindJSON(&user)
	log.Println("created user: " + user.ToString())
	store.Users[user.LoginName] = &user
	c.IndentedJSON(http.StatusCreated, user)
}
