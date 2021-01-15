package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"strings"
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


type resolve struct {
	Key string
	Index int
	Field domain.PostField
	Fields []domain.PostField
	*Service
}

func (r *resolve) fieldAppender(appender func(domain.PostField)) {

	pipe := r.Key + r.Field.Name + SEPARATOR + cast.ToString(r.Index)

	for _, v := range r.Fields {

		pipeLen := strings.Split(pipe, SEPARATOR)
		keyLen := strings.Split(v.Key, SEPARATOR)

		if strings.HasPrefix(v.Key, pipe) && len(pipeLen) + 1 == len(keyLen) {

			fieldType := v.Type
			if fieldType == "repeater" {
				v.Value = r.resolveRepeater(pipe+SEPARATOR, v, r.Fields)
				appender(v)
				return
			}

			if fieldType == "flexible" {
				v.Value = r.resolveFlexible(pipe+SEPARATOR, v, r.Fields)
				appender(v)
				return
			}

			appender(r.resolveField(v))
		}
	}
}