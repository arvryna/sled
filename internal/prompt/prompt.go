package prompt

import (
	"fmt"
	"os"

	"github.com/arvryna/sled/internal/handler"
	"github.com/arvryna/sled/internal/repository"
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

func showOptions() {
	fmt.Println("\n" + getOptions())
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
		case 9:
			p.handler.ResumeTask()
		case 0:
			p.handler.PerformCleanup()
			os.Exit(0)
		}
	}
}

func getOptions() string {
	return `Welcome to Sled! Please choose an option:
1) Track task
2) List tasks
3) Create category
4) Export
5) Send timesheet by Email
6) Manual
7) Stop timer
8) Pause timer
9) Resume task
0) Exit `
}
