package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"reflect"
)

// getMedia obtains the media by ID and returns a domain.Media type
// or nil if not found.
func (t *TemplateFunctions) getMedia(i interface{}) *domain.Media {
	var id int

	fmt.Println(reflect.TypeOf(i))

	switch i.(type) {
	default:
		id = i.(int)
	case *int:
		p := i.(*int)
		if p != nil {
			id = *p
		}
	case *float64: {
		p := i.(*float64)
		if p != nil {
			id = int(*p)
		}
	}
	case float64:
		id = int(i.(float64))
	}

	m, err := t.store.Media.GetById(id)

	if err != nil {
		return nil
	}
	return &m
}
