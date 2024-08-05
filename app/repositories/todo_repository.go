package repositories

import (
	"gotodo/app/models"
	"gotodo/app/models/requests"
	"gotodo/app/repositories/connection"
	"log"
)

type ITodoRepository interface {
	GetAll() (*[]models.Todo, error)
	Save(request requests.NewToDoRequest) (int64, error)
}

type TodoRepository struct {
	connector *connection.DbConnector
}

func NewToDoRepository(connector *connection.DbConnector) *TodoRepository {
	return &TodoRepository{connector: connector}
}

func (repo *TodoRepository) GetAll() (*[]models.Todo, error) {
	connect, err := repo.connector.DbConnect()
	if err != nil {
		return nil, err
	}

	defer connect.Close()
	rows, err := connect.Query("SELECT Id, Title FROM ToDo ORDER BY id DESC")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	todos := make([]models.Todo, 0)
	for rows.Next() {
		var (
			id    int64
			title string
		)

		if err := rows.Scan(&id, &title); err != nil {
			log.Println(err)
			return nil, err
		}

		todos = append(todos, models.Todo{Id: id, Title: title})
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &todos, nil
}

func (repo *TodoRepository) Save(request requests.NewToDoRequest) (int64, error) {
	connect, err := repo.connector.DbConnect()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	defer connect.Close()
	query := connect.QueryRow("INSERT INTO ToDo (title, body)"+
		"VALUES ($1, $2) RETURNING id", request.Title, request.Body)

	if query.Err() != nil {
		log.Println(err)
		return 0, err
	}

	var id int64
	if err = query.Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}
