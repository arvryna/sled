package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/arvryna/golib/pkg/utils"
	"github.com/arvryna/sled/internal/model"
	"github.com/arvryna/sled/internal/repository"
	"github.com/arvryna/sled/internal/service"
	"github.com/schollz/progressbar/v3"
)

type Handler interface {
	HandleTaskCreation()
	HandleCategoryCreation()
	PrintTasks()
	PerformCleanup()
	ResumeTask()
}

type handler struct {
	dao repository.DAO
}

func NewHandler(dao repository.DAO) Handler {
	return &handler{dao: dao}
}

func (h *handler) HandleTaskCreation() {

	title := utils.GetUserInput("Enter Task title: ")
	desc := utils.GetUserInput("Enter Task description: ")
	categoryId, _ := strconv.Atoi(utils.GetUserInput("Enter Category ID: (For now enter 1 here): "))
	duration, _ := strconv.Atoi(utils.GetUserInput("Enter task Duration (in minutes): "))

	startTimer(duration)

	// consider the task completed and create task after timer is done...
	taskService := service.NewTaskService(h.dao)

	task := model.Task{
		Title:       title,
		Description: desc,
		Duration:    duration,
		CategoryId:  categoryId,
	}

	taskService.CreateTask(&task)
}

func (h *handler) HandleCategoryCreation() {
	title := utils.GetUserInput("Enter Category name: ")
	desc := utils.GetUserInput("Enter Category description: ")

	category := model.Category{
		Title:       title,
		Description: desc,
	}

	// consider the task completed and create task after timer is done...
	categoryService := service.NewCategoryService(h.dao)
	categoryService.CreateCategory(&category)
}

func (h *handler) PrintTasks() {
	results := service.NewTaskService(h.dao).ListTask()
	if len(results) > 0 {
		fmt.Println("\n----- Task List -----")
	} else {
		return
	}
	for _, task := range results {
		fmt.Print(fmt.Sprintf("\n%d-%s", task.Id, task.Title))
	}
	fmt.Println("\n----- End -----")
}

func (h *handler) ResumeTask() {
	taskId, _ := strconv.Atoi(utils.GetUserInput("Please enter taskID: "))
	taskService := service.NewTaskService(h.dao)
	task := taskService.GetTask(taskId)
	duration, _ := strconv.Atoi(utils.GetUserInput("How long you want to work more on this task ? (in minutes)"))
	startTimer(duration)
	task.Duration = task.Duration + duration
	taskService.UpdateTask(&task)
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
