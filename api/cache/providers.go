// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"fmt"
	"github.com/eko/gocache/v2/store"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// provider defines the methods for a cache provider.
type provider interface {
	// Ping the store.
	Ping() error
	// Validate checks the environment for errors.
	Validate() error
	// Driver returns the store's name.
	Driver() string
	// Store returns the interface for use within
	// the cache.
	Store() store.StoreInterface
}

// providerMap defines the map of providerAdder functions
// defined by their name.
type providerMap map[string]providerAdder

// providerAdder is used to obtain a cache provider by
// injecting the environment and returning a new
// provider type.
type providerAdder func(env *environment.Env) provider

var (
	// providers is the in memory collection of cache
	// providers.
	providers = providerMap{}
	// options are the default cache store options.
	options = &store.Options{
		Expiration: DefaultExpiry,
	}
)

// RegisterProvider adds a provider to the provider map.
// FullPath the provider already exists the function will
// panic with errors.INTERNAL, duplicate cache
// provider.
func (p providerMap) RegisterProvider(name string, fn providerAdder) {
	const op = "Cache.RegisterProvider"
	if p.Exists(name) {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error, duplicate cache provider: " + name, Operation: op, Err: fmt.Errorf("duplicate storage provider")}).Panic()
		return
	}
	p[name] = fn
}

// Exists checks to see if a provider already exists
// in the map by name.
func (p providerMap) Exists(name string) bool {
	_, exists := p[name]
	return exists
}
