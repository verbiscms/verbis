// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/spf13/cast"
)

// Query defines the map of arguments passed to
// list functions in templates.
type Query map[string]interface{}

var (
	// Defaults represents the default params if
	// none were passed for templates.
	Defaults = params.Defaults{
		Page:           params.DefaultPage,
		Limit:          params.DefaultLimit,
		OrderBy:        "updated_at",
		OrderDirection: "desc",
	}
)

// Get
//
// Returns parameters for the store to used for obtaining
// multiple entities. If the orderBy or orderDirection
// arguments are not passed, defaults will be used.
func (q Query) Get(orderBy, orderDirection string) params.Params {
	def := Defaults
	if orderBy != "" {
		def.OrderBy = orderBy
	}
	if orderDirection != "" {
		def.OrderDirection = orderDirection
	}
	return params.New(q, def).Get()
}

// Param
//
// Is an implementation of a stringer to return
// parameters from the Query map.
func (q Query) Param(param string) string {
	val, ok := q[param]
	if !ok {
		return ""
	}
	s, err := cast.ToStringE(val)
	if err != nil {
		return ""
	}
	return s
}

// Default
//
// Sets or gets default parameters for the Query.
// If the parameter is not found, it will
// return the default string passed.
func (q Query) Default(param, def string) interface{} {
	val, ok := q[param]
	if !ok {
		return def
	}
	return val
}
