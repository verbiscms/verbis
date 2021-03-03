// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HandlerSuite represents the suite of testing methods for controllers.
type DBSuite struct {
	suite.Suite
	DB   *sqlx.DB
	Mock sqlmock.Sqlmock
}

type DBMockResultErr struct{}

func (m DBMockResultErr) LastInsertId() (int64, error) {
	return 0, fmt.Errorf("error")
}

func (m DBMockResultErr) RowsAffected() (int64, error) {
	return 0, fmt.Errorf("error")
}

// NewDBSuite
//
// New recorder for testing
// controllers, initialises gin & sets gin mode.
func NewDBSuite(t *testing.T) DBSuite {
	cache.Init()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	return DBSuite{
		DB:   sqlxDB,
		Mock: mock,
	}
}

// RunT
//
// Run the DB test.
func (t *DBSuite) RunT(want, actual interface{}) {
	err := t.Mock.ExpectationsWereMet()
	if err != nil {
		t.Fail("expectations were not met for mock call: ", err)
	}
	t.Equal(want, actual)
}
