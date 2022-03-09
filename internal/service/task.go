package service

type TaskService interface {
	CreateTask()
}

type typeService struct {
	// Get copy of repository
}

func NewTaskService() TaskService {
	return &typeService{}
}

func (t *typeService) CreateTask() {

}
