// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Value defines the methods used for resolving
// domain.FieldValue's. The store is required
// for use with DB calls such as Posts.
type Value struct {
	deps *deps.Deps
}

var (
	// Iterable defines the types that can be a slice,
	// by a comma delimited Original Value (1,2,3).
	Iterable = []string{
		"category",
		"image",
		"post",
		"user",
		"tags",
	}
	// Choice defines the types that can be transformed
	// into a `choice` struct.
	Choice = []string{
		"button_group",
		"radio",
		"select",
	}
)

// valuer defines the function for resolving domain.FieldValue's
type valuer func(field domain.FieldValue) (interface{}, error)

// valueMap represents a map of field types with a valuer used
// to resolve field values.
type valueMap map[string]valuer

// Field
//
// Resolve's a field value.
func Field(field domain.PostField, d *deps.Deps) domain.PostField {
	exec := &Value{
		deps: d,
	}
	resolved := exec.resolve(field)
	return resolved
}

// getMap
//
// Returns the map of functions for resolving values.
func (v *Value) getMap() valueMap {
	return valueMap{
		"button_group": v.choice,
		"category":     v.category,
		"checkbox":     v.checkbox,
		"image":        v.media,
		"number":       v.number,
		"post":         v.post,
		"radio":        v.choice,
		"range":        v.number,
		"select":       v.choice,
		"user":         v.user,
	}
}

// resolve
//
// This function is the core for resolving the fields value
// for use with templates. It determines if the given
// field values is a slice or array or singular and
// returns a resolved field value or a slice of
// interfaces.
func (v *Value) resolve(field domain.PostField) domain.PostField {
	original := field.OriginalValue

	if original.IsEmpty() {
		field.Value = field.OriginalValue.String()
		return field
	}

	if field.TypeIsInSlice(Choice) && field.Key != "map" {
		field.Value = field.OriginalValue.String()
		return field
	}

	if !field.TypeIsInSlice(Iterable) {
		field.Value = v.execute(field.OriginalValue.String(), field.Type)
		return field
	}

	var items []interface{}
	for _, f := range original.Slice() {
		res := v.execute(f, field.Type)
		if res != nil {
			items = append(items, res)
		}
	}
	field.Value = items

	if len(items) == 1 {
		field.Value = items[0]
	}

	return field
}

// execute
//
// Executes the function based on the fields type.
// If the function is not within the valueMap,
// the original value will be returned.
func (v *Value) execute(value, typ string) interface{} {
	fn, ok := v.getMap()[typ]
	if !ok {
		return value
	}

	val, err := fn(domain.FieldValue(value))
	if err != nil {
		logger.WithError(err).Error()
	}

	return val
}
