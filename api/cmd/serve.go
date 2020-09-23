package cmd

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/cron"
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
				return
			}

			// Init logging
			err = logger.Init()
			if err != nil {
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
				return
			}

			// Load cron jobs
			scheduler := cron.New(store)
			go scheduler.Run()

			// Load app
			app = App{
				db: db,
				store: store,
			}

			// Set up the router
			serve, err := server.New()
			if err != nil {
				printError(err.Error())
				return
			}

			// Pass the stores to the controllers
			controllers, err := controllers.New(app.store)
			if err != nil {
				printError(err.Error())
				return
			}

			// Load the routes
			routes.Load(serve, controllers, app.store)

			// Listen & serve.
			err = serve.ListenAndServe(8080)
			if err != nil {
				printError(err.Error())
			}
		},
	}
)


// Bootstrap the application
func bootstrap() {


}
