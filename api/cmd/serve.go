package cmd

import (
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/routes"
	"github.com/ainsleyclark/verbis/api/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve the CMS.",
		Long:  `Serve will serve the system dependant on port number passed.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			bootstrap()

			// Set up the router
			serve, err := server.New()
			if err != nil {
				log.Panic(err)
			}

			// Pass the stores to the controllers
			controllers, err := controllers.New(app.store)
			if err != nil {
				log.Panic(err)
			}

			// Load the routes
			routes.Load(serve, controllers, app.store)

			// Listen & serve.
			err = serve.ListenAndServe(8080)
			if err != nil {
				return err
			}
			return nil
		},
	}
)

