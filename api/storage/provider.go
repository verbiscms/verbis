// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

type provider struct {
	config
}

func (p *provider) Name() string {
	return p.Name()
}

func (p *provider) Set(provider domain.StorageProvider) error {
	const op = "Storage.SetProvider"

	prov, err := p.GetProvider(provider)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting provider", Operation: op, Err: err}
	}

	err = p.OptionsRepo.Update("storage_provider", provider)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new provider", Operation: op, Err: err}
	}

	p.SetLocation(prov)

	return nil
}
