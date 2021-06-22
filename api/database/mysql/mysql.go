// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysql

import (
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/database/internal"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/version"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	//
	MaxIdleConns = 5
	MaxOpenConns = 100
)

// MySQL defines the implementation of the
// MySQL.Driver if MySQL is selected
// as the main driver.
type MySQL struct {
	driver   *sqlx.DB
	env      *environment.Env
	migrator internal.Migrator
}

// Setup
//
// New - Creates a new MySQL instance and returns
// a new database driver.
// Returns errors.INVALID if there was an error establishing a connection or pinging.
func Setup(env *environment.Env) (*MySQL, error) {
	const op = "MySQL.Setup"

	m := MySQL{
		env: env,
	}

	driver, err := sqlx.Connect("mysql", m.connectString())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: internal.ErrDBConnectionMessage, Operation: op, Err: err}
	}

	err = driver.Ping()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error pinging database", Operation: op, Err: err}
	}

	driver.SetMaxIdleConns(MaxIdleConns)
	driver.SetMaxOpenConns(MaxOpenConns)

	migrator, err := internal.NewMigrator(internal.MySQLDriver, driver)
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
func (m *MySQL) DB() *sqlx.DB {
	return m.driver
}

// Schema
//
// Returns the schema (blank for MySQL),
func (m *MySQL) Schema() string {
	return ""
}

// Builder
//
// Returns a new query builder instance.
func (m *MySQL) Builder() *builder.Sqlbuilder {
	return builder.New("mysql")
}

// Close
//
// Closes the MySQL connection.
func (m *MySQL) Close() error {
	const op = "MySQL.Close"

	err := m.driver.Close()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error closing database", Operation: op, Err: err}
	}

	return nil
}

// Install
//
// Migrate the db by executing the MySQL migration file.
// Returns errors.INVALID if the sql file could not be located.
// Returns errors.INTERNAL if the exec command could not be ran.
func (m *MySQL) Install() error {
	const op = "MySQL.Install"
	err := m.migrator.Migrate(version.SemVer)
	if err != nil {
		return err
	}
	return nil
}

// Tables
//
// Runs checks on the Verbis DB installation to see if all
// the tables exist. If some are missing, a slice of
// strings will be returned containing the table
// name.
// Returns internal.ErrTableNotFound if a table wasn't found.
func (m *MySQL) Tables() ([]string, error) {
	const op = "MySQL.Tables"

	var failedTables []string
	for _, table := range internal.Tables {
		q := m.Builder().
			From("information_schema.tables").
			Where("table_schema", "=", m.env.DbDatabase).
			Where("table_name", "=", table).
			Limit(1).
			Exists()

		var exists bool
		err := m.driver.QueryRow(q).Scan(&exists)
		if err != nil {
			failedTables = append(failedTables, table)
		}
	}

	if len(failedTables) > 0 {
		return failedTables, &errors.Error{Code: errors.INVALID, Message: internal.ErrTableNotFoundMessage, Operation: op, Err: internal.ErrTableNotFound}
	}

	return nil, nil
}

// Exists
//
// CheckExists check's if the database exists.
// Returns errors.INVALID if the database was not found.
func (m *MySQL) Exists() error {
	const op = "MySQL.Exists"

	q := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"

	_, err := m.driver.Exec(q, m.env.DbDatabase)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "No database found with the name: " + m.env.DbDatabase, Operation: op, Err: err}
	}

	return nil
}

// execCommand func for mysqldump
var execCommand = exec.Command

// Dump
//
// Dump the database to file with the given path and
// file name.
// Returns errors.INTERNAL if the connection, dump failed.
func (m *MySQL) Dump(path, filename string) error {
	const op = "MySQL.Dump"

	cmd := execCommand("mysqldump", "-P"+m.env.DbPort, "-h"+m.env.DbHost, "-u"+m.env.DbUser, "-p"+m.env.DbPassword, m.env.DbDatabase)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	p := path + string(os.PathSeparator) + filename + ".sql"

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error writing SQL to file with the path: " + p, Operation: op, Err: err}
	}

	err = ioutil.WriteFile(p, bytes, 0644)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "No file or directory with the path: " + p, Operation: op, Err: err}
	}

	return nil
}

// Drop
//
// Drop deletes the database with the environments database name.
// Returns errors.INTERNAL if the exec command could not be ran.
func (m *MySQL) Drop() error {
	const op = "MySQL.Drop"

	_, err := m.driver.Exec("DROP DATABASE " + m.env.DbDatabase + ";")
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error dropping the database with the name: " + m.env.DbDatabase, Operation: op, Err: err}
	}
	return nil
}

func (m *MySQL) Migrate(version *sm.Version) error {
	err := m.migrator.Migrate(version)
	if err != nil {
		return err
	}
	return nil
}

// connectString
//
// Returns the MySQL database connection string.
func (m *MySQL) connectString() string {
	return m.env.DbUser + ":" + m.env.DbPassword + "@tcp(" + m.env.DbHost + ":" + m.env.DbPort + ")/" + m.env.DbDatabase + "?tls=false&parseTime=true&multiStatements=true"
}
