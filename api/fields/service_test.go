package fields

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
)

func (t *FieldTestSuite) TestNewService() {
	m := &models.Store{}

	var l = make([]domain.FieldGroup, 0)
	var f = make([]domain.PostField, 0)

	pd := &domain.PostData{
		Post: domain.Post{
			Id: 1,
		},
		Layout: l,
		Fields: f,
	}

	deps := &deps.Deps{
		Store:   m,
		Config:  config.Configuration{},
	}

	service := &Service{
		deps:   deps,
		postId: 1,
		fields: f,
		layout: l,
	}

	t.Equal(NewService(deps, pd), service)
}
