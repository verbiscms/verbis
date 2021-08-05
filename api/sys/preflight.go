// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
)

// newDB is an alias for creating a new database.
var newDB = database.New

// Preflight checks to see if the database is valid before
// proceeding with the installation.
func (s *Sys) Preflight(db domain.InstallPreflight) error {
	logger.Info("Attempting to connect to database")

	env := &environment.Env{
		DbDriver:   database.MySQLDriver,
		DbHost:     db.DbHost,
		DbPort:     db.DbPort,
		DbDatabase: db.DbDatabase,
		DbUser:     db.DbUser,
		DbPassword: db.DbPassword,
	}

	_, err := newDB(env)
	if err != nil {
		return err
	}

	logger.Info("Successfully connected to the database: " + db.DbDatabase)
	return nil
}
