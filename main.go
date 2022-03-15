package main

import (
	"ArSearch/pkg/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/query",controller.SearchData)
	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}