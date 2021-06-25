// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/categories"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/fields"
	users "github.com/ainsleyclark/verbis/api/mocks/store/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *FieldTestSuite) TestForms_Create() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository, u *users.Repository, c *categories.Repository)
	}{
		"No User": {
			fieldGroups,
			http.StatusOK,
			"Successfully created form with ID: 123",
			fieldGroups,
			func(m *mocks.Repository, u *users.Repository, c *categories.Repository) {
				m.On("Create", &domain.PostDatum{}).Return(fieldGroups, nil)
				u.On()
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/forms", "/forms", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).List(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
