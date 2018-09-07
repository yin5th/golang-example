package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	//多文件上传
	r.POST("uploads", uploadMultiHandle)
	//单文件上传
	r.POST("upload", uploadHandle)
	r.Run(":8088")
}

func uploadMultiHandle(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("form request failed, err:%v\n", err)
		return
	}
	//中括号中file为上传的字段名  多个文件上传时 字段名不变即一个字段传多个文件
	files := form.File["file"]
	//遍历所有文件 逐个保存
	var succNum int
	var failedNum int
	for index, file := range files {
		savepath := fmt.Sprintf("/go/src/data/%s_%d", file.Filename, index)
		err = c.SaveUploadedFile(file, savepath)
		if err != nil {
			fmt.Printf("file save failed, err:%v\n", err)
			failedNum++
			continue
		}
		succNum++
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": fmt.Sprintf("total %d files uploads, %d success, %d failed", len(files), succNum, failedNum),
	})
}

func uploadHandle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("file upload failed, err:%v\n", err)
		return
	}
	savepath := fmt.Sprintf("/go/src/data/%s", file.Filename)
	c.SaveUploadedFile(file, savepath)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "file upload success",
	})
}
