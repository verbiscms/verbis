package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/cron"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/routes"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Running start will start Verbis project from the current directory and run the CMS project.",
		Long:  `This command will start Verbis from the current directory. First it will
run Verbis doctor to see if the environment is configured correctly. It will then start
up the server on the port specified in the .env file.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError(err.Error())
			}

			// Init Config
			cfg, err := config.New()
			if err != nil {
				printError(errors.Message(err))
			}

			// Set up stores & pass the database.
			store := models.New(db, *cfg)
			if err != nil {
				printError(err.Error())
			}

			// Load cron jobs
			scheduler := cron.New(store)
			go scheduler.Run()

			// Set up the router & pass logger
			serve := server.New(store.Options)

			// Pass the stores to the controllers
			controllers, err := controllers.New(store, *cfg)
			if err != nil {
				fmt.Println(err)
				printError(err.Error())
			}

			// Load the routes
			routes.Load(serve, controllers, store, *cfg)

			// Print listening success
			printSuccess(fmt.Sprintf("Verbis listening on port: %d", environment.GetPort()))
			fmt.Println()

			// Listen & serve.
			err = serve.ListenAndServe(environment.GetPort())
			if err != nil {
				printError(err.Error())
			}
		},
	}
)