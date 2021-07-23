// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	location "github.com/verbiscms/verbis/api/mocks/services/fields/location"
	categories "github.com/verbiscms/verbis/api/mocks/store/categories"
	users "github.com/verbiscms/verbis/api/mocks/store/users"
	"net/http"
)

func (t *FieldTestSuite) TestFields_List() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(l *location.Finder, u *users.Repository, c *categories.Repository)
		url     string
	}{
		"User": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:   domain.Post{UserId: 123},
					Author: user.HideCredentials(),
				}
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 0).Return(domain.Category{}, fmt.Errorf("error"))
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms?user_id=123",
		},
		"No User": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:   domain.Post{UserId: 123},
					Author: user.HideCredentials(),
				}
				u.On("Owner").Return(user)
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 0).Return(domain.Category{}, fmt.Errorf("error"))
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms",
		},
		"Category": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:     domain.Post{UserId: 123},
					Category: &category,
					Author:   user.HideCredentials(),
				}
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 123).Return(category, nil)
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms?category_id=123&user_id=123",
		},
		"Page Template": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:   domain.Post{UserId: 123, PageTemplate: "template"},
					Author: user.HideCredentials(),
				}
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 0).Return(domain.Category{}, fmt.Errorf("error"))
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms?page_template=template&user_id=123",
		},
		"Page Layout": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:   domain.Post{UserId: 123, PageLayout: "layout"},
					Author: user.HideCredentials(),
				}
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 0).Return(domain.Category{}, fmt.Errorf("error"))
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms?layout=layout&user_id=123",
		},
		"Resource": {
			fieldGroups,
			http.StatusOK,
			"Successfully obtained fields",
			func(l *location.Finder, u *users.Repository, c *categories.Repository) {
				post := domain.PostDatum{
					Post:   domain.Post{UserId: 123, Resource: "resource"},
					Author: user.HideCredentials(),
				}
				u.On("Find", 123).Return(user, nil)
				c.On("Find", 0).Return(domain.Category{}, fmt.Errorf("error"))
				l.On("Layout", mock.Anything, post, false).Return(fieldGroups, nil)
			},
			"/forms?resource=resource&user_id=123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/forms", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).List(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
