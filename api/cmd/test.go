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
			//env, err := environment.Load()
			//if err != nil {
			//	color.Red.Println(err)
			//}
			//
			//driver := database.TestPostgres(env)
			//err = driver.Ping()
			//if err != nil {
			//	color.Red.Println(err)
			//}
			//
			//store := categories.New(&models.StoreCfgOld{
			//	DB: driver,
			//})
			//
			//cat, err := store.Find(2)
			//if err != nil {
			//	color.Red.Println(err)
			//}
			//
			//color.Green.Println(cat)
		},
	}
)
