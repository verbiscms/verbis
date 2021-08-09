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

func (t *SysTestSuite) TestSys_Validate() {
	tt := map[string]struct {
		input domain.InstallVerbis
		step  int
		fn    func(env *environment.Env) (database.Driver, error)
		want  interface{}
	}{
		"Database Success": {
			install,
			InstallDatabaseStep,
			func(env *environment.Env) (database.Driver, error) {
				return &mocks.Driver{}, nil
			},
			nil,
		},
		"Database Validation Failed": {
			installBadValidation,
			InstallDatabaseStep,
			nil,
			"'db_host' failed on the 'required' tag",
		},
		"Database Ping Error": {
			install,
			InstallDatabaseStep,
			func(env *environment.Env) (database.Driver, error) {
				return nil, fmt.Errorf("ping error")
			},
			"ping error",
		},
		"User Success": {
			install,
			InstallUserStep,
			nil,
			nil,
		},
		"User Validation Failed": {
			installBadValidation,
			InstallUserStep,
			func(env *environment.Env) (database.Driver, error) {
				return &mocks.Driver{}, nil
			},
			"'user_first_name' failed on the 'required' tag",
		},
		"Site Success": {
			install,
			InstallSiteStep,
			nil,
			nil,
		},
		"Site Validation Failed": {
			installBadValidation,
			InstallSiteStep,
			func(env *environment.Env) (database.Driver, error) {
				return &mocks.Driver{}, nil
			},
			"'site_title' failed on the 'required' tag",
		},
		"Invalid Step": {
			installBadValidation,
			100,
			nil,
			"invalid step provided: 100",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			if test.fn != nil {
				orig := newDB
				defer func() { newDB = orig }()
				newDB = test.fn
			}
			s := Sys{}

			err := s.ValidateInstall(test.step, test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
