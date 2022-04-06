package main

import (
	"github.com/arvpyrna/sled/internal/db"
	"github.com/arvpyrna/sled/internal/prompt"
	"github.com/arvpyrna/sled/internal/repository"
)

func main() {
	db := db.Init()
	dao := repository.NewDao(db)
	prompt := prompt.NewPrompt(dao)
	prompt.ProcessUserInput()
}
