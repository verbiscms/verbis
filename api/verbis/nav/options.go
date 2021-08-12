// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"encoding/json"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
)

// Args represents the arguments for obtaining a
// Verbis navigation menu.
type Args map[string]interface{}

// ToOptions parses the arguments to the nav.Options{}
// struct. Returns an error on failed marshal
// or unmarshal.
func (a Args) ToOptions() (Options, error) {
	const op = "Nav.Args.ToOptions"
	m, err := json.Marshal(a)
	if err != nil {
		return Options{}, &errors.Error{Code: errors.INVALID, Message: "Error converting arguments to navigation options", Operation: op, Err: err}
	}
	opts := Options{}
	err = json.Unmarshal(m, &opts)
	if err != nil {
		return Options{}, &errors.Error{Code: errors.INVALID, Message: "Error converting arguments to navigation options", Operation: op, Err: err}
	}
	return opts, nil
}

// Options represents the options for obtaining a
// navigational menu.
type Options struct {
	Menu      string `json:"menu" binding:"required"`
	MenuClass string `json:"menu_class"`
	LiClass   string `json:"li_class"`
	LinkClass string `json:"link_class"`
	Depth     int    `json:"depth"`
}

// Validate validates the options struct to ensure
// the correct fields are parsed.
func (o *Options) Validate() error {
	return validation.Validator().Struct(o)
}
