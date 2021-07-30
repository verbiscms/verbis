// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mocks

import "github.com/eko/gocache/v2/store"

type storeInterface interface {
	store.StoreInterface
}

type provider interface {
	Ping() error
	Validate() error
	Driver() string
	Store() store.StoreInterface
}
