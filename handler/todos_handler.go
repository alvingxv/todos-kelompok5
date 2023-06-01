package handler

import (
	"net/http"

	"github.com/alvingxv/todos-kelompok5/dto"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
	"github.com/alvingxv/todos-kelompok5/pkg/helpers"
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

func (th *todoHandler) CreateTodo(ctx *gin.Context) {

	var todoRequest dto.CreateTodoRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := th.todoService.CreateTodo(todoRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

func (th *todoHandler) GetTodoById(ctx *gin.Context) {

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	result, err := th.todoService.GetTodoById(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
