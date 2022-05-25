package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	//c.HTML(200, "hello")
	c.JSON(200, gin.H{
		"message": "hello this is golang",
	})
}

type library struct {
	name    string
	bookNum int
	finance float64
}

func main() {
	l := new(library)
	l.bookNum = 200
	l.finance = 45000
	l.name = "华江图书馆"
	r := gin.Default() // return default route engine
	r.GET("/hello", sayHello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method":  "Get",
			"bookNum": l.bookNum,
		})
	})

	//start service
	r.Run()
}
