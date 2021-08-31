// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	categories "github.com/verbiscms/verbis/api/mocks/store/categories"
	fields "github.com/verbiscms/verbis/api/mocks/store/fields"
	media "github.com/verbiscms/verbis/api/mocks/store/media"
	posts "github.com/verbiscms/verbis/api/mocks/store/posts"
	users "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store"
	"testing"
)

// FieldTestSuite defines the helper used for field
// testing.
type FieldTestSuite struct {
	suite.Suite
	LogWriter bytes.Buffer
}

// TestFields asserts testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}

// Cannot parse helper.
type noStringer struct{}

// BeforeTest assign the logger to a buffer.
func (t *FieldTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(&t.LogWriter)
}

// Reset the log writer.
func (t *FieldTestSuite) Reset() {
	t.LogWriter.Reset()
}

// GetMockService mock service for testing.
func (t *FieldTestSuite) GetMockService(f domain.PostFields, fnc func(f *fields.Repository, c *categories.Repository, ca *cache.Store)) *Service {
	fieldsMock := &fields.Repository{}
	categoryMock := &categories.Repository{}
	cacheMock := &cache.Store{}

	if fnc != nil {
		fnc(fieldsMock, categoryMock, cacheMock)
	}

	s := t.GetService(f)
	s.deps = &deps.Deps{
		Store: &store.Repository{
			Categories: categoryMock,
			Fields:     fieldsMock,
		},
		Cache: cacheMock,
	}

	return s
}

// GetPostsMockService mock posts service for testing.
func (t *FieldTestSuite) GetPostsMockService(f domain.PostFields, fnc func(p *posts.Repository)) *Service {
	postsMocks := &posts.Repository{}

	if fnc != nil {
		fnc(postsMocks)
	}

	s := t.GetService(f)
	s.deps = &deps.Deps{
		Store: &store.Repository{
			Posts: postsMocks,
		},
	}

	return s
}

// GetTypeMockService mock store service for testing.
func (t *FieldTestSuite) GetTypeMockService(fnc func(c *categories.Repository, m *media.Repository, p *posts.Repository, u *users.Repository)) *Service {
	categoryMock := &categories.Repository{}
	mediaMock := &media.Repository{}
	postsMock := &posts.Repository{}
	userMock := &users.Repository{}

	if fnc != nil {
		fnc(categoryMock, mediaMock, postsMock, userMock)
	}

	s := t.GetService(nil)
	s.deps = &deps.Deps{
		Store: &store.Repository{
			Categories: categoryMock,
			Media:      mediaMock,
			Posts:      postsMock,
			User:       userMock,
		},
	}

	return s
}

// GetService mock service for testing.
func (t *FieldTestSuite) GetService(f domain.PostFields) *Service {
	c := &cache.Store{}
	CacheFieldError(c)
	return &Service{
		fields: f,
		deps: &deps.Deps{
			Cache: c,
		},
	}
}

var CacheFieldError = func(ca *cache.Store) {
	ca.On("Get", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("error"))
	ca.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func (t *FieldTestSuite) TestNewService() {
	s := &store.Repository{}

	var l = make(domain.FieldGroups, 0)
	var f = make(domain.PostFields, 0)

	pd := &domain.PostDatum{
		Post: domain.Post{
			ID: 1,
		},
		Layout: l,
		Fields: f,
	}

	d := &deps.Deps{
		Store:  s,
		Config: &domain.ThemeConfig{},
	}

	service := &Service{
		deps:   d,
		postID: 1,
		fields: f,
		layout: l,
	}

	t.Equal(NewService(d, pd), service)
}
