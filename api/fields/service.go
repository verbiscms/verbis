package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/google/uuid"
	"sort"
)

type Service struct {
	store  *models.Store
	postId int
	Fields []domain.PostField
}

func (s *Service) findByKey(key string) (domain.PostField, error) {
	const op = "Fields.findByKey"
	for _, field := range s.Fields {
		if key == field.Key {
			return field, nil
		}
	}
	return domain.PostField{}, &errors.Error{Code: errors.NOTFOUND, Message: "Field does not exist", Operation: op, Err: fmt.Errorf("no field exists with the key: %s", key)}
}

func (s *Service) getChildren(uniq uuid.UUID) []domain.PostField {
	var pf []domain.PostField
	for _, field := range s.Fields {
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