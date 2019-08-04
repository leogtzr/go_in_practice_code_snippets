package main

import (
	"fmt"
	"net/http"
	"os"
)

func shutdown(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Bye .... ")
	os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintf(res, "The homepage.")
}

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
