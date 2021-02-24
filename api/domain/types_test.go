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

func TestDBMap_Scan(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			[]byte(`{"test": "test"}`),
			nil,
		},
		"Bad Unmarshal": {
			[]byte(`{"test": wrong}`),
			"Error unmarshalling into DBMap",
		},
		"Nil": {
			nil,
			DBMap{},
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for DBMap",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := DBMap{}
			err := m.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestDBMap_Value(t *testing.T) {

	tt := map[string]struct {
		input DBMap
		want  interface{}
	}{
		"Success": {
			DBMap{"test": "test"},
			DBMap{"test": "test"},
		},
		"Bad Unmarshal": {
			DBMap{"test": make(chan int)},
			"Error marshalling DBMap",
		},
		"Nil Length": {
			nil,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			value, err := test.input.Value()

			if test.input == nil {
				assert.Nil(t, value)
				return
			}

			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}

			got, err := cast.ToStringE(value)
			assert.NoError(t, err)

			want, err := json.Marshal(test.input)
			assert.NoError(t, err)

			assert.Equal(t, string(want), got)
		})
	}
}
