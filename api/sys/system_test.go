// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/database"
	"io/ioutil"
	"testing"
)

// SysTestSuite defines the helper used for system
// testing.
type SysTestSuite struct {
	suite.Suite
}

// TestSys asserts testing has begun.
func TestSys(t *testing.T) {
	suite.Run(t, new(SysTestSuite))
}

// SetupSuite discards the logger
func (t *SysTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
}

var (
	// The default install verbis used for testing.
	install = domain.InstallVerbis{
		InstallDatabase: domain.InstallDatabase{
			DBHost:     "host",
			DBPort:     "port",
			DBDatabase: "database",
			DBUser:     "user",
			DBPassword: "password",
		},
		InstallUser: domain.InstallUser{
			UserFirstName:       "verbis",
			UserLastName:        "cms",
			UserEmail:           "hello@verbiscms.com",
			UserPassword:        "password",
			UserConfirmPassword: "password",
		},
		InstallSite: domain.InstallSite{
			SiteTitle:           "title",
			SiteURL:             "http://127.0.0.1",
			Robots:              false,
		},
	}
	// The default install verbis with wrong validation
	// used for testing.
	installBadValidation = domain.InstallVerbis{
		InstallDatabase: domain.InstallDatabase{
			DBPort:     "port",
			DBDatabase: "database",
			DBUser:     "user",
			DBPassword: "password",
		},
		InstallUser: domain.InstallUser{
			UserLastName:        "cms",
			UserEmail:           "hello@verbiscms.com",
			UserPassword:        "password",
			UserConfirmPassword: "password",
		},
		InstallSite: domain.InstallSite{
			SiteURL:             "http://127.0.0.1",
			Robots:              false,
		},
	}
)

func (t *SysTestSuite) TestNew() {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		exec  func() (string, error)
		bin   string
		panic bool
		want  interface{}
	}{
		"Success": {
			func() (s string, err error) {
				return "exec", nil
			},
			"test",
			false,
			"exec",
		},
		"Error": {
			func() (s string, err error) {
				return "", fmt.Errorf("error")
			},
			"test",
			true,
			"cannot get path to binary",
		},
		"Absolute": {
			func() (s string, err error) {
				return "exec", nil
			},
			"/test",
			false,
			"/test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			if test.exec == nil {
				t.Fail("exec function cannot be nil")
				return
			}

			origExec := exec
			origBin := bin

			defer func() {
				bin = origBin
				exec = origExec
			}()

			exec = test.exec
			bin = test.bin

			if test.panic {
				t.Panics(func() {
					New(&mocks.Driver{}, true)
				})
				return
			}

			got := New(&mocks.Driver{}, true)
			t.Equal(test.want, got.ExecutablePath)
		})
	}
}
