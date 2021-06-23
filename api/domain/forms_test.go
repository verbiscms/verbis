// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/ainsleyclark/verbis/api/errors"
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

func TestFormValues_JSON(t *testing.T) {
	tt := map[string]struct {
		input FormValues
		want  interface{}
	}{
		"Success": {
			FormValues{
				"test": 1,
			},
			"{\"test\":1}",
		},
		"Error": {
			FormValues{
				"test": make(chan []byte, 0),
			},
			"Error processing the form fields for storing",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := test.input.JSON()
			if err != nil {
				assert.Equal(t, test.want, errors.Message(err))
				return
			}
			assert.Equal(t, test.want, string(got))
		})
	}
}

func TestFormValues_Scan(t *testing.T) {
	UtilTestScanner(&FormValues{}, t)
}

func TestFormValues_Value(t *testing.T) {
	UtilTestValue(FormValues{
		"key": "val",
	}, t)
	UtilTestValueNil(FormValues{}, t)
}
