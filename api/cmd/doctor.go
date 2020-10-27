package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
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

			return
		},
	}
)

// doctor checks if the environment is validated and checks
// to see if there is a valid database connection and the
// database exists before proceeding.
func doctor() (*database.MySql, error) {

	printSpinner("Running doctor...")

	// Check paths are correct
	if err := paths.BaseCheck(); err != nil {
		printError(err.Error())
		return nil, err
	}

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
			var msg = fmt.Sprintf("Obtaining environment variable: %s", strings.ToUpper(v.Key))
			if v.Type == "ip" {
				msg = fmt.Sprintf("Obtaining environment variable: %s must be a valid IP address", strings.ToUpper(v.Key))
			}
			printError(msg)
		}
		return nil, fmt.Errorf("Validation failed for the enviroment")
	}

	// Get the database and ping
	db, err := database.New()
	if err != nil {
		printError("Establishing database connection, are the credentials in the .env file correct?")
		return nil, fmt.Errorf("Error establishing database connection")
	}

	// Check if the database exists
	if err := db.CheckExists(); err != nil {
		printError("Establishing database connection, are the credentials in the .env file correct?")
		return nil, fmt.Errorf("error establishing database connection")
	}

	// Init Cache
	cache.Init()

	// Init Config
	con, err := config.New()
	if err != nil {
		printError(errors.Message(err))
	}

	// Init logging
	if err := logger.Init(*con); err != nil {
		printError(err.Error())
	}

	printSuccess("All checks passed.")

	return db, nil
}