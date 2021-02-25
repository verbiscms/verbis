// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/publisher"
	"github.com/jasonlvhit/gocron"
)

// Scheduler defines all necessary cron jobs for the application
// when running.
type Scheduler struct {
	store *models.Store
}

// New - Construct for a new Scheduler
func New(m *models.Store) *Scheduler {
	return &Scheduler{
		store: m,
	}
}

// Run all cron jobs
func (s *Scheduler) Run() {

	// Clean password resets table every 15 minutes
	err := gocron.Every(15).Minutes().Do(func() {
		logger.Info("Cleaning password resets table")
		err := s.store.Auth.CleanPasswordResets()
		if err != nil {
			logger.WithError(err).Error()
		}
	})
	if err != nil {
		logger.WithError(err).Error()
	}

	// Clean sitemap cache
	err = gocron.Every(6).Hours().Do(func() {
		publisher.NewSitemap(s.store).ClearCache()
	})
	if err != nil {
		logger.WithError(err).Error()
	}

	// Start all the pending jobs
	<-gocron.Start()
}
