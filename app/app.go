package app

import (
	"fmt"
	"net/http"

	"github.com/axelldev/todo-api/repository"
	"github.com/gorilla/mux"
)

type routerBinder func(a *App, r *mux.Router)

// App is the main struct of the application
type App struct {
	Addr           string                     // Addr is the address of the server
	TodoRepository *repository.TodoRepository // TodoRepository is the repository for the todo entity
	Router         *mux.Router                // Router is the router of the application
}

// NewApp creates a new App instance
func NewApp(addr string, Router *mux.Router, todoRepository *repository.TodoRepository) *App {
	return &App{
		Addr:           addr,
		TodoRepository: todoRepository,
		Router:         Router,
	}
}

// Run runs the server
func (a *App) Run() error {
	fmt.Println("Server running on port", a.Addr)
	return http.ListenAndServe(a.Addr, a.Router)
}

// BindRoutes binds the routes to the router
// The binder is a function that receives the App and the Router
func (a *App) BindRoutes(binder routerBinder) {
	binder(a, a.Router)
}
