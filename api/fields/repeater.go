package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Repeater represents the collection of fields fused
// for the repeater function in templates.
type Repeater []domain.PostField

// TODO: We no longer need the format paramater
// TODO: The repeater needs to be an array of arrays.
// Get repeater
// if (type is repeater)
// return repeater
// All values must be resolved beforr passing back.

// GetRepeater
//
// Returns the collection of children from the given key and returns
// a new Repeater.
// Returns errors.INVALID if the field type is not a repeater.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetRepeater(name string, args ...interface{}) (Repeater, error) {
	const op = "FieldsService.GetRepeater"

	fields, format := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil, err
	}

	if field.Type != "repeater" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Field is not a repeater", Operation: op, Err: fmt.Errorf("field with the name: %s, is not a repeater", name)}
	}

	return s.getFieldChildren(field.UUID, fields, format), nil
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
func (r Repeater) SubField(name string) interface{} {
	for _, sub := range r {
		if name == sub.Name {
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
