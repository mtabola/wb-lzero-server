package db

import (
	"fmt"
	"log"
	"server/internal/models"
)

func (s *Store) ReadAllOrders(lim uint8) ([]models.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders LIMIT $1", lim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ords []models.Order
	var order models.Order
	for rows.Next() {
		err := rows.Scan(&order.Id, &order.Data)
		if err != nil {
			log.Panic("row scan error:", err.Error())
			continue
		}
		ords = append(ords, order)
	}

	return ords, nil
}

func (s *Store) ReadOrder(oid string) (models.Order, error) {
	var ord models.Order
	err := s.db.QueryRow(fmt.Sprintf("SELECT * FROM orders WHERE order_info->>'order_uid' = '%s'", oid)).Scan(&ord.Id, &ord.Data)
	if err != nil {
		return models.Order{}, err
	}
	return ord, nil
}

func (s *Store) CreateOrder(order models.OrderStruct) error {
	_, err := s.db.Exec(`INSERT INTO orders (order_info) VALUES($1)`, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CheckOrderExists(oid string) bool {
	var exists bool
	_ = s.db.QueryRow(`SELECT EXISTS(SELECT * FROM orders WHERE order_info->>'order_uid' = $1)`, oid).Scan(&exists)

	return exists
}
