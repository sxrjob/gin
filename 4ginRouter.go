package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v2 := r.Group("/v2")

	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "sxr")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "hyx")
	c.String(http.StatusOK, "hello "+name)
}
