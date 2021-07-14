// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package environment

import (
	"fmt"
	validation "github.com/ainsleyclark/verbis/api/common/vaidation"
	"github.com/ainsleyclark/verbis/api/errors"
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
	TestDataPath string
	OriginalEnv  []byte
}

// TestEnv asserts testing has begun.
func TestEnv(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

// SetupSuite reassigns API path for testing.
func (t *EnvTestSuite) SetupSuite() {
	logrus.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestDataPath = filepath.Join(wd, "testdata")
}

// ChangePath Assigns a new path to the test suite and
// returns a teardown function to set the original
// back to what is was before testing.
func (t *EnvTestSuite) ChangePath(path string) func() {
	basePathOrig := basePath
	fn := func() {
		basePath = basePathOrig
	}
	basePath = path
	return fn
}

// Original saves the original .env test file in bytes.
func (t *EnvTestSuite) Original() {
	file, err := ioutil.ReadFile(filepath.Join(t.TestDataPath, ".env"))
	if err != nil {
		t.Fail("Error reading test env path")
	}
	t.OriginalEnv = file
}

// Overwrite the original .env test file.
func (t *EnvTestSuite) Overwrite() {
	file, err := os.Create(filepath.Join(t.TestDataPath, ".env"))
	if err != nil {
		t.Fail("Error removing original test env")
	}
	_, err = file.WriteString(string(t.OriginalEnv))
	if err != nil {
		t.Fail("Error creating original test env")
	}
}

func (t *EnvTestSuite) TestLoad() {
	tt := map[string]struct {
		path string
		want interface{}
	}{
		"Success": {
			t.TestDataPath,
			&Env{
				AppEnv:          "dev",
				AppDebug:        "true",
				AppPort:         "8080",
				DbHost:          "127.0.0.1",
				DbPort:          "3306",
				DbDatabase:      "verbis",
				DbUser:          "root",
				DbPassword:      "password",
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
				DbDriver:   "mysql",
				DbHost:     "127.0.0.1",
				DbPort:     "3306",
				DbDatabase: "verbis",
				DbUser:     "root",
				DbPassword: "password",
			},
			ve,
		},
		"Bad Validation": {
			Env{},
			validation.Errors{
				validation.Error{Key: "app_port", Type: "required", Message: "App Port is required."},
				validation.Error{Key: "db_driver", Type: "required", Message: "Db Driver is required."},
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

func (t *EnvTestSuite) TestEnv_Set() {
	tt := map[string]struct {
		key   string
		value interface{}
		path  string
		want  interface{}
	}{
		"Success": {
			"key",
			"value",
			t.TestDataPath,
			nil,
		},
		"Bad Value": {
			"",
			make(chan int),
			"/test/testdata/env",
			"Error casting value to string",
		},
		"Error": {
			"",
			"",
			"wrongpath",
			"Error reading env file with the path",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			teardown := t.ChangePath(test.path)
			defer teardown()

			t.Original()
			defer t.Overwrite()

			env := &Env{}
			err := env.Set(test.key, test.value)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Nil(err)
		})
	}
}

func (t *EnvTestSuite) TestEnv_SetError() {
	teardown := t.ChangePath(t.TestDataPath)
	defer teardown()

	orig := write
	defer func() { write = orig }()
	write = func(envMap map[string]string, filename string) error {
		return fmt.Errorf("error")
	}

	env := &Env{}
	err := env.Set("key", "value")
	if err == nil {
		t.Fail("error should not be nil")
		return
	}
	t.Contains(errors.Message(err), "Error writing env file with the path")
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
