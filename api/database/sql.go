package database

import (
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySql struct {
	Sqlx 	*sqlx.DB
}

// Create a new database instance.
func New() (*MySql, error) {
	db := MySql{}

	sql, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}
	db.Sqlx = sql

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &db, nil
}

// Open sql database connection
func (db *MySql) GetDatabase() (*sqlx.DB, error) {
	var driver *sqlx.DB
	driver, err := sqlx.Connect("mysql", environment.ConnectString())
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return driver, nil
}

// Check if database exists
func (db *MySql) CheckExists() error {
	_, err := db.Sqlx.Exec("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", environment.GetDatabaseName())
	if err != nil {
		return fmt.Errorf("No database found with the name: %s", environment.GetDatabaseName())
	}
	return nil
}

//Ping database to check connection
func (db *MySql) Ping() error {
	if err := db.Sqlx.Ping(); err != nil {
		return fmt.Errorf("Error pinging the database")
	}
	return nil
}

// Install the cms
func (db *MySql) Install() error {
	path := paths.Migration() + "/schema.sql"
	sql, err := files.GetFileContents(path)
	if err != nil {
		return fmt.Errorf("Could not get the schema sql file from the path: %s", path)
	}

	if _, err := db.Sqlx.Exec(sql); err != nil {
		return err
	}

	return nil
}

// Drop database
func (db *MySql) Drop() error {
	_, err := db.Sqlx.Exec("DROP DATABASE " + environment.GetDatabaseName() + ";")
	if err != nil {
		return fmt.Errorf("Error dropping database: %w", err)
	}
	return nil
}

// Create database
func (db *MySql) Create() error {
	_, err := db.Sqlx.Exec("CREATE DATABASE " + environment.GetDatabaseName() + ";")
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}
	return nil
}

// Dump database
func (db *MySql) Dump(path string, filename string) error {
	dumper, err := mysqldump.Register(db.Sqlx.DB, path, filename)
	if err != nil {
		return fmt.Errorf("could not conenct to the database %v", err)
	}

	_, err = dumper.Dump()
	if err != nil {
		return fmt.Errorf("errror dumping the database: %v", err)
	}

	if err := dumper.Close(); err != nil {
		return fmt.Errorf("errror closing database: %v", err)
	}

	return nil
}



