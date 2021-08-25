// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"context"
	"fmt"
	"github.com/verbiscms/verbis/api/cache"
)

const (
	// FieldCacheKey is the key stored in the cache
	// for standard fields.
	FieldCacheKey = "standard"
	// RepeaterCacheKey is the key stored in the cache
	// for repeater fields.
	RepeaterCacheKey = "repeater"
	// FlexibleCacheKey is the key stored in the cache
	// for flexible fields.
	FlexibleCacheKey = "flexible"
)

// getCacheField obtains a cached field by key. If the field
// does not exist in the cache, false will be returned.
// If it is found, the field and true will be returned
// to the caller.
func (s *Service) getCacheField(name, key string) (interface{}, bool) {
	field, err := s.deps.Cache.Get(context.Background(), s.getCacheKey(name, key))
	if err != nil {
		return nil, false
	}
	return field, true
}

// setCacheField sets a cache field by name, key and the value
// of the field.
func (s *Service) setCacheField(val interface{}, name, key string) {
	s.deps.Cache.Set(context.Background(), s.getCacheKey(name, key), val, cache.Options{
		Expiration: 0,
		Tags:       []string{"fields", PostCacheTag(s.postID)},
	})
}

// getCacheKey returns a unique cache key for a singular field.
func (s *Service) getCacheKey(name, typ string) string {
	return fmt.Sprintf("field-%d-%s-%s", s.postID, name, typ)
}

// PostCacheTag is the obtains the unique key that fields are
// attached to. This will be used for setting and clearing
// fields that are attached to a specific post.
func PostCacheTag(postID int) string {
	return fmt.Sprintf("fields-post-%d", postID)
}
