package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError(err.Error())
			}

			// Set up stores & pass the database.
			store, err := models.New(db)
			if err != nil {
				printError(err.Error())
			}

			opts, err := store.Options.GetStruct()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(opts)


		},
	}
)
