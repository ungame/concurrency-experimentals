package db

import (
	"concurrency-experimentals/configs"
	"concurrency-experimentals/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type UserPersistence interface {
	Create(*models.User) error
	Get(id string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteAll() error
}

type ProductsPersistence interface {
	Create(*models.Product) error
	Get(id string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	DeleteAll() error
}

func GetPostgresConnection() *sql.DB {
	conn, err := sql.Open("postgres", configs.GetPostgresDsn())
	if err != nil {
		panic(err)
	}
	return conn
}
