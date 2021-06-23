// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates the Verbis core.",
		Long: `This command will update the current version of Verbis from the 
update server. If the update failed, please rename the .old file the update
generates to roll back`,
		Run: func(cmd *cobra.Command, args []string) {
			config, _, err := doctor(false)
			if err != nil {
				return
			}

			update, err := config.System.Update(false)
			if err != nil {
				printError(err.Error())
			}

			printSuccess("Updated to Version: " + update)
		},
	}
)
