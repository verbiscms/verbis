// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
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

type AnyUUID struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyUUID) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

// NewDBSuite
//
// New recorder for testing
// controllers, initialises gin & sets gin mode.
func NewDBSuite(t *testing.T) DBSuite {
	cache.Init()
	logger.SetOutput(ioutil.Discard)

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
	defer func() {
		db, mock, err := sqlmock.New()
		t.NoError(err)
		t.DB = sqlx.NewDb(db, "sqlmock")
		t.Mock = mock
	}()
	err := t.Mock.ExpectationsWereMet()
	if err != nil {
		t.Fail("expectations were not met for mock call: ", err)
	}
	t.Equal(want, actual)
}
