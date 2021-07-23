// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
	"github.com/verbiscms/verbis/api/version"
	"path/filepath"
)

func (t *InternalTestSuite) TestNewMigrator() {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"MySQL": {
			MySQLDriver,
			nil,
		},
		"Postgres": {
			PostgresDriver,
			nil,
		},
		"Wrong Driver": {
			"wrong",
			"wrong driver",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			db := &sqlx.DB{}
			_, err := NewMigrator(test.input, db)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

func (t *InternalTestSuite) TestMigrator_Migrate() {
	tt := map[string]struct {
		version *sm.Version
		input   MigrationRegistry
		mock    func(m sqlmock.Sqlmock)
		panics  bool
		want    interface{}
	}{
		"Simple": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
			false,
			nil,
		},
		"Major Versions Mismatch": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql")},
				&Migration{Version: "v1.0.0", Stage: version.Major, SemVer: version.Must("v1.0.0"), SQLPath: filepath.Join("v1", "v1.0.0.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
			false,
			nil,
		},
		"Can't Migrate": {
			sm.Must(sm.NewVersion("v1.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.0", Stage: version.Major, SemVer: version.Must("v0.0.0"), SQLPath: filepath.Join("v0", "v0.0.0.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectCommit()
			},
			false,
			nil,
		},
		"CallBack": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql"), CallBackUp: func() error {
					return nil
				}, CallBackDown: func() error {
					return nil
				}},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
			false,
			nil,
		},
		"Begin Error": {
			sm.Must(sm.NewVersion("v0.0.2")),
			MigrationRegistry{},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin().
					WillReturnError(fmt.Errorf("error"))
			},
			false,
			fmt.Errorf("error"),
		},
		"RollBack Error": {
			sm.Must(sm.NewVersion("v0.0.1")),
			MigrationRegistry{
				&Migration{Version: "v0.0.0", Stage: version.Patch, SemVer: version.Must("v0.0.0"), SQLPath: filepath.Join("v0", "v0.0.0.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnError(fmt.Errorf("error"))
				m.ExpectRollback().
					WillReturnError(fmt.Errorf("error"))
			},
			true,
			nil,
		},
		"Commit Error": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql")},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit().
					WillReturnError(fmt.Errorf("error"))
			},
			true,
			nil,
		},
		"CallBackUp Error": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql"), CallBackUp: func() error {
					return fmt.Errorf("error")
				}, CallBackDown: func() error {
					return nil
				}},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectRollback()
			},
			false,
			fmt.Errorf("error"),
		},
		"CallBackDown Error": {
			sm.Must(sm.NewVersion("v0.0.0")),
			MigrationRegistry{
				&Migration{Version: "v0.0.1", Stage: version.Major, SemVer: version.Must("v0.0.1"), SQLPath: filepath.Join("v0", "v0.0.0.sql"), CallBackUp: func() error {
					return nil
				}, CallBackDown: func() error {
					return fmt.Errorf("error")
				}},
				&Migration{Version: "v0.0.2", Stage: version.Major, SemVer: version.Must("v0.0.2"), SQLPath: filepath.Join("v0", "v0.0.1.sql"), CallBackUp: func() error {
					return fmt.Errorf("error")
				}, CallBackDown: func() error {
					return nil
				}},
			},
			func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec("UPDATE posts SET title = 'v0.0.0' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec("UPDATE posts SET title = 'v0.0.1' WHERE id = 1;").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectRollback()
			},
			true,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)

			migrations = test.input
			defer func() {
				migrations = make(MigrationRegistry, 0)
			}()

			if test.panics {
				t.Panics(func() {
					m.Migrate(test.version) //nolint
				})
				return
			}

			err := m.Migrate(test.version)
			t.RunExpectationsT(err, test.want)
		})
	}
}

func (t *InternalTestSuite) TestMigrator_CanMigrate() {
	tt := map[string]struct {
		base    *sm.Version
		migrate *sm.Version
		want    interface{}
	}{
		"Can Migrate": {
			sm.Must(sm.NewVersion("v0.0.2")),
			sm.Must(sm.NewVersion("v0.0.5")),
			true,
		},
		"Can Migrate Major": {
			sm.Must(sm.NewVersion("v1.0.0")),
			sm.Must(sm.NewVersion("v1.0.1")),
			true,
		},
		"Major Version": {
			sm.Must(sm.NewVersion("v0.0.1")),
			sm.Must(sm.NewVersion("v1.0.2")),
			false,
		},
		"Major Version Zero": {
			sm.Must(sm.NewVersion("v0.0.1")),
			sm.Must(sm.NewVersion("v1.0.0")),
			false,
		},
		"Equal": {
			sm.Must(sm.NewVersion("v0.0.9")),
			sm.Must(sm.NewVersion("v0.0.9")),
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := migrate{}
			canMigrate, err := m.canMigrate(test.base, test.migrate)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, canMigrate)
		})
	}
}
