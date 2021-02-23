// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldValue_Array(t *testing.T) {

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
			f := FieldValue(test.input)
			got := f.Slice()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFieldValue_IsEmpty(t *testing.T) {
	f := FieldValue("")
	assert.Equal(t, true, f.IsEmpty())
}

func TestFieldValue_String(t *testing.T) {
	f := FieldValue("test")
	assert.Equal(t, "test", f.String())
}

func TestFieldValue_Int(t *testing.T) {

	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Success": {
			"1",
			1,
		},
		"Error": {
			"wrongval",
			"Unable to cast FieldValue to an integer",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := FieldValue(test.input)
			got, err := f.Int()
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
