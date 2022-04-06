package db

import (
	"github.com/arvryna/sled/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("build/sled_db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Task{})
	return db
}
