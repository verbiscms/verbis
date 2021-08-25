// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/deps"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
)

func (t *FieldTestSuite) TestService_GetCacheField() {
	tt := map[string]struct {
		name  string
		key   string
		cache func(c *cache.Store)
		found bool
		want  interface{}
	}{
		"Found": {
			"name",
			"field1",
			func(c *cache.Store) {
				c.On("Get", mock.Anything, "field-0-0-name-field1").
					Return("test", nil)
			},
			true,
			"test",
		},
		"Not Found": {
			"name",
			"field1",
			func(c *cache.Store) {
				c.On("Get", mock.Anything, "field-0-0-name-field1").
					Return(nil, fmt.Errorf("error"))
			},
			false,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := &cache.Store{}
			if test.cache != nil {
				test.cache(c)
			}
			s := Service{
				deps: &deps.Deps{Cache: c},
			}
			got, ok := s.getCacheField(test.name, test.key, 0)
			t.Equal(test.found, ok)
			t.Equal(test.want, got)
		})
	}
}
