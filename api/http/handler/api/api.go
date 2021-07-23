// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/gin-gonic/gin"
)

const (
	// DefaultPage is the default page number when none is
	// passed.
	DefaultPage = 1
	// DefaultLimit is the default limit when none is
	// passed.
	DefaultLimit = 15
)

var (
	// DefaultParams represents the default params if
	// none were passed for the API.
	DefaultParams = params.Defaults{
		Page:           DefaultPage,
		Limit:          DefaultLimit,
		OrderBy:        "created_at",
		OrderDirection: "desc",
	}
)

// Params sets up a new Params struct with context.
func Params(g *gin.Context) *params.Params {
	return params.New(&apiParams{ctx: g}, DefaultParams)
}

// apiParams defines the helper for returning context parameters.
type apiParams struct {
	ctx *gin.Context
}

// Param satisfies the Stringer interface by returning a
// query parameters to pass information to models.
func (a *apiParams) Param(q string) string {
	return a.ctx.Query(q)
}
