package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

type Fields map[string]interface{}

// GetFields
//
//
func (s *Service) GetFields(args ...interface{}) (Fields, error) {
	fields, format := s.handleArgs(args)

	f := make(Fields, len(fields))
	s.mapper(fields, func(field domain.PostField) {
		if !format {
			f[field.Name] = field.Value
			return
		}
		f[field.Name] = s.resolveField(field).Value
	})

	return f, nil
}

type WalkerFunc func(field domain.PostField)

func (s *Service) mapper(fields []domain.PostField, walkerFunc WalkerFunc) {
	for _, field := range fields {

		if field.Parent != nil || field.Layout != nil {
			continue
		}

		if field.Type == "repeater" {
			if repeater, err := s.GetRepeater(field.Name); err == nil {
				field.Value = repeater
				walkerFunc(field)
				continue
			}
		}

		if field.Type == "flexible" {
			repeater, err := s.GetFlexible(field.Name); if err == nil {
				field.Value = repeater
				walkerFunc(field)
				continue
			}
		}

		walkerFunc(field)
	}
	return
}