package fields

import "github.com/ainsleyclark/verbis/api/domain"

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

	return s.resolveField(field).Value, nil
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

	return s.resolveField(field), nil
}
