// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
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
func (t *FieldTestSuite) GetMockService(fields domain.PostFields, fnc func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)) *Service {
	fieldsMock := &mocks.FieldsRepository{}
	categoryMock := &mocks.CategoryRepository{}

	if fnc != nil {
		fnc(fieldsMock, categoryMock)
	}

	s := t.GetService(fields)
	s.deps = &deps.Deps{
		Store: &models.Store{
			Categories: categoryMock,
			Fields:     fieldsMock,
		},
	}

	return s
}

// GetPostsMockService
//
// Mock posts service for testing.
func (t *FieldTestSuite) GetPostsMockService(fields domain.PostFields, fnc func(p *mocks.PostsRepository)) *Service {
	postsMocks := &mocks.PostsRepository{}

	if fnc != nil {
		fnc(postsMocks)
	}

	s := t.GetService(fields)
	s.deps = &deps.Deps{
		Store: &models.Store{
			Posts: postsMocks,
		},
	}

	return s
}

// GetTypeMockService
//
// Mock store service for testing.
func (t *FieldTestSuite) GetTypeMockService(fnc func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)) *Service {
	categoryMock := &mocks.CategoryRepository{}
	mediaMock := &mocks.MediaRepository{}
	postsMock := &mocks.PostsRepository{}
	userMock := &mocks.UserRepository{}

	if fnc != nil {
		fnc(categoryMock, mediaMock, postsMock, userMock)
	}

	s := t.GetService(nil)
	s.deps = &deps.Deps{
		Store: &models.Store{
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
