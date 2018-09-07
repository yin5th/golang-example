package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//GET 获取
	r.GET("user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "get user info success",
		})
	})
	//POST 新增
	r.POST("user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "create new user info success",
		})
	})

	//PUT 修改
	r.PUT("user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "update user info success",
		})
	})

	//DELETE 删除
	r.DELETE("user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "delete user info success",
		})
	})

	r.Run(":8088")
}
