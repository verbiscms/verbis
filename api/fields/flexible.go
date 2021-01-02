package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
)

type Flexible []Layout

type Layout struct {
	Name      string
	SubFields SubFields
}

type SubFields []domain.PostField

func (s *Service) GetFlexible(name string, args ...interface{}) (Flexible, error) {
	const op = "FieldsService.GetFlexible"

	fields, format := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil, err
	}

	if field.Type != "flexible" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Field is not flexible content", Operation: op, Err: fmt.Errorf("field with the name: %s, is not flexible content", name)}
	}

	layouts, err := cast.ToStringSliceE(field.Value)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to obtain field layouts", Operation: op, Err: err}
	}

	return s.getLayouts(fields, layouts, format), nil
}

func (s *Service) getLayouts(fields []domain.PostField, layouts []string, format bool) Flexible {
	var flexible []Layout
	for _, v := range layouts {
		var sub SubFields
		for _, field := range fields {
			if field.Layout != nil && *field.Layout == v {
				if !format {
					sub = append(sub, field)
					continue
				}
				sub = append(sub, s.resolveField(field))
			}
		}
		flexible = append(flexible, Layout{Name: v, SubFields: sub})
	}
	return flexible
}

// HasRows
//
// Determines if the Flexible content has any rows.
func (f Flexible) HasRows() bool {
	return len(f) != 0
}

// SubField
//
// Returns a sub field by key or nil if it wasn't found.
func (s SubFields) SubField(name string) interface{} {
	for _, sub := range s {
		if name == sub.Name {
			return sub.Value
		}
	}
	return nil
}

// First
//
// Returns the first element in the sub fields, or nil if
// the length of the sub fields is zero.
func (s SubFields) First() interface{} {
	if len(s) == 0 {
		return nil
	}
	return s[0]
}

// Last
//
// Returns the last element in the sub fields, or nil if
// the length of the sub fields is zero.
func (s SubFields) Last() interface{} {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}
