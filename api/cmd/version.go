// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verbiscms/verbis/api/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Obtains the current version ofr Verbis",
		Long: `This command will obtain the current version of verbis that is
installed on the operating system`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Version)
		},
	}
)
