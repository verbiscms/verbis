// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import "github.com/eko/gocache/v2/cache"

func (t *CacheTestSuite) TestMemory() {
	m := memory{}
	t.Nil(m.Validate())
	t.Equal(MemoryStore, m.Driver())
	store := m.Store()
	t.IsType(&cache.Cache{}, store)
	t.Nil(m.Ping())
}
