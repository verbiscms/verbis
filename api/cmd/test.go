// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			env, err := environment.Load()
			if err != nil {
				printError(err.Error())
				return
			}

			client, err := storage.New(env, &domain.Options{
				StorageProvider: domain.StorageAWS,
				StorageBucket:   "reddicotest",
			})

			if err != nil {
				printError(err.Error())
				return
			}


			_, err = client.Find("test.txt")
			if err != nil {
				printError(err.Error())
				return
			}



			//
			//upload, err := client.Upload("test.txt", strings.NewReader("this is a test"))
			//if err != nil {
			//	fmt.Println("her")
			//	printError(err.Error())
			//}
			//
			//fmt.Println(upload)

		},
	}
)
