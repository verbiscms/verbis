package cron

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
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
	const op = "Scheduler.Run"

	// Clean password resets table every 15 minutes
	if err := gocron.Every(15).Minutes().Do(s.store.Auth.CleanPasswordResets); err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Could not set up the password resets cron", Operation: "op", Err: err},
		}).Error()
	}

	// Start all the pending jobs
	<- gocron.Start()
}


