// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/graymeta/stow"
	stowLocal "github.com/graymeta/stow/local"
)

type local struct{}

const LocalName = "Local Storage"

func (l *local) Dial(env *environment.Env) (stow.Location, error) {
	return dialler(stowLocal.Kind, stow.ConfigMap{
		stowLocal.ConfigKeyPath: paths.Get().Storage,
	})
}

func (l *local) Info(env *environment.Env) domain.StorageProviderInfo {
	return domain.StorageProviderInfo{
		Order:          1,
		Name:           LocalName,
		Connected:      true,
		Error:          false,
		EnvironmentSet: true,
	}
}
