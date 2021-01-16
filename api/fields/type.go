package fields

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type fieldValue struct {
	*Service
}

type valuer func(field domain.FieldValue) (interface{}, error)

type fieldValueMap map[string]valuer

// resolveField
//
// Determines if the given field values is a slice or array or singular
// and returns a resolved field value or a slice of interfaces.
func (s *Service) resolveField(field domain.PostField) domain.PostField {
	exec := &fieldValue{s}
	resolved := exec.Resolve(field)
	return resolved
}

func (f *fieldValue) GetMap() fieldValueMap {
	return fieldValueMap{
		"category": f.Category,
		"checkbox": f.Checkbox,
		"choice":   f.Choice,
		"image":    f.Media,
		"number":   f.Number,
		"post":     f.Post,
		"range":    f.Number,
		"user":     f.User,
	}
}

func (f *fieldValue) Resolve(field domain.PostField) domain.PostField  {
	original := field.OriginalValue

	if original.IsEmpty() && field.Key != "map" {
		field.Value = field.OriginalValue.String()
		return field
	}

	if !original.IsArray() {
		field.Value = f.Execute(field.OriginalValue.String(), field.Type)
		return field
	}

	var items []interface{}
	for _, v := range original.Array() {
		items = append(items, f.Execute(v, field.Type))
	}
	field.Value = items

	return field
}

func (f *fieldValue) Execute(value string, typ string) interface{} {
	fn, ok := f.GetMap()[typ]
	if !ok {
		return value
	}

	val, err := fn(domain.FieldValue(value))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
	}

	return val
}

// Categories
//
//
func (f *fieldValue) Category(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Category"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast category ID to an integer", Operation: op, Err: err}
	}

	category, err := f.store.Categories.GetById(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// Checkbox
//
//
func (f *fieldValue) Checkbox(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Checkbox"

	check, err := cast.ToBoolE(value.String())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast checkbox field to an bool", Operation: op, Err: err}
	}

	return check, nil
}

// Choice
//
//
func (f *fieldValue) Choice(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Choice"

	var c = struct {
		key   string
		value string
	}{}

	err := json.Unmarshal([]byte(value), &c)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Unable to unmarshal to choice map"), Operation: op, Err: err}
	}

	return c, nil
}

// Media
//
//
func (f *fieldValue) Media(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Media"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
	}

	media, err := f.store.Media.GetById(id)
	if err != nil {
		return nil, err
	}

	return media, nil
}

// Number
//
//
func (f *fieldValue) Number(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Number"

	number, err := cast.ToInt64E(value.String())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast field to an integer", Operation: op, Err: err}
	}

	return number, nil
}

// Post
//
//
func (f *fieldValue) Post(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Post"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast post ID to an integer", Operation: op, Err: err}
	}

	post, err := f.store.Posts.GetById(id)
	if err != nil {
		return nil, err
	}

	formatPost, err := f.store.Posts.Format(post)
	if err != nil {
		return nil, err
	}

	return formatPost, nil
}

// User
//
//
func (f *fieldValue) User(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.User"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
	}

	user, err := f.store.User.GetById(id)
	if err != nil {
		return nil, err
	}

	return *user.HideCredentials(), nil
}