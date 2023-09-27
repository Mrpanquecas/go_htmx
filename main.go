package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templ *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	templ = template.Must(template.ParseFiles("static/templates/index.html"))
	templ.Execute(w, nil)
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "install go", Done: true},
			{Item: "learn go", Done: false},
			{Item: "learn more go", Done: false},
		},
	}
	fmt.Print(data)
}

func main() {

	http.HandleFunc("/", todo)
	log.Fatal(http.ListenAndServe(":80", nil))
}
