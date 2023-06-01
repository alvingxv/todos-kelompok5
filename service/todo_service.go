package service

import (
	"github.com/alvingxv/todos-kelompok5/dto"
	"github.com/alvingxv/todos-kelompok5/entity"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository"
	"github.com/asaskevich/govalidator"
)

type todoService struct {
	todoRepository todo_repository.TodoRepository
}

type TodoService interface {
	GetAllTodos() (*[]dto.GetAllTodoResponse, errs.MessageErr)
	CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr)
	GetTodoById(id int) (*dto.GetTodoById, errs.MessageErr)
	UpdateTodo(payload dto.UpdateRequest, id int) (*dto.UpdateResponse, errs.MessageErr)
	DeleteTodo(id int) errs.MessageErr
}

func NewTodoService(todoRepository todo_repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (ts *todoService) GetAllTodos() (*[]dto.GetAllTodoResponse, errs.MessageErr) {

	todos, err := ts.todoRepository.GetAllTodos()

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllTodoResponse

	for _, todo := range todos {
		todosResponse := todo.TodoToTodoResponses()
		response = append(response, todosResponse)
	}

	return &response, nil

}

func (ts *todoService) CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	todos := &entity.Todo{
		Title: payload.Title,
	}

	err := ts.todoRepository.CreateTodo(todos)

	if err != nil {
		return nil, err
	}

	response := dto.CreateTodoResponse{
		Id:        todos.Id,
		Title:     todos.Title,
		Completed: todos.Completed,
		CreatedAt: todos.CreatedAt,
		UpdatedAt: todos.UpdatedAt,
	}

	return &response, nil

}

func (ts *todoService) GetTodoById(id int) (*dto.GetTodoById, errs.MessageErr) {

	todo := &entity.Todo{
		Id: id,
	}

	err := ts.todoRepository.GetTodoById(todo)

	if err != nil {
		return nil, err
	}

	response := dto.GetTodoById{
		Id:        todo.Id,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &response, nil

}

func (ts *todoService) UpdateTodo(payload dto.UpdateRequest, id int) (*dto.UpdateResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	todo := &entity.Todo{
		Id:    id,
		Title: payload.Title,
	}

	err := ts.todoRepository.UpdateTodo(todo)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateResponse{
		Id:        todo.Id,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &response, nil
}

func (ts *todoService) DeleteTodo(id int) errs.MessageErr {
	err := ts.todoRepository.DeleteTodo(id)

	if err != nil {
		return err
	}

	return nil
}

// func (cs *categoryService) GetCategory(userId uint) (*[]dto.GetCategoryResponse, errs.MessageErr) {
// 	categories, err := cs.categoryRepository.GetAllCategory(userId)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var responses []dto.GetCategoryResponse

// 	for _, category := range categories {

// 		var itemsResponses []dto.CategoryTask
// 		if len(category.Tasks) == 0 {
// 			itemsResponses = []dto.CategoryTask{}
// 		} else {
// 			for _, eachTask := range category.Tasks {
// 				itemResponse := eachTask.TaskToCategoryTaskResponse()
// 				itemsResponses = append(itemsResponses, itemResponse)
// 			}
// 		}

// 		response := dto.GetCategoryResponse{
// 			Id:        category.ID,
// 			Type:      category.Type,
// 			UpdatedAt: category.UpdatedAt,
// 			CreatedAt: category.CreatedAt,
// 			Tasks:     itemsResponses,
// 		}
// 		responses = append(responses, response)
// 	}

// 	return &responses, nil
// }

// func (cs *categoryService) CreateCategory(payload dto.CategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr) {

// 	_, errv := govalidator.ValidateStruct(payload)

// 	if errv != nil {
// 		return nil, errs.NewBadRequest(errv.Error())
// 	}

// 	category := &entity.Category{
// 		Type: payload.Type,
// 	}

// 	err := cs.categoryRepository.CreateCategory(category)

// 	if err != nil {
// 		return nil, err
// 	}

// 	response := dto.CreateCategoryResponse{
// 		Id:        category.ID,
// 		Type:      category.Type,
// 		CreatedAt: category.CreatedAt,
// 	}

// 	return &response, nil
// }

// func (cs *categoryService) UpdateCategory(payload dto.CategoryRequest, id uint) (*dto.UpdateCategoryResponse, errs.MessageErr) {

// 	_, errv := govalidator.ValidateStruct(payload)

// 	if errv != nil {
// 		return nil, errs.NewBadRequest(errv.Error())
// 	}

// 	category := &entity.Category{
// 		ID:   id,
// 		Type: payload.Type,
// 	}

// 	err := cs.categoryRepository.UpdateCategory(category)

// 	if err != nil {
// 		return nil, err
// 	}

// 	response := dto.UpdateCategoryResponse{
// 		Id:        category.ID,
// 		Type:      category.Type,
// 		UpdatedAt: category.UpdatedAt,
// 	}

// 	return &response, nil
// }

// func (cs *categoryService) DeleteCategory(id uint) errs.MessageErr {

// 	err := cs.categoryRepository.DeleteCategory(id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
