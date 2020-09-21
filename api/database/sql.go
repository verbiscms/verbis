package database

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Sqlx 	*sqlx.DB
}

func newInstance() *DB {
	return &DB{}
}

// Create a new database instance.
func New() (*DB, error) {
	db := newInstance()

	// Open sql database connection
	sqlxInstance, err := sqlx.Open("mysql", environment.ConnectString())
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	db.Sqlx = sqlxInstance

	defer func(){
		db.Sqlx.Close()
	}()

	// Check if database exists
	_, err = db.Sqlx.Exec("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", environment.Env.DbDatabase)
	if err != nil {
		return nil, fmt.Errorf("unknown database: %s", environment.Env.DbDatabase)
	}

	//Ping database to check connection
	err = db.Sqlx.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

// Drop database
func (db *DB) drop() error {
	_, err := db.Sqlx.Exec("DROP DATABASE " + environment.Env.DbDatabase + ";")
	if err != nil {
		return fmt.Errorf("error dropping database: %w", err)
	}
	return nil
}

// Create database
func (db *DB) create() error {
	_, err := db.Sqlx.Exec("CREATE DATABASE " + environment.Env.DbDatabase + ";")
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}
	return nil
}


