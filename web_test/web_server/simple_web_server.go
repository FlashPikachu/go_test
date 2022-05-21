package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is go")
}

func main() {
	http.HandleFunc("/hello", myFunc)
	http.ListenAndServe(":9090", nil)
}
