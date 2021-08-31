// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/services/media"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			// Run doctor
			cfg, _, err := doctor(true)
			if err != nil {
				printError(err.Error())
			}

			d, err := deps.New(*cfg)
			if err != nil {
				printError(err.Error())
			}

			m := media.New(d.Store.Media, d.Storage, d.Store.Options, d.Theme)

			p, err := m.ReGenerateWebP()
			if err != nil {
				return
			}

			color.Green.Println(p)
		},
	}
)
