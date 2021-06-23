// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/database/internal"
	"github.com/ainsleyclark/verbis/api/database/mysql"
	"github.com/ainsleyclark/verbis/api/database/postgres"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
)

// Driver represents the functions and methods for
// interacting with the Database, which could
// be MySQL (tested) Postgres (experimental).
type Driver interface {
	DB() *sqlx.DB
	Schema() string
	Builder() *builder.Sqlbuilder
	Install() error
	Migrate(version *version.Version) error
	Tables() ([]string, error)
	Dump(path, filename string) error
	Drop() error
}

// New creates a new database driver dependant on the
// environment.
// Returns errors.INTERNAL if there there was an error setting up the driver.
// Returns errors.INVALID if the environment us invalid or the DB could not be pinged.
func New(env *environment.Env) (Driver, error) {
	const op = "Database.New"

	var (
		db  Driver
		err error
	)

	switch env.DbDriver {
	case internal.MySQLDriver:
		db, err = mysql.Setup(env)
	case internal.PostgresDriver:
		db, err = postgres.Setup(env)
	default:
		return nil, &errors.Error{Code: errors.INVALID, Message: "DB Driver invalid in environment must be 'mysql' or 'postgres", Operation: op, Err: fmt.Errorf("invalid database driver")}
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
