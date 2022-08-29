package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "白日依山尽",
		"msg":  "教练我想打篮球",
	})

}
func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("hello", helloHandler)

	return r
}
