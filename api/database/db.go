// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

//nolint
import (
	_ "embed"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/database/mysql"
	"github.com/ainsleyclark/verbis/api/database/postgres"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Driver
type Driver interface {
	DB() *sqlx.DB
	Schema() string
	Builder() *builder.Sqlbuilder
	Install() error
	Dump(path, filename string) error
	Drop() error
}

const (
	// MySQLDriver driver is represented under DB_DRIVER
	// for MySQL.
	MySQLDriver = "mysql"
	// PostgresDriver driver is represented under
	// DB_DRIVER for postgres.
	PostgresDriver = "postgres"
)

// TODO
//
// establish what drier it is and do a switch
func New(env *environment.Env) (Driver, error) {
	const op = "Database.New"

	var (
		db  Driver
		err error
	)

	switch env.DbDriver {
	case MySQLDriver:
		db, err = mysql.Setup(env)
	case PostgresDriver:
		db, err = postgres.Setup(env)
	default:
		return nil, &errors.Error{Code: errors.INVALID, Message: "DB Driver invalid in environment must be 'mysql' or 'postgres", Operation: op, Err: fmt.Errorf("invalid database driver")}
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
