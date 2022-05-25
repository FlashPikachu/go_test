package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板 parse template
	t, err := template.ParseFiles("web_test/gin_learn/template_1/resource/hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template error,err %v\n", err)
		return
	}
	//渲染模板 render template
	err = t.Execute(w, "just a test")
	if err != nil {
		fmt.Printf("Render template error,err %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err %v\n", err)
		return
	}
}
