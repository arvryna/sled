package service

import (
	"github.com/arvpyrna/sled/internal/repository"
	"github.com/transip/gotransip/v6/repository"
)

type TaskService interface {
	CreateTask()
	ListTask()
}

// Service holds the reference to DAO to access database
type taskService struct {
	dao *repository.DAO
}

func NewTaskService(dao *repository.DAO) TaskService {
	return &taskService{dao: dao}
}

func (t *taskService) CreateTask() {

}

func (t *taskService) ListTask() {

}
