package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/fields/resolve"
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

	if len(args) == 0 {
		return s.fields
	}

	switch f := args[0].(type) {
	case domain.PostData:
		return f.Fields
	case []domain.PostField:
		return f
	case domain.PostTemplate:
		return f.Fields
	}

	id, err := cast.ToIntE(args[0])
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INVALID, Message: "Invalid argument passed to ", Operation: op, Err: fmt.Errorf("unable to cast post id to integer")},
		}).Error()
		return nil
	}

	return s.getFieldsByPost(id)

}

// getFieldsByPost
//
// Returns the fields by Post with the given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining the post.
func (s *Service) getFieldsByPost(id int) []domain.PostField {

	fields, err := s.deps.Store.Fields.GetByPost(id)
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

// walker represents the struct to be passed when resolving
// repeaters and flexible content types.
type walker struct {
	Key    string
	Index  int
	Field  domain.PostField
	Fields []domain.PostField
	*Service
}

// Walk
//
// Constructs a pipe based on the key, name, SEPARATOR and the index
// in order to look up dynamic Flexible content and Repeater
// types. The key `flex|0|repeater|0|text` will be split
// and looked up. If the child value is of type Repeater
// or Flexible the function will call itself meaning
// all values will be resolved.
// The appender func outputs the field to the caller once resolved.
func (r *walker) Walk(appender func(domain.PostField)) {

	pipe := r.Key + r.Field.Name + SEPARATOR + cast.ToString(r.Index)

	for _, v := range r.Fields {

		pipeLen := strings.Split(pipe, SEPARATOR)
		keyLen := strings.Split(v.Key, SEPARATOR)

		if strings.HasPrefix(v.Key, pipe) && len(pipeLen)+1 == len(keyLen) {

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

			appender(resolve.Field(v, r.deps))
		}
	}
}
