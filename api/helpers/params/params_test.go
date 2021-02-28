// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"encoding/json"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockStringer struct {
	def map[string]interface{}
}

func Setup(t *testing.T, def Defaults) *mockStringer {
	data, err := json.Marshal(def) // Convert to a json string
	assert.NoError(t, err)
	var m = make(map[string]interface{}, 0)
	err = json.Unmarshal(data, &m) // Convert to a map

	ms := mockStringer{def: m}

	return &ms
}

func (m *mockStringer) Param(param string) string {
	val, ok := m.def[param]
	if !ok {
		return ""
	}
	s, err := cast.ToStringE(val)
	if err != nil {
		return ""
	}
	return s
}

func TestParams_Get(t *testing.T) {

	tt := map[string]struct {
		def  Defaults
		want Params
	}{
		"No Defaults": {
			Defaults{},
			Params{Page: 1, Limit: DefaultLimit, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Nil Page": {
			Defaults{Page: -1},
			Params{Page: 1, Limit: DefaultLimit, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Limit": {
			Defaults{Limit: 20},
			Params{Page: 1, Limit: 20, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Zero Limit": {
			Defaults{Limit: 0},
			Params{Page: 1, Limit: DefaultLimit, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Nil Limit": {
			Defaults{Limit: nil},
			Params{Page: 1, Limit: DefaultLimit, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Limit All": {
			Defaults{Limit: "all"},
			Params{Page: 1, Limit: 0, LimitAll: true, OrderBy: DefaultOrderBy, OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Order By": {
			Defaults{OrderBy: "name"},
			Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "name", OrderDirection: DefaultOrderDirection, Filters: nil},
		},
		"Order Direction": {
			Defaults{OrderDirection: "ASC"},
			Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: DefaultOrderBy, OrderDirection: "ASC", Filters: nil},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			mock := Setup(t, test.def)
			p := New(mock, test.def)
			got := p.Get()
			assert.Equal(t, test.want, got)
		})
	}
}

// Limit Error
type mockLimit struct{}

func (m *mockLimit) Param(param string) string {
	return "0"
}

func TestParams_LimitError(t *testing.T) {
	m := &mockLimit{}
	p := Params{Stringer: m, defaults: Defaults{Limit: 20}}
	limit, limitAll := p.limit()
	assert.Equal(t, 20, limit)
	assert.False(t, limitAll)
}

// Page Error
type mockPage struct{}

func (m *mockPage) Param(param string) string {
	return "99999999999999999999999999999"
}

func TestParams_PageError(t *testing.T) {
	m := &mockPage{}
	p := Params{Stringer: m}
	assert.Equal(t, 1, p.page())
}

// Filter testing
type mockFilter struct {
	str string
}

func (m *mockFilter) Param(param string) string {
	return m.str
}

func TestParams_Filter(t *testing.T) {

	tt := map[string]struct {
		filter string
		want   map[string][]Filter
	}{
		"Filter": {
			`{"resource":[{"operator":"=", "value":"verbis"}]}`,
			Filters{
				"resource": {
					{Operator: "=", Value: "verbis"},
				},
			},
		},
		"Failed Filter": {
			`{"resource":[, "value":"verbis"}]}`,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &mockFilter{str: test.filter}
			p := Params{Stringer: m}
			got := p.filter()
			assert.Equal(t, test.want, got)
		})
	}
}
