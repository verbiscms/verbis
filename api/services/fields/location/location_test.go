// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package location

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// LocationTestSuite defines the helper used for location
// field testing.
type LocationTestSuite struct {
	suite.Suite
	TestPath          string
	OriginalFieldPath string
}

// TestLocation asserts testing has begun.
func TestLocation(t *testing.T) {
	suite.Run(t, new(LocationTestSuite))
}

// SetupSuite Discard the logger on setup and
// init caching.
func (t *LocationTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.OriginalFieldPath = FieldPath
	FieldPath = ""
	t.TestPath = filepath.Join(wd, "testdata")
}

// TearDownSuite reassigns field file after the
// tests.
func (t *LocationTestSuite) TearDownSuite() {
	FieldPath = t.OriginalFieldPath
}

func (t *LocationTestSuite) TestLocation_Layout() {
	tt := map[string]struct {
		cacheable bool
		file      string
		want      interface{}
	}{
		"Bad Path": {
			false,
			"wrongval",
			domain.FieldGroups{},
		},
		"Not Cached": {
			false,
			"location.json",
			domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Cacheable Nil": {
			true,
			"location.json",
			domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Cacheable": {
			true,
			"location.json",
			domain.FieldGroups{{Title: "title", Fields: domain.Fields{{Name: "test"}}}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{}
			t.Equal(test.want, l.Layout(filepath.Join(t.TestPath, test.file), domain.PostDatum{}, test.cacheable))
		})
	}
}

//func (t *LocationTestSuite) TestLocation_Layout_CacheError() {
//	buf := &bytes.Buffer{}
//	logger.SetOutput(buf)
//	c := &mockCache.Cacher{}
//	c.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
//	c.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Error"))
//	cache.SetDriver(c)
//	l := &Location{}
//	l.Layout(filepath.Join(t.TestPath, "location.json"), domain.PostDatum{}, true)
//	t.Contains(buf.String(), "error")
//}

func (t *LocationTestSuite) TestLocation_GroupResolver() {
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
		"post": {
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
		"Page template": {
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
			post: domain.PostDatum{Post: domain.Post{Resource: "resource"}},
			groups: domain.FieldGroups{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "resource", Operator: "==", Value: "resource"}},
					},
				},
			},
			want: domain.FieldGroups{{Title: "post"}},
		},
		"Nil Resource": {
			post: domain.PostDatum{Post: domain.Post{Resource: ""}},
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

func (t *LocationTestSuite) TestLocation_FieldGroupWalker() {
	var fg domain.FieldGroups

	id, err := uuid.Parse("6a4d7442-1020-490f-a3e2-436f9135bc24")
	t.NoError(err)

	// For bad file
	openErrorPath := filepath.Join(t.TestPath, "open-error.json")
	err = os.Chmod(openErrorPath, 000)
	t.NoError(err)
	defer func() {
		err = os.Chmod(openErrorPath, os.ModePerm)
		t.NoError(err)
	}()

	tt := map[string]struct {
		file string
		want interface{}
	}{
		"Success": {
			"success.json",
			domain.FieldGroups{{UUID: id, Title: "Title", Fields: domain.Fields{{Name: "test"}}}},
		},
		"Bad TestPath": {
			"wrongval",
			"no such file or directory",
		},
		"Unmarshal Error": {
			"unmarshal.json",
			fg,
		},
		"Open Error": {
			"open-error.json",
			fg,
		},
		"Empty Fields": {
			"empty.json",
			fg,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{}
			got, err := l.fieldGroupWalker(filepath.Join(t.TestPath, test.file))
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
