package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
)

// checkbox
//
// Casts the domain.FieldValue to a boolean.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an bool.
func (v *Value) checkbox(value domain.FieldValue) (interface{}, error) {
	const op = "Value.Checkbox"

	check, err := cast.ToBoolE(value.String())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast checkbox field to an bool", Operation: op, Err: err}
	}

	return check, nil
}