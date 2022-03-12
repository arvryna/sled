package repository

import (
	"gorm.io/gorm"
)

type DAO interface {
	NewTaskQuery() TaskQuery
}

// DAO - Data/Database access object

var DB *gorm.DB

type dao struct{}

func NewDao(db *gorm.DB) DAO {
	DB = db
	return &dao{}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{}
}
