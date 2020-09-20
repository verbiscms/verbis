package cron

import (
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

type Scheduler struct {
	store *models.Store
}

//Construct
func New(m *models.Store) *Scheduler {
	return &Scheduler{
		store: m,
	}
}

// Run all cron jobs
func (s *Scheduler) Run() {

	// Clean password resets table every minute
	if err := gocron.Every(5).Seconds().Do(s.store.Auth.CleanPasswordResets); err != nil {
		log.Error(err)
	}

	//_, time := gocron.NextRun()
	//fmt.Println(time)

	// Start all the pending jobs
	// TODO: Ask Kirk, not working! Blocking other operations.
	//<- gocron.Start()
}


