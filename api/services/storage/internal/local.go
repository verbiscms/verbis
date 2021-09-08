// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/graymeta/stow"
	stowLocal "github.com/graymeta/stow/local"
	"github.com/verbiscms/verbis/api/common/files"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"os"
)

// local satisfies the provider interface by implementing
// dial and info.
type local struct {
	path string
}

// LocalName is the friendly name for the provider
// passed back from info()
const LocalName = "Local Storage"

// Dial returns a new stow.Location by calling the
// dialler.
func (l *local) Dial(env *environment.Env) (stow.Location, error) {
	if l.path == "" {
		l.path = paths.Get().Storage
	}
	if !files.DirectoryExists(l.path) {
		err := os.MkdirAll(l.path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return dialler(stowLocal.Kind, stow.ConfigMap{
		stowLocal.ConfigKeyPath: l.path,
	})
}

// Info returns information about the local storage
// provider.
func (l *local) Info(env *environment.Env) domain.StorageProviderInfo {
	return domain.StorageProviderInfo{
		Order:          1,
		Name:           LocalName,
		Connected:      true,
		Error:          false,
		EnvironmentSet: true,
	}
}
