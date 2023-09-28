package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
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

func show_todo(w http.ResponseWriter, r *http.Request) {
	templ = template.Must(template.ParseFiles("static/templates/index.html"))
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "install go", Done: true},
			{Item: "learn go", Done: false},
			{Item: "learn more go", Done: false},
		},
	}
	templ.Execute(w, data)
}

func add_todo(w http.ResponseWriter, r *http.Request) {
	// To emulate latency
	time.Sleep(1 * time.Second)
	todo_status, _ := strconv.ParseBool(r.PostFormValue("todo-done"))
	todo_title := r.PostFormValue("todo-item")
	templ = template.Must(template.ParseFiles("static/templates/index.html"))
	templ.ExecuteTemplate(w, "todo-list-item", Todo{Item: todo_title, Done: todo_status})
}
func main() {

	http.HandleFunc("/add-todo/", add_todo)
	http.HandleFunc("/", show_todo)
	log.Fatal(http.ListenAndServe(":80", nil))
}
