package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

func (t *TemplateFunctions) getMedia(id float64) *domain.Media {
	m, err := t.store.Media.GetById(int(id))
	if err != nil {
		return nil
	}
	return &m
}
