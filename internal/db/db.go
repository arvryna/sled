package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Just need DB url to init the database and also perform necessary
// migrations
func Init() *gorm.DB {
	// TODO: move this to config
	db, err := gorm.Open(sqlite.Open("build/sled_db"), &gorm.Config{})

	// open connection:
	if err != nil {
		panic(err)
	}

	// WARN: remove this in prod
	// db.AutoMigrate(&models.Contact{})
	return db
}
