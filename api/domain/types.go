// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
)

type (
	// DBMap defines the helper for unmarshalling into a
	// map directly from the database.
	DBMap map[string]interface{}
)

// Scan
//
// Scanner for DBMap. unmarshal the DBMap when
// the entity is pulled from the database.
func (m DBMap) Scan(value interface{}) error {
	const op = "Domain.DBMap.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for DBMap", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into DBMap", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for DBMap. marshal the DBMap when
// the entity is inserted to the database.
func (m DBMap) Value() (driver.Value, error) {
	const op = "Domain.MediaSizes.DBMap"
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling DBMap", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
