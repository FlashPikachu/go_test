package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func m1(ctx *gin.Context) {
	fmt.Println("m1 in")
	start := time.Now()
	ctx.Next()
	//ctx.Abort()  阻止调用后续的函数
	//end := time.Now()

	cost := time.Since(start)
	fmt.Printf("cost time : %v\n", cost)
	value, exists := ctx.Get("name")
	if exists {
		fmt.Println(value)
	}
	fmt.Println("m1 out")
}

func m2(ctx *gin.Context) {
	fmt.Println("m2 in")
	//ctx.Next()
	//ctx.Abort()  阻止调用后续的函数
	//end := time.Now()
	//ctx.Abort()
	ctx.Set("name", "pikachu")
	fmt.Println("m2 out")
}

func main() {
	r := gin.Default()
	r.Use(m1, m2)
	r.GET("/index", func(context *gin.Context) {
		time.Sleep(time.Second * 1)
		context.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})

	r.GET("/book", func(context *gin.Context) {
		time.Sleep(time.Second * 1)
		value, exists := context.Get("name")
		if !exists {
			value = "匿名用户"

		}
		context.JSON(http.StatusOK, gin.H{
			"msg":  "book",
			"name": value,
		})
	})

	r.Run(":9000")
}
