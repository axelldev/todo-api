package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/axelldev/todo-api/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/axelldev/todo-api/app"
	"github.com/axelldev/todo-api/response"
)

// GetAllTodos is a handler function that returns all todos
func GetAllTodos(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		todos, err := a.TodoRepository.FindAll()

		if err != nil {
			log.Println(err.Error())
			response.JSON(w, http.StatusInternalServerError, err)
			return
		}

		response.JSON(w, http.StatusOK, todos)
	}
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateTodo is a handler function that creates a new todo
func CreateTodo(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateTodoRequest
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&req); err != nil {
			response.JSON(w, http.StatusBadRequest, err)
			return
		}

		if req.Title == "" {
			response.JSON(w, http.StatusBadRequest, "Title is required")
			return
		}

		params := repository.CreateTodoParams{
			Title:       req.Title,
			Description: req.Description,
		}

		if err := a.TodoRepository.Save(params); err != nil {
			response.JSON(w, http.StatusInternalServerError, err)
			return
		}

		response.JSON(w, http.StatusCreated, nil)
	}
}

// GetTodoById is a handler function that returns a todo by id
func GetTodoById(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		id, err := strconv.Atoi(idString)

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Invalid id")
			return
		}

		todo, err := a.TodoRepository.FindById(id)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "Todo not found")
				return
			}
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, *todo)
	}
}

// DeleteTodo is a handler function that deletes a todo by id
func DeleteTodo(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString := vars["id"]
		id, err := strconv.Atoi(idString)

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Invalid id")
			return
		}

		todo, err := a.TodoRepository.FindById(id)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "Todo not found")
				return
			}
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := a.TodoRepository.Delete(todo.ID); err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

type UpdateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func UpdateTodo(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idString, ok := vars["id"]

		if !ok {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		todo, err := a.TodoRepository.FindById(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				response.JSON(w, http.StatusNotFound, "Todo not found")
				return
			}
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var req UpdateTodoRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			response.JSON(w, http.StatusBadRequest, err)
			return
		}

		if req.Title != nil {
			if *req.Title == "" {
				response.JSON(w, http.StatusBadRequest, "Title is required")
				return
			}
			todo.Title = *req.Title
		}

		if req.Description != nil {
			todo.Description = *req.Description
		}

		if req.Completed != nil {
			todo.Completed = *req.Completed
		}

		if err := a.TodoRepository.Update(*todo); err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
