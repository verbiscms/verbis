// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/spf13/cast"
)

// Repeater represents the collection of rows used
// for the repeater function in templates.
type Repeater []Row

// Row represents the collection of the repeaters
// containing `sub_fields.
type Row domain.PostFields

// GetRepeater
//
// Returns the collection of children from the given key and returns
// a new Repeater.
//
// Returns errors.NOTFOUND if the field was not found by the given key.
// Returns errors.INVALID if the field type is not a repeater or the name could not be cast.
func (s *Service) GetRepeater(input interface{}, args ...interface{}) Repeater {
	const op = "FieldsService.GetRepeater"

	repeater, ok := input.(Repeater)
	if ok {
		return repeater
	}

	name, err := cast.ToStringE(input)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Could not cast input to string", Operation: op, Err: err}).Error()
		return nil
	}

	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil
	}

	if field.Type != "repeater" {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Field is not a repeater", Operation: op, Err: fmt.Errorf("field with the name: %s, is not a repeater", name)}).Error()
		return nil
	}

	return s.resolveRepeater("", field, fields)
}

// getFieldChildren
//
// Loops through the given slice of domain.PostField and compares the
// uuid passed with the field's parent UUID.
// It's not necessary to use a database call for this look up, as we will
// be looping through them anyway to append and format the fields.
// Returns the sorted slice of fields.
func (s *Service) resolveRepeater(key string, field domain.PostField, fields domain.PostFields) Repeater {
	const op = "FieldsService.resolveRepeater"

	amount, err := field.OriginalValue.Int()
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast repeater value to integer", Operation: op, Err: err})
		return Repeater{}
	}

	var repeater = make(Repeater, amount)
	for index := 0; index < len(repeater); index++ {
		r := walker{
			Key:     key,
			Index:   index,
			Field:   field,
			Fields:  fields,
			Service: s,
		}

		var row Row
		r.Walk(func(f domain.PostField) {
			row = append(row, f)
		})

		repeater[index] = row
	}

	return repeater
}

// HasRows
//
// Determines if the Repeater has any rows.
func (r Repeater) HasRows() bool {
	return len(r) != 0
}

// SubField
//
// Returns a sub field by key or nil if it wasn't found.
func (r Row) SubField(name string) interface{} {
	for _, sub := range r {
		if name == sub.Name {
			return sub.Value
		}
	}
	return nil
}

// HasField
//
// Returns true if a field exists within the row.
func (r Row) HasField(name string) bool {
	for _, sub := range r {
		if name == sub.Name {
			return true
		}
	}
	return false
}

// First
//
// Returns the first element in the repeater, or nil if
// the length of the repeater is zero.
func (r Row) First() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[0]
}

// Last
//
// Returns the last element in the repeater, or nil if
// the length of the repeater is zero.
func (r Row) Last() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[len(r)-1]
}
