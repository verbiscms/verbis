// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
		Use:   "Verbis",
		Short: "Verbis CLI",
		//Long:              `Verbis - CHANGE.`,
		DisableAutoGenTag: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	// Pass the super admin to bool (ldflags)
	admin, _ := strconv.ParseBool(api.ProductionString)
	api.Production = admin

	// Execute the main command
	if err := rootCmd.Execute(); err != nil {
		printError(fmt.Sprintf("Could not start Verbis: %s", err.Error()))
	}
}

// Add child commands and bootstrap
func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(importCmd)

	// Test routes
	if !api.Production {
		rootCmd.AddCommand(testCmd)
	}
}
