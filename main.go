package main

import (
	"ArSearch/pkg/controller"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pprof.Register(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/query",controller.SearchData)
	r.GET("/query_mirror",controller.SearchOnMirror)
	r.Run("0.0.0.0:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}