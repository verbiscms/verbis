// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eko/gocache/v2/store"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"sync"
	"time"
)

// Store defines methods for interacting with the
// caching system.
type Store interface {
	// Get retrieves a specific item from the cache by key.
	// Returns errors.NOTFOUND if it could not be found.
	Get(ctx context.Context, key interface{}, returnObj interface{}) (interface{}, error)
	// Set set's a singular item in memory by key, value
	// and options (tags and expiration time).
	// Logs errors.INTERNAL if the item could not be set.
	Set(ctx context.Context, key interface{}, value interface{}, options Options)
	// Delete removes a singular item from the cache by
	// a specific key.
	// Logs errors.INTERNAL if the item could not be deleted.
	Delete(ctx context.Context, key interface{})
	// Invalidate removes items from the cache via the
	// InvalidateOptions passed.
	// Returns errors.INVALID if the cache could not be invalidated.
	Invalidate(ctx context.Context, options InvalidateOptions) error
	// Clear removes all items from the cache.
	// Returns errors.INTERNAL if the cache could not be cleared.
	Clear(ctx context.Context) error
	// Driver returns the current store being used, it can be
	// MemoryStore, RedisStore or MemcachedStore.
	Driver() string
}

// Cache defines the methods for interacting with the
// cache layer.
type Cache struct {
	// store is the package store interface used for interacting
	// with the cache store.
	store store.StoreInterface
	// driver is the current store being used, it can be
	// MemoryStore, RedisStore or MemcachedStore.
	driver string
}

const (
	// MemoryStore is the Redis Driver, depicted
	// in the environment.
	MemoryStore = "memory"
	// RedisStore is the Redis Driver, depicted
	// in the environment.
	RedisStore = "redis"
	// MemcacheStore is the Memcached Driver, depicted
	// in the environment.
	MemcacheStore = "memcache"
	// DefaultExpiry defines how many minutes the item
	// lasts for in the cache by default.
	DefaultExpiry = -1
	// DefaultCleanup defines the clean up interval of
	// the cache.
	DefaultCleanup = 5 * time.Minute
	// RememberForever is an alias for setting the
	// cache item to never be removed.
	RememberForever = -1
)

var (
	// ErrInvalidDriver is returned by Load when an
	// invalid driver was passed via the env.
	ErrInvalidDriver = errors.New("invalid cache Driver")
	mtx              = sync.Mutex{}
)

// Load initialises the cache store by the environment.
// It will load a driver into memory ready for setting
// getting setting and deleting. Drivers supported are Memory
// Redis and MemCached.
// Returns ErrInvalidDriver if the driver passed does not exist.
func Load(env *environment.Env) (*Cache, error) {
	const op = "Cache.Load"

	if env == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error loading cache", Operation: op, Err: fmt.Errorf("nil environment")}
	}

	driver := env.CacheDriver
	if env.CacheDriver == "" {
		driver = MemoryStore
	}

	if !providers.Exists(driver) {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error loading cache, invalid driver: " + env.CacheDriver, Operation: op, Err: ErrInvalidDriver}
	}

	prov := providers[driver](env)

	err := prov.Validate()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error loading cache, validation failed", Operation: op, Err: ErrInvalidDriver}
	}

	err = prov.Ping()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error error pinging cache store: " + prov.Driver(), Operation: op, Err: err}
	}

	return &Cache{
		store:  prov.Store(),
		driver: prov.Driver(),
	}, nil
}

// Get retrieves a specific item from the cache by key.
// Returns errors.NOTFOUND if it could not be found.
func (c *Cache) Get(ctx context.Context, key interface{}, returnObj interface{}) (interface{}, error) {
	const op = "Cache.Get"
	mtx.Lock()
	defer mtx.Unlock()
	result, err := c.store.Get(ctx, key)

	switch v := result.(type) {
	case []byte:
		err = json.Unmarshal(v, returnObj)
	case string:
		err = json.Unmarshal([]byte(v), returnObj)
	}

	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "Error getting item with key: " + cast.ToString(key), Operation: op, Err: err}
	}
	return result, nil
}

// Set set's a singular item in memory by key, value
// and options (tags and expiration time).
// Logs errors.INTERNAL if the item could not be set.
func (c *Cache) Set(ctx context.Context, key interface{}, value interface{}, options Options) {
	const op = "Cache.Set"
	mtx.Lock()
	defer mtx.Unlock()

	marshal, err := json.Marshal(value)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error marshalling cache item", Operation: op, Err: err}).Error()
	}
	value = marshal

	str := cast.ToString(key)
	err = c.store.Set(ctx, key, value, options.toStore())
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error setting cache key: " + str, Operation: op, Err: err}).Error()
		return
	}

	logger.Trace("Successfully set cache item with key: " + str)
}

// Delete removes a singular item from the cache by
// a specific key.
// Logs errors.INTERNAL if the item could not be deleted.
func (c *Cache) Delete(ctx context.Context, key interface{}) {
	const op = "Cache.Delete"
	mtx.Lock()
	defer mtx.Unlock()
	str := cast.ToString(key)
	err := c.store.Delete(ctx, key)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting cache key: " + str, Operation: op, Err: err}).Error()
		return
	}
	logger.Trace("Successfully deleted cache item with key: " + str)
}

// Invalidate removes items from the cache via the
// InvalidateOptions passed.
// Returns errors.INVALID if the cache could not be invalidated.
func (c *Cache) Invalidate(ctx context.Context, options InvalidateOptions) error {
	const op = "Cache.Invalidate"
	mtx.Lock()
	defer mtx.Unlock()
	err := c.store.Invalidate(ctx, options.toStore())
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error invalidating cache", Operation: op, Err: err}
	}
	logger.Trace(fmt.Sprintf("Successfully invalidated cache with tags: %v", options.Tags))
	return nil
}

// Clear removes all items from the cache.
// Returns errors.INTERNAL if the cache could not be cleared.
func (c *Cache) Clear(ctx context.Context) error {
	const op = "Cache.Clear"
	mtx.Lock()
	defer mtx.Unlock()
	err := c.store.Clear(ctx)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error clearing cache", Operation: op, Err: err}
	}
	logger.Trace("Successfully cleared cache")
	return nil
}

// Driver returns the current store being used, it can be
// MemoryStore, RedisStore or MemcachedStore.
func (c *Cache) Driver() string {
	return c.driver
}
