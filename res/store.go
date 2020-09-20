package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Store store

type store struct {
	User UserRepository
}

// Create a new database instance, connect to database.
func NewStore() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", DBConnectString())

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	Store = store{
		User: NewUserStore(db),
	}

	return db, nil
}

