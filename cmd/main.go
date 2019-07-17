package main

import (
	"golang-todo-list/app/handlers"
	"golang-todo-list/app/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repositories.Initialize("postgres://dbuser:sNChv71OhNvILU@localhost/teachingstrategies-sample?sslmode=disable")

	router := mux.NewRouter()

	subrouter := router.PathPrefix("/v1/api").Subrouter()

	// todoListsHandler := new(handlers.TodoListsHandler)
	todoListHandler := new(handlers.TodoListHandler)
	// todoItemHander := new(handlers.TodoItemHandler)

	subrouter.HandleFunc("/lists", todoListHandler.Get).Methods("GET")

	// subrouter.HandleFunc("/lists", todoListsHandler.Post).Methods("POST")
	// subrouter.HandleFunc("/lists", todoListsHandler.Patch).Methods("PATCH")
	// subrouter.HandleFunc("/lists/{listId}", todoListsHandler.Delete).Methods("DELETE")

	// subrouter.HandleFunc("/lists/{listId}/todos", todoListHandler.Get).Methods("GET")
	// subrouter.HandleFunc("/lists/{listId}/todos", todoListHandler.Post).Methods("POST")

	// subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Get).Methods("GET")
	// subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Patch).Methods("PATCH")
	// subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Delete).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
