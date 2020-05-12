package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	type User struct {
		Name string
		Email string
		Age int
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content Type", "text/html")
		user := User{
			Name: "Akul Srivastava",
			Email: "akulsr0@gmail.com",
			Age: 21,
		}
		err = tpl.Execute(w, user)
		if err != nil {
			log.Fatalln(err)
		}
	})
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
