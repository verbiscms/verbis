package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Repeater represents the collection of fields for use
// with templates.
type Repeater []domain.PostField

// GetRepeater
//
// Obtains the collection of children from the given key and returns
// a new Repeater.
// Returns errors.INVALID if the field type is not a repeater.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetRepeater(key string, args ...interface{}) (Repeater, error) {
	const op = "FieldService.GetRepeater"

	fields := s.handleArgs(args)

	field, err := s.findFieldByKey(key, fields)
	if err != nil {
		return nil, err
	}

	if field.Type != "repeater" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Field is not a repeater", Operation: op, Err: fmt.Errorf("field with the key: %s, is not a repeater", key)}
	}

	return s.getFieldChildren(field.UUID, fields), nil
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
func (r Repeater) SubField(key string) interface{} {
	for _, sub := range r {
		if key == sub.Key {
			return sub.Value
		}
	}
	return nil
}

// First
//
// Returns the first element in the repeater, or nil if
// the length of the repeater is zero.
func (r Repeater) First() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[0]
}

// Last
//
// Returns the last element in the repeater, or nil if
// the length of the repeater is zero.
func (r Repeater) Last() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[len(r)-1]
}
