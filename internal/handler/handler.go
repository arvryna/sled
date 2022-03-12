package handler

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/arvpyrna/sled/internal/repository"
	"github.com/schollz/progressbar/v3"
)

type Handler interface {
	HandleTaskCreation()
	PerformCleanup()
}

type handler struct{ dao *repository.DAO }

func NewHandler(dao *repository.DAO) Handler {
	return &handler{dao: dao}
}

func (h *handler) HandleTaskCreation() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter task Description: ")
	desc, _ := reader.ReadString('\n')

	fmt.Print("Enter task Duration (in minutes): ")
	var duration int
	fmt.Scanf("%d", &duration)

	// get input for task
	startTimer(duration)
	// consider the task completed...
	// create task after timer is done...
	fmt.Println("Task is marked completed", title, desc)
}

func (h *handler) PerformCleanup() {
	fmt.Println("Closing app...")
}

// Timer with interactive progress bar
// Gets input in minute but shows progress bar in the unit of seconds
func startTimer(minutes int) {
	durationInSec := minutes * 60
	fmt.Println(fmt.Sprintf("Timer starting for %d minute(s)", minutes))
	bar := progressbar.Default(int64(durationInSec))
	for i := 0; i < durationInSec; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}
