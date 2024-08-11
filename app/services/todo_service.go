package services

import (
	"context"
	"gotodo/app/models/dtos"
	"gotodo/app/models/requests"
	"gotodo/app/repositories"
	"gotodo/app/repositories/transaction"
)

type IToDoService interface {
	SaveToDo(ctx context.Context, request requests.NewToDoRequest) (int64, error)
	GetAll(ctx context.Context) (*[]*dtos.TodoDto, error)
}

type ToDoService struct {
	todoRepository     repositories.ITodoRepository
	transactionManager *transaction.Manager
}

func NewToDoService(repository repositories.ITodoRepository, transactionManager *transaction.Manager) *ToDoService {
	return &ToDoService{
		todoRepository:     repository,
		transactionManager: transactionManager,
	}
}

func (service *ToDoService) SaveToDo(ctx context.Context, request requests.NewToDoRequest) (int64, error) {
	tx, err := service.transactionManager.BeginReadCommited(ctx)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()
	id, err := service.todoRepository.Save(tx, request)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *ToDoService) GetAll(ctx context.Context) (*[]*dtos.TodoDto, error) {
	todos, err := service.todoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
