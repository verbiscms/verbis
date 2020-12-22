package templates

import (
	"fmt"
)

// dict allows to pass multiple values to templates to use inside a
// template call for use with the post loop.
func (t *TemplateFunctions) dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("Invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("Dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
