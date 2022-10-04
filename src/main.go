package main

import (
	"fmt"
	"net/http"
)

func handler(web http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(web, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
