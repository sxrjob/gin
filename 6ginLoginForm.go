package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.POST("/loginForm", loginForm)

	r.Run(":8888")
}

func loginForm(context *gin.Context) {

	var login Login
	err := context.Bind(&login)
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
