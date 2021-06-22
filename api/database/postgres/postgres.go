// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postgres

import (
	_ "embed" // Embed Migration
	"fmt"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/database/internal"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres Driver
)

// Postgres defines the implementation of the
// database.Driver if Postgres is selected
// as the main driver.
type Postgres struct {
	driver   *sqlx.DB
	env      *environment.Env
	schema   string
	migrator internal.Tester
}

// Setup
//
// New - Creates a new mySql instance and returns
// a new database driver.
// Returns errors.INVALID if there was an error establishing a connection or pinging.
func Setup(env *environment.Env) (*Postgres, error) {
	const op = "Database.Setup"

	m := Postgres{
		env:    env,
		schema: env.DbSchema,
	}

	driver, err := sqlx.Connect("Postgres", m.connectString())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error establishing a database connection", Operation: op, Err: err}
	}

	err = driver.Ping()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error pinging database", Operation: op, Err: err}
	}

	//driver.SetMaxIdleConns(database.MaxIdleConns)
	//driver.SetMaxOpenConns(database.MaxOpenConns)

	migrator, err := internal.NewMigrator(internal.PostgresDriver, driver)
	if err != nil {
		return nil, err
	}

	m.driver = driver
	m.migrator = migrator

	return &m, nil
}

// DB
//
// Returns the sqlx driver.
func (p *Postgres) DB() *sqlx.DB {
	return p.driver
}

// Schema
//
// Returns the schema (blank for MySQL),
func (p *Postgres) Schema() string {
	return p.schema
}

// Close
//
// Closes the MySQL connection.
func (p *Postgres) Close() error {
	return p.driver.Close()
}

// Builder
//
// Returns a new query builder instance.
func (p *Postgres) Builder() *builder.Sqlbuilder {
	return builder.New("Postgres")
}

// Install
//
// Migrate the db by executing the Postgres migration file.
// Returns errors.INVALID if the sql file could not be located.
// Returns errors.INTERNAL if the exec command could not be ran.
func (p *Postgres) Install() error {
	//const op = "Database.Install"
	//_, err := p.driver.Exec(migration)
	//if err != nil {
	//	return &errors.Error{Code: errors.INTERNAL, Message: "Error executing migration file", Operation: op, Err: err}
	//}
	return nil
}

func (p *Postgres) Tables() ([]string, error) {
	panic("Implement me")
}

// Exists
//
// CheckExists check's if the database exists.
// Returns errors.INVALID if the database was not found.
func (p *Postgres) Exists() error {
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
func (p *Postgres) Dump(path, filename string) error {
	const op = "Database.Dump"
	// TODO: Implement!
	return &errors.Error{Code: errors.INTERNAL, Message: "Not yet implemented", Operation: op, Err: fmt.Errorf("function not available")}
}

// Drop
//
// Drop deletes the database with the environments database name.
// Returns errors.INTERNAL if the exec command could not be ran.
func (p *Postgres) Drop() error {
	const op = "Database.Drop"
	_, err := p.driver.Exec("DROP DATABASE [IF EXISTS] " + p.env.DbDatabase + ";")
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error dropping the database with the name: " + p.env.DbDatabase, Operation: op, Err: err}
	}
	return nil
}

func (p *Postgres) Migrate(version *sm.Version) error {
	err := p.migrator.Migrate(version)
	if err != nil {
		return err
	}
	return nil
}

// connectString
//
// Returns the Postgres database connection string.
func (p *Postgres) connectString() string {
	return "postgresql://" + p.env.DbHost + ":" + p.env.DbPort + "/" + p.env.DbDatabase + "?user=" + p.env.DbUser + "&password=" + p.env.DbPassword + "&statement_cache_mode=describe"
}
