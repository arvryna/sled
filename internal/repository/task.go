package repository

import (
	"fmt"

	"github.com/arvryna/sled/internal/model"
)

type TaskQuery interface {
	CreateTask(task *model.Task)
	ListTasks() []model.Task
	GetTask(taskId int) model.Task
	UpdateTask(task *model.Task)
}

type taskQuery struct{}

func (t *taskQuery) CreateTask(task *model.Task) {
	if res := DB.Create(&task); res.Error != nil {
		fmt.Println("Can't create task", res.Error)
	} else {
		fmt.Println("Task Created Successfully!")
	}
}

func (t *taskQuery) GetTask(taskId int) model.Task {
	task := model.Task{}
	if res := DB.First(&task, taskId); res.Error != nil {
		fmt.Println("Can't fetch task with Id", taskId)
	}
	return task
}

func (t *taskQuery) ListTasks() []model.Task {
	var tasks []model.Task
	DB.Find(&tasks)
	return tasks
}

func (t *taskQuery) UpdateTask(task *model.Task) {
	DB.Save(task)
}
