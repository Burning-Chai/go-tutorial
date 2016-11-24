package main

import (
	"net/http"
	"text/template"
	"time"
)

type Page struct {
	Time   string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{time.Now().Format(time.RFC3339), 1}

	tmpl, err := template.ParseFiles("./templates/sample.html")
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(w, page); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":8080", nil)
}

