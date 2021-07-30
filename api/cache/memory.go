// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	gocache "github.com/patrickmn/go-cache"
	"github.com/verbiscms/verbis/api/environment"
)

// memory defines the data stored for the gocache
// client.
type memory struct {
	client *gocache.Cache
	env    *environment.Env
}

// init adds the memory store to the the providerMap
// on initialisation of the app.
func init() {
	providers.RegisterProvider(MemoryStore, func(env *environment.Env) provider {
		return &memory{gocache.New(DefaultExpiry, DefaultCleanup), env}
	})
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (m *memory) Validate() error {
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory driver name.
func (m *memory) Driver() string {
	return MemoryStore
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (m *memory) Store() store.StoreInterface {
	return cache.New(store.NewGoCache(m.client, options))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (m *memory) Ping() error {
	return nil
}
