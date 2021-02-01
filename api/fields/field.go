package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/fields/resolve"
	log "github.com/sirupsen/logrus"
)

// GetField
//
// Returns the value of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetField(name string, args ...interface{}) interface{} {
	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return nil
	}

	resolved := resolve.Field(field, s.deps)

	return resolved.Value
}

// GetFieldObject
//
// Returns the raw object of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetFieldObject(name string, args ...interface{}) domain.PostField {
	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return domain.PostField{}
	}

	resolved := resolve.Field(field, s.deps)

	return resolved
}
