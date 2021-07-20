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

// azure satisfies the provider interface by implementing
// dial and info.
type azure struct{}

// AzureName is the friendly name for the provider
// passed back from info()
const AzureName = "Microsoft Azure"

var (
	// AzureEnvKeys defines the environment keys needed in
	// order to dial the azure provider.
	AzureEnvKeys = []string{
		"STORAGE_AZURE_ACCOUNT",
		"STORAGE_AZURE_KEY",
	}
)

// Dial returns a new stow.Location by calling the
// dialler.
func (a *azure) Dial(env *environment.Env) (stow.Location, error) {
	return dialler(az.Kind, stow.ConfigMap{
		az.ConfigAccount: env.AzureAccount,
		az.ConfigKey:     env.AzureKey,
	})
}

// Info returns information about the azure storage
// provider.
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
