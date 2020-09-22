package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"strings"
)

var (
	doctorCmd = &cobra.Command{
		Use:   "doctor",
		Short: "Running verbis doctor will check the system for any potential hiccups when installing or updating Verbis.",
		Long:  `This command is a diagnostic tool to find any potential issues for your
Verbis install. It will check if the database has been set up correctly as well as the
environment.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// TODO:
			// This needs to be a function that returns an error and let cobra handle the errors
			// No lowercase

			// Load the environment (.env file)
			err := environment.Load()
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			// Check if the environment values are valid
			vErrors := environment.Validate()
			if vErrors != nil {
				for _, v := range vErrors {
					var msg = fmt.Sprintf("-> Error obtaining environment variable: %s", strings.ToUpper(v.Key))
					if v.Type == "ip" {
						msg = fmt.Sprintf("-> Error obtaining environment variable: %s must be a valid IP address", strings.ToUpper(v.Key))
					}
					color.Red.Printf(msg)
					fmt.Println()
				}
				return fmt.Errorf("Validation failed for the enviroment")
			}

			// Get the database and ping
			if _, err := app.db.GetDatabase(); err != nil {
				color.Red.Println("-> Error establishing database connection, are the credentials in the .env file correct?")
				return fmt.Errorf("Error establishing database connection")
			}

			// Check if the database exists
			if err := app.db.CheckExists(); err != nil {
				color.Red.Println("-> Error establishing database connection, are the credentials in the .env file correct?")
				return fmt.Errorf("Error establishing database connection")
			}

			return nil
		},
	}
)