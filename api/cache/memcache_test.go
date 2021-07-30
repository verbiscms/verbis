// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	pkg "github.com/bradfitz/gomemcache/memcache"
	"github.com/verbiscms/verbis/api/environment"
)

func (t *CacheTestSuite) TestMemcache() {
	t.UtilTestProvider_Success(&memcache{
		client: pkg.New(""),
		env:    &environment.Env{MemCachedHosts: "127.0.0.1"},
	}, MemcacheStore)
	t.UtilTestProvider_Error(&memcache{
		client: pkg.New(""),
		env:    &environment.Env{MemCachedHosts: ""},
	})
}
