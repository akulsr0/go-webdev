package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello/", slugCheck)
	http.ListenAndServe(":3001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>Hey there!</h1>
		<code>Goto /hello/yourname </code>
	`)
}

func slugCheck(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(r.URL.Path, "/hello/", "", 1)
	fmt.Fprintf(w, "Hello %s", name)
}
