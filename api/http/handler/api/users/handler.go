// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for users to interact with the server
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

// Users defines the handler for all user routes.
type Users struct {
	*deps.Deps
}

// New
//
// Creates a new users handler.
func New(d *deps.Deps) *Users {
	return &Users{
		Deps: d,
	}
}

// clearCache
//
// TODO: This needs to be in a model, or the cache store.
//
// Clear the post cache that have the given user ID
// attached to it.
//func (u *Users) clearCache(id int) {
//	go func() {
//		p, _, err := u.Store.Posts.List(params.Params{LimitAll: true}, false, posts.ListConfig{})
//		if err != nil {
//			logger.WithError(err).Error()
//		}
//		cache.ClearUserCache(id, p)
//	}()
//}
