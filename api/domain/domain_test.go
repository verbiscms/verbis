// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/errors"
	"testing"
)

func UtilTestScanner(scanner sql.Scanner, t *testing.T) {
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
			"Error unmarshalling into",
		},
		"Nil": {
			nil,
			nil,
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			err := scanner.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func UtilTestValue(valuer driver.Valuer, t *testing.T) {
	tt := map[string]struct {
		marshal func(v interface{}) ([]byte, error)
		want    interface{}
	}{
		"Success": {
			func(v interface{}) ([]byte, error) {
				return []byte(`{"key":"value"}`), nil
			},
			`{"key":"value"}`,
		},
		"Error": {
			func(v interface{}) ([]byte, error) {
				return nil, fmt.Errorf("error")
			},
			"error",
		},
		"Nil Length": {
			func(v interface{}) ([]byte, error) {
				return nil, nil
			},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			orig := marshaller
			defer func() { marshaller = orig }()
			marshaller = test.marshal

			value, err := valuer.Value()
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, cast.ToString(value))
		})
	}
}

func UtilTestValueNil(valuer driver.Valuer, t *testing.T) {
	got, err := valuer.Value()
	assert.Nil(t, got)
	assert.NoError(t, err)
}
