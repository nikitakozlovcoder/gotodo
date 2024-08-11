package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DbConnector struct {
	connectionString string
}

func NewDbConnector(connectionString string) *DbConnector {
	return &DbConnector{connectionString: connectionString}
}

func (connector *DbConnector) DbConnect() (*DbConnection, error) {
	db, err := sql.Open("postgres", connector.connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DbConnection{DB: db}, nil
}
