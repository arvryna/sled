package db

import (
	"github.com/arvpyrna/sled/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// TODO: move this to config
	db, err := gorm.Open(sqlite.Open("build/sled_db"), &gorm.Config{})

	// open connection:
	if err != nil {
		panic(err)
	}

	// WARN: remove this in prod
	db.AutoMigrate(&model.Task{})
	return db
}
