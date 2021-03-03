// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

//nolint
import (
	_ "embed"
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	MaxIdleConns = 5
	MaxOpenConns = 100
)

type DBCfg struct {
	Sqlx  *sqlx.DB
	En    environment.Env
	Paths paths.Paths
}

var (
	//go:embed schema.sql
	migration string
)

// MySQL defines the driver for the database
type MySQL struct {
	Sqlx     *sqlx.DB
	env      *environment.Env
	database string
	paths    paths.Paths
}

// New - Creates a new MySQL instance.
func New(env *environment.Env) (*MySQL, error) {
	db := MySQL{
		env:      env,
		database: env.DBDatabase,
		paths:    paths.Get(),
	}

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

// Get Database open's sql database connection
// Returns errors.INVALID if the the connection string or database is invalid.
func (db *MySQL) GetDatabase() (*sqlx.DB, error) {
	const op = "Database.GetDatabase"
	var driver *sqlx.DB
	driver, err := sqlx.Connect("mysql", db.env.ConnectString())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Could not establish a database connection", Operation: op, Err: err}
	}
	driver.SetMaxIdleConns(MaxIdleConns)
	driver.SetMaxOpenConns(MaxOpenConns)
	return driver, nil
}

// CheckExists check's if database exists with a given name
// Returns errors.INVALID if the database was not found.
func (db *MySQL) CheckExists() error {
	const op = "Database.CheckExists"
	_, err := db.Sqlx.Exec("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", db.database)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("No database found with the name: %s", db.database), Operation: op, Err: err}
	}
	return nil
}

// Ping database to check connection
// Returns errors.INVALID if the ping was unsuccessful.
func (db *MySQL) Ping() error {
	const op = "Database.Ping"
	if err := db.Sqlx.Ping(); err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Pinging the database was unsuccessful", Operation: op, Err: err}
	}
	return nil
}

// Install Verbis by executing the migration file
// Returns errors.INVALID if the sql file could not be located.
// Returns errors.INTERNAL if the exec command could not be ran.
func (db *MySQL) Install() error {
	const op = "Database.Install"
	if _, err := db.Sqlx.Exec(migration); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could execute the migration file", Operation: op, Err: err}
	}
	return nil
}

// Drop deletes the database with the environments database name.
// Returns errors.INTERNAL if the exec command could not be ran.
func (db *MySQL) Drop() error {
	const op = "Database.Drop"
	_, err := db.Sqlx.Exec("DROP DATABASE " + db.database + ";")
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not drop the database with the name: %s", db.database), Operation: op, Err: err}
	}
	return nil
}

// Create the database with the environments database name.
// Returns errors.INTERNAL if the exec command could not be ran.
func (db *MySQL) Create() error {
	const op = "Database.Create"
	_, err := db.Sqlx.Exec("CREATE DATABASE " + db.database + ";")
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the database with the name: %s", db.database), Operation: op, Err: err}
	}
	return nil
}

// Dump the database to file with the given path and file name.
// Returns errors.INTERNAL if the connection, dump failed as well as closing
// the database.
func (db *MySQL) Dump(path string, filename string) error {
	const op = "Database.Dump"
	dumper, err := mysqldump.Register(db.Sqlx.DB, path, filename)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Unable to register with mysqldump", Operation: op, Err: err}
	}

	_, err = dumper.Dump()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not dump the database with the path and filename: %s", path+filename), Operation: op, Err: err}
	}

	if err := dumper.Close(); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not close the database connection", Operation: op, Err: err}
	}

	return nil
}
