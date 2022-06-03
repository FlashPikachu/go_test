package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func sayHello(c *gin.Context) {
	stockInfo := getStockInfo("http://suggest3.sinajs.cn/suggest/type=&key=60&name=suggestdata_1429775785401")
	//c.HTML(200, "hello")
	c.JSON(200, gin.H{
		"message": "hello this is golang",
		"stock":   stockInfo,
	})
}

type library struct {
	name    string
	bookNum int
	finance float64
}

func getStockInfo(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
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
