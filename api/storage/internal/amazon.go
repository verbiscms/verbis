// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/s3"
)

type amazon struct{}

const (
	AmazonName = "Amazon S3"
)

var (
	AmazonEnvKeys = []string{
		"STORAGE_AWS_ACCESS_KEY",
		"STORAGE_AWS_SECRET",
	}
)

func (a *amazon) Dial(env *environment.Env) (stow.Location, error) {
	return dialler(s3.Kind, stow.ConfigMap{
		s3.ConfigAccessKeyID: env.AWSAccessKey,
		s3.ConfigSecretKey:   env.AWSSecret,
	})
}

func (a *amazon) Info(env *environment.Env) domain.StorageProviderInfo {
	sp := domain.StorageProviderInfo{
		Name:            AmazonName,
		Order:           3,
		Connected:       false,
		Error:           false,
		EnvironmentSet:  false,
		EnvironmentKeys: AmazonEnvKeys,
	}

	if env.AWSSecret == "" && env.AWSAccessKey == "" {
		sp.Error = ErrMessageConfigNotSet + domain.StorageAWS.TitleCase().String()
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
