// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	pkg "github.com/go-redis/redis/v8"
	"github.com/verbiscms/verbis/api/environment"
)

func (t *CacheTestSuite) TestRedis() {
	t.UtilTestProviderSuccess(&redis{
		client: pkg.NewClient(&pkg.Options{Addr: "", Password: ""}),
		env:    &environment.Env{RedisAddress: "127.0.0.1", RedisPassword: "password"},
	}, RedisStore)
	t.UtilTestProviderError(&redis{
		client: pkg.NewClient(&pkg.Options{Addr: "", Password: ""}),
		env:    &environment.Env{RedisAddress: "127.0.0.1", RedisPassword: ""},
	})
	t.UtilTestProviderError(&redis{
		client: pkg.NewClient(&pkg.Options{Addr: "", Password: ""}),
		env:    &environment.Env{RedisAddress: "", RedisPassword: "password"},
	})
}
