package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		_, header, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}

		if header.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}

		if header.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(header, header.Filename)
		c.String(http.StatusOK, header.Filename)
	})

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
