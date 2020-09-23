/*
Copyright © 2020 NAME HERE ainsley@reddico.co.uk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
)

// TODO: Change db and store to local variables
type App struct {
	db *database.MySql
	store *models.Store
}

// Root represents the base command when called without any subcommands
var (
	app App
	rootCmd = &cobra.Command{
		Use:   "Verbis",
		Short: "Verbis CLI",
		Long: `Verbis - CHANGE.`,
		DisableAutoGenTag: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		printError(err.Error())
	}
}

// Add child commands and bootstrap
func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(testCmd)
}

