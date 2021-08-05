// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
)

const (
	InstallDatabaseStep = iota
	InstallUserStep
	InstallSiteStep
)

// newDB is an alias for creating a new database.
var newDB = database.New

func (s *Sys) Validate(step int, data interface{}) error {
	switch step {
	case InstallDatabaseStep:
		db, ok := data.(domain.InstallDatabase)
		if !ok {
			return fmt.Errorf("error")
		}
		fmt.Println("got here")
		err := validation.Validator()(data)
		if err != nil {
			return err
		}
		return s.Preflight(db)
	default:
		return validation.Validator().Struct(data)
	}
}


func (s *Sys) ValidateDatabaseStep(db domain.InstallDatabase) error {
	logger.Info("Attempting to connect to database")

	env := &environment.Env{
		DbDriver:   database.MySQLDriver,
		DbHost:     db.DBHost,
		DbPort:     db.DBPort,
		DbDatabase: db.DBDatabase,
		DbUser:     db.DBUser,
		DbPassword: db.DBPassword,
	}

	_, err := newDB(env)
	if err != nil {
		return err
	}

	logger.Info("Successfully connected to the database: " + db.DBDatabase)
	return nil
}

func (s *Sys) ValidateUserStep(user domain.InstallUser) error {
	return validation.Validator().Struct(user)
}

func (s *Sys) ValidateSiteStep(site domain.InstallSite) error {
	return validation.Validator().Struct(site)
}

// Preflight checks to see if the database is valid before
// proceeding with the installation.
func (s *Sys) Preflight(db domain.InstallDatabase) error {
	logger.Info("Attempting to connect to database")

	env := &environment.Env{
		DbDriver:   database.MySQLDriver,
		DbHost:     db.DBHost,
		DbPort:     db.DBPort,
		DbDatabase: db.DBDatabase,
		DbUser:     db.DBUser,
		DbPassword: db.DBPassword,
	}

	_, err := newDB(env)
	if err != nil {
		return err
	}

	logger.Info("Successfully connected to the database: " + db.DBDatabase)
	return nil
}
