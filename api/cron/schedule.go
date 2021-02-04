// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cron

import (
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
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
	if err := gocron.Every(15).Minutes().Do(s.store.Auth.CleanPasswordResets); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error()
	}

	// Clean sitemap cache
	if err := gocron.Every(6).Hours().Do(render.NewSitemap(s.store).ClearCache); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error()
	}

	// Start all the pending jobs
	<-gocron.Start()
}
