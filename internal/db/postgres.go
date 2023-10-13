package db

import (
	"database/sql"
	"fmt"
	"server/config"
)

type Store struct {
	DB *sql.DB
}

func New(dbcp config.DBParams) (*Store, error) {
	connlink := fmt.Sprintf("host = %s port = %d user=%s password=%s dbname=%s sslmode=disable", dbcp.Host, dbcp.Port, dbcp.User, dbcp.Password, dbcp.DBName)
	db, err := sql.Open("postgres", connlink)

	if err != nil {
		return nil, err
	}
	err = prepare(db, ` 
	CREATE TABLE IF NOT EXISTS orders (
		order_id serial PRIMARY KEY,
		order_info jsonb
	)`,
		"db.postgres.New()")
	if err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}
