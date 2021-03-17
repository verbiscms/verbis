// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
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
	r := "news"

	tt := map[string]struct {
		input *string
		want  bool
	}{
		"Resource": {
			&r,
			true,
		},
		"No Resource": {
			nil,
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
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := PostDatum{
				Post: Post{Id: test.id},
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

func TestPostMeta_Scan(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			[]byte(`{"title": "test"}`),
			nil,
		},
		"Bad Unmarshal": {
			[]byte(`{"title": wrong}`),
			"Error unmarshalling into PostMeta",
		},
		"Nil": {
			nil,
			PostMeta{},
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for PostMeta",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := PostMeta{}
			err := m.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestPostMeta_Value(t *testing.T) {
	tt := map[string]struct {
		input PostMeta
		want  interface{}
	}{
		"Success": {
			PostMeta{Title: "title"},
			PostMeta{Title: "title"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			value, err := test.input.Value()
			assert.NoError(t, err)

			got, err := cast.ToStringE(value)
			assert.NoError(t, err)

			want, err := json.Marshal(test.input)
			assert.NoError(t, err)

			assert.Equal(t, string(want), got)
		})
	}
}

func TestPostSeo_Scan(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			[]byte(`{"canonical": "test"}`),
			nil,
		},
		"Bad Unmarshal": {
			[]byte(`{"canonical": wrong}`),
			"Error unmarshalling into PostSeo",
		},
		"Nil": {
			nil,
			PostSeo{},
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for PostSeo",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := PostSeo{}
			err := m.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestPostSeo_Value(t *testing.T) {
	tt := map[string]struct {
		input PostSeo
		want  interface{}
	}{
		"Success": {
			PostSeo{Canonical: "test"},
			PostSeo{Canonical: "test"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			value, err := test.input.Value()
			assert.NoError(t, err)

			got, err := cast.ToStringE(value)
			assert.NoError(t, err)

			want, err := json.Marshal(test.input)
			assert.NoError(t, err)

			assert.Equal(t, string(want), got)
		})
	}
}
