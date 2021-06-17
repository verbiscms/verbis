// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _, err := doctor(false)
			if err != nil {
				printError(err.Error())
				os.Exit(1)
			}

			d := deps.New(*cfg)

			fmt.Println(version.Version)

			update, err := d.Updater.Update()
			if err != nil {
				color.Blue.Println(update, err)
			}
			fmt.Println(update)
		},
	}
)
