// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/storage/internal"
)

type provider struct {
	*internal.Config
}

func (p *provider) Set(provider domain.StorageProvider) error {
	panic("implement me")
}

func (p *provider) Name() string {
	return p.Name()
}

func (p *provider) SetProvider(provider domain.StorageProvider) error {
	const op = "Storage.SetProvider"

	p, err := p.GetProvider(provider)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting provider", Operation: op, Err: err}
	}

	err = p.OptionsRepo.Update("storage_provider", provider)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new provider", Operation: op, Err: err}
	}

	p.SetProvider(p, provider)

	return nil
}
