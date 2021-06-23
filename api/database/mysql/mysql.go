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
	// MaxIdleConnections represents the maximum amount
	// of idle connections for MySQL.
	MaxIdleConnections = 5
	// MaxOpenConnections represents the maximum amount
	// of open connections for MySQL.
	MaxOpenConnections = 100
)

// MySQL defines the implementation of the
// MySQL.Driver if MySQL is selected
// as the main driver.
type MySQL struct {
	driver   *sqlx.DB
	env      *environment.Env
	migrator internal.Migrator
}

// Setup creates a new MySQL instance and returns
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

	driver.SetMaxIdleConns(MaxIdleConnections)
	driver.SetMaxOpenConns(MaxOpenConnections)

	migrator, err := internal.NewMigrator(internal.MySQLDriver, driver)
	if err != nil {
		return nil, err
	}

	m.driver = driver
	m.migrator = migrator

	return &m, nil
}

// DB returns the sqlx driver.
func (m *MySQL) DB() *sqlx.DB {
	return m.driver
}

// Schema returns the schema (blank for MySQL),
func (m *MySQL) Schema() string {
	return ""
}

// Builder returns a new query builder instance.
func (m *MySQL) Builder() *builder.Sqlbuilder {
	return builder.New("mysql")
}

// Close closes the MySQL connection.
func (m *MySQL) Close() error {
	const op = "MySQL.Close"
	err := m.driver.Close()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error closing database", Operation: op, Err: err}
	}
	return nil
}

// Install migrates the the database by executing the MySQL
// migration file. The migrator is used to traverse the
// migrations and install the database so it is
// full up to date.
// Returns errors.INTERNAL if the migration failed.
func (m *MySQL) Install() error {
	const op = "MySQL.Install"
	err := m.migrator.Migrate(version.SemVer)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error installing Verbis", Operation: op, Err: err}
	}
	return nil
}

// Migrate migrates and updates the database and the file
// system to the most recent version.
// Returns errors.INTERNAL if the migration failed.
func (m *MySQL) Migrate(ver *sm.Version) error {
	const op = "MySQL.Migrate"
	err := m.migrator.Migrate(ver)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error migrating", Operation: op, Err: err}
	}
	return nil
}

// Tables Runs checks on the Verbis DB installation to see if
// all the tables exist. If some are missing, a slice of
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

// Exists check's if the database exists.
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

	err = ioutil.WriteFile(p, bytes, os.ModePerm)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "No file or directory with the path: " + p, Operation: op, Err: err}
	}

	return nil
}

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

// connectString Returns the MySQL database connection
// string.
func (m *MySQL) connectString() string {
	return m.env.DbUser + ":" + m.env.DbPassword + "@tcp(" + m.env.DbHost + ":" + m.env.DbPort + ")/" + m.env.DbDatabase + "?tls=false&parseTime=true&multiStatements=true"
}
