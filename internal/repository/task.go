package repository

import (
	"fmt"

	"github.com/arvpyrna/sled/internal/model"
)

type TaskQuery interface {
	CreateTask(task model.Task)
}

type taskQuery struct{}

func (t *taskQuery) CreateTask(task model.Task) {
	if res := DB.Create(&task); res.Error != nil {
		fmt.Println("Can't create task", res.Error)
	} else {
		fmt.Println("Task Created Successfully!")
	}
}
