package cmd

import (
	"github.com/ainsleyclark/verbis/api/database/seeds"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install will first run Verbis doctor and then install the database and insert any data dependant on Verbis.",
		Long:  `This command will install first run Verbis doctor to see if the database,
exists and is passable. Install will then run the migration to insert into the schema.
Seeds are also run, inserting options and any necessary configuration into the 
database.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Start the spinner
			printSpinner("Installing Verbis...")

			// Run doctor
			db, err := doctor()
			if err != nil {
				return
			}

			// Install the database
			if err := db.Install(); err != nil {
				printError(err.Error())
				return
			}

			// Set up stores & pass the database.
			store, err := models.New(db)
			if err != nil {
				printError(err.Error())
				return
			}

			// Run the seeds
			seeder := seeds.New(db.Sqlx, store)
			if err := seeder.Seed(); err != nil {
				printError(err.Error())
				return
			}

			// Print success
			printSuccess("Successfully installed verbis")
		},
	}
)

func init() {

}
