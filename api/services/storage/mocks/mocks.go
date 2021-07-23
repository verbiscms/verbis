// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mocks

import (
	"github.com/graymeta/stow"
	"github.com/verbiscms/verbis/api/domain"
)

type StowLocation interface {
	stow.Location
}

type StowContainer interface {
	stow.Container
}

type StowItem interface {
	stow.Item
}

type Service interface {
	Provider(provider domain.StorageProvider) (stow.Location, error)
	Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error)
	BucketByFile(file domain.File) (stow.Container, error)
	Config() (domain.StorageProvider, string, error)
}
