package prompt

import (
	"fmt"
	"os"

	"github.com/arvpyrna/sled/internal/handler"
	"github.com/arvpyrna/sled/internal/repository"
)

type Prompt interface {
	ProcessUserInput()
}

type prompt struct {
	handler handler.Handler
	dao     repository.DAO
}

func NewPrompt(dao repository.DAO) Prompt {
	return &prompt{
		dao:     dao,
		handler: handler.NewHandler(dao)}
}

const OPTIONS_PATH = "res/options.txt"

func showOptions() {
	options, err := os.ReadFile(OPTIONS_PATH)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(options))
	fmt.Printf("\nChoose an option: ")
}

func (p *prompt) ProcessUserInput() {
	for {
		showOptions()
		var userInput int
		_, err := fmt.Scanf("%d", &userInput)
		if err != nil {
			fmt.Println("Error: Please enter a valid digit")
		}

		switch userInput {
		case 1:
		case 2:
			p.handler.HandleTaskCreation()
		case 3:
		case 4:
		case 5:
		case 6:
		case 0:
			p.handler.PerformCleanup()
			os.Exit(0)
		}
	}
}
