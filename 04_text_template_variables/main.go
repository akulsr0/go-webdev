package main

import (
	"html/template"
	"net/http"
)

func main() {
	tpl := template.Must(template.ParseFiles("index.gohtml"))

	type Project struct {
		Name string
		Language string
	}

	type User struct {
		Id int
		Name string
		Username string
		Projects []Project
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Project1 := Project{Name: "Calcy", Language: "Golang"}
		Project2 := Project{Name: "Shitter", Language: "Javascript"}
		Project3 := Project{Name: "botbyakul", Language: "Python"}
		user := User{
			Id: 1,
			Name: "Akul Srivastava",
			Username: "akulsr0",
			Projects: []Project{Project1, Project2, Project3},
		}
		tpl.Execute(w, user)
	})
	http.ListenAndServe(":3000", nil)
}
