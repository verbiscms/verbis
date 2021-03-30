// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
)

// unmarshal
//
// Unmarshal's the option value.
// Returns errors.INTERNAL if the unmarshalling failed
func (s *Store) unmarshal(value json.RawMessage) (interface{}, error) {
	const op = "OptionStore.unmarshal"

	var v interface{}
	err := json.Unmarshal(value, &v)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling the option value", Operation: op, Err: err}
	}

	return value, nil
}

// marshal
//
// Marshals the option value.
// Returns errors.INTERNAL if the marshalling failed
func (s *Store) marshal(value interface{}) (json.RawMessage, error) {
	const op = "OptionStore.marshal"

	m, err := json.Marshal(value)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling the option value", Operation: op}
	}

	return m, nil
}
