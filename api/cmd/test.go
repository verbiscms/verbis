// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/version/updates"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			env, err := environment.Load()
			if err != nil {
				must(err)
			}

			db, err := database.New(env)
			if err != nil {
				must(err)
			}

			u, err := updates.New(db)
			if err != nil {
				must(err)
			}

			err = u.Run()
			if err != nil {
				must(err)
			}

			color.Green.Println("Updated successfully")
		},
	}
)

func must(err error) {
	color.Red.Println(err)
	os.Exit(1)
}
