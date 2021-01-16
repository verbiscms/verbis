package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// media
//
// Resolves a media from the given value.
// Returns the domain.Media if it was found and no error occurred.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an integer.
func (v *Value) media(value domain.FieldValue) (interface{}, error) {
	const op = "Value.Media"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
	}

	media, err := v.store.Media.GetById(id)
	if err != nil {
		return nil, err
	}

	return media, nil
}