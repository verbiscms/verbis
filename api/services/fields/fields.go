// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/services/fields/resolve"
	"strings"
)

// Fields defines the map of fields to be returned to the template.
type Fields map[string]interface{}

// GetFields
//
// Returns all of the fields for the current post, or post ID given.
func (s *Service) GetFields(args ...interface{}) Fields {
	fields := s.handleArgs(args)

	var f = make(Fields)
	s.mapper(fields, func(field domain.PostField) {
		f[field.Name] = resolve.Field(field, s.deps).Value
	})

	return f
}

// WalkerFunc defines the function for walking the slice of domain.PostField
// when being mapped. It send the field back to the calling function for
// processing.
type WalkerFunc func(field domain.PostField)

// mapper
//
// Ranges over the fields and resolves all of the values from the given
// slice. If the field has a parent of field layout, the field will
// be skipped.
func (s *Service) mapper(fields domain.PostFields, walkerFunc WalkerFunc) {
	for _, field := range fields {
		if field.Type == "repeater" {
			repeater := s.GetRepeater(field.Name)
			if repeater != nil {
				field.Value = repeater
				walkerFunc(field)
			}
			continue
		}

		if field.Type == "flexible" {
			flexible := s.GetFlexible(field.Name)
			if flexible != nil {
				field.Value = flexible
				walkerFunc(field)
			}
			continue
		}

		if field.Key == "" || len(strings.Split(field.Key, SEPARATOR)) == 0 {
			walkerFunc(field)
		}
	}
}
