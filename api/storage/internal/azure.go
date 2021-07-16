// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/graymeta/stow"
	az "github.com/graymeta/stow/azure"
)

type azure struct{}

const AzureName = "Microsoft Azure"

var (
	AzureEnvKeys = []string{
		"STORAGE_AZURE_ACCOUNT",
		"STORAGE_AZURE_KEY",
	}
)

func (a *azure) Dial(env *environment.Env) (stow.Location, error) {
	return dialler(az.Kind, stow.ConfigMap{
		az.ConfigAccount: env.AzureAccount,
		az.ConfigKey:     env.AzureKey,
	})
}

func (a *azure) Info(env *environment.Env) domain.StorageProviderInfo {
	sp := domain.StorageProviderInfo{
		Name:            AzureName,
		Order:           4,
		Connected:       false,
		Error:           false,
		EnvironmentSet:  false,
		EnvironmentKeys: AzureEnvKeys,
	}

	if env.AzureAccount == "" && env.AzureKey == "" {
		sp.Error = ErrMessageConfigNotSet + domain.StorageAzure.TitleCase().String()
		return sp
	}

	sp.EnvironmentSet = true

	_, err := a.Dial(env)
	if err != nil {
		sp.Error = ErrMessageDial + err.Error()
		return sp
	}

	sp.Connected = true
	sp.Error = false

	return sp
}
