package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	name := "Akul Srivastava"

	string_html :=
`
<html>
	<head><title>Welcome</title></head>
	<body>
		<h1>Welcome ` + name + ` </h1> 
	</body>
</html>
`

	indexHtml, err := os.Create("index.html")

	if err!=nil {
		log.Fatal("Error creating index.html")
	}
	defer indexHtml.Close()

	io.Copy(indexHtml, strings.NewReader(string_html))

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":3000", nil)

}
