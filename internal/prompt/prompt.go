package prompt

import (
	"fmt"
	"os"

	"github.com/arvryna/sled/internal/handler"
	"github.com/arvryna/sled/internal/repository"
)

const OPTIONS_PATH = "res/options.txt"
const USER_MANUAL_PATH = "res/manual.txt"

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

func getContentsOfFile(filePath string) string {
	options, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(options)
}

func showOptions() {
	fmt.Println("\n" + getContentsOfFile(OPTIONS_PATH))
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
			p.handler.HandleTaskCreation()
		case 2:
			p.handler.PrintTasks()
		case 3:
			p.handler.HandleCategoryCreation()
		case 4:
		case 5:
		case 6:
			fmt.Println(getContentsOfFile(USER_MANUAL_PATH))
		case 9:
			p.handler.ResumeTask()
		case 0:
			p.handler.PerformCleanup()
			os.Exit(0)
		}
	}
}
