// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package location

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// LocationTestSuite defines the helper used for location
// field testing.
type LocationTestSuite struct {
	suite.Suite
	Path string
}

// TestLocation
//
// Assert testing has begun.
func TestLocation(t *testing.T) {
	suite.Run(t, new(LocationTestSuite))
}

// SetupSuite
//
// Discard the logger on setup and init caching.
func (t *LocationTestSuite) SetupSuite() {
	cache.Init()

	logger.Init(&environment.Env{})
	logger.SetOutput(ioutil.Discard)

	wd, err := os.Getwd()
	t.NoError(err)
	t.Path = filepath.Join(filepath.Dir(wd)+"./..") + "/test/testdata/fields"
}

func TestNewLocation(t *testing.T) {
	assert.Equal(t, &Location{JSONPath: "test/fields"}, NewLocation("test"))
}

func (t *LocationTestSuite) TestLocation_GetLayout() {
	tt := map[string]struct {
		cacheable bool
		jsonPath  string
		want      interface{}
	}{
		"Bad Path": {
			cacheable: false,
			jsonPath:  "wrongval",
			want:      domain.FieldGroups{},
		},
		"Not Cached": {
			cacheable: false,
			jsonPath:  "/test-get-layout",
			want:      domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Cacheable Nil": {
			cacheable: true,
			jsonPath:  "/test-get-layout",
			want:      domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Cacheable": {
			cacheable: true,
			jsonPath:  "/test-get-layout",
			want:      domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{JSONPath: t.Path + test.jsonPath}
			t.Equal(test.want, l.GetLayout(domain.PostDatum{}, test.cacheable))
		})
	}
}

func (t *LocationTestSuite) TestLocation_GroupResolver() {
	r := "resource"
	uu := uuid.New()

	tt := map[string]struct {
		post   domain.PostDatum
		groups domain.FieldGroups
		want   interface{}
	}{
		"None": {
			want: domain.FieldGroups{},
		},
		"Already Added": {
			post: domain.PostDatum{Post: domain.Post{Id: 1, Title: "title", Status: "published"}},
			groups: domain.FieldGroups{
				{Title: "status", UUID: uu},
				{Title: "status", UUID: uu},
			},
			want: domain.FieldGroups{{Title: "status", UUID: uu}},
		},
		"Status": {
			post: domain.PostDatum{Post: domain.Post{Id: 1, Title: "title", Status: "published"}},
			groups: domain.FieldGroups{
				{
					Title: "status",
					Locations: [][]domain.FieldLocation{
						{{Param: "status", Operator: "==", Value: "published"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "status"}},
		},
		"Post": {
			post: domain.PostDatum{Post: domain.Post{Id: 1}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "post", Operator: "==", Value: "1"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Page Template": {
			post: domain.PostDatum{Post: domain.Post{PageTemplate: "template"}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "page_template", Operator: "==", Value: "template"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Layout": {
			post: domain.PostDatum{Post: domain.Post{PageLayout: "layout"}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "page_layout", Operator: "==", Value: "layout"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Resource": {
			post: domain.PostDatum{Post: domain.Post{Resource: &r}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "resource", Operator: "==", Value: r}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Nil Resource": {
			post: domain.PostDatum{Post: domain.Post{Resource: nil}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "resource", Operator: "==", Value: "false"}},
					},
				},
			},
			want: domain.FieldGroups{},
		},
		"Category": {
			post: domain.PostDatum{Category: &domain.Category{Id: 1}},
			groups: domain.FieldGroups{
				{
					Title: "category",
					Locations: [][]domain.FieldLocation{
						{{Param: "category", Operator: "==", Value: "1"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "category"}},
		},
		"Nil Category": {
			post: domain.PostDatum{Category: nil},
			groups: domain.FieldGroups{
				{
					Title: "category",
					Locations: [][]domain.FieldLocation{
						{{Param: "category", Operator: "==", Value: "false"}},
					},
				},
			},
			want: domain.FieldGroups{},
		},
		"Author": {
			post: domain.PostDatum{Author: domain.UserPart{Id: 1}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "author", Operator: "==", Value: "1"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Role": {
			post: domain.PostDatum{Author: domain.UserPart{
				Role: domain.Role{
					Id: 1,
				},
			},
			},
			groups: domain.FieldGroups{
				{
					Title: "role",
					Locations: [][]domain.FieldLocation{
						{{Param: "role", Operator: "==", Value: "1"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "role"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{Groups: test.groups}
			t.Equal(test.want, l.groupResolver(test.post))
		})
	}
}

func (t *LocationTestSuite) TestLocation_fieldGroupWalker() {
	testPath := "/test-field-groups/"

	var fg domain.FieldGroups

	id, err := uuid.Parse("6a4d7442-1020-490f-a3e2-436f9135bc24")
	t.NoError(err)

	// For bad path
	err = os.Chmod(t.Path+testPath+"open-error/location.json", 000)
	t.NoError(err)
	defer func() {
		err = os.Chmod(t.Path+testPath+"open-error/location.json", 777)
		t.NoError(err)
	}()

	tt := map[string]struct {
		path string
		want interface{}
	}{
		"Success": {
			path: testPath + "/success",
			want: domain.FieldGroups{{UUID: id, Title: "Title", Fields: domain.Fields{{Name: "test"}}}, {UUID: id, Title: "Title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Bad Path": {
			path: testPath + "/wrongval",
			want: "no such file or directory",
		},
		"Unmarshal Error": {
			path: testPath + "/unmarshal",
			want: fg,
		},
		"Open Error": {
			path: testPath + "/open-error",
			want: fg,
		},
		"Empty Fields": {
			path: testPath + "/empty",
			want: fg,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{
				JSONPath: t.Path + test.path,
			}
			got, err := l.fieldGroupWalker()

			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *LocationTestSuite) Test_CheckLocation() {
	tt := map[string]struct {
		check    string
		location domain.FieldLocation
		want     bool
	}{
		"Equal Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "==", Value: "val"},
			want:     true,
		},
		"Equal Not Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "==", Value: "wrongval"},
			want:     false,
		},
		"Not Equal Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "!=", Value: "val"},
			want:     false,
		},
		"Not Equal Not Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "!=", Value: "wrongval"},
			want:     true,
		},
		"Wrong Operator": {
			check:    "val",
			location: domain.FieldLocation{Operator: "wrong", Value: "val"},
			want:     false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, checkLocation(test.check, test.location))
		})
	}
}

func (t *LocationTestSuite) Test_CheckMatch() {
	tt := map[string]struct {
		matches []bool
		want    bool
	}{
		"Matches": {
			matches: []bool{true, false, true, false},
			want:    false,
		},
		"Not Matched": {
			matches: []bool{true, true, true, true},
			want:    true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, checkMatch(test.matches))
		})
	}
}

func (t *LocationTestSuite) Test_HasBeenAdded() {
	key := uuid.New()

	tt := map[string]struct {
		fg   domain.FieldGroups
		key  string
		want bool
	}{
		"Added": {
			fg:   domain.FieldGroups{{UUID: key}},
			key:  key.String(),
			want: true,
		},
		"Not Added": {
			fg:   domain.FieldGroups{{UUID: uuid.New()}},
			key:  key.String(),
			want: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, hasBeenAdded(test.key, test.fg))
		})
	}
}
