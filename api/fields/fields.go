package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// Fields defines the map of fields to be returned to the template.
type Fields map[string]interface{}

// GetFields
//
// Returns all of the fields for the current post, or post ID given.
func (s *Service) GetFields(args ...interface{}) (Fields, error) {
	fields, format := s.handleArgs(args)

	f := make(Fields, len(fields))
	s.mapper(fields, func(field domain.PostField) {
		if !format {
			f[field.Name] = field.OriginalValue.String()
			return
		}
		if field.Type == "repeater" || field.Type == "flexible" {
			f[field.Name] = field.Value
			return
		}
		f[field.Name] = s.resolveField(field).Value
	})

	return f, nil
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

		//if field.Parent != nil || field.Layout != nil {
		//	continue
		//}

		if field.Type == "repeater" {
			if repeater, err := s.GetRepeater(field.Name); err == nil {
				field.Value = repeater
				walkerFunc(field)
				continue
			}
		}

		if field.Type == "flexible" {
			flexible, err := s.GetFlexible(field.Name)
			if err == nil {
				field.Value = flexible
				walkerFunc(field)
				continue
			}
		}

		walkerFunc(field)
	}
	return
}
