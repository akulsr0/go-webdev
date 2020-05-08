package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":3001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>This is the index page.</h1>
		<a href="/about">About</a> <br>
		<a href="/contact">Contact</a> 
	`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>This is the about page.</h1>
		<a href="/">Index</a> <br>
		<a href="/contact">Contact</a> 
	`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>This is the contact page.</h1>
		<a href="/">Index</a> <br>
		<a href="/about">About</a> 
	`)
}
