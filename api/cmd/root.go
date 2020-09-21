/*
Copyright © 2020 NAME HERE ainsley@reddico.co.uk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/cron"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/logger"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// TODO: Change db and store to local variables
type App struct {
	db *database.DB
	store *models.Store
}

// Root represents the base command when called without any subcommands
var (
	app App
	rootCmd = &cobra.Command{
		Use:   "Verbis",
		Short: "Verbis - Command Shell for Serving, Installing & Migrating.",
		Long: `Verbis - Command Shell for Serving, Installing & Migrating.`,
		DisableAutoGenTag: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Add child commands and bootstrap
func init() {
	bootstrap()
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(usersCmd)
	rootCmd.AddCommand(postsCmd)
	rootCmd.AddCommand(seedCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(testCmd)
}

// Bootstrap the application
func bootstrap() {

	// Load ENV
	err := environment.Load()
	if err != nil {
		log.Panic(err)
	}

	// Init logging
	err = logger.Init()
	if err != nil {
		log.Panic(err)
	}

	// Init Cache
	cache.Init()

	// Init Config
	config.Init()

	// Load Database
	db, err := database.New()
	if err != nil {
		log.Panic(err)
	}

	// Set up stores & pass the database.
	store, err := models.New(db)
	if err != nil {
		log.Panic(err)
	}

	// Load cron jobs
	scheduler := cron.New(store)
	go scheduler.Run()

	// Load app
	app = App{
		db: db,
		store: store,
	}
}
