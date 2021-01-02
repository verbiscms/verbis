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
}

func (s *Service) handleArgs(args ...interface{}) []domain.PostField {
	const op = "FieldsService.handleArgs"

	if len(args) == 0 {
		return s.Fields
	}

	if len(args) == 1 {
		id, err := cast.ToIntE(args[0])
		if err != nil {
			return s.Fields
		}

		fields, err := s.store.Fields.GetByPost(id)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error()
			return s.Fields
		}
		return fields
	}

	return s.Fields
}

func (s *Service) findFieldByKey(key string, fields []domain.PostField) (domain.PostField, error) {
	const op = "FieldsService.findByKey"
	for _, field := range fields {
		if key == field.Key {
			return field, nil
		}
	}
	return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: "Field does not exist", Operation: op, Err: fmt.Errorf("no field exists with the key: %s", key)}
}

func (s *Service) getFieldChildren(uniq uuid.UUID, fields []domain.PostField) []domain.PostField {
	var pf []domain.PostField
	for _, field := range fields {
		parent := field.Parent
		if parent != nil && uniq == *parent {
			f := s.resolveField(field)
			if f.Value != nil {
				pf = append(pf, f)
			}
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
