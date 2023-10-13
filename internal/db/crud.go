package db

import "server/internal/models"

func (s *Store) ReadAllOrders() error {
	return nil
}

func (s *Store) ReadArticle(orderId string) error {
	return nil
}

func (s *Store) CreateOrder(order models.OrderStruct) error {
	_, err := s.DB.Exec(`INSERT INTO orders (order_info) VALUES($1)`, order)
	if err != nil {
		return err
	}
	return nil
}
