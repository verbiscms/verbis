// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ainsleyclark/verbis/api/watchers"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			w := watchers.New("/Users/ainsley/Desktop/Reddico/apis/verbis/themes/Verbis")

			go func() {
				for {
					select {
					case event := <-w.Event:
						color.Green.Println(event.Path)
						color.Green.Println(event.Extension)
						color.Green.Println(event.Mime)
					case err := <-w.Error:
						color.Red.Println(err)
					}
				}
			}()

			w.Start()
		},
	}
)
