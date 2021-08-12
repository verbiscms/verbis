// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"time"
)

// Preview
//
// Returns http.StatusOK if the roles were obtained successfully.
// Returns http.StatusInternalServerError if there was an error getting the roles.
func (e *Editor) Preview(ctx *gin.Context) {
	const op = "EditorHandler.Preview"

	var post domain.PostCreate
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var cat *domain.Category
	if post.Category != nil {
		find, err := e.Store.Categories.Find(*post.Category)
		if err == nil {
			cat = &find
		}
	}

	author, err := e.Store.User.Find(post.Author)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Error obtaining authro information", nil)
	}

	datum := domain.PostDatum{
		Post:     post.Post,
		Author:   author.UserPart,
		Category: cat,
		Fields:   post.Fields,
	}

	e.Cache.Set(ctx, "editor-preview-"+post.Slug, datum, cache.Options{Expiration: time.Hour * 1})

	url := e.Options.SiteURL + "/" + post.Slug + "?preview=true"

	api.Respond(ctx, http.StatusOK, "Successfully generated preview link", url)
}
