// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// choice defines the properties to be sent to the template
// if the return format is set as a 'map' for choice
// fields (select, radio group etc).
type choice struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// choice
//
// Unmarshalls the domain.FieldValue into a choice type.
// Returns errors.INVALID if the unmarshal was not successful.
func (v *Value) choice(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.choice"

	var c choice
	err := json.Unmarshal([]byte(value), &c)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Unable to unmarshal to choice map"), Operation: op, Err: err}
	}

	return c, nil
}
