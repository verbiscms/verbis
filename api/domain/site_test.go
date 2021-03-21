// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResources_Clean(t *testing.T) {
	tt := map[string]struct {
		input Resources
		want  Resources
	}{
		"Cleaned": {
			Resources{
				"news": Resource{
					Name:               "News",
					FriendlyName:       "News",
					SingularName:       "News Item",
					Slug:               "/news////",
					Icon:               "icon",
					Hidden:             false,
					HideCategorySlug:   false,
					AvailableTemplates: []string{"archive"},
				},
			},
			Resources{
				"news": Resource{
					Name:               "News",
					FriendlyName:       "News",
					SingularName:       "News Item",
					Slug:               "news",
					Icon:               "icon",
					Hidden:             false,
					HideCategorySlug:   false,
					AvailableTemplates: []string{"archive"},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Clean()
			assert.Equal(t, test.want, got)
		})
	}
}
