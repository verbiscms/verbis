// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cron"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/server/routes"
	"github.com/ainsleyclark/verbis/api/tpl/tplimpl"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Running start will start Verbis project from the current directory and run the CMS project.",
		Long: `This command will start Verbis from the current directory. First it will
run Verbis doctor to see if the environment is configured correctly. It will then start
up the server on the port specified in the .env file.`,
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			cfg, _, err := doctor(true)
			if err != nil {
				printError(err.Error())
			}

			cfg.Running = true
			d := deps.New(*cfg)

			// Set dependencies
			d.SetTmpl(tplimpl.New(d))

			// Set up the router
			serve := server.New(d)

			// Load the routes
			routes.Load(d, serve)

			// Print listening success
			printSuccess(fmt.Sprintf("Verbis listening on port: %d \n", cfg.Env.Port()))
			emoji.Printf(":backhand_index_pointing_right: Visit your site at:          %s \n", d.Options.SiteURL)
			emoji.Printf(":key: Or visit the admin area at:  %s \n", d.Options.SiteURL+"/admin")
			fmt.Println()

			// Load cron jobs
			scheduler := cron.New(d)
			go scheduler.Run()

			// Listen & serve.
			err = serve.ListenAndServe(cfg.Env.Port())
			if err != nil {
				printError(err.Error())
			}
		},
	}
)
