package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/spf13/cobra"
	"time"
)

var (
	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dumps the Verbis database to the storage dumps directory using the database name provided in the .env file.",
		Long:  `This command will dump the database to the dumps directory,
located in ./storage/dumps. First the export command runs Verbis doctor to
check if the database exists connection is passable. Then dump the
database to file`,
		Run: func(cmd *cobra.Command, args []string) {

			// Start the spinner
			printSpinner("Dumping database...")

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError("Could not dump the database, is your database connection valid?")
				return
			}

			// Dump the database
			time := time.Now().Format(time.RFC3339)
			fileName := fmt.Sprintf("%s-dump-%v", environment.GetDatabaseName(), time)
			if err := db.Export(paths.Storage() + "/dumps", fileName); err != nil {
				printError("Could not dump the database, is your database connection valid?")
				return
			}

			// Print success
			printSuccess(fmt.Sprintf("Successfully exported database to filename: %s", fileName))
		},
	}
)
