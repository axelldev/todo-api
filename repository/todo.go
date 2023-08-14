package repository

import (
	"database/sql"
	"github.com/axelldev/todo-api/model"
)

// TodoRepository is a repository for todos.
type TodoRepository struct {
	db *sql.DB
}

// NewTodoRepository creates a new TodoRepository
// with the given database connection.
func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

type CreateTodoParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

const createTodoQuery = `
INSERT INTO todos (title, description)
VALUES ($1, $2);
`

// Save saves a todo in the database.
func (r *TodoRepository) Save(todo CreateTodoParams) error {
	_, err := r.db.Exec(createTodoQuery, todo.Title, todo.Description)
	return err
}

const getAllTodosQuery = `
SELECT * FROM todos;
`

// FindAll returns all todos from the database.
func (r *TodoRepository) FindAll() ([]model.Todo, error) {
	rows, err := r.db.Query(getAllTodosQuery)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()
	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

const findTodoByIdQuery = `
SELECT id, title, description, created_at, completed FROM todos
WHERE id = $1;
`

// FindById returns a todo with the given id.
func (r *TodoRepository) FindById(id int) (*model.Todo, error) {
	row := r.db.QueryRow(findTodoByIdQuery, id)
	var todo model.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.Completed)
	return &todo, err
}

const deleteTodoQuery = `
DELETE FROM todos WHERE id = $1;
`

// Delete deletes a todo with the given id.
func (r *TodoRepository) Delete(id int) error {
	_, err := r.db.Exec(deleteTodoQuery, id)
	return err
}

const updateTodoQuery = `
UPDATE todos
SET title = $2, description = $3, completed = $4
WHERE id = $1;
`

// Update updates a todo with the given id.
func (r *TodoRepository) Update(todo model.Todo) error {
	_, err := r.db.Exec(updateTodoQuery, todo.ID, todo.Title, todo.Description, todo.Completed)
	return err
}

const completeTodoQuery = `
UPDATE todos 
SET completed = true 
WHERE id = $1;
`

// CompleteTodo completes a todo with the given id.
func (r *TodoRepository) CompleteTodo(id int) error {
	_, err := r.db.Exec(completeTodoQuery, id)
	return err
}
