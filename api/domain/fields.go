// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type (
	// FieldGroup defines a group of JSON fields.
	FieldGroup struct {
		UUID      uuid.UUID         `json:"uuid"`
		Title     string            `json:"title"`
		Fields    []Field           `json:"fields,omitempty"`
		Locations [][]FieldLocation `json:"location,omitempty"`
	}
	// Field defines an individual field type.
	Field struct {
		UUID         uuid.UUID                  `json:"uuid"`
		Label        string                     `json:"label"`
		Name         string                     `json:"name"`
		Type         string                     `json:"type"`
		Instructions string                     `json:"instructions"`
		Required     bool                       `json:"required"`
		Logic        *[][]FieldConditionalLogic `json:"conditional_logic"`
		Wrapper      *FieldWrapper              `json:"wrapper"`
		Options      map[string]interface{}     `json:"options"`
		SubFields    *[]Field                   `json:"sub_fields,omitempty"`
		Layouts      map[string]FieldLayout     `json:"layouts,omitempty"`
	}

	// FieldLayout defines the structure of fields for
	// individual pages and resources.
	FieldLayout struct {
		UUID      uuid.UUID `json:"uuid"`
		Name      string    `json:"name"`
		Label     string    `json:"label"`
		Display   string    `json:"didpslay"`
		SubFields *[]Field  `json:"sub_fields,omitempty"`
	}
	// FieldLocation defines where the FieldGroup will appear.
	FieldLocation struct {
		Param    string
		Operator string
		Value    string
	}
	// FieldWrapper defines the container for field objects on
	// the front end.
	FieldWrapper struct {
		Width int `json:"width"`
	}
	// FieldConditionalLogic defines the logic used to process
	// a field and if one can be shown.
	FieldConditionalLogic struct {
		Field    string `json:"field"`
		Operator string `json:"operator"`
		Value    string `json:"value"`
	}
	// FieldValue defines the original value of the field in
	// string form.
	FieldValue string
)

// Slice
//
// Returns a slice of split field values by comma.
func (f FieldValue) Slice() []string {
	return strings.FieldsFunc(f.String(), func(c rune) bool {
		return c == ','
	})
}

// IsEmpty
//
// Determines if the field is an empty string.
func (f FieldValue) IsEmpty() bool {
	return string(f) == ""
}

// String
//
// Stringer on the FieldValue type.
func (f FieldValue) String() string {
	return string(f)
}

// Int
//
// Converts the field value to a string.
//
// Returns errors.INVALID if the string convert failed.
func (f FieldValue) Int() (int, error) {
	const op = "FieldValue.Int"
	i, err := strconv.Atoi(f.String())
	if err != nil {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Unable to cast FieldValue to an integer", Operation: op, Err: err}
	}
	return i, nil
}
