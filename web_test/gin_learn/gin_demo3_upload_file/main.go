package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("web_test/gin_learn/gin_demo3_upload_file/resource/upload.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})

	router.GET("/redirect", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//})
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	router.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("f1")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			log.Println(file.Filename)

			dst := path.Join("F:/temp/go_files", file.Filename)

			context.SaveUploadedFile(file, dst)
			log.Println(dst)
			context.JSON(http.StatusOK, gin.H{
				"message:":  fmt.Sprintf("'%s' uploaded!", file.Filename),
				"file_path": dst,
			})
		}
	})

	router.Run(":9000")
}
