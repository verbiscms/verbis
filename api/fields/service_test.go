package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
)

func (t *FieldTestSuite) TestNewService() {
	m := &models.Store{}
	var l []domain.FieldGroup
	var f []domain.PostField

	pd := domain.PostData{
		Post:     domain.Post{
			Id: 1,
		},
		Layout:   &l,
		Fields:   &f,
	}

	service := &Service{
		store:  m,
		postId: 1,
		fields: f,
		layout: l,
	}

	t.Equal(NewService(m, pd), service)
}