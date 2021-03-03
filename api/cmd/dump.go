// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dumps the Verbis database to the storage dumps directory using the database name provided in the .env file.",
		Long: `This command will dump the database to the dumps directory,
located in ./storage/dumps. First the export command runs Verbis doctor to
check if the database exists connection is passable. Then dump the
database to file`,
		Run: func(cmd *cobra.Command, args []string) {
			printSpinner("Dumping database...")

			cfg, db, err := doctor(false)
			if err != nil {
				printError("Could not dump the database, is your database connection valid?")
			}

			time := time.Now().Format(time.RFC3339)
			fileName := fmt.Sprintf("%s-dump-%v", cfg.Env.DbDatabase, time)
			if err := db.Dump(cfg.Paths.Storage+"/dumps", fileName); err != nil {
				printError("Could not dump the database, is your database connection valid?")
			}

			printSuccess(fmt.Sprintf("Successfully exported database to filename: %s", fileName))
		},
	}
)
