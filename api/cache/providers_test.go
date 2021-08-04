// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/verbiscms/verbis/api/environment"
)

var adder = func(env *environment.Env) provider { return nil }

func (t *CacheTestSuite) TestProviderMap_RegisterProvider() {
	tt := map[string]struct {
		input  providerMap
		name   string
		panics bool
	}{
		"Added": {
			map[string]providerAdder{},
			MemoryStore,
			false,
		},
		"Panics": {
			map[string]providerAdder{MemoryStore: adder},
			MemoryStore,
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			if test.panics {
				t.Panics(func() {
					test.input.RegisterProvider(test.name, adder)
				})
				return
			}
			test.input.RegisterProvider(test.name, adder)
			t.NotNil(test.input[test.name])
		})
	}
}

func (t *CacheTestSuite) TestProviderMap_Exists() {
	tt := map[string]struct {
		input    string
		provider providerMap
		want     bool
	}{
		"True": {
			MemoryStore,
			map[string]providerAdder{MemoryStore: adder},
			true,
		},
		"False": {
			RedisStore,
			map[string]providerAdder{MemoryStore: adder},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := test.provider.Exists(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *CacheTestSuite) TestInit() {
	got := providers
	t.True(got.Exists(MemoryStore))
	t.True(got.Exists(RedisStore))
	t.True(got.Exists(MemcacheStore))
	t.NotNil(got[MemoryStore](&environment.Env{}))
	t.NotNil(got[RedisStore](&environment.Env{}))
	t.NotNil(got[MemcacheStore](&environment.Env{}))
}

func (t *CacheTestSuite) UtilTestProviderSuccess(p provider, name string) {
	// Test Driver() method
	t.Equal(name, p.Driver())

	// Test Validate() method
	err := p.Validate()
	t.NoErrorf(err, "expecting provider to pass validation")

	// Test Store() method
	store := p.Store()
	t.NotNil(store)
	t.IsType(&cache.Cache{}, store)
}

func (t *CacheTestSuite) UtilTestProviderError(p provider) {
	// Test Validate() method
	err := p.Validate()
	t.Errorf(err, "expecting provider to fail validation")

	// Test Ping() method
	pingErr := p.Ping()
	t.Errorf(pingErr, "expecting provider have ping error")
}
