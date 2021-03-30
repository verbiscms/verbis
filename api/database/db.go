// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

//nolint
import (
	_ "embed"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/database/mysql"
	"github.com/ainsleyclark/verbis/api/environment"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Driver interface {
	DB() *sqlx.DB
	Schema() string
	Builder() *builder.Sqlbuilder
	Install() error
	Dump(path, filename string) error
	Drop() error
}

// TODO
// establish what drier it is and do a switch
func New(env *environment.Env) (Driver, error) {
	db, err := mysql.Setup(env)
	if err != nil {
		return nil, err
	}

	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}

	return db, nil
}
