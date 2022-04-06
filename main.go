package main

import (
	"github.com/arvryna/sled/internal/db"
	"github.com/arvryna/sled/internal/prompt"
	"github.com/arvryna/sled/internal/repository"
)

func main() {
	db := db.Init()
	dao := repository.NewDao(db)
	prompt := prompt.NewPrompt(dao)
	prompt.ProcessUserInput()
}
