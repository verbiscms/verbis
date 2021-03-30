// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	categories "github.com/ainsleyclark/verbis/api/mocks/store/categories"
	fields "github.com/ainsleyclark/verbis/api/mocks/store/fields"
	media "github.com/ainsleyclark/verbis/api/mocks/store/media"
	posts "github.com/ainsleyclark/verbis/api/mocks/store/posts"
	users "github.com/ainsleyclark/verbis/api/mocks/store/users"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FieldTestSuite defines the helper used for field
// testing.
type FieldTestSuite struct {
	suite.Suite
	logWriter bytes.Buffer
}

// TestFields
//
// Assert testing has begun.
func TestFields(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}

// Cannot parse helper.
type noStringer struct{}

// BeforeTest
//
// Assign the logger to a buffer.
func (t *FieldTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.logWriter = b
	logger.SetOutput(&t.logWriter)
}

// Reset
//
// Reset the log writer.
func (t *FieldTestSuite) Reset() {
	t.logWriter.Reset()
}

// GetMockService
//
// Mock service for testing.
func (t *FieldTestSuite) GetMockService(f domain.PostFields, fnc func(f *fields.Repository, c *categories.Repository)) *Service {
	fieldsMock := &fields.Repository{}
	categoryMock := &categories.Repository{}

	if fnc != nil {
		fnc(fieldsMock, categoryMock)
	}

	s := t.GetService(f)
	s.deps = &deps.Deps{
		Store: &store.Repository{
			Categories: categoryMock,
			Fields:     fieldsMock,
		},
	}

	return s
}

// GetPostsMockService
//
// Mock posts service for testing.
func (t *FieldTestSuite) GetPostsMockService(fields domain.PostFields, fnc func(p *posts.Repository)) *Service {
	postsMocks := &posts.Repository{}

	if fnc != nil {
		fnc(postsMocks)
	}

	s := t.GetService(fields)
	s.deps = &deps.Deps{
		Store: &store.Repository{
			Posts: postsMocks,
		},
	}

	return s
}

// GetTypeMockService
//
// Mock store service for testing.
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

// GetService
//
// Mock service for testing.
func (t *FieldTestSuite) GetService(fields domain.PostFields) *Service {
	return &Service{
		fields: fields,
		deps:   &deps.Deps{},
	}
}
