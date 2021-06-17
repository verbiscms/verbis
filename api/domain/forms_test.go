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

func TestForm_GetRecipients(t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Singular": {
			"test",
			[]string{"test"},
		},
		"Multiple": {
			"test,test,test",
			[]string{"test", "test", "test"},
		},
		"Trailing Comma": {
			"test,",
			[]string{"test"},
		},
		"Leading Comma": {
			",test",
			[]string{"test"},
		},
		"Commas": {
			",,test,,",
			[]string{"test"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := Form{Recipients: test.input}
			got := f.GetRecipients()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFormLabel_Name(t *testing.T) {
	f := FormLabel("   test   ")
	got := f.Name()
	assert.Equal(t, "Test", got)
}

func TestFormLabel_String(t *testing.T) {
	f := FormLabel("test")
	assert.Equal(t, "test", f.String())
}

// TODO, Make scan and value test helper functions
func TestFormValues_Scan(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			[]byte(`{"large": {"url": "/test"}}`),
			nil,
		},
		"Bad Unmarshal": {
			[]byte(`{"large": wrong}`),
			"Error unmarshalling into FormValues",
		},
		"Nil": {
			nil,
			FormValues{},
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for FormValues",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := FormValues{}
			err := m.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestFormValues_Value(t *testing.T) {
	tt := map[string]struct {
		input FormValues
		want  interface{}
	}{
		"Success": {
			FormValues{
				"key": "val",
			},
			FormValues{
				"key": "val",
			},
		},
		"Nil Length": {
			nil,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			value, _ := test.input.Value()

			if test.input == nil {
				assert.Nil(t, value)
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
