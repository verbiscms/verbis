// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
)

func (t *FieldTestSuite) TestNewService() {
	m := &models.Store{}

	var l = make(domain.FieldGroups, 0)
	var f = make(domain.PostFields, 0)

	pd := &domain.PostDatum{
		Post: domain.Post{
			Id: 1,
		},
		Layout: l,
		Fields: f,
	}

	deps := &deps.Deps{
		Store:  m,
		Config: &config.Configuration{},
	}

	service := &Service{
		deps:   deps,
		postId: 1,
		fields: f,
		layout: l,
	}

	t.Equal(NewService(deps, pd), service)
}
