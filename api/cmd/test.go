// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	v0 "github.com/ainsleyclark/verbis/api/database/mysql/migrations/v0"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(v0.Version)

		},
	}
)

func must(err error) {
	color.Red.Println(err)
	os.Exit(1)
}
