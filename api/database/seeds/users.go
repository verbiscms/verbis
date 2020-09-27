package seeds

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// runUsers will insert an owner user into the database when installing
// if api.SuperAdmin is set to true
func (s *Seeder) runUsers() error {
	u := domain.User{
		FirstName:   "Ainsley",
		LastName:    "Clark",
		Email:       "ainsley@reddico.co.uk",
		Password:    "password",
		Role: 		  domain.UserRole{
			Id:       6,
		},
	}

	if _, err := s.models.User.Create(&u); err != nil {
		return err
	}

	return nil
}

