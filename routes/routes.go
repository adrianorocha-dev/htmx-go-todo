package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/adrianorocha-dev/html-deez-nuts/model"
	"github.com/gorilla/mux"
)

type PageData struct {
	Todos []model.Todo
}

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()

	// fmt.Println("todos", todos)

	if err != nil {
		fmt.Println("Could not get todos")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.ExecuteTemplate(w, "Todos", PageData{Todos: todos})

	if err != nil {
		fmt.Println("Could not send todos", err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos()

	if err != nil {
		fmt.Println("Could not get todos")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, PageData{Todos: todos})

	if err != nil {
		fmt.Println("Could not send todos", err.Error())
	}
}

func toggleTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		fmt.Println("Could not parse id")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = model.MarkDone(id)

	if err != nil {
		fmt.Println("Could not toggle todo", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Could not parse form")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = model.CreateTodo(r.FormValue(("todo")))

	if err != nil {
		fmt.Println("Could not create todo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		fmt.Println("Could not parse id")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = model.Delete(id)

	if err != nil {
		fmt.Println("Could not delete todo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todos/{id}", toggleTodo).Methods("PUT")
	mux.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/todos", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
