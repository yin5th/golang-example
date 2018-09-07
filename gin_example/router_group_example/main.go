package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//v1版本
	//路由： /v1/login
	//      /v1/user
	v1 := r.Group("v1")
	{
		v1.POST("/login", loginV1)
		v1.GET("/user", userV1)
	}

	//v2版本
	//路由： /v2/login
	//      /v2/user
	v2 := r.Group("v2")
	{
		v2.POST("/login", loginV2)
		v2.GET("/user", userV2)
	}
	r.Run(":8088")
}

func loginV1(c *gin.Context) {
	c.String(http.StatusOK, "this is login handler of v1")
}

func loginV2(c *gin.Context) {
	c.String(http.StatusOK, "this is login handler of v2")
}

func userV1(c *gin.Context) {
	c.String(http.StatusOK, "this is user handler of v1")
}

func userV2(c *gin.Context) {
	c.String(http.StatusOK, "this is user handler of v2")
}
