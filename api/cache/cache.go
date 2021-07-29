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

// Store defines methods for interacting with the
// caching system.
type Store interface {
	// Get retrieves a specific item from the cache by key.
	// Returns errors.NOTFOUND if it could not be found.
	Get(ctx context.Context, key interface{}) (interface{}, error)
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

type Pinger interface {
	Ping() error
}

func Ping(p Pinger) error {
	return p.Ping()
}

type mem struct {
	memcache.Client
}

func (m *mem) Ping() error {
	return m.Ping()
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
)

//type APinger interface {
//	Ping(context.Context) (error)
//}
//
//type BPinger interface {
//	Ping() (error)
//}
//
//switch ping.(type) {
//case APinger:
//...
//case BPinger:
//...
//}

// Load initialises the cache store by the environment.
// It will load a driver into memory ready for setting
// getting and deleting. Drivers supported are Memory
// Redis and MemCached.
// Returns ErrInvalidDriver if the driver passed does not exist.
func Load(env *environment.Env) (*Cache, error) {
	const op = "Cache.Load"

	if env == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error loading cache", Operation: op, Err: fmt.Errorf("nil environment")}
	}

	c := Cache{}

	switch env.CacheDriver {
	case MemoryStore, "":
		client := gocache.New(DefaultExpiry, DefaultCleanup)
		cacheStore := store.NewGoCache(client, nil)
		c = Cache{
			store:  cache.New(cacheStore),
			driver: MemoryStore,
		}
	case RedisStore:
		opts := &redis.Options{
			Addr:     env.RedisAddress,
			Password: env.RedisPassword,

			DB: cast.ToInt(env.RedisDb),
		}
		cacheStore := store.NewRedis(redis.NewClient(opts), nil)

		fmt.Println(cacheStore.Set(context.Background(), "hhh", "", &store.Options{}))

		c = Cache{
			store:  cache.New(cacheStore),
			driver: RedisStore,
		}
	case MemcacheStore:
		memcacheStore := store.NewMemcache(memcache.New(env.MemCachedHosts), &store.Options{
			Expiration: DefaultExpiry,
		})
		c = Cache{
			store:  cache.New(memcacheStore),
			driver: MemcacheStore,
		}
	default:
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error loading cache, invalid Driver: " + env.CacheDriver, Operation: op, Err: ErrInvalidDriver}
	}

	return &c, nil
}

// Get retrieves a specific item from the cache by key.
// Returns errors.NOTFOUND if it could not be found.
func (c *Cache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	const op = "Cache.Get"
	i, err := c.store.Get(ctx, key)
	str := cast.ToString(key)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "Error getting item with key: " + str, Operation: op, Err: err}
	}
	return i, nil
}

// Set set's a singular item in memory by key, value
// and options (tags and expiration time).
// Logs errors.INTERNAL if the item could not be set.
func (c *Cache) Set(ctx context.Context, key interface{}, value interface{}, options Options) {
	const op = "Cache.Set"
	str := cast.ToString(key)
	err := c.store.Set(ctx, key, value, options.toStore())
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error setting cache key: " + str, Operation: op, Err: err}).Error()
		return
	}
	logger.Debug("Successfully set cache item with key: " + str)
}

// Delete removes a singular item from the cache by
// a specific key.
// Logs errors.INTERNAL if the item could not be deleted.
func (c *Cache) Delete(ctx context.Context, key interface{}) {
	const op = "Cache.Delete"
	str := cast.ToString(key)
	err := c.store.Delete(ctx, key)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting cache key: " + str, Operation: op, Err: err}).Error()
		return
	}
	logger.Debug("Successfully deleted cache item with key: " + str)
}

// Invalidate removes items from the cache via the
// InvalidateOptions passed.
// Returns errors.INVALID if the cache could not be invalidated.
func (c *Cache) Invalidate(ctx context.Context, options InvalidateOptions) error {
	const op = "Cache.Invalidate"
	err := c.store.Invalidate(ctx, options.toStore())
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error invalidating cache", Operation: op, Err: err}
	}
	return nil
}

// Clear removes all items from the cache.
// Returns errors.INTERNAL if the cache could not be cleared.
func (c *Cache) Clear(ctx context.Context) error {
	const op = "Cache.Clear"
	err := c.store.Clear(ctx)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error clearing cache", Operation: op, Err: err}
	}
	return nil
}

// Driver returns the current store being used, it can be
// MemoryStore, RedisStore or MemcachedStore.
func (c *Cache) Driver() string {
	return c.driver
}
