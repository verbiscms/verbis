package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
)

// dict
//
// Allows to pass multiple values to templates to use inside a
// template call for use with the post loop or partial
// calls.
//
// Returns errors.TEMPLATE if the dict values are not divisible by two or
// any dict keys were not strings.
//
// Example: {{ $map := dict "colour" "green" "height" 20 }}
func (t *TemplateManager) dict(values ...interface{}) (map[string]interface{}, error) {
	const op = "Templates.dict"

	if len(values)%2 != 0 {
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Invalid dict call", Operation: op, Err: fmt.Errorf("dict values are not divisable by two")}
	}
	dict := make(map[string]interface{}, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Dict keys must be strings", Operation: op, Err: fmt.Errorf("dict keys passed are not strings")}
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
