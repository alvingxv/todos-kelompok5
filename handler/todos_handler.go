package handler

import (
	"net/http"

	"github.com/alvingxv/todos-kelompok5/service"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) todoHandler {
	return todoHandler{
		todoService: todoService,
	}
}

func (th *todoHandler) GetAllTodos(ctx *gin.Context) {

	result, err := th.todoService.GetAllTodos()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
