package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func webServer() {

	r := gin.Default()

	r.POST("/register", registerUser)
	r.GET("/queryUser", queryUser)
	r.Run(":9000")

}

func registerUser(c *gin.Context) {

	var user Person
	err := c.ShouldBind(&user)
	fmt.Printf("user: %v\n", user)
	if err != nil {
		fmt.Printf("参数错误err: %v\n", err)
		c.JSON(400, gin.H{
			"code": 400,
			"data": err,
			"msg":  "参数错误",
		})
		return
	}

	if err = addUser(user); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "插入数据库失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": nil,
		"msg":  "success",
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
		return
	}

	err, user := queryUserWithID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": err,
			"msg":  "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  "success",
	})
}
