package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchData(c *gin.Context){

	param := c.Query("q")
	fmt.Println(param)

	c.JSON(http.StatusOK,gin.H{
		"data":"hello",
	})
}