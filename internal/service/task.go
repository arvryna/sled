/*
   Service Layer
   - Contains access to DAO
   - Businesss logic is written here
*/

package service

import (
	"github.com/arvpyrna/sled/internal/model"
	"github.com/arvpyrna/sled/internal/repository"
)

type TaskService interface {
	CreateTask(task model.Task)
	ListTask() []model.Task
	GetTask(taskId int) model.Task
	UpdateTask(task model.Task)
}

type taskService struct {
	dao repository.DAO
}

func NewTaskService(dao repository.DAO) TaskService {
	return &taskService{dao: dao}
}

func (t *taskService) CreateTask(task model.Task) {
	t.dao.NewTaskQuery().CreateTask(task)
}

func (t *taskService) ListTask() []model.Task {
	return t.dao.NewTaskQuery().ListTasks()
}

func (t *taskService) GetTask(taskId int) model.Task {
	return t.dao.NewTaskQuery().GetTask(taskId)
}

func (t *taskService) UpdateTask(task model.Task) {
	t.dao.NewTaskQuery().UpdateTask(task)
}
