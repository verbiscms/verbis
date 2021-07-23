// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dict

import (
	"fmt"
	"github.com/verbiscms/verbis/api/errors"
)

// Dict
//
// Allows to pass multiple values to templates to use inside a
// template call for use with the post loop or partial
// calls.
//
// Returns errors.TEMPLATE if the dict values are not divisible by two or
// any dict keys were not strings.
//
// Example: {{ dict "colour" "green" "height" 20 }}
// Returns: map[string]interface{}{"colour":"green", "height":20}
func (ns *Namespace) Dict(values ...interface{}) (map[string]interface{}, error) {
	const op = "Templates.Dict"

	if len(values)%2 != 0 {
		return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Invalid dict call", Operation: op, Err: fmt.Errorf("dict values are not divisable by two")}
	}
	dict := make(map[string]interface{}, len(values)/2) //nolint

	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, &errors.Error{Code: errors.TEMPLATE, Message: "Dict keys must be strings", Operation: op, Err: fmt.Errorf("dict keys passed are not strings")}
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
