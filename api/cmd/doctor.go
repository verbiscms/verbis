// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
	"strings"
)

var (
	doctorCmd = &cobra.Command{
		Use:   "doctor",
		Short: "Running doctor will check the system for any potential hiccups when installing, updating or running Verbis.",
		Long: `This command is a diagnostic tool to find any potential issues for your
Verbis install. It will check if the database has been set up correctly as well as the
environment.`,
		Run: func(cmd *cobra.Command, args []string) {
			if _, _, err := doctor(); err != nil {
				return
			}
			return
		},
	}
)

// doctor checks if the environment is validated and checks
// to see if there is a valid database connection and the
// database exists before proceeding.
func doctor() (*deps.DepsConfig, *database.MySql, error) {

	printSpinner("Running doctor...")

	// Check paths are correct
	if err := paths.BaseCheck(); err != nil {
		printError(err.Error())
		return nil, nil, err
	}

	// Load the environment (.env file)
	env, err := environment.Load()
	if err != nil {
		printError(err.Error())
		return nil, nil, err
	}

	// Init logging
	logger.Init(env)

	// Check if the environment values are valid
	vErrors := env.Validate()
	if vErrors != nil {
		for _, v := range vErrors {
			printError(fmt.Sprintf("Obtaining environment variable: %s", strings.ToUpper(v.Key)))
		}
		return nil, nil, fmt.Errorf("Validation failed for the enviroment")
	}

	// Get the database and ping
	db, err := database.New(env)
	if err != nil {
		printError(fmt.Sprintf("Establishing database connection, are the credentials in the .env file correct? %s", err.Error()))
		return nil, nil, fmt.Errorf("Error establishing database connection")
	}

	// Check if the database exists
	if err := db.CheckExists(); err != nil {
		printError(fmt.Sprintf("Establishing database connection, are the credentials in the .env file correct? %s", err.Error()))
		return nil, nil, fmt.Errorf("error establishing database connection")
	}

	// Init Cache
	cache.Init()

	paths := paths.Get()

	// Init Config
	// TODO: We need pass the default theme (Verbis 2021)

	//if err != nil {
	//	printError(errors.Message(err))
	//}

	// Set up stores & pass the database.
	store := models.New(&models.StoreConfig{
		DB:    db.Sqlx,
		Paths: paths,
	})

	printSuccess("All checks passed.")

	return &deps.DepsConfig{
		Store:  store,
		Env:    env,
		Config: config.Get(),
		Paths:  paths,
	}, db, nil
}
