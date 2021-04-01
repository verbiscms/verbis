// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			// Run doctor
			cfg, _, err := doctor(false)
			if err != nil {
				printError(err.Error())
			}
			d := deps.New(*cfg)

			color.Green.Println(d.WebP.Install())

			//p := paths.Get()
			//err := os.Rename(p.Base+"/verbis", p.Base+"/verbis.bak")
			//fmt.Println(err)

			// download new version
			// unpack exec
			// platform check
			// download the zip file from git

			// https://github.com/ainsleyclark/TryVerbis/archive/refs/tags/0.0.1.zip

			// get list of available tags

			// curl https://api.github.com/repos/ainsleyclark/TryVerbis/releases
		},
	}
)
