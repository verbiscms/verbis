package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

// getMedia obtains the media by ID and returns a domain.Media type
// or nil if not found.
func (t *TemplateFunctions) getMedia(i interface{}) *domain.Media {
	var id int
	switch i.(type) {
	case *int, *float64:
		p := i.(*int)
		if p != nil {
			id = *p
		}
		break
	default:
		id = i.(int)
		break
	}
	m, err := t.store.Media.GetById(id)

	if err != nil {
		return nil
	}
	return &m
}
