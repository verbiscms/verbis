// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gookit/color"
)

var (
	// opts is an alias for a singleton for the main Verbis
	// options.
	opts = &domain.Options{}
	//// once ensures the options are set only once upon
	//// initialisation.
	//once = &sync.Once{}
	// TODO check this only returns once!
)

// Struct
//
// Returns the options struct for use in the API.
// Logs errors.INTERNAL and panics if any condition failed.
func (s *Store) Struct() *domain.Options {
	const op = "OptionStore.Struct"

	m, err := s.Map()
	if err != nil {
		color.Red.Println(err)
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
	}

	mOpts, err := json.Marshal(m)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
	}

	err = json.Unmarshal(mOpts, opts)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error getting options", Operation: op, Err: err}).Panic()
	}

	return opts
}
