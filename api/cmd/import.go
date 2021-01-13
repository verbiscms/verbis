package cmd

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "Import",
		Short: "Import from Wordpress",
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
			_ = models.New(db, *cfg)
			if err != nil {
				printError(err.Error())
			}
		},
	}
)
