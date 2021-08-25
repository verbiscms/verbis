// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/services/fields/resolve"
)

// GetField returns the value of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetField(name string, args ...interface{}) interface{} {
	f, ok := s.getCacheField(name, FieldCacheKey)
	if ok {
		return f
	}

	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil
	}

	resolved := resolve.Field(field, s.deps)

	s.setCacheField(resolved.Value, name, FieldCacheKey)

	return resolved.Value
}

// GetFieldObject returns the raw object of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetFieldObject(name string, args ...interface{}) domain.PostField {
	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		logger.WithError(err).Error()
		return domain.PostField{}
	}

	resolved := resolve.Field(field, s.deps)

	return resolved
}
