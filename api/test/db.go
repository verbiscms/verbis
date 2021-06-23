// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"bytes"
	"database/sql/driver"
	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/database/builder"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/mocks/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HandlerSuite represents the suite of testing methods for controllers.
type DBSuite struct {
	suite.Suite
	DB         *sqlx.DB
	Driver     *mocks.Driver
	Mock       sqlMock.Sqlmock
	Logger     *bytes.Buffer
	mockDriver *mocks.Driver
}

// Any for mock args.
type DBAny struct{}

// Match satisfies sqlmock.Argument interface
// for any arg.
func (a DBAny) Match(v driver.Value) bool {
	return true
}

// Any string for mock string args.
type DBAnyString struct{}

// Match satisfies sqlmock.Argument interface
// for any strings.
func (a DBAnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

// Any string for mock string args.
type DBAnyJSONMessage struct{}

// Match satisfies sqlmock.Argument interface
// for any json raw messages.
func (a DBAnyJSONMessage) Match(v driver.Value) bool {
	_, ok := v.([]byte)
	return ok
}

// NewDBSuite
//
// New recorder for testing
// controllers, initialises gin & sets gin mode.
func NewDBSuite(t *testing.T) DBSuite {
	cache.Init()

	buf := &bytes.Buffer{}
	logger.SetOutput(buf)

	db, m, err := sqlMock.New()
	assert.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockDriver := &mocks.Driver{}
	mockDriver.On("DB").Return(sqlxDB)

	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()
	mockDriver.On("Builder").Return(builder.New("mysql")).Once()

	mockDriver.On("Schema").Return("")

	return DBSuite{
		DB:         sqlxDB,
		Driver:     mockDriver,
		Mock:       m,
		Logger:     buf,
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
	t.Logger = db.Logger
	t.mockDriver = db.mockDriver
}

// RunT runs the the DB test for the store with Schema
// being tested.
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

// RunExpectationsT runs the DB test and asserts the
// expectations were met.
func (t *DBSuite) RunExpectationsT(want, actual interface{}) {
	err := t.Mock.ExpectationsWereMet()
	if err != nil {
		t.Fail("expectations were not met for mock call: ", err)
	}
	t.Equal(want, actual)
}
