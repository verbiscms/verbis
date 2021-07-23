// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"time"
)

// Cacher defines methods for interacting with the
// caching system.
type Cacher interface {
	store.StoreInterface
}

var (
	// c is an alias for a Cacher or store.StoreInterface
	// used by go-cache.
	c Cacher
	// Driver is the current driver being used.
	Driver string
)

const (
	// MemoryStore is the Redis Driver, depicted
	// in the environment.
	MemoryStore = "memory"
	// RedisStore is the Redis Driver, depicted
	// in the environment.
	RedisStore = "redis"
	// MemcachedStore is the Memcached Driver, depicted
	// in the environment.
	MemcachedStore = "memcached"
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
)

// Load initialises the cache store by the environment.
// It will load a driver into memory ready for setting
// getting and deleting. Drivers supported are Memory
// Redis and MemCached.
// Returns ErrInvalidDriver if the driver passed does not exist.
func Load(env *environment.Env) error {
	const op = "Cache.Load"

	if env == nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error loading cache", Operation: op, Err: fmt.Errorf("nil environment")}
	}

	switch env.CacheDriver {
	case MemoryStore:
		client := gocache.New(DefaultExpiry, DefaultCleanup)
		cacheStore := store.NewGoCache(client, nil)
		c = cache.New(cacheStore)
		Driver = MemoryStore
	case RedisStore:
		opts := &redis.Options{
			Addr:     env.RedisAddress,
			Password: env.RedisPassword,
			DB:       cast.ToInt(env.RedisDb),
		}
		cacheStore := store.NewRedis(redis.NewClient(opts), nil)
		c = cache.New(cacheStore)
		Driver = RedisStore
	case MemcachedStore:
		memcacheStore := store.NewMemcache(memcache.New(env.MemCachedHosts), &store.Options{
			Expiration: DefaultExpiry,
		})
		c = cache.New(memcacheStore)
		Driver = MemcachedStore
	default:
		return &errors.Error{Code: errors.INVALID, Message: "Error loading cache, invalid Driver: " + env.CacheDriver, Operation: op, Err: ErrInvalidDriver}
	}

	return nil
}

// Get retrieves a specific item from the cache by key.
// Returns errors.NOTFOUND if it could not be found.
func Get(ctx context.Context, key interface{}) (interface{}, error) {
	const op = "Cache.Get"
	i, err := c.Get(ctx, key)
	str := cast.ToString(key)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "Error getting item with key: " + str, Operation: op, Err: err}
	}
	return i, nil
}

// Set set's a singular item in memory by key, value
// and options (tags and expiration time).
// Logs errors.INTERNAL if the item could not be set.
func Set(ctx context.Context, key interface{}, value interface{}, options Options) {
	const op = "Cache.Set"
	str := cast.ToString(key)
	err := c.Set(ctx, key, value, options.toStore())
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error setting cache key: " + str, Operation: op, Err: err}).Error()
	}
	logger.Trace("Successfully set cache item with key: " + str)
}

// Delete removes a singular item from the cache by
// a specific key.
// Logs errors.INTERNAL if the item could not be deleted.
func Delete(ctx context.Context, key interface{}) {
	const op = "Cache.Delete"
	str := cast.ToString(key)
	err := c.Delete(ctx, key)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting cache key: " + str, Operation: op, Err: err}).Error()
	}
	logger.Trace("Successfully deleted cache item with key: " + str)
}

// Invalidate removes items from the cache via the
// InvalidateOptions passed.
// Returns errors.INVALID if the cache could not be invalidated.
func Invalidate(ctx context.Context, options InvalidateOptions) error {
	const op = "Cache.Invalidate"
	err := c.Invalidate(ctx, options.toStore())
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error invalidating cache", Operation: op, Err: err}
	}
	return nil
}

// Clear removes all items from the cache.
// Returns errors.INTERNAL if the cache could not be cleared.
func Clear(ctx context.Context) error {
	const op = "Cache.Clear"
	err := c.Clear(ctx)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error clearing cache", Operation: op, Err: err}
	}
	return nil
}

// SetDriver updates the cache to a new driver. If
// the cacher passed to the function is nil, it
// will panic.
func SetDriver(cacher Cacher) {
	if cacher == nil {
		logger.Panic("Nil cacher")
		return
	}
	c = cacher
}
