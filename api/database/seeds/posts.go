package seeds

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// runPosts will insert all demo psots for the user.
func (s *Seeder) runPosts() error {

	p := domain.PostCreate{
		Post: domain.Post{
			Slug:         "/",
			Title:        "Welcome to Verbis",
			Status:       "published",
			PageTemplate: "",
			PageLayout:   "",
			UserId:       0,
		},
		Author:   0,
		Category: nil,
	}

	_, err := s.models.Posts.Create(&p)
	if err != nil {
		return err
	}

	return nil
}
