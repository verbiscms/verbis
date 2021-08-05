// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	mocks "github.com/verbiscms/verbis/api/mocks/database"
)

func (t *SysTestSuite) TestSys_Preflight() {
	tt := map[string]struct {
		fn   func(env *environment.Env) (database.Driver, error)
		want interface{}
	}{
		"Success": {
			func(env *environment.Env) (database.Driver, error) {
				return &mocks.Driver{}, nil
			},
			nil,
		},
		"Error": {
			func(env *environment.Env) (database.Driver, error) {
				return nil, fmt.Errorf("error")
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			orig := newDB
			defer func() { newDB = orig }()
			newDB = test.fn
			s := Sys{}

			err := s.Preflight(domain.InstallPreflight{})
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
