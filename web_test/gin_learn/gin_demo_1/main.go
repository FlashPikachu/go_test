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

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err :%v\n", err)
		return
	}
}
