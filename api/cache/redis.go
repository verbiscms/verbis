// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	pkg "github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
)

// redis defines the data stored for the redis
// client.
type redis struct {
	client *pkg.Client
	env    *environment.Env
}

// init adds the redis store to the the providerMap
// on initialisation of the app.
func init() {
	providers.RegisterProvider(RedisStore, func(env *environment.Env) provider {
		return &redis{pkg.NewClient(&pkg.Options{
			Network:  "",
			Addr:     env.RedisAddress,
			Username: "",
			Password: env.RedisPassword,
			DB:       cast.ToInt(env.RedisDb),
		}), env}
	})
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (r *redis) Validate() error {
	if r.env.RedisAddress == "" {
		return errors.New("no redis address defined in env")
	}
	if r.env.RedisPassword == "" {
		return errors.New("no redis password defined in ev")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory driver name.
func (r *redis) Driver() string {
	return RedisStore
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (r *redis) Store() store.StoreInterface {
	return cache.New(store.NewRedis(r.client, options))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (r *redis) Ping() error {
	return r.client.Ping(context.Background()).Err()
}
