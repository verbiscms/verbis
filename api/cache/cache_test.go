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
	mocks "github.com/verbiscms/verbis/api/mocks/cache/mocks"
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

// Setup assigns a mock Store to c.
func (t *CacheTestSuite) Setup(mf func(m *mocks.StoreInterface)) *Cache {
	m := &mocks.StoreInterface{}
	if mf != nil {
		mf(m)
	}
	return &Cache{
		store: m,
	}
}

func (t *CacheTestSuite) TestLoad() {
	tt := map[string]struct {
		mock  func(m *mocks.Provider)
		input *environment.Env
		want  interface{}
	}{
		"Success": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(nil)
				m.On("Ping").Return(nil)
				m.On("Driver").Return(MemoryStore)
				m.On("Store").Return(nil)
			},
			&environment.Env{CacheDriver: MemoryStore},
			MemoryStore,
		},
		"Default": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(nil)
				m.On("Ping").Return(nil)
				m.On("Driver").Return(MemoryStore)
				m.On("Store").Return(nil)
			},
			&environment.Env{CacheDriver: ""},
			MemoryStore,
		},
		"Nil Env": {
			nil,
			nil,
			"Error loading cache",
		},
		"Invalid Driver": {
			nil,
			&environment.Env{CacheDriver: "wrong"},
			"Error loading cache, invalid driver",
		},
		"Validate Error": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(fmt.Errorf("error"))
			},
			&environment.Env{CacheDriver: MemoryStore},
			"Error loading cache, validation failed",
		},
		"Ping Error": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(nil)
				m.On("Ping").Return(fmt.Errorf("error"))
				m.On("Driver").Return(MemoryStore)
			},
			&environment.Env{CacheDriver: MemoryStore},
			"Error error pinging cache store",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Provider{}
			if test.mock != nil {
				test.mock(m)
			}
			providers = providerMap{MemoryStore: func(env *environment.Env) provider {
				return m
			}}
			c, err := Load(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			if c == nil {
				t.Fail("nil Driver")
			}
			t.Equal(test.want, c.Driver())
		})
	}
}

func (t *CacheTestSuite) TestGet() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).Return("item", nil)
			},
			"item",
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
			},
			"Error getting item with key",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got, err := c.Get(context.Background(), "key")
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
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			"Successfully set cache item with key",
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(fmt.Errorf("set error"))
			},
			"set error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			c.Set(context.Background(), "key", "key", Options{})
			t.Contains(t.LogWriter.String(), test.want)
			t.Reset()
		})
	}
}

func (t *CacheTestSuite) TestDelete() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Delete", mock.Anything, mock.Anything).
					Return(nil)
			},
			"Successfully deleted cache item with key",
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Delete", mock.Anything, mock.Anything).
					Return(fmt.Errorf("delete error"))
			},
			"delete error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			c.Delete(context.Background(), "key")
			t.Contains(t.LogWriter.String(), test.want)
			t.Reset()
		})
	}
}

func (t *CacheTestSuite) TestInvalidate() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Invalidate", mock.Anything, mock.Anything).
					Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Invalidate", mock.Anything, mock.Anything).
					Return(fmt.Errorf("invalidate error"))
			},
			"Error invalidating cache",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Invalidate(context.Background(), InvalidateOptions{})
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
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Clear", mock.Anything).
					Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Clear", mock.Anything).
					Return(fmt.Errorf("error"))
			},
			"Error clearing cache",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Clear(context.Background())
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
