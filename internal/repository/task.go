package repository

import "fmt"

type TaskQuery interface {
	CreateTask()
}

type taskQuery struct{}

func (t *taskQuery) CreateTask() {
	// create a task in DB
	fmt.Println("Task created")
}
