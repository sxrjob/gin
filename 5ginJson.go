package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"user" url:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" url:"password" xml:"password" binding:"required"`
}

//func main() {
//
//	r := gin.Default()
//
//	r.POST("loginJson", loginJson)
//
//	r.Run(":8888")
//}

func loginJson(context *gin.Context) {

	var login Login
	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login.User != "root" || login.Password != "admin" {
		context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "200"})

}
