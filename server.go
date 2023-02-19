package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hkm007/hello_api/utils"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": utils.APP_NAME,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}