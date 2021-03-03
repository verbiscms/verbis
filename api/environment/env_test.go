// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package environment

import (
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// EnvTestSuite defines the helper used for environment
// testing.
type EnvTestSuite struct {
	suite.Suite
	apiPath string
}

// TestEnv
//
// Assert testing has begun.
func TestEnv(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

// SetupSuite
//
// Reassign API path for testing.
func (t *EnvTestSuite) SetupSuite() {
	logrus.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "")
}

// ChangePath
//
// Assigns a new path to the test suite and returns a
// teardown function to set the original back to
// what is was before testing.
func (t *EnvTestSuite) ChangePath(path string) func() {
	basePathOrig := basePath
	fn := func() {
		basePath = basePathOrig
	}
	basePath = t.apiPath + path
	return fn
}

func (t *EnvTestSuite) TestLoad() {
	tt := map[string]struct {
		path string
		want interface{}
	}{
		"Success": {
			"/test/testdata/env",
			&Env{
				AppEnv:          "dev",
				AppDebug:        "true",
				AppPort:         "8080",
				DBHost:          "127.0.0.1",
				DBPort:          "3306",
				DBDatabase:      "verbis",
				DBUser:          "root",
				DBPassword:      "password",
				SparkpostAPIKey: "key",
				SparkpostURL:    "url",
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "Verbis",
			},
		},
		"Error": {
			"wrongpath",
			"no such file or directory",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			teardown := t.ChangePath(test.path)
			defer teardown()

			got, err := Load()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *EnvTestSuite) TestEnv_Validate() {
	var ve validation.Errors

	tt := map[string]struct {
		input Env
		want  interface{}
	}{
		"No Errors": {
			Env{
				AppPort:    "8080",
				DBHost:     "127.0.0.1",
				DBPort:     "3306",
				DBDatabase: "verbis",
				DBUser:     "root",
				DBPassword: "password",
			},
			ve,
		},
		"Bad Validation": {
			Env{},
			validation.Errors{
				validation.Error{Key: "app_port", Type: "required", Message: "App Port is required."},
				validation.Error{Key: "db_host", Type: "required", Message: "Db Host is required."},
				validation.Error{Key: "db_port", Type: "required", Message: "Db Port is required."},
				validation.Error{Key: "db_database", Type: "required", Message: "Db Database is required."},
				validation.Error{Key: "db", Type: "required", Message: "Db is required."},
				validation.Error{Key: "db_password", Type: "required", Message: "Db Password is required."},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := test.input.Validate()
			t.Equal(test.want, got)
		})
	}
}

func (t *EnvTestSuite) TestEnv_Port() {
	tt := map[string]struct {
		port string
		want int
	}{
		"Success": {
			"8000",
			8000,
		},
		"Error": {
			"prod",
			5000,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			e := Env{AppPort: test.port}
			t.Equal(test.want, e.Port())
		})
	}
}

func (t *EnvTestSuite) TestEnv_ConnectString() {
	e := Env{
		DBHost:     "127.0.0.1",
		DBPort:     "3000",
		DBDatabase: "verbis",
		DBUser:     "verbis",
		DBPassword: "password",
	}
	want := "verbis:password@tcp(127.0.0.1:3000)/verbis?tls=false&parseTime=true&multiStatements=true"
	t.Equal(want, e.ConnectString())
}

func (t *EnvTestSuite) TestEnv_MailConfig() {
	e := Env{
		SparkpostAPIKey: "key",
		SparkpostURL:    "sparkpost",
		MailFromAddress: "hello@verbiscms.com",
		MailFromName:    "verbis",
	}
	want := Mail{
		SparkpostAPIKey: "key",
		SparkpostURL:    "sparkpost",
		FromAddress:     "hello@verbiscms.com",
		FromName:        "verbis",
	}
	t.Equal(want, e.MailConfig())
}

func (t *EnvTestSuite) TestEnv_IsProduction() {
	tt := map[string]struct {
		env  string
		want bool
	}{
		"Production": {
			"production",
			true,
		},
		"Prod": {
			"prod",
			true,
		},
		"Dev": {
			"dev",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			e := Env{AppEnv: test.env}
			t.Equal(test.want, e.IsProduction())
		})
	}
}

func (t *EnvTestSuite) TestEnv_IsDebug() {
	tt := map[string]struct {
		debug string
		want  bool
	}{
		"Debug": {
			"true",
			true,
		},
		"Not Debug": {
			"false",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			e := Env{AppDebug: test.debug}
			t.Equal(test.want, e.IsDebug())
		})
	}
}
