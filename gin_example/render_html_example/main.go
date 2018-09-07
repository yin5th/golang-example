package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//将所有模板先加载到内存中
	r.LoadHTMLGlob("templates/*")

	r.GET("/test/html", testHtml)
	r.Run(":8088")
}

func testHtml(c *gin.Context) {
	//渲染到html模板
	c.HTML(http.StatusOK, "test/index.html", gin.H{
		"title":   "test",
		"content": "this is a test of render html",
	})
}
