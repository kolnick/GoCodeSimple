package _3_filetransfer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"testing"
)

func TestMultipartFormFile(t *testing.T) {
	engine := gin.Default()
	engine.Static("/static", "../static")
	engine.MaxMultipartMemory = 8 << 20 // 最大8MB
	engine.POST("/upload", func(c *gin.Context) {
		// 获取表单参数
		username := c.PostForm("username")
		fmt.Println("username:", username)
		files := c.Request.MultipartForm.File["files"]

		ok := true
		for _, fileHeader := range files {
			filename := fileHeader.Filename
			dst := path.Join("../static/upload", filename)
			err := c.SaveUploadedFile(fileHeader, dst)
			if err != nil {
				ok = false
			}
		}
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"success":  true,
				"username": username,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success":  false,
				"username": username,
			})
		}

	})
	engine.Run()
}

func TestSingleFileTransfer(t *testing.T) {
	engine := gin.Default()
	engine.Static("/static", "../static")
	engine.MaxMultipartMemory = 8 << 20 // 最大8MB
	engine.POST("/upload", func(c *gin.Context) {
		// 获取表单参数
		username := c.PostForm("username")
		fmt.Println("username:", username)
		file, err := c.FormFile("file")

		//		file, header, err := c.Request.FormFile("file")
		//if err != nil {
		//	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		//	return
		//}
		//defer file.Close()
		//// 创建目标文件
		//out, err := os.Create(header.Filename)
		//if err != nil {
		//	c.String(http.StatusInternalServerError, fmt.Sprintf("create file err: %s", err.Error()))
		//	return
		//}
		//defer out.Close()
		//// 将上传的文件内容复制到目标文件
		//_, err = io.Copy(out, file)
		dst := path.Join("../static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("copy file err: %s", err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"username": username,
			"dst":      dst,
		})
	})
	engine.Run()
}
