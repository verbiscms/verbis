// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
)

// Config represents the configuration parsed to the
// store.
type Config struct {
	database.Driver
	Paths   paths.Paths
	Owner   *domain.User
	Theme   config.Provider
	Running bool
}
