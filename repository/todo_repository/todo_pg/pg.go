package todo_pg

import (
	"errors"

	"github.com/alvingxv/todos-kelompok5/entity"
	"github.com/alvingxv/todos-kelompok5/pkg/errs"
	"github.com/alvingxv/todos-kelompok5/repository/todo_repository"
	"gorm.io/gorm"
)

type todoPG struct {
	db *gorm.DB
}

func NewTodoPG(db *gorm.DB) todo_repository.TodoRepository {
	return &todoPG{
		db: db,
	}
}

func (t *todoPG) GetAllTodos() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo
	result := t.db.Find(&todos).Error

	if result != nil {
		return nil, errs.NewInternalServerError("something Went Wrong")
	}

	return todos, nil
}

func (t *todoPG) CreateTodo(todo *entity.Todo) errs.MessageErr {

	err := t.db.Create(&todo).Error

	if err != nil {
		return errs.NewInternalServerError("something Went Wrong")
	}

	return nil
}

func (t *todoPG) GetTodoById(todo *entity.Todo) errs.MessageErr {
	err := t.db.First(&todo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("Not found")
		}
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *todoPG) UpdateTodo(todo *entity.Todo) errs.MessageErr {
	title := todo.Title

	result := t.db.First(&todo, todo.Id)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = t.db.Model(&todo).Update("title", title)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *todoPG) DeleteTodo(id int) errs.MessageErr {
	result := t.db.First(&entity.Todo{}, id)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = t.db.Delete(&entity.Todo{}, id)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

// func (c *categoryPG) GetAllCategory(userId uint) ([]entity.Category, errs.MessageErr) {
// 	var categories []entity.Category

// 	result := c.db.Model(&entity.Category{}).Preload("Tasks", "user_id = ?", userId).Find(&categories).Error

// 	if result != nil {
// 		return nil, errs.NewInternalServerError("something Went Wrong")
// 	}

// 	return categories, nil
// }

// func (c *categoryPG) CreateCategory(category *entity.Category) errs.MessageErr {

// 	err := c.db.Create(&category).Error

// 	if err != nil {
// 		return errs.NewInternalServerError("something Went Wrong")
// 	}

// 	return nil

// }

// func (c *categoryPG) UpdateCategory(category *entity.Category) errs.MessageErr {
// 	result := c.db.Select("id").First(&category, category.ID)
// 	if result.Error != nil {
// 		return errs.NewNotFoundError("not found")
// 	}

// 	result = c.db.Model(&category).Update("type", category.Type)

// 	if result.Error != nil {
// 		return errs.NewInternalServerError("Internal Server Error")
// 	}

// 	return nil
// }

// func (c *categoryPG) DeleteCategory(id uint) errs.MessageErr {
// 	result := c.db.Select("id").First(&entity.Category{}, id)
// 	if result.Error != nil {
// 		return errs.NewNotFoundError("not found")
// 	}

// 	result = c.db.Delete(&entity.Category{}, id)

// 	if result.Error != nil {
// 		return errs.NewInternalServerError("Internal Server Error")
// 	}

// 	return nil
// }

// func (c *categoryPG) GetCategoryById(id uint) errs.MessageErr {
// 	err := c.db.Debug().First(&entity.Category{}, id).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return errs.NewNotFoundError("Category didn't exist")
// 		}
// 		return errs.NewInternalServerError("Internal Server Error")
// 	}

// 	return nil
// }
