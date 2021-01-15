package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// resolveField
//
// Determines if the given field values is a slice or array or singular
// and returns a resolved field value or a slice of interfaces.
func (s *Service) resolveField(field domain.PostField) domain.PostField {
	original := field.OriginalValue

	if original.IsEmpty() {
		return field
	}

	if !original.IsArray() {
		field.Value = s.resolveValue(original.String(), field.Type)
		return field
	}

	var items []interface{}
	for _, v := range original.Array() {
		items = append(items, s.resolveValue(v, field.Type))
	}
	field.Value = items

	return field
}




type fieldValue struct {
	Service
}

type valuer func(value domain.FieldValue) (interface{}, error)

type fieldValueMap map[string]valuer


func getMap() {
	v := &fieldValue{}

	// TODO: Posts, Choice (Tags, Radio, Button, Select)
	// Checkbox

	mm := fieldValueMap{
		"number" : v.Number,
		"media": v.Media,
		"range": v.Number,
		"checkbox": v.Checkbox,
		"categories": v.Categories,
		"user": v.User,
	}

	fmt.Println(mm["number"]("ff"))
}

// Number
//
//
func (f *fieldValue) Number(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Number"
	number, err := cast.ToInt64E(value)
	if err != nil {
		// TODO CHANGE TYPE HERE DYNAMIC
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast TODO CHANGE TYPE HERE DYNAMIC field to an integer", Operation: op, Err: err}
	}
	return number, nil
}

// Checkbox
//
//
func (f *fieldValue) Checkbox(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Checkbox"
	check, err := cast.ToBoolE(value)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast checkbox field to an bool", Operation: op, Err: err}
	}
	return check, nil
}

// Media
//
//
func (f *fieldValue) Media(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.User"

	var mediaItems []domain.Media
	for _, v := range value.Array() {

		id, err := cast.ToIntE(v)
		if err != nil {
			return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
		}

		media, err := f.store.Media.GetById(id)
		if err != nil {
			return nil, err
		}

		mediaItems = append(mediaItems, media)
	}

	if len(mediaItems) == 1 {
		return mediaItems[0], nil
	}

	return mediaItems, nil
}

// Categories
//
//
func (f *fieldValue) Categories(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.User"

	var categories []domain.Category
	for _, v := range value.Array() {

		id, err := cast.ToIntE(v)
		if err != nil {
			return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast category ID to an integer", Operation: op, Err: err}
		}

		category, err := f.store.Categories.GetById(id)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if len(categories) == 1 {
		return categories[0], nil
	}

	return categories, nil
}

// User
//
//
func (f *fieldValue) User(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.User"

	var users []domain.UserPart
	for _, v := range value.Array() {

		id, err := cast.ToIntE(v)
		if err != nil {
			return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
		}

		user, err := f.store.User.GetById(id)
		if err != nil {
			return nil, err
		}

		users = append(users, *user.HideCredentials())
	}

	if len(users) == 1 {
		return users[0], nil
	}

	return users, nil
}


// resolveValue
//
// Switches the fields type and resolves the field value accordingly.
// If there was an error resolving the value, the original field
// will be returned.
// Resolves categories, images, posts and users from the ID
// to the type.
func (s *Service) resolveValue(value string, typ string) interface{} {
	var e error
	var r interface{} = value

	switch typ {
	case "number", "range":
		number, err := cast.ToInt64E(value)
		if err != nil {
			e = err
		}
		r = number
		// TODO: Need to cast numbers to integers

	case "category":
		category, err := s.store.Categories.GetById(cast.ToInt(value))
		if err != nil {
			e = err
		}
		r = category
	case "image":
		media, err := s.store.Media.GetById(cast.ToInt(value))
		if err != nil {
			e = err
		}
		r = media
	case "post":
		post, err := s.store.Posts.GetById(cast.ToInt(value))
		if err != nil {
			e = err
			break
		}
		formatPost, err := s.store.Posts.Format(post)
		if err != nil {
			e = err
		}
		r = formatPost
	case "user":
		user, err := s.store.User.GetById(cast.ToInt(value))
		if err != nil {
			e = err
		}
		r = *user.HideCredentials()
	default:
		return value
	}

	if e != nil {
		log.WithFields(log.Fields{"error": e}).Error()
		return nil
	}

	return r
}
