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
func (s *Service) GetRepeater(key string) (Repeater, error) {
	const op = "Fields.GetRepeater"

	field, err := s.findByKey(key)
	if err != nil {
		return nil, err
	}

	if field.Type != "repeater" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Field is not a repeater", Operation: op, Err: fmt.Errorf("field with the key: %s, is not a repeater", key)}
	}

	var repeater Repeater = s.getChildren(field.UUID)

	return repeater, nil
}


// HasRows
//
func (r Repeater) HasRows() bool {
	return len(r) != 0
}

// SubField
//
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
func (r Repeater) First() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[0]
}

// Last
//
func (r Repeater) Last() interface{} {
	if len(r) == 0 {
		return nil
	}
	return r[len(r)-1]
}
