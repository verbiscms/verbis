// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			//// Run doctor
			//cfg, _, err := doctor(true)
			//if err != nil {
			//	printError(err.Error())
			//}
			//
			//d, err := deps.New(*cfg)
			//if err != nil {
			//	printError(err.Error())
			//}
		},
	}
)
