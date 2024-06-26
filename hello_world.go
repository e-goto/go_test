// Hello World with go as a simple dive into the language
package main

import (
	"fmt"
	"net/http"
)

func hello(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprint(responseWriter, "hello world")
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":9999", nil)
}
