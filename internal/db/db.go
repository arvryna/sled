package db

import (
	"log"

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
	} else {
		// Use a better logging tool, eg: Logrus
		log.Println("DB is initialized!")
	}

	// WARN: remove this in prod
	// db.AutoMigrate(&models.Contact{})
	return db
}
