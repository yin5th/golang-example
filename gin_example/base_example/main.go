package main

import "github.com/gin-gonic/gin"

func pingHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    "ok",
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.GET("/ping", pingHandle)

	r.Run(":8088")
}
