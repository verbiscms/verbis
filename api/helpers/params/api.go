// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import "github.com/gin-gonic/gin"

// ApiParams
//
// Sets up a new Params struct with context.
func ApiParams(g *gin.Context, def Defaults) *Params {
	p := &Params{
		Stringer: &apiParams{ctx: g},
		defaults: def,
	}
	return p
}

// apiParams defines the helper for returning context parameters.
type apiParams struct {
	ctx *gin.Context
}

// Param
//
// Satisfies the Stringer interface by returning a query
// parameters to pass information to models.
func (a *apiParams) Param(q string) string {
	return a.ctx.Query(q)
}
