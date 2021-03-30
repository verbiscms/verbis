// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/services/theme"
)

type Config struct {
	database.Driver
	Theme        *domain.ThemeConfig
	Options      *domain.Options
	Paths        paths.Paths
	Owner        *domain.User
	ThemeService theme.Repository
	Running      bool
}

//type Store struct {
//	Auth       auth.Repository
//	Categories categories.Repository
//	Forms      forms.Repository
//	Media      media.Repository
//	Options    options.Repository
//	Posts      posts.Repository
//	Redirects  redirects.Repository
//	Roles      roles.Repository
//	User       users.Repository
//}
