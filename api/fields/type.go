package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"reflect"
)

func (s *Service) resolveField(field domain.PostField) domain.PostField {
	typ := reflect.TypeOf(field.Value)

	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		field.Value = s.resolveValue(field.Value, field.Type)
		return field
	}

	val := reflect.ValueOf(field.Value)
	var items []interface{}
	for i := 0; i < val.Len(); i++ {
		element := val.Index(i)
		items = append(items, s.resolveValue(element.Interface(), field.Type))
	}
	field.Value = items

	return field
}

func (s *Service) resolveValue(value interface{}, typ string) interface{} {
	var e error
	var r = value

	switch typ {
	case "category":
		category, err := s.store.Categories.GetById(cast.ToInt(value))
		if err != nil {
			e = err
			break
		}
		r = category
		break
	case "image":
		media, err := s.store.Media.GetById(cast.ToInt(value))
		if err != nil {
			e = err
			break
		}
		r = media
		break
	case "post":
		post, err := s.store.Posts.GetById(cast.ToInt(value))
		if err != nil {
			e = err
			break
		}
		formatPost, err := s.store.Posts.Format(post)
		if err != nil {
			e = err
			break
		}
		r = formatPost
		break
	case "user":
		user, err := s.store.User.GetById(cast.ToInt(value))
		if err != nil {
			e = err
			break
		}
		r = *user.HideCredentials()
		break
	default:
		return value
	}

	if e != nil {
		log.WithFields(log.Fields{"error": e}).Error()
		return nil
	}

	return r
}
