// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/events"
	"github.com/ainsleyclark/verbis/api/tpl/tplimpl"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			// Run doctor
			cfg, _, err := doctor(true)
			if err != nil {
				printError(err.Error())
			}

			d := deps.New(*cfg)

			// Set dependencies
			d.SetTmpl(tplimpl.New(d))

			dispatch := events.NewResetPassword(d)

			err = dispatch.Dispatch(events.ResetPassword{
				User: domain.UserPart{
					FirstName: "Ainsley",
					LastName:  "Clark",
					Email:     "ainsley@reddico.co.uk",
				},
				Url: "http://127.0.0.1",
			}, []string{"ainsley@reddico.co.uk"}, nil)

			if err != nil {
				color.Red.Println(err)
				return
			}
		},
	}
)
