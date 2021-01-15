package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"strings"
)

// Fields defines the map of fields to be returned to the template.
type Fields map[string]interface{}

// GetFields
//
// Returns all of the fields for the current post, or post ID given.
func (s *Service) GetFields(args ...interface{}) Fields {
	fields := s.handleArgs(args)

	var f = make(Fields, 0)
	s.mapper(fields, func(field domain.PostField) {
		f[field.Name] = field.Value
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
func (s *Service) mapper(fields []domain.PostField, walkerFunc WalkerFunc) {
	for _, field := range fields {

		if field.Type == "repeater" {
			repeater, err := s.GetRepeater(field.Name)
			if err == nil {
				field.Value = repeater
				walkerFunc(field)
			}
			continue
		}

		if field.Type == "flexible" {
			flexible, err := s.GetFlexible(field.Name)
			if err == nil {
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
