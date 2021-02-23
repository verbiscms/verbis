// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
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
