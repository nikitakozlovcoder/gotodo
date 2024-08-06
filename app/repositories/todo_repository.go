package repositories

import (
	"gotodo/app/models/dtos"
	"gotodo/app/models/requests"
	"gotodo/app/repositories/connection"
	"log"
)

type ITodoRepository interface {
	GetAll() (*[]dtos.TodoDto, error)
	Save(request requests.NewToDoRequest) (int64, error)
}

type TodoRepository struct {
	connector *connection.DbConnector
}

func NewToDoRepository(connector *connection.DbConnector) *TodoRepository {
	return &TodoRepository{connector: connector}
}

func (repo *TodoRepository) GetAll() (*[]*dtos.TodoDto, error) {
	connect, err := repo.connector.DbConnect()
	if err != nil {
		return nil, err
	}

	defer connect.Close()

	rows, err := connect.Query(`SELECT td.id todo_id, td.title todo_title, tg.id tag_id, tg.name tag_name FROM todo td
    LEFT JOIN todo_tag tt ON tt.tag_id = td.id
    LEFT JOIN tag tg ON tg.id = tt.tag_id ORDER BY td.id DESC`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	todos := make([]*dtos.TodoDto, 0)
	todosMap := make(map[int64]*dtos.TodoDto)
	tagsMap := make(map[int64]*dtos.TagDto)

	for rows.Next() {
		var (
			todoId    int64
			todoTitle string
			tagId     int64
			tagName   string
		)

		if err := rows.Scan(&todoId, &todoTitle, &tagId, &tagName); err != nil {
			log.Println(err)
			return nil, err
		}

		todo := dtos.TodoDto{Id: todoId, Title: todoTitle}
		todoFromMap, todoExists := todosMap[todoId]
		if !todoExists {
			todosMap[todoId] = &todo
		}

		if tagId != 0 {
			tag := dtos.TagDto{Id: tagId, Name: tagName}
			tagFromMap, tagExists := tagsMap[tagId]
			if !tagExists {
				tagsMap[tagId] = &tag
			}

			todoFromMap.Tags = append(todoFromMap.Tags, tagFromMap)
		}

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
