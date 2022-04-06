/* DAO - Data/Database access object
 * We don't access DB directly, it will be used through DAO
 * It is important to hide this detail and use DAO to manage interaction
 * other classes to keep the DB loosely coupled
 * We have used SQlite in this project, but can be safely moved to any database
 * by just changing the Init function of the db.go file
 * DAO does not know about Service layer
 */

package repository

import (
	"gorm.io/gorm"
)

type DAO interface {
	NewTaskQuery() TaskQuery
}

var DB *gorm.DB

type dao struct{}

func NewDao(db *gorm.DB) DAO {
	DB = db
	return &dao{}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{}
}
