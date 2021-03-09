// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysql

import (
	_ "embed"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// postgres defines the implementation of the
// database.Driver if Postgres is selected
// as the main driver.
type postgres struct {
	driver *sqlx.DB
	env    *environment.Env
	schema string
}

var (
	//go:embed schema.sql
	migration string
)

// Setup
//
// New - Creates a new mySql instance and returns
// a new database driver.
// Returns errors.INVALID if there was an error establishing a connection or pinging.
func Setup(env *environment.Env) (database.Driver, error) {
	const op = "Database.Setup"

	m := postgres{
		env:    env,
		schema: env.DbSchema,
	}

	driver, err := sqlx.Connect("postgres", m.connectString())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error establishing a database connection", Operation: op, Err: err}
	}

	err = driver.Ping()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error pinging database", Operation: op, Err: err}
	}

	driver.SetMaxIdleConns(database.MaxIdleConns)
	driver.SetMaxOpenConns(database.MaxOpenConns)

	m.driver = driver

	return &m, nil
}

// DB
//
// Returns the sqlx driver.
func (p *postgres) DB() *sqlx.DB {
	return p.driver
}

// Schema
//
// Returns the schema (blank for MySQL),
func (p *postgres) Schema() string {
	return p.schema
}

// Close
//
// Closes the MySQL connection.
func (p *postgres) Close() error {
	return p.driver.Close()
}

// Builder
//
// Returns a new query builder instance.
func (p *postgres) Builder() *builder.Sqlbuilder {
	return builder.New("postgres")
}

// Install
//
// Migrate the db by executing the Postgres migration file.
// Returns errors.INVALID if the sql file could not be located.
// Returns errors.INTERNAL if the exec command could not be ran.
func (p *postgres) Install() error {
	const op = "Database.Install"
	_, err := p.driver.Exec(migration)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error executing migration file", Operation: op, Err: err}
	}
	return nil
}

// Exists
//
// CheckExists check's if the database exists.
// Returns errors.INVALID if the database was not found.
func (p *postgres) Exists() error {
	const op = "Database.CheckExists"
	_, err := p.driver.Exec("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", p.env.DbDatabase)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "No database found with the name: " + p.env.DbDatabase, Operation: op, Err: err}
	}
	return nil
}

// Dump
//
// Dump the database to file with the given path and
// file name.
// Returns errors.INTERNAL if the connection, dump failed.
func (p *postgres) Dump(path, filename string) error {
	const op = "Database.Dump"
	// TODO: Implement!
	return &errors.Error{Code: errors.INTERNAL, Message: "Not yet implemented", Operation: op, Err: fmt.Errorf("function not available")}
}

// Drop
//
// Drop deletes the database with the environments database name.
// Returns errors.INTERNAL if the exec command could not be ran.
func (p *postgres) Drop() error {
	const op = "Database.Drop"
	_, err := p.driver.Exec("DROP DATABASE [IF EXISTS] " + p.env.DbDatabase + ";")
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error dropping the database with the name: " + p.env.DbDatabase, Operation: op, Err: err}
	}
	return nil
}

// connectString
//
// Returns the Postgres database connection string.
func (p *postgres) connectString() string {
	return "postgresql://" + p.env.DbHost + ":" + p.env.DbPort + "/" + p.env.DbDatabase + "?user=" + p.env.DbUser + "&password=" + p.env.DbPassword + "&statement_cache_mode=describe"
}
