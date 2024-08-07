package services

import (
	"gotodo/app/models/dtos"
	"gotodo/app/models/requests"
	"gotodo/app/repositories"
)

type IToDoService interface {
	SaveToDo(request requests.NewToDoRequest) (int64, error)
	GetAll() (*[]*dtos.TodoDto, error)
}

type ToDoService struct {
	todoRepository repositories.ITodoRepository
}

func NewToDoService(repository repositories.ITodoRepository) *ToDoService {
	return &ToDoService{
		todoRepository: repository,
	}
}

func (service *ToDoService) SaveToDo(request requests.NewToDoRequest) (int64, error) {
	id, err := service.todoRepository.Save(request)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *ToDoService) GetAll() (*[]*dtos.TodoDto, error) {
	todos, err := service.todoRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
