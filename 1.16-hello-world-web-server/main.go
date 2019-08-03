package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(res, "Hello, my name is Leo")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}
