// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mocks

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/graymeta/stow"
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
