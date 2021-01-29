package media

import (
	"github.com/spf13/cast"
)

// Find
//
// Obtains the media by ID and returns a domain.Media type
// or nil if not found or the ID parameter failed to be
// parsed.
//
// Example:
// {{ $image := media 10 }}
// {{ $image.Url }}
func (ns *Namespace) Find(i interface{}) interface{} {
	if i == nil {
		return nil
	}

	id, err := cast.ToIntE(i)
	if err != nil {
		return nil
	}

	m, err := ns.deps.Store.Media.GetById(id)
	if err != nil {
		return nil
	}

	return m
}
