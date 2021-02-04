// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/stretchr/testify/assert"
	"testing"
)

type noStringer struct{}

func TestQuery_Get(t *testing.T) {

	q := Query{}

	tt := map[string]struct {
		orderBy        string
		orderDirection string
		want           params.Params
	}{
		"Empty": {
			"",
			"",
			params.Params{
				Page:           1,
				Limit:          15,
				LimitAll:       false,
				OrderDirection: Defaults.OrderDirection,
				OrderBy:        Defaults.OrderBy,
				Filters:        nil,
			},
		},
		"Order By": {
			"test",
			"",
			params.Params{
				Page:           1,
				Limit:          15,
				LimitAll:       false,
				OrderDirection: Defaults.OrderDirection,
				OrderBy:        "test",
				Filters:        nil,
			},
		},
		"Order Direction": {
			"",
			"test",
			params.Params{
				Page:           1,
				Limit:          15,
				LimitAll:       false,
				OrderDirection: "test",
				OrderBy:        Defaults.OrderBy,
				Filters:        nil,
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := q.Get(test.orderBy, test.orderDirection)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestQuery_Param(t *testing.T) {

	tt := map[string]struct {
		query Query
		param string
		want  string
	}{
		"Simple": {
			Query{"test": "1"},
			"test",
			"1",
		},
		"Int": {
			Query{"test": 1},
			"test",
			"1",
		},
		"Not Found": {
			Query{"wrongval": ""},
			"test",
			"",
		},
		"Bad Cast": {
			Query{"test": noStringer{}},
			"test",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.query.Param(test.param)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestQuery_Default(t *testing.T) {

	tt := map[string]struct {
		query Query
		param string
		def   string
		want  interface{}
	}{
		"Found": {
			Query{"test": "1"},
			"test",
			"",
			"1",
		},
		"Not Found": {
			Query{},
			"wrongval",
			"default",
			"default",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.query.Default(test.param, test.def)
			assert.Equal(t, test.want, got)
		})
	}
}
