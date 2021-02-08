// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

type ErrorHandler interface {
	NotFound(g *gin.Context)
}

type Errors struct {
	*deps.Deps
}

func (r *Render) NotFound(g *gin.Context) {

	exec := r.Tmpl().Prepare(tpl.Config{
		Root:      paths.Theme(),
		Extension: r.Theme.FileExtension,
		Master:    "",
	})

	err := exec.ExecutePost(g.Writer, "404", g, &domain.PostData{})

	if err != nil {
		color.Green.Println(err)
	}

	return
}
