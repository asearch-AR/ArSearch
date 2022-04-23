package controller

import (
	"ArSearch/pkg/service"
	"github.com/gin-gonic/gin"
)

func SearchData(c *gin.Context) {
	param := c.Query("q")
	res, err := service.SearchInEs(param)

	if err != nil {
		c.JSON(500, gin.H{
			"err_msg":err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}

func SearchOnMirror(c *gin.Context) {
	param := c.Query("q")
	res, err := service.SearchMirrorData(param)

	if err != nil {
		c.JSON(500, gin.H{
			"err_msg":err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}

