// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	Store *cache.Cache
)

type Cacher interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Flush()
}

const (
	// For use with functions that take an expiration time.
	RememberForever   time.Duration = -1
	postIDKey         string        = "post-id-"
	DefaultExpiration               = 5 * time.Minute
	DefaultCleanup                  = 10 * time.Minute
)

// Init set-ups go-cache with defaults
func Init() {
	Store = cache.New(DefaultExpiration, DefaultCleanup)
}

func ClearPostCache(id int) {
	Store.Delete(GetPostKey(id))
}

func ClearUserCache(userID int, posts domain.PostData) {
	for _, v := range posts {
		if v.UserId == userID {
			ClearPostCache(v.Id)
		}
	}
}

func ClearCategoryCache(categoryID int, posts domain.PostData) {
	for _, v := range posts {
		if v.Category.Id == categoryID {
			ClearPostCache(v.Id)
		}
	}
}

func GetPostKey(id int) string {
	return fmt.Sprintf("%s%d", postIDKey, id)
}
