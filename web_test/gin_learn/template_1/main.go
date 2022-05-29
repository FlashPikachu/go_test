package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string
	Age    int32
}

var htmlPath1 = "web_test/gin_learn/template_1/resource/hello.tmpl"

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板 parse template
	t, err := template.ParseFiles("web_test/gin_learn/template_1/resource/hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template error,err %v\n", err)
		return
	}
	//渲染模板 render template
	u1 := User{
		Name:   "法莎莉亚",
		Gender: "女",
		Age:    100,
	}
	m1 := map[string]interface{}{
		"name":   "pharsalia",
		"gender": "女",
		"age":    18,
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Printf("Render template error,err %v\n", err)
		return
	}
}

func c1(w http.ResponseWriter, r *http.Request) {
	privateProcess := func(name string) string {
		return "Cool guy!" + name
	}

	t := template.New("hello.tmpl").Funcs(template.FuncMap{
		"admire": privateProcess,
	})
	_, err := t.ParseFiles(htmlPath1)
	if err != nil {
		fmt.Printf("Render template error,err %v\n", err)
		return
	}
	t.Execute(w, "pharsalia")

}

func main() {
	//http.HandleFunc("/", sayHello)
	http.HandleFunc("/customize", c1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err %v\n", err)
		return
	}
}
