package db

import (
	"concurrency-experimentals/models"
	"database/sql"
	"fmt"
	"strings"
)

const ProductsTable = "products"

type postgresProductsPersistence struct {
	conn *sql.DB
}

func NewPostgresProductsPersistence(conn *sql.DB) ProductsPersistence {
	return &postgresProductsPersistence{conn}
}

func (p *postgresProductsPersistence) Create(product *models.Product) error {
	columns := strings.Join(GetProductColumns(), ",")
	values := ToPostgresValues(GetProductColumns())
	query := fmt.Sprintf("insert into %s (%s) values (%s)", ProductsTable, columns, values)

	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.ID,
		product.Name,
		product.Description,
		product.Quantity,
		product.Price,
		product.Available,
		product.PhotoURL,
		product.Ratings,
		product.Category,
		product.Manufacturer,
		product.IsDeleted,
		product.CreatedAt,
		product.UpdatedAt,
		product.DeletedAt)

	if err != nil {
		fmt.Println("sql exec statement error: ", err.Error())
		return tx.Rollback()
	}
	return tx.Commit()
}
func (p *postgresProductsPersistence) Get(id string) (*models.Product, error) {
	query := fmt.Sprintf("select * from %s where id = $1", ProductsTable)
	var product *models.Product
	err := p.conn.QueryRow(query, id).Scan(&product)
	return product, err
}
func (p *postgresProductsPersistence) GetAll() ([]*models.Product, error) {
	query := fmt.Sprintf("select * from %s", ProductsTable)
	rows, err := p.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product *models.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.Available,
			&product.PhotoURL,
			&product.Ratings,
			&product.Category,
			&product.Manufacturer,
			&product.IsDeleted,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt)
		if err != nil {
			fmt.Println("sql row scan error: ", err.Error())
			continue
		}
		products = append(products, product)
	}
	return products, nil
}
func (p *postgresProductsPersistence) DeleteAll() error {
	query := fmt.Sprintf("delete from %s", ProductsTable)
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec(query)
	if err != nil {
		fmt.Println("Sql exec transaction error: ", err.Error())
		return tx.Rollback()
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Sql rows affected error: ", err.Error())
	} else {
		fmt.Println("Deleted rows: ", rows)
	}
	return tx.Commit()
}

func GetProductColumns() []string {
	columns := []string{
		"id",
		"name",
		"description",
		"quantity",
		"price",
		"available",
		"photo_url",
		"ratings",
		"category",
		"manufacturer",
		"is_deleted",
		"created_at",
		"updated_at",
		"deleted_at",
	}
	return columns
}

func ToPostgresValues(columns []string) string {
	result := make([]string, 0, len(columns))
	for index := range columns {
		result = append(result, fmt.Sprintf("$%d", index+1))
	}
	return strings.Join(result, ",")
}
