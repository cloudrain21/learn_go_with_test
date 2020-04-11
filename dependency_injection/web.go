package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func myOwnGreeHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Chris")
}

func main() {
	http.ListenAndServe(":9000", http.HandlerFunc(myOwnGreeHandler))
}
