// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "github.com/verbiscms/verbis/api/errors"

var (
	// Tables define the current database tables within
	// Verbis.
	Tables = []string{
		"categories",
		"files",
		"forms",
		"media",
		"media_sizes",
		"options",
		"post_categories",
		"post_fields",
		"post_options",
		"posts",
		"redirects",
		"roles",
		"user_roles",
		"users",
	}
)

var (
	// ErrTableNotFound is returned by the driver if there are
	// tables missing from the installation.
	ErrTableNotFound = errors.New("database tables missing from Verbis")
)

const (
	// MySQLDriver driver is represented under DB_DRIVER
	// for MySQL.
	MySQLDriver = "mysql"
	// PostgresDriver driver is represented under
	// DB_DRIVER for postgres.
	PostgresDriver = "postgres"
	// ErrDBConnectionMessage is used as an error message when
	// no database connection could be established.
	ErrDBConnectionMessage = "Error establishing database connection"
	// ErrTableNotFoundMessage is used as an error message
	// when a table is missing from the installation.
	ErrTableNotFoundMessage = "Verbis database tables missing"
)
