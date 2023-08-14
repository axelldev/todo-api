package main

import (
	"log"

	"github.com/axelldev/todo-api/app"
	"github.com/axelldev/todo-api/db"
	"github.com/axelldev/todo-api/handler"
	"github.com/axelldev/todo-api/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()
	database := db.NewDatabase()
	todoRepo := repository.NewTodoRepository(database)
	newApp := app.NewApp(":8080", r, todoRepo)
	newApp.BindRoutes(binder)
	log.Fatal(newApp.Run())
}

func binder(a *app.App, r *mux.Router) {
	r.HandleFunc("/api/v1/health", handler.HealthCheckHandler)
	r.HandleFunc("/api/v1/todos", handler.GetAllTodos(a)).Methods("GET")
	r.HandleFunc("/api/v1/todos", handler.CreateTodo(a)).Methods("POST")
	r.HandleFunc("/api/v1/todos/{id}", handler.GetTodoById(a)).Methods("GET")
	r.HandleFunc("/api/v1/todos/{id}", handler.DeleteTodo(a)).Methods("DELETE")
	r.HandleFunc("/api/v1/todos/{id}", handler.UpdateTodo(a)).Methods("PUT")
}
