package repository

import (
	"fmt"

	"github.com/arvryna/sled/internal/model"
)

type CategoryQuery interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() []model.Category
}

type categoryQuery struct{}

func (c *categoryQuery) CreateCategory(category *model.Category) error {
	if res := DB.Create(&category); res.Error != nil {
		fmt.Println("Can't create task", res.Error)
		return res.Error
	} else {
		fmt.Println("Task Created Successfully!")
	}
	return nil
}

func (c *categoryQuery) GetAllCategories() []model.Category {
	var results []model.Category
	DB.Order("created_at desc").Find(&results)
	return results
}
