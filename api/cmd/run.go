package cmd

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/cron"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/logger"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/routes"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Serve the CMS.",
		Long:  `Serve will serve the system dependant on port number passed.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError(err.Error())
			}

			// Init logging
			if err := logger.Init(); err != nil {
				printError(err.Error())
			}

			// Init Cache
			cache.Init()

			// Init Config
			config.Init()

			// Set up stores & pass the database.
			store, err := models.New(db)
			if err != nil {
				printError(err.Error())
			}

			// Load cron jobs
			scheduler := cron.New(store)
			go scheduler.Run()

			// Set up the router & pass logger
			serve, err := server.New()
			if err != nil {
				printError(err.Error())
			}

			// Pass the stores to the controllers
			controllers, err := controllers.New(store)
			if err != nil {
				printError(err.Error())
			}

			// Load the routes
			routes.Load(serve, controllers, store)

			// Listen & serve.
			err = serve.ListenAndServe(environment.GetPort())
			if err != nil {
				printError(err.Error())
			}
		},
	}
)