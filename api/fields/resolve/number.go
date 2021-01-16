package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
)

// number
//
// Casts the domain.FieldValue to int64.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an int64.
func (v *Value) number(value domain.FieldValue) (interface{}, error) {
	const op = "Value.Number"

	number, err := cast.ToInt64E(value.String())
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast field to an integer", Operation: op, Err: err}
	}

	return number, nil
}