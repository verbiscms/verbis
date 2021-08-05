// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/database/seeds"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/store"
)

func (s *Sys) Install(db domain.InstallVerbis, restart bool) error {
	defer func() {
		if !restart {
			return
		}
		logger.Info("Restarting Verbis")
		err := s.Restart()
		if err != nil {
			logger.WithError(err).Panic()
		}
	}()

	env := &environment.Env{
		DbHost:     db.DBHost,
		DbPort:     db.DBPort,
		DbDatabase: db.DBDatabase,
		DbUser:     db.DBUser,
		DbPassword: db.DBPassword,
	}

	// Connect to database
	logger.Info("Attempting to connect to database")
	driver, err := database.New(env)
	if err != nil {
		return err
	}
	logger.Info("Successfully connected to the database: " + db.DBDatabase)

	// Migrate
	logger.Info("Migrating database")
	err = driver.Install()
	if err != nil {
		return err
	}

	// Obtain the store
	repository, err := store.New(driver, false)
	if err != nil {
		return err
	}

	// Run the seeds
	logger.Info("Attempting to run seeds")
	seeder := seeds.New(driver, repository)
	err = seeder.Seed()
	if err != nil {
		return err
	}
	logger.Info("Successfully ran seeds")

	// Write to the env
	logger.Info("Attempting to write to env file")
	err = env.Install()
	if err != nil {
		return err
	}
	logger.Info("Successfully wrote to .env file")

	return nil
}
