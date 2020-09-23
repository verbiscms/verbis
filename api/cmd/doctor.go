package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/spf13/cobra"
	"strings"
)

var (
	doctorCmd = &cobra.Command{
		Use:   "doctor",
		Short: "Running doctor will check the system for any potential hiccups when installing, updating or running Verbis.",
		Long:  `This command is a diagnostic tool to find any potential issues for your
Verbis install. It will check if the database has been set up correctly as well as the
environment.`,
		Run: func(cmd *cobra.Command, args []string)  {
			if _, err := doctor(); err != nil {
				return
			}
			printSuccess("All checks passed.")
			return
		},
	}
)

// doctor checks if the environment is validated and checks
// to see if there is a valid database connection and the
// database exists before proceeding.
func doctor() (*database.MySql, error) {

	// Load the environment (.env file)
	err := environment.Load()
	if err != nil {
		printError(err.Error())
		return nil, err
	}

	// Check if the environment values are valid
	vErrors := environment.Validate()
	if vErrors != nil {
		for _, v := range vErrors {
			var msg = fmt.Sprintf("-> Error obtaining environment variable: %s", strings.ToUpper(v.Key))
			if v.Type == "ip" {
				msg = fmt.Sprintf("-> Error obtaining environment variable: %s must be a valid IP address", strings.ToUpper(v.Key))
			}
			printError(msg)
		}
		return nil, fmt.Errorf("validation failed for the enviroment")
	}

	// Get the database and ping
	sqlx, err := app.db.GetDatabase()
	if err != nil {
		printError("Error establishing database connection, are the credentials in the .env file correct?")
		return nil, fmt.Errorf("error establishing database connection")
	}

	db := database.MySql{
		Sqlx: sqlx,
	}

	// Check if the database exists
	if err := db.CheckExists(); err != nil {
		fmt.Println("here")
		printError("Error establishing database connection, are the credentials in the .env file correct?")
		return nil, fmt.Errorf("error establishing database connection")
	}

	return &db, nil
}