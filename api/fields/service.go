package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"sort"
)

type Service struct {
	store  *models.Store
	postId int
	Fields []domain.PostField
	Layout *[]domain.FieldGroup
}

func (s *Service) handleArgs(args []interface{}) ([]domain.PostField, bool) {

	switch len(args) {
	case 1:
		fields := s.getFieldsByPost(args[0])
		if fields == nil {
			return s.Fields, true
		}
		return fields, true
	case 2:
		format, err := cast.ToBoolE(args[1])
		if err != nil {
			format = true
		}
		fields := s.getFieldsByPost(args[0])
		if fields == nil {
			return s.Fields, true
		}
		return fields, format
	default:
		return s.Fields, true
	}
}

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

func (s *Service) getLayoutByPost(id interface{}) []domain.FieldGroup {
	const op = "FieldsService.getFieldsByPost"

	i, err := cast.ToIntE(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INVALID, Message: "Unable to cast Post ID to integer", Operation: op, Err: err},
		}).Error()
	}

	p, err := s.store.Posts.GetById(i)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return nil
	}

	t, err := s.store.Posts.Format(p)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return nil
	}

	return *t.Layout
}

func (s *Service) findFieldByName(name string, fields []domain.PostField) (domain.PostField, error) {
	const op = "FieldsService.findFieldByName"
	for _, field := range fields {
		if name == field.Name {
			return field, nil
		}
	}
	return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: "Field does not exist", Operation: op, Err: fmt.Errorf("no field exists with the name: %s", name)}
}

func (s *Service) getFieldChildren(uniq uuid.UUID, fields []domain.PostField, format bool) []domain.PostField {
	var pf []domain.PostField
	for _, field := range fields {
		parent := field.Parent
		if parent != nil && uniq == *parent {
			if !format {
				pf = append(pf, field)
				continue
			}
			pf = append(pf, s.resolveField(field))
		}
	}
	return s.sortFields(pf)
}

func (s *Service) sortFields(pf []domain.PostField) []domain.PostField {
	sort.SliceStable(pf, func(i, j int) bool {
		return pf[i].Index < pf[j].Index
	})
	return pf
}
