/*
Copyright © 2020 Verbis ainsley@reddico.co.uk

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
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/spf13/cobra"
	"strconv"
)

// Root represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:               "Verbis",
		Short:             "Verbis CLI",
		Long:              `Verbis - CHANGE.`,
		DisableAutoGenTag: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	// Pass the super admin to bool (ldflags)
	admin, _ := strconv.ParseBool(api.SuperAdminString)
	api.SuperAdmin = admin

	// Execute the main command
	if err := rootCmd.Execute(); err != nil {
		printError(fmt.Sprintf("Could not start Verbis: %s", err.Error()))
	}
}

// Add child commands and bootstrap
func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(importCmd)
}
