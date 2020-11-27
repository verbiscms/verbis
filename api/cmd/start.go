package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/cron"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/routes"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Running start will start Verbis project from the current directory and run the CMS project.",
		Long: `This command will start Verbis from the current directory. First it will
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
			controllers, err := handler.New(store, *cfg)
			if err != nil {
				fmt.Println(err)
				printError(err.Error())
			}

			// Load the routes
			routes.Load(serve, controllers, store, *cfg)

			// Get options
			opts, err := store.Options.GetStruct()
			if err != nil {
				printError(err.Error())
			}

			// Print listening success
			printSuccess(fmt.Sprintf("Verbis listening on port: %d \n", environment.GetPort()))
			emoji.Printf(":backhand_index_pointing_right: Visit your site at:          %s \n", opts.SiteUrl)
			emoji.Printf(":key: Or visit the admin area at:  %s \n", opts.SiteUrl+"/admin")
			fmt.Println()

			// Listen & serve.
			err = serve.ListenAndServe(environment.GetPort())
			if err != nil {
				printError(err.Error())
			}
		},
	}
)
