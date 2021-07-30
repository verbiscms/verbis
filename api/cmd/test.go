// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/logger"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			err := filepath.Walk(paths.Get().Storage, func(path string, info fs.FileInfo, err error) error {
				ext := filepath.Ext(path)
				if ext == ".webp" {
					err := os.Remove(path)
					if err != nil {
						logger.Trace(err)
					}
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)
