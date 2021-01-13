package cmd

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/importer/wordpress"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "import",
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
			store := models.New(db, *cfg)
			if err != nil {
				printError(err.Error())
			}

			wp, err := wordpress.New("/Users/ainsley/Desktop/Reddico/websites/reddico-website/theme/res/import-xml/test.xml", store)
			if err != nil {
				printError(err.Error())
			}

			wp.Import()
		},
	}
)
