// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "github.com/ainsleyclark/verbis/api/errors"

var (
	Tables = []string{
		"categories",
		"form_fields",
		"form_submissions",
		"forms",
		"media",
		"options",
		"password_resets",
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
	ErrTableNotFound = errors.New("database tables missing from verbis installation")
)

const (
	// MySQLDriver driver is represented under DB_DRIVER
	// for MySQL.
	MySQLDriver = "mysql"
	// PostgresDriver driver is represented under
	// DB_DRIVER for postgres.
	PostgresDriver = "postgres"

	ErrDBConnectionMessage  = "Error establishing database connection"
	ErrTableNotFoundMessage = "Verbis database tables missing"
)
