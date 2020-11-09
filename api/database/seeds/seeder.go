package seeds

import (
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/jmoiron/sqlx"
)

type Seeder struct {
	db *sqlx.DB
	models *models.Store
}

// Construct
func New(db *sqlx.DB, s *models.Store) *Seeder {
	return &Seeder{
		db: db,
		models: s,
	}
}

// Seed
func (s *Seeder) Seed() error {
	// IMPORTANT: Run roles before inserting the user.
	if err := s.runRoles(); err != nil {
		return err
	}
	if err := s.runOptions(); err != nil {
		return err
	}
	//if err := s.runPosts(); err != nil {
	//	return err
	//}
	return nil
}