package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//web, _ := ioutil.ReadFile("./resource/hello.txt")
	web, _ := ioutil.ReadFile("web_test/gin_learn/gin_demo_1/resource/hello.txt")
	println(string(web))
	println(web)
	_, _ = fmt.Fprintln(w, string(web))
}

func getStockInfo(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	getStockInfo("http://suggest3.sinajs.cn/suggest/type=&key=60&name=suggestdata_1429775785401")
	//http.HandleFunc("/hello", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Printf("http server failed, err :%v\n", err)
	//	return
	//}
}
