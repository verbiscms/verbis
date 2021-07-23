// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/cache"
	"testing"
)

const (
	CacheKey = "key"
)

// CacheTestSuite defines the helper used for cache
// testing.
type CacheTestSuite struct {
	suite.Suite
	LogWriter bytes.Buffer
}

// TestCache asserts testing has begun.
func TestCache(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}

// BeforeTest assign the logger to a buffer.
func (t *CacheTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(&t.LogWriter)
	logger.SetLevel(logrus.TraceLevel)
}

// Reset the log writer.
func (t *CacheTestSuite) Reset() {
	t.LogWriter.Reset()
}

// Setup assigns a mock Cacher to c.
func (t *CacheTestSuite) Setup(mf func(m *mocks.Cacher)) {
	m := &mocks.Cacher{}
	if mf != nil {
		mf(m)
	}
	c = m
}

func (t *CacheTestSuite) TestLoad() {
	tt := map[string]struct {
		input *environment.Env
		want  interface{}
	}{
		"Nil Env": {
			nil,
			"nil environment",
		},
		"Invalid Driver": {
			&environment.Env{CacheDriver: "wrong"},
			ErrInvalidDriver.Error(),
		},
		"Memory": {
			&environment.Env{CacheDriver: MemoryStore},
			MemoryStore,
		},
		"Redis": {
			&environment.Env{CacheDriver: RedisStore},
			RedisStore,
		},
		"Memcached": {
			&environment.Env{CacheDriver: MemcachedStore},
			MemcachedStore,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer func() { c = nil }()
			err := Load(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			if c == nil {
				t.Fail("nil Driver")
			}
			t.Equal(test.want, Driver)
		})
	}
}

func (t *CacheTestSuite) TestGet() {
	tt := map[string]struct {
		mock func(m *mocks.Cacher)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Cacher) {
				m.On("Get", mock.Anything, mock.Anything).Return("item", nil)
			},
			"item",
		},
		"Error": {
			func(m *mocks.Cacher) {
				m.On("Get", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
			},
			"Error getting item with key",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Setup(test.mock)
			got, err := Get(context.Background(), "key")
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *CacheTestSuite) TestSet() {
	tt := map[string]struct {
		mock func(m *mocks.Cacher)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Cacher) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			"Successfully set cache item with key",
		},
		"Error": {
			func(m *mocks.Cacher) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Setup(test.mock)
			Set(context.Background(), "key", "key", Options{})
			t.Contains(t.LogWriter.String(), test.want)
			t.Reset()
		})
	}
}

func (t *CacheTestSuite) TestDelete() {
	tt := map[string]struct {
		mock func(m *mocks.Cacher)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Cacher) {
				m.On("Delete", mock.Anything, mock.Anything).Return(nil)
			},
			"Successfully deleted cache item with key",
		},
		"Error": {
			func(m *mocks.Cacher) {
				m.On("Delete", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Setup(test.mock)
			Delete(context.Background(), "key")
			t.Contains(t.LogWriter.String(), test.want)
			t.Reset()
		})
	}
}

func (t *CacheTestSuite) TestInvalidate() {
	tt := map[string]struct {
		mock func(m *mocks.Cacher)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Cacher) {
				m.On("Invalidate", mock.Anything, mock.Anything).Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.Cacher) {
				m.On("Invalidate", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			"Error invalidating cache",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Setup(test.mock)
			err := Invalidate(context.Background(), InvalidateOptions{})
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

func (t *CacheTestSuite) TestClear() {
	tt := map[string]struct {
		mock func(m *mocks.Cacher)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Cacher) {
				m.On("Clear", mock.Anything).Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.Cacher) {
				m.On("Clear", mock.Anything).Return(fmt.Errorf("error"))
			},
			"Error clearing cache",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Setup(test.mock)
			err := Clear(context.Background())
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

func (t *CacheTestSuite) TestSetDriver() {
	tt := map[string]struct {
		input  Cacher
		panics bool
	}{
		"Success": {
			&mocks.Cacher{},
			false,
		},
		"Error": {
			nil,
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			if test.panics {
				t.Panics(func() {
					SetDriver(test.input)
				})
				return
			}
			SetDriver(test.input)
			t.Equal(test.input, c)
		})
	}
}
