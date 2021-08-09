// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/mocks/common/vaidation/mocks"
	"reflect"
	"testing"
)

func TestValidator(t *testing.T) {
	assert.Equal(t, v.validate, Validator())
}

func TestProcess(t *testing.T) {
	var ve Errors
	tt := map[string]struct {
		input func() error
		want  interface{}
	}{
		"Normal Error": {
			func() error {
				return fmt.Errorf("error")
			},
			ve,
		},
		"Simple": {
			func() error {
				item := &mocks.FieldError{}
				item.On("Tag").Return("required").Twice()
				item.On("Field").Return("field").Twice()
				item.On("Param").Return(mock.Anything)
				return validator.ValidationErrors{
					item,
				}
			},
			Errors{{Key: "field", Type: "required", Message: "Field is required."}},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Process(test.input())
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMessage(t *testing.T) {
	tt := map[string]struct {
		kind  string
		field string
		param string
		want  interface{}
	}{
		"Required": {"required", "password", "", "Password is required."},
		"Email":    {"email", "", "", "Enter a valid email address."},
		"Min":      {"min", "", "3", "Enter a minimum of 3 characters."},
		"Max":      {"max", "", "3", "Enter a maximum of 3 characters."},
		"Alpha":    {"alpha", "Password", "", "Password must be alpha."},
		"AlphaNum": {"alphanum", "Password", "", "Password must be alphanumeric."},
		"IP":       {"ip", "address", "", "Address must be valid IP address."},
		"URL":      {"url", "url", "", "Enter a valid URL."},
		"EQ Field": {"eqfield", "url", "param", "URL must equal the param."},
		"Password": {"password", "password", "", "Password doesn't match our records."},
		"Default":  {"", "url", "", "Validation failed on the URL field."},
		"id":       {"", "id", "", "Validation failed on the ID field."},
		"Id":       {"", "id", "", "Validation failed on the ID field."},
		"Url":      {"", "url", "", "Validation failed on the URL field."},
		"url":      {"", "url", "", "Validation failed on the URL field."},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := message(test.kind, test.field, test.param)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestComparePassword(t *testing.T) {
	fl := &mocks.FieldLevel{}
	pr := &domain.UserPasswordReset{DBPassword: "wrong"}
	fl.On("Field").Return(reflect.ValueOf("assord"))
	fl.On("Parent").Return(reflect.ValueOf(pr))
	got := comparePassword(fl)
	assert.False(t, got)
}

func TestDefaultValidator_Engine(t *testing.T) {
	d := defaultValidator{}
	assert.Equal(t, d.validate, d.Engine())
}

func TestDefaultValidator_ValidateStruct(t *testing.T) {
	type str struct {
		Test string `binding:"required"`
	}
	tt := map[string]struct {
		input interface{}
		error bool
	}{
		"String": {
			"test",
			false,
		},
		"Pass": {
			&str{Test: "test"},
			false,
		},
		"Error": {
			&str{},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := v.ValidateStruct(test.input)
			if !test.error {
				assert.Nil(t, got)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestKindOfData(t *testing.T) {
	ptr := 1
	tt := map[string]struct {
		input interface{}
		want  string
	}{
		"String": {
			"test",
			"string",
		},
		"Int": {
			1,
			"int",
		},
		"Pointer": {
			&ptr,
			"int",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := kindOfData(test.input)
			assert.Equal(t, test.want, got.String())
		})
	}
}

func TestTagNameFunc(t *testing.T) {
	tt := map[string]struct {
		input reflect.StructField
		want  string
	}{
		"Key": {
			reflect.StructField{
				Tag: `validation_key:"key"`,
			},
			"key",
		},
		"JSON": {
			reflect.StructField{
				Tag: `json:"json"`,
			},
			"json",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := tagNameFunc(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
