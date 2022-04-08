package db

import (
	"fmt"

	"github.com/arvryna/sled/internal/model"
	"github.com/arvryna/sled/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const SLED_DIR_PATH = "/home/arv/.sled"
const SLED_DB_PATH = "sled_db"

func Init() *gorm.DB {
	path := SLED_DIR_PATH

	if err := utils.CreateFolderIfNotExist(path); err != nil {
		panic(err)
	} else {
		fmt.Println("Initializing Sled files here: (path with DB) \n", SLED_DIR_PATH)
	}

	dbPath := SLED_DIR_PATH + "/" + SLED_DB_PATH

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Task{}, &model.Category{})
	return db
}
