// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

//nolint
import (
	_ "embed"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	MaxIdleConns = 5
	MaxOpenConns = 100
)

type Driver interface {
	DB() *sqlx.DB
	Schema() string
	Install() error
	Dump(path, filename string) error
	Drop() error
}

// establish what drier it is and do a switch
