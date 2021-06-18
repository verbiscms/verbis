// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/version/updates"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			color.Green.Println(updates.UpdateRegistry)

			for _, v := range updates.UpdateRegistry {
				fmt.Println(v.Version)
				fmt.Println(v.Migration)
			}
		},
	}
)
