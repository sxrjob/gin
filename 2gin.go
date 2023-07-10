package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "sxr")
		c.String(http.StatusOK, name+" is "+name)
	})

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
