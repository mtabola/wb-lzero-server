package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func prepare(db *sql.DB, query string, callFunc string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: %w", callFunc, err)
	}

	if _, err = stmt.Exec(); err != nil {
		return fmt.Errorf("%s: %w", callFunc, err)
	}
	return nil
}
