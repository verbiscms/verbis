package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/fields/resolve"
)

// GetField
//
// Returns the value of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetField(name string, args ...interface{}) (interface{}, error) {
	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil, err
	}

	resolved := resolve.Field(field, s.store)

	return resolved.Value, nil
}

// GetFieldObject
//
// Returns the raw object of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetFieldObject(name string, args ...interface{}) (domain.PostField, error) {
	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return domain.PostField{}, err
	}

	resolved := resolve.Field(field, s.store)

	return resolved, nil
}
