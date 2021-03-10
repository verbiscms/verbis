// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"database/sql/driver"
	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/mocks/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

// HandlerSuite represents the suite of testing methods for controllers.
type DBSuite struct {
	suite.Suite
	DB         *sqlx.DB
	Driver     database.Driver
	Mock       sqlMock.Sqlmock
	mockDriver *mocks.Driver
}

// Any string for mock string args.
type DBAnyString struct{}

// Match satisfies sqlmock.Argument interface
// for any strings.
func (a DBAnyString) Match(v driver.Value) bool {
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

	db, m, err := sqlMock.New()
	assert.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockDriver := &mocks.Driver{}
	mockDriver.On("DB").Return(sqlxDB)

	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()

	mockDriver.On("Schema").Return("")

	return DBSuite{
		DB:         sqlxDB,
		Driver:     mockDriver,
		Mock:       m,
		mockDriver: mockDriver,
	}
}

// Reset
//
// Sets up a new mock, driver and database upon
// test completion.
func (t *DBSuite) Reset() {
	db := NewDBSuite(t.T())
	t.Driver = db.Driver
	t.DB = db.DB
	t.Mock = db.Mock
	t.mockDriver = db.mockDriver
}

// RunT
//
// Run the DB test.
func (t *DBSuite) RunT(want, actual interface{}, times ...int) {
	if len(times) == 1 {
		t.mockDriver.AssertNumberOfCalls(t.T(), "Schema", times[0])
	}

	t.mockDriver.AssertCalled(t.T(), "Schema")

	err := t.Mock.ExpectationsWereMet()
	if err != nil {
		t.Fail("expectations were not met for mock call: ", err)
	}
	t.Equal(want, actual)
}
