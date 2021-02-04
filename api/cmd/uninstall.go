// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "Use with caution! This will drop the Verbis database entirely, it should only be used in development.",
		Long: `This command will drop the database that is provided in the environment file.
It will first run Verbis doctor to ensure the database connection is passable and then 
continue to drop the database. Use with caution!`,
		Run: func(cmd *cobra.Command, args []string) {

			// Message
			fmt.Println()
			fmt.Println("This will drop the entire Verbis database, are you sure you want to continue?")
			fmt.Println("Are you sure you want to continue? (yes/no)")
			fmt.Println()

			// Check for user input
			reader := bufio.NewReader(os.Stdin)
			for {

				fmt.Print("-> ")
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)

				if strings.Contains(text, "yes") {
					break
				} else if strings.Contains(text, "no") {
					fmt.Println()
					fmt.Println("Bye!")
					os.Exit(0)
				} else {
					printErrorNoExit("Please enter yes or no")
					fmt.Println()
				}
			}

			// Start the spinner
			printSpinner("Uninstalling Verbis...")

			// Run doctor
			db, err := doctor()
			if err != nil {
				printError(err.Error())
			}

			// Drop the database
			if err := db.Drop(); err != nil {
				printError(err.Error())
			}

			// Print success
			printSuccess("Successfully uninstalled verbis")

			return
		},
	}
)

func init() {

}
