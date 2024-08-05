package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DbConnector struct {
	connectionString string
}

func NewDbConnector(connectionString string) *DbConnector {
	return &DbConnector{connectionString: connectionString}
}

func (connection *DbConnector) DbConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", connection.connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
