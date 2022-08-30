package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func webServer() {

	r := gin.Default()

	r.GET("/register", registerUser)
	r.GET("/queryUser", queryUser)
	r.Run(":9000")

}

func registerUser(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func queryUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "参数错误",
		})
	}

	err, user := queryUserWithID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  "success",
	})
}
