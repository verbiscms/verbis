// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/sys"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/spf13/cobra"
	"runtime"
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
			if _, _, err := doctor(false); err != nil {
				return
			}
		},
	}
)

// doctor checks if the environment is validated and checks
// to see if there is a valid database connection and the
// database exists before proceeding.
func doctor(running bool) (*deps.Config, database.Driver, error) {
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
		return nil, nil, fmt.Errorf("validation failed for the environment")
	}

	// Get the database and ping
	db, err := database.New(env)
	if err != nil {
		return nil, nil, err
	}

	// Init Cache
	cache.Init()

	p := paths.Get()
	sys := sys.New(db)

	logger.Info(fmt.Sprintf("Verbis Version: %s, %s", version.Version, version.Prerelease))
	logger.Info(fmt.Sprintf("Go runtime version: %s", runtime.Version()))

	if sys.HasUpdate() {
		logger.Warn(fmt.Sprintf("Verbis outdated, please visit the dashboard to update to version: %s", sys.LatestVersion()))
	}

	// Init Theme
	// TODO: We need pass the default theme (Verbis 2021)
	s, c, err := store.New(db, running)
	if err != nil {
		printError(err.Error())
	}

	printSuccess("All checks passed.")

	return &deps.Config{
		Store:   s,
		Env:     env,
		Config:  c,
		Paths:   p,
		Running: running,
		System:  sys,
	}, db, nil
}
