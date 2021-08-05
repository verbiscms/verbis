// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostDatum_IsPublic(t *testing.T) {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"Public": {
			"published",
			true,
		},
		"Not Public": {
			"private",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := PostDatum{
				Post: Post{Status: test.input},
			}
			got := p.IsPublic()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostDatum_HasResource(t *testing.T) {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"Resource": {
			"resource",
			true,
		},
		"No Resource": {
			"",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := PostDatum{
				Post: Post{Resource: test.input},
			}
			got := p.HasResource()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostDatum_HasCategories(t *testing.T) {
	tt := map[string]struct {
		input *Category
		want  bool
	}{
		"Categories": {
			&Category{Name: "test"},
			true,
		},
		"No Category": {
			nil,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := PostDatum{Category: test.input}
			got := p.HasCategory()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostDatum_IsHomepage(t *testing.T) {
	tt := map[string]struct {
		input int
		id    int
		want  bool
	}{
		"Resource": {
			1,
			1,
			true,
		},
		"No Resource": {
			1,
			2,
			false,
		},
		"Zero FullPath": {
			0,
			2,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := PostDatum{
				Post: Post{ID: test.id},
			}
			got := p.IsHomepage(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostDatum_Tpl(t *testing.T) {
	got := PostDatum{
		Post:   Post{Title: "title"},
		Author: UserPart{},
		Fields: PostFields{{Name: "test"}},
	}
	want := PostTemplate{
		Post:   got.Post,
		Author: got.Author,
		Fields: got.Fields,
	}
	assert.Equal(t, want, got.Tpl())
}

func TestPostField_TypeIsInSlice(t *testing.T) {
	tt := map[string]struct {
		input []string
		field *PostField
		want  bool
	}{
		"Found": {
			[]string{"test"},
			&PostField{Type: "test"},
			true,
		},
		"Not Found": {
			[]string{"wrong"},
			&PostField{Type: "test"},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.field.TypeIsInSlice(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostField_IsValueJSON(t *testing.T) {
	tt := map[string]struct {
		field *PostField
		want  bool
	}{
		"No Keywords": {
			&PostField{OriginalValue: `none`},
			false,
		},
		"Invalid JSON": {
			&PostField{OriginalValue: `{"key": "key1", "value": }`},
			false,
		},
		"Not Found": {
			&PostField{OriginalValue: `{"key": "key1", "value": "value1"}`},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.field.IsValueJSON()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPostMeta_Scan(t *testing.T) {
	UtilTestScanner(&PostMeta{Title: "title"}, t)
}

func TestPostMeta_Value(t *testing.T) {
	UtilTestValue(&PostMeta{Title: "title"}, t)
}

func TestPostSeo_Scan(t *testing.T) {
	UtilTestScanner(&PostSeo{Canonical: "test"}, t)
}

func TestPostSeo_Value(t *testing.T) {
	UtilTestValue(&PostSeo{Canonical: "test"}, t)
}
