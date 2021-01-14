package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// handleArgs
//
// The array of interfaces are presumed to be the following:
// [0] for Post ID, fields are obtained by the post given.
//
// Returns the fields to be modified & processed.
func (s *Service) handleArgs(args []interface{}) []domain.PostField {
	const op = "FieldsService.handleArgs"

	if len(args) == 1 {
		fields := s.getFieldsByPost(args[0])
		return fields
	}

	return s.fields
}

// getFieldsByPost
//
// Returns the fields by Post with the given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining the post.
func (s *Service) getFieldsByPost(id interface{}) []domain.PostField {
	const op = "FieldsService.getFieldsByPost"

	i, err := cast.ToIntE(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INVALID, Message: "Unable to cast Post ID to integer", Operation: op, Err: err},
		}).Error()
		return nil
	}

	fields, err := s.store.Fields.GetByPost(i)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return nil
	}

	return fields
}

// findFieldByName
//
// Returns a singular domain.PostField by the given name.
// Returns errors.NOTFOUND if the field does not exist.
func (s *Service) findFieldByName(name string, fields []domain.PostField) (domain.PostField, error) {
	const op = "FieldsService.findFieldByName"
	for _, field := range fields {
		if name == field.Name {
			return field, nil
		}
	}
	return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: "Field does not exist", Operation: op, Err: fmt.Errorf("no field exists with the name: %s", name)}
}

// getFieldChildren
//
// Loops through the given slice of domain.PostField and compares the
// uuid passed with the field's parent UUID.
// It's not necessary to use a database call for this look up, as we will
// be looping through them anyway to append and format the fields.
// Returns the sorted slice of fields.
//func (s *Service) getFieldChildren(uuid uuid.UUID, fields []domain.PostField, format bool) []domain.PostField {
//var pf []domain.PostField
//for _, field := range fields {
//	parent := field.Parent
//	if parent != nil && uuid == *parent {
//		if !format {
//			pf = append(pf, field)
//			continue
//		}
//		pf = append(pf, s.resolveField(field))
//	}
//}
//return nil
//return s.sortFields(pf)
//}

// sortFields
//
// Sort's the slice of domain.PostFields by Index, used for
// the service.GetRepeater function
//func (s *Service) sortFields(pf []domain.PostField) []domain.PostField {
//	sort.SliceStable(pf, func(i, j int) bool {
//		lt := pf[i].Index
//		gt := pf[j].Index
//		zero := 0
//		if lt == nil {
//			lt = &zero
//		}
//		if gt == nil {
//			gt = &zero
//		}
//		return *lt < *gt
//	})
//	return pf
//}
