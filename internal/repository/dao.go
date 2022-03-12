package repository

import (
	"gorm.io/gorm"
)

type DAO interface {
	NewTaskQuery() TaskQuery
}

// DAO - Data/Database access object

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) DAO {
	return &dao{db: db}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{}
}
