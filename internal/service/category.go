package service

import (
	"github.com/arvryna/sled/internal/model"
	"github.com/arvryna/sled/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() []model.Category
}

type categoryservice struct {
	dao repository.DAO
}

func NewCategoryService(dao repository.DAO) CategoryService {
	return &categoryservice{dao: dao}
}

func (c *categoryservice) CreateCategory(category *model.Category) error {
	if err := c.dao.NewCategoryQuery().CreateCategory(category); err != nil {
		return err
	}
	return nil
}

func (c *categoryservice) GetAllCategories() []model.Category {
	return c.dao.NewCategoryQuery().GetAllCategories()
}
