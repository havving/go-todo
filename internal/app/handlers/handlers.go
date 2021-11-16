package handlers

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/app/handlers/todo"
)

type APIHandlerBlock struct{}

// GetTodoListHandler GET /api/todos
func (h *APIHandlerBlock) GetTodoListHandler(ctx echo.Context) error {
	return todo.GetTodoList(ctx)
}

// PostTodoHandler POST /api/todos
func (h *APIHandlerBlock) PostTodoHandler(ctx echo.Context) error {
	return todo.PostTodo(ctx)
}

// RemoveTodoHandler DELETE /api/todos/:id
func (h *APIHandlerBlock) RemoveTodoHandler(ctx echo.Context) error {
	return todo.RemoveTodo(ctx)
}

// UpdateTodoHandler PUT /api/todos/:id
func (h *APIHandlerBlock) UpdateTodoHandler(ctx echo.Context) error {
	return todo.UpdateTodo(ctx)
}
