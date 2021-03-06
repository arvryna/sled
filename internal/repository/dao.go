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
	"fmt"

	"gorm.io/gorm"
)

// Interface that holds functions that returns object that implements those interfaces
type DAO interface {
	NewTaskQuery() TaskQuery
	NewCategoryQuery() CategoryQuery
	CloseDBConnection()
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

func (d *dao) NewCategoryQuery() CategoryQuery {
	return &categoryQuery{}
}

func (d *dao) CloseDBConnection() {
	db, err := DB.DB()
	if err == nil {
		fmt.Println("DB Connection is closed..")
	}
	db.Close()
}
