// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Struct
//
// Returns the options struct for use in the API.
// Logs errors.INTERNAL and panics if any condition failed.
func (s *Store) Struct() domain.Options {
	const op = "OptionStore.GetStruct"

	m, err := s.Map()
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
		return domain.Options{}
	}

	mOpts, err := json.Marshal(m)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
		return domain.Options{}
	}

	var options domain.Options
	if err := json.Unmarshal(mOpts, &options); err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
		return domain.Options{}
	}

	return options
}
