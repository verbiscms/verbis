// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/publisher"
)

// Scheduler defines all necessary cron jobs for the application
// when running.
type Scheduler struct {
	*deps.Deps
}

// New - Construct for a new Scheduler
func New(d *deps.Deps) *Scheduler {
	return &Scheduler{
		Deps: d,
	}
}

// Run all cron jobs
func (s *Scheduler) Run() {
	// Clean password resets table every 15 minutes
	err := gocron.Every(15).Minutes().Do(func() { //nolint
		logger.Info("Cleaning password resets table")
		err := s.Store.Auth.CleanPasswordResets()
		if err != nil {
			logger.WithError(err).Error()
		}
	})
	if err != nil {
		logger.WithError(err).Error()
	}

	// Clean sitemap cache
	err = gocron.Every(6).Hours().Do(func() { //nolint
		logger.Info("Clearing sitemap cache")
		publisher.NewSitemap(s.Deps).ClearCache()
	})
	if err != nil {
		logger.WithError(err).Error()
	}

	// Start all the pending jobs
	<-gocron.Start()
}
