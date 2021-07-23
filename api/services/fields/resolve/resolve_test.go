// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/store"
	"io/ioutil"
	"testing"
)

// ResolverTestSuite defines the helper used for resolver
// field testing.
type ResolverTestSuite struct {
	suite.Suite
}

// TestResolver
//
// Assert testing has begun.
func TestResolver(t *testing.T) {
	suite.Run(t, new(ResolverTestSuite))
}

// SetupSuite
//
// Discard the logger on setup.
func (t *ResolverTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
}

// GetValue returns a default value.
func (t *ResolverTestSuite) GetValue() *Value {
	return &Value{
		&deps.Deps{
			Store: &store.Repository{},
		},
	}
}

func (t *ResolverTestSuite) Test_Field() {
	d := &deps.Deps{}
	field := domain.PostField{Type: "text", OriginalValue: "test"}

	got := Field(field, d)

	t.Equal(domain.PostField{Type: "text", OriginalValue: "test", Value: "test"}, got)
}

func (t *ResolverTestSuite) TestValue_Resolve() {
	tt := map[string]struct {
		field domain.PostField
		want  domain.PostField
	}{
		"Empty": {
			field: domain.PostField{OriginalValue: ""},
			want:  domain.PostField{OriginalValue: "", Value: ""},
		},
		"Not Iterable": {
			field: domain.PostField{OriginalValue: "999", Type: "number"},
			want:  domain.PostField{OriginalValue: "999", Type: "number", Value: int64(999)},
		},
		"Iterable": {
			field: domain.PostField{OriginalValue: "1,2,3,4,5", Type: "tags"},
			want:  domain.PostField{OriginalValue: "1,2,3,4,5", Type: "tags", Value: []interface{}{"1", "2", "3", "4", "5"}},
		},
		"Length of One": {
			field: domain.PostField{OriginalValue: "1", Type: "tags"},
			want:  domain.PostField{OriginalValue: "1", Type: "tags", Value: "1"},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetValue().resolve(test.field))
		})
	}
}

func (t *ResolverTestSuite) TestValue_Execute() {
	tt := map[string]struct {
		value string
		typ   string
		want  interface{}
	}{
		"Not found": {
			value: "test",
			typ:   "wrongval",
			want:  "test",
		},
		"Found": {
			value: "999",
			typ:   "number",
			want:  int64(999),
		},
		"Error": {
			value: "wrongval",
			typ:   "number",
			want:  nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetValue().execute(test.value, test.typ))
		})
	}
}
