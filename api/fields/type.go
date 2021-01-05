package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
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
