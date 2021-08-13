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
	const op = "Menus.Args.ToOptions"

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
	// The menu ID defined in the theme configuration file
	// this is the only parameter that is required
	// for lookup of a navigation menu.
	Menu string `json:"menu" binding:"required"`
	// MenuClass is the CSS class that is wrapped by the
	// <uL> HTML element.
	MenuClass string `json:"menu_class"`
	// LiClass is the CSS class that is wrapped by the
	// <li> HTML element.
	LiClass string `json:"li_class"`
	// LinkClass is the CSS class that is wrapped by the
	// <a> HTML element.
	LinkClass string `json:"link_class"`
	// Depth defines how many levels of the hierarchy are
	// to be included. 0 means all, which is the default.
	// However, the depth must be greater than 0 or
	// the options will fail validation.
	Depth int `json:"depth" binding:"gte=0"`
	// Partial is a file name of a partial to be used to
	// execute the template.
	Partial string `json:"partial"`
}

// Validate validates the options' struct to ensure
// the correct fields are parsed.
func (o *Options) Validate() error {
	return validation.Validator().Struct(o)
}
