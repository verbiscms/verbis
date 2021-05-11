// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			env, err := environment.Load()
			if err != nil {
				fmt.Println(err)
			}

			// Get the database and ping
			db, err := database.New(env)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(db.Tables())

		},
	}
)
