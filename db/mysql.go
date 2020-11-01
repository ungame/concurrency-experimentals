package db

import (
	"concurrency-experimentals/models"
	"database/sql"
	"fmt"
	"strings"
)

const OrdersTable = "orders"

type mysqlDbOrdersPersistence struct {
	conn *sql.DB
}

func NewMySqlDbOrdersPersistence(conn *sql.DB) OrdersPersistence {
	return &mysqlDbOrdersPersistence{conn}
}

func (p *mysqlDbOrdersPersistence) Create(order *models.Order) error {
	columns := strings.Join(GetOrderColumns(), ",")
	values := ToMysqlValues(GetOrderColumns())
	query := fmt.Sprintf("insert into %s (%s) values (%s)", OrdersTable, columns, values)

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
		order.ID,
		order.UserID,
		order.ProductID,
		order.Description,
		order.Quantity,
		order.UnitPrice,
		order.Amount,
		order.Status,
		order.ReasonReject,
		order.Paid,
		order.PaymentType,
		order.IsDeleted,
		order.CreatedAt,
		order.UpdatedAt)

	if err != nil {
		fmt.Println("sql exec statement error: ", err.Error())
		return tx.Rollback()
	}
	return tx.Commit()
}
func (p *mysqlDbOrdersPersistence) Get(id string) (*models.Order, error) {
	query := fmt.Sprintf("select * from %s where id = ?", OrdersTable)
	var order *models.Order
	err := p.conn.QueryRow(query, id).Scan(&order)
	return order, err
}
func (p *mysqlDbOrdersPersistence) GetAll() ([]*models.Order, error) {
	query := fmt.Sprintf("select * from %s", OrdersTable)
	rows, err := p.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order

	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductID,
			&order.Description,
			&order.Quantity,
			&order.UnitPrice,
			&order.Amount,
			&order.Status,
			&order.ReasonReject,
			&order.Paid,
			&order.PaymentType,
			&order.IsDeleted,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.DeletedAt,
			&order.CanceledAt)
		if err != nil {
			fmt.Println("sql row scan error: ", err.Error())
			continue
		}
		orders = append(orders, order)
	}
	return orders, nil
}
func (p *mysqlDbOrdersPersistence) DeleteAll() error {
	query := fmt.Sprintf("delete from %s", OrdersTable)
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

func GetOrderColumns() []string {
	return []string{
		"id",
		"user_id",
		"product_id",
		"description",
		"quantity",
		"unit_price",
		"amount",
		"status",
		"reason_reject",
		"paid",
		"payment_type",
		"is_deleted",
		"created_at",
		"updated_at",
		//"deleted_at",
		//"canceled_at",
	}
}


func ToMysqlValues(columns []string) string {
	result := make([]string, 0, len(columns))
	for _ = range columns {
		result = append(result, fmt.Sprint("?"))
	}
	return strings.Join(result, ",")
}