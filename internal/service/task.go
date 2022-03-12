package service

import (
	"github.com/arvpyrna/sled/internal/model"
	"github.com/arvpyrna/sled/internal/repository"
)

type TaskService interface {
	CreateTask(task model.Task)
	ListTask()
}

// Service holds the reference to DAO to access database
type taskService struct {
	dao repository.DAO
}

func NewTaskService(dao repository.DAO) TaskService {
	return &taskService{dao: dao}
}

func (t *taskService) CreateTask(task model.Task) {
	t.dao.NewTaskQuery().CreateTask(task)
}

func (t *taskService) ListTask() {

}
