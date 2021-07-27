// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
	"github.com/verbiscms/livereload"
	"github.com/verbiscms/verbis/api/cron"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/http/sockets"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/server"
	"github.com/verbiscms/verbis/api/server/routes"
	"github.com/verbiscms/verbis/api/tpl/tplimpl"
	"github.com/verbiscms/verbis/api/watcher"
	"syscall"
)

type serverCmd struct {
	liveReload bool
}

func newStartCmd() *cobra.Command {
	sc := serverCmd{}

	cmd := &cobra.Command{
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
			emoji.Printf(":backhand_index_pointing_right: Visit your site at:          %s \n", d.Options.SiteUrl)
			emoji.Printf(":key: Or visit the admin area at:  %s \n", d.Options.SiteUrl+"/admin")
			fmt.Println()

			// Load cron jobs
			scheduler := cron.New(d)
			go scheduler.Run()

			w := watcher.New()
			handleFileEvents(w, d, serve, sc)
			defer w.Close()
			defer sockets.Close()

			// Listen & serve.
			err = serve.ListenAndServe(cfg.Env.Port())
			if err != nil {
				printError(err.Error())
			}
		},
	}
	cmd.Flags().BoolVarP(&sc.liveReload, "watch", "w", false, "Enables live reload")
	return cmd
}

func handleFileEvents(w watcher.FileWatcher, d *deps.Deps, s *server.Server, cmd serverCmd) {
	go func() {
		err := w.Watch(d.Paths.Themes, watcher.PollingDuration)
		if err != nil {
			logger.WithError(err).Error()
		}
	}()

	if cmd.liveReload {
		livereload.Initialize()
		s.GET("/livereload.js", gin.WrapF(livereload.ServeJS))
		s.GET("/livereload", gin.WrapF(livereload.Handler))
	}

	go func() {
		for {
			select {
			case event := <-w.Events():
				if event.IsDir() {
					continue
				}

				if !event.IsPath(d.ThemePath()) {
					continue
				}

				if cmd.liveReload {
					livereload.ForceRefresh()
				}

				// Check if it's the configuration file
				//sockets.AdminHub.Broadcast <- sockets.AdminData{Theme: domain.ThemeConfig{
				//	Theme: domain.Theme{
				//		Title: event.Path,
				//	},
				//}}
			case err := <-w.Errors():
				if err.Err != syscall.EPIPE {
					logger.WithError(err).Error()
				}
			}
		}
	}()
}
