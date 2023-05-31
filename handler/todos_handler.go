package handler

import "github.com/alvingxv/todos-kelompok5/service"

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) todoHandler {
	return todoHandler{
		todoService: todoService,
	}
}
