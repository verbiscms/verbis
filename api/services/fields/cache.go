// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
)

const (
	// FieldCacheTag is the global tag given for caching
	// fields.
	FieldCacheTag = "fields"
	// standardCacheKey is the key stored in the cache
	// for standard fields.
	standardCacheKey = "standard"
	// repeaterCacheKey is the key stored in the cache
	// for repeater fields.
	repeaterCacheKey = "repeater" //nolint
	// flexibleCacheKey is the key stored in the cache
	// for flexible fields.
	flexibleCacheKey = "flexible" //nolint
)

// getCacheField obtains a cached field by key. If the field
// does not exist in the cache, false will be returned.
// If it is found, the field and true will be returned
// to the caller.
func (s *Service) getCacheField(name, key string, id int) (interface{}, bool) {
	return nil, false
	//field, err := s.deps.Cache.Get(context.Background(), s.getCacheKey(name, key, id))
	//if err != nil {
	//	return nil, false
	//}
	//return field, true
}

// setCacheField sets a cache field by name, key and the value
// of the field.
func (s *Service) setCacheField(val interface{}, name, key string, id int) {
	//s.deps.Cache.Set(context.Background(), s.getCacheKey(name, key, id), val, cache.Options{
	//	Expiration: cache.RememberForever,
	//	Tags:       []string{FieldCacheTag},
	//})
}

// getCacheKey returns a unique cache key for a singular field.
func (s *Service) getCacheKey(name, typ string, id int) string {
	return fmt.Sprintf("field-%d-%d-%s-%s", s.postID, id, name, typ)
}
