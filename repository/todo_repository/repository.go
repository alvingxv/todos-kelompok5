package todo_repository

import (
	"github.com/alvingxv/todos-kelompok5/entity"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
)

type TodoRepository interface {
	GetAllTodos() ([]entity.Todo, errs.MessageErr)
	CreateTodo(todo *entity.Todo) errs.MessageErr
}
