// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Handler defines methods for users to interact with the server
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Roles(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

// Users defines the handler for all user routes.
type Users struct {
	*deps.Deps
}

// clearCache
//
// TODO: This needs to be in a model, or the cache store.
//
// Clear the post cache that have the given user ID
// attached to it.
func (u *Users) clearCache(id int) {
	go func() {
		posts, _, err := u.Store.Posts.Get(params.Params{LimitAll: true}, false, "", "")
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error()
		}
		cache.ClearUserCache(id, posts)
	}()
}
