// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/update"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			u := update.New(paths.Get())
			files, err := u.Update()
			if err != nil {
				color.Red.Println(err)
				u.RollBack()
			}
			fmt.Printf("Updated %d files.\n", files)
			//go func() {
			//	time.Sleep(time.Second * 2)
			//	logger.Info("Restarting Verbis...")
			//	reload.Exec()
			//}()
		},
	}
)
