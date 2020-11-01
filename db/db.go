package db

import (
	"concurrency-experimentals/configs"
	"concurrency-experimentals/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

type OrdersPersistence interface {
	Create(*models.Order) error
	Get(id string) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	DeleteAll() error
}

func GetPostgresConnection() *sql.DB {
	return open("postgres", configs.GetPostgresDsn())
}

func GetMysqlConnection() *sql.DB {
	return open("mysql", configs.GetMysqlDsn())
}

func open(driver, dsn string) *sql.DB {
	conn, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	return conn
}
