// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	NotFound(g *gin.Context)
}

type Errors struct {
	*deps.Deps
}

func (e *Errors) NotFound(g *gin.Context) {

	exec := e.Tmpl().Prepare(tpl.Config{
		Root:      paths.Theme(),
		Extension: e.Theme.FileExtension,
		Master:    "",
	})

	err := exec.Execute(g.Writer, "404", nil)

	if err != nil {
		panic(err)
	}

	return
}
