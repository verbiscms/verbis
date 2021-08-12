// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"io/ioutil"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input  *domain.Options
		panics bool
	}{
		"Success": {
			&domain.Options{NavMenus: map[string]interface{}{}},
			false,
		},
		"Panics": {
			&domain.Options{NavMenus: map[string]interface{}{"test": make(chan bool)}},
			true,
		},
	}

	for _, test := range tt {
		t.Run(name, func(t *testing.T) {
			d := &deps.Deps{Options: test.input}
			td := &internal.TemplateDeps{}
			if test.panics {
				assert.Panics(t, func() {
					Init(d, td)
				})
				return
			}
			ns := Init(d, td)

			assert.Equal(t, ns.Name, name)
			assert.NotNil(t, ns.Context())
		})
	}
}
