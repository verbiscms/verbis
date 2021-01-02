package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

type Flexible []Layout

type Layout struct {
	Name string
	SubFields []domain.PostField
}

func (s *Service) GetFlexible(key string, args ...interface{}) (Flexible, error) {
	const op = "FieldService.GetFlexible"

	fields := s.handleArgs(args)

	field, err := s.findFieldByKey(key, fields)
	if err != nil {
		return Flexible{}, err
	}

	if field.Type != "Flexible" {
		return Flexible{}, &errors.Error{Code: errors.INVALID, Message: "Field is not flexible content", Operation: op, Err: fmt.Errorf("field with the key: %s, is not flexible content", key)}
	}

	var flexible []Layout
	// Layouts
	for _, l := range s.getFieldChildren(field.UUID, fields) {
		// Sub Fields
		var sub []domain.PostField
		for _, s := range s.getFieldChildren(l.UUID, fields) {
			sub = append(sub, s)
		}
		flexible = append(flexible, Layout{
			Name:      l.Key,
			SubFields: sub,
		})
	}

	return flexible, nil
}