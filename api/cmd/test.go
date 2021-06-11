// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/update/github"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			git := github.Github{
				RepoURL:     api.Repo,
				ArchiveName: fmt.Sprintf("verbis_%s_%s_%s.zip", version.Version, runtime.GOOS, runtime.GOARCH),
			}

			err := git.Open()
			if err != nil {
				fmt.Println(err)
			}

			//git.Walk(func(info *github.FileInfo) error {
			//	if info.Mode.IsRegular() {
			//		if !strings.Contains(info.Path, "node_modules") {
			//			fmt.Println(info.Path)
			//		}
			//	}
			//
			//	return nil
			//})

			//u := update.New(paths.Get())
			//_, files, err := u.Update()
			//if err != nil {
			//	color.Red.Println(err)
			//	u.RollBack()
			//}
			//fmt.Printf("Updated %d files.\n", files)
			//go func() {
			//	time.Sleep(time.Second * 2)
			//	logger.Info("Restarting Verbis...")
			//	reload.Exec()
			//}()
		},
	}
)
