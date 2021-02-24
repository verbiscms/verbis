// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/sockets"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ghodss/yaml"
	"github.com/radovskyb/watcher"
	"io/ioutil"
	"sync"
	"time"
)

var (
	cfg  = &domain.ThemeConfig{}
	once = sync.Once{}
)

const (
	// The configuration file path within the theme.
	Path = "/config.yml"
)

// Config
//
//
func Config() *domain.ThemeConfig {
	once.Do(func() {
		watch()
		fetch()
	})
	return cfg
}

// Watch
//
//
func watch() {

	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case _ = <-w.Event:
				logger.Info("Updating theme configuration file, sending websocket")
				fetch()
				err := sockets.Broadcast(cfg)
				if err != nil {
					logger.WithError(err).Error()
				}
			case err := <-w.Error:
				logger.Error(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(paths.Theme()); err != nil {
		logger.Error(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	go func() {
		if err := w.Start(time.Millisecond * 100); err != nil {
			logger.Error(err)
		}
	}()
}

// Fetch
//
// Get"s the themes configuration from the themes path
// Returns errors.INTERNAL if the unmarshalling was unsuccessful.
func fetch() *domain.ThemeConfig {
	const op = "SiteRepository.GetThemeConfig"

	var d = defaults()
	theme, err := ioutil.ReadFile(paths.Theme() + Path)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Unable to get retrieve theme config file", Operation: op, Err: err}).Error()
	}

	err = yaml.Unmarshal(theme, &d)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Syntax error in theme config file", Operation: op, Err: err}).Error()
	}

	cfg = &d

	return cfg
}

// defaults
//
// Global Configuration, sets defaults to ensure that there
// are no empty values within the themes config to
// prevent any errors.
func defaults() domain.ThemeConfig {
	return domain.ThemeConfig{
		Theme:         domain.Theme{},
		Resources:     nil,
		AssetsPath:    "assets",
		FileExtension: ".cms",
		TemplateDir:   "templates",
		LayoutDir:     "layouts",
		Admin: domain.AdminConfig{
			Path:                "/admin",
			InactiveSessionTime: 60,
		},
		Media: domain.MediaConfig{
			UploadPath: "/uploads",
			AllowedFileTypes: []string{
				"image/png",
				"image/jpeg",
				"image/gif",
				"image/webp",
				"image/bmp",
				"image/svg+xml",
				"video/mpeg",
				"video/mp4",
				"video/webm",
				"application/pdf",
				"application/msword",
			},
		},
		Editor: domain.Editor{
			Modules: []string{
				"blockquote",
				"code_block",
				"code_block_highlight",
				"hardbreak",
				"h1",
				"h2",
				"h3",
				"h4",
				"h5",
				"h6",
				"paragraph",
				"hr",
				"ul",
				"ol",
				"bold",
				"code",
				"italic",
				"link",
				"strike",
				"underline",
				"history",
				"search",
				"trailing_node",
				"color",
			},
			Options: map[string]interface{}{
				"palette": []string{
					"#4D4D4D", "#999999", "#FFFFFF", "#F44E3B", "#FE9200", "#FCDC00",
					"#DBDF00", "#A4DD00", "#68CCCA", "#73D8FF", "#AEA1FF", "#FDA1FF",
					"#333333", "#808080", "#CCCCCC", "#D33115", "#E27300", "#FCC400",
					"#B0BC00", "#68BC00", "#16A5A5", "#009CE0", "#7B64FF", "#FA28FF",
					"#000000", "#666666", "#B3B3B3", "#9F0500", "#C45100", "#FB9E00",
					"#808900", "#194D33", "#0C797D", "#0062B1", "#653294", "#AB149E",
				},
			},
		},
	}
}
