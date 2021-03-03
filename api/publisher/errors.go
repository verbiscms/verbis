// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	NotFound(g *gin.Context)
}

type Errors struct {
	*deps.Deps
}

func (r *publish) NotFound(g *gin.Context) {
	exec := r.Tmpl().Prepare(tpl.Config{
		Root:      r.ThemePath(),
		Extension: r.Config.FileExtension,
		Master:    "",
	})

	_, err := exec.ExecutePost(g.Writer, "404", g, &domain.PostDatum{})

	if err != nil {
		logger.WithError(err).Error()
	}
}
