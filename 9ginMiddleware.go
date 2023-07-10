package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	r := gin.Default()

	//注册中间件
	r.Use(MiddleWare())

	{
		r.GET("/ce", func(c *gin.Context) {
			value, _ := c.Get("request")
			fmt.Println("request:", value)
			// 页面接收
			c.JSON(200, gin.H{"request": value})
		})
	}

	r.Run(":8888")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
