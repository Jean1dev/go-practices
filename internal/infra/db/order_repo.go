package db

import (
	"database/sql"
	"jeanlucafp/go-intesivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func buildRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("Insert int orders (id, price, tax, final_price) Values(?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
