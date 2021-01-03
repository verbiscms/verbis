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
	Walker(s, fields, func(field domain.PostField) {
		if !format {
			f[field.Name] = field.Value
			return
		}
		f[field.Name] = s.resolveField(field)
	})

	return f, nil
}

type WalkerFunc func(field domain.PostField)

func Walker(service *Service, fields []domain.PostField, walkerFunc WalkerFunc) {
	for _, field := range fields {
		// The type is not a repeater of flexible content
		if field.Parent == nil || field.Layout != nil {
			walkerFunc(field)
		}
		// Check repeater values
		if repeater, err := service.GetRepeater(field.Name); err != nil {
			field.Value = repeater
			walkerFunc(field)
		}
		// Check flexible content
		if flexible, err := service.GetFlexible(field.Name); err != nil {
			field.Value = flexible
			walkerFunc(field)
		}
	}
}