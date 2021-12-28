package handler

import (
	"github.com/gin-gonic/gin"
	"hyper-gateway/hypercloud/gateway/data"
	"log"
	"net/http"
)

var users = make(map[string]*data.User)

func GetUser(c *gin.Context) {
	user := users[c.Param("loginName")]
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
	users[user.LoginName] = &user
	c.IndentedJSON(http.StatusCreated, user)
}
