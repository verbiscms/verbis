package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// category
//
// Resolves a category from the given value.
// Returns the domain.Category if it was found and no error occurred.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an integer.
func (v *Value) category(value domain.FieldValue) (interface{}, error) {
	const op = "fieldValue.Category"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast category ID to an integer", Operation: op, Err: err}
	}

	category, err := v.store.Categories.GetById(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}