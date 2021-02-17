// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import "github.com/gin-gonic/gin"

func ApiParams(g *gin.Context, def Defaults) *Params {
	p := &Params{
		Stringer: &apiParams{ctx: g},
		defaults: def,
	}
	return p
}

type apiParams struct {
	ctx *gin.Context
}

func (a *apiParams) Param(q string) string {
	return a.ctx.Query(q)
}
