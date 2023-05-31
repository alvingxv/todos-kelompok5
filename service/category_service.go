package service

import (
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository"
)

type todoService struct {
	todoRepository todo_repository.TodoRepository
}

type TodoService interface {
}

func NewTodoService(todoRepository todo_repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
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
