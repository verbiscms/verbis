// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"testing"
)

func TestMigrator_Migrate(t *testing.T) {
	tt := map[string]struct {
		input Migration
		want  interface{}
	}{
		"Success": {
			Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1")},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			defer func() {
				migrations = make(MigrationRegistry, 0)
			}()
			err := AddMigration(&test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.input, *migrations[0])
		})
	}
}

// CategoriesTestSuite defines the helper used for
// category testing.
type CategoriesTestSuite struct {
	test.DBSuite
}

// TestCategories
//
// Assert testing has begun.
func TestCategories(t *testing.T) {
	suite.Run(t, &CategoriesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock categories database
// for testing.
func (t *CategoriesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) Migrator {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return &migrate{
		down:   nil,
		Driver: MySQLDriver,
		DB:     t.DB,
	}
}

func (t *CategoriesTestSuite) TestMigrator_Migrate() {
	tt := map[string]struct {
		input  MigrationRegistry
		mock   func(m sqlmock.Sqlmock)
		panics bool
		want   interface{}
	}{
		"Simple": {
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "mysql_schema.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("hello").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
			false,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)
			var err error
			if test.panics {
				t.Panics(func() {
					err := m.Migrate(version.Must("v0.0.1"))
					fmt.Println(err)
					t.RunT(err, test.want)
				})
				return
			}
			t.RunT(err, test.want)
		})
	}
}
