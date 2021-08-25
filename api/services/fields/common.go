// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/services/fields/resolve"
	"strings"
)

// handleArgs processes the args for handling fields. The
// slice of interfaces are presumed to be the
// following:
// [0] for post ID, fields are obtained by the post given.
// Returns the fields to be modified & processed.
func (s *Service) handleArgs(args []interface{}) (domain.PostFields, int) {
	const op = "FieldsService.HandleArgs"

	if len(args) == 0 {
		return s.fields, 0
	}

	switch f := args[0].(type) {
	case domain.PostDatum:
		return f.Fields, f.Post.ID
	case domain.PostFields:
		return f, 0
	case domain.PostTemplate:
		return f.Fields, f.Post.ID
	}

	id, err := cast.ToIntE(args[0])
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Invalid argument passed to field", Operation: op, Err: fmt.Errorf("unable to cast post id to integer")}).Error()
		return nil, 0
	}

	return s.getFieldsByPost(id), id
}

// getFieldsByPost Returns the fields by post with the
// given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining the post.
func (s *Service) getFieldsByPost(id int) domain.PostFields {
	fields, err := s.deps.Store.Fields.Find(id)
	if err != nil {
		logger.WithError(err).Error()
		return nil
	}
	return fields
}

// findFieldByName returns a singular domain.PostField by
// the given name.
// Returns errors.NOTFOUND if the field does not exist.
func (s *Service) findFieldByName(name string, fields domain.PostFields) (domain.PostField, error) {
	const op = "FieldsService.findFieldByName"
	for _, field := range fields {
		if name == field.Name {
			return field, nil
		}
	}
	return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: "Field does not exist", Operation: op, Err: fmt.Errorf("no field exists with the name: %s", name)}
}

// walker represents the struct to be passed when resolving
// repeaters and flexible content types.
type walker struct {
	Key    string
	Index  int
	Field  domain.PostField
	Fields domain.PostFields
	*Service
}

// Walk constructs a pipe based on the key, name, SEPARATOR and the
// index in order to look up dynamic Flexible content and Repeater
// types. The key `flex|0|repeater|0|text` will be split
// and looked up. If the child value is of type Repeater
// or Flexible the function will call itself meaning
// all values will be resolved.
// The appender func outputs the field to the caller once resolved.
func (r *walker) Walk(appender func(domain.PostField)) {
	pipe := r.Key + r.Field.Name + SEPARATOR + cast.ToString(r.Index)

	for _, v := range r.Fields {
		pipeLen := strings.Split(pipe, SEPARATOR)
		keyLen := strings.Split(v.Key, SEPARATOR)

		if strings.HasPrefix(v.Key, pipe) && len(pipeLen)+1 == len(keyLen) {
			fieldType := v.Type
			if fieldType == "repeater" {
				v.Value = r.resolveRepeater(pipe+SEPARATOR, v, r.Fields)
				appender(v)
				continue
			}

			if fieldType == "flexible" {
				v.Value = r.resolveFlexible(pipe+SEPARATOR, v, r.Fields)
				appender(v)
				continue
			}

			appender(resolve.Field(v, r.deps))
		}
	}
}
