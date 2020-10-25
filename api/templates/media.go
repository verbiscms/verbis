package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// getMedia obtains the media by ID and returns a domain.Media type
// or nil if not found.
func (t *TemplateFunctions) getMedia(id float64) *domain.Media {
	m, err := t.store.Media.GetById(int(id))
	if err != nil {
		return nil
	}
	return &m
}
