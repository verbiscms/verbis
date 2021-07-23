// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	"io/ioutil"
	"testing"
)

func TestProviderMap_RegisterProvider(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input  ProviderMap
		name   domain.StorageProvider
		panics bool
	}{
		"Added": {
			ProviderMap{},
			domain.StorageAzure,
			false,
		},
		"Panics": {
			ProviderMap{domain.StorageAWS: &amazon{}},
			domain.StorageAWS,
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.panics {
				assert.Panics(t, func() {
					test.input.RegisterProvider(test.name, &amazon{})
				})
				return
			}
			test.input.RegisterProvider(test.name, &amazon{})
			assert.True(t, test.input.Exists(test.name))
		})
	}
}

func TestProviderMap_Exists(t *testing.T) {
	tt := map[string]struct {
		input    ProviderMap
		provider domain.StorageProvider
		want     bool
	}{
		"True": {
			ProviderMap{domain.StorageAWS: &amazon{}},
			domain.StorageAWS,
			true,
		},
		"False": {
			ProviderMap{},
			domain.StorageAWS,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Exists(test.provider)
			assert.Equal(t, test.want, got)
		})
	}
}

var (
	stowLocation = &mocks.StowLocation{}
	dialSuccess  = func(kind string, config stow.Config) (stow.Location, error) {
		return stowLocation, nil
	}
	dialErr = func(kind string, config stow.Config) (stow.Location, error) {
		return nil, fmt.Errorf("error")
	}
)

func UtilTestProviderDial(env *environment.Env, p Provider, t *testing.T) {
	tt := map[string]struct {
		dial func(kind string, config stow.Config) (stow.Location, error)
		want interface{}
	}{
		"Success": {
			dialSuccess,
			nil,
		},
		"Error": {
			dialErr,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			orig := dialler
			defer func() { dialler = orig }()
			dialler = test.dial

			dial, err := p.Dial(env)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.NotNil(t, dial)
		})
	}
}

func UtilTestProviderInfo(env *environment.Env, p Provider, t *testing.T) {
	tt := map[string]struct {
		env  *environment.Env
		dial func(kind string, config stow.Config) (stow.Location, error)
		want interface{}
	}{
		"Success": {
			env,
			dialSuccess,
			nil,
		},
		"Empty Env": {
			&environment.Env{},
			dialSuccess,
			"error",
		},
		"Dial Error": {
			env,
			dialErr,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			orig := dialler
			defer func() { dialler = orig }()
			dialler = test.dial

			got := p.Info(test.env)

			if name == "Empty Env" {
				assert.False(t, got.EnvironmentSet)
				assert.False(t, got.Connected)
				assert.Contains(t, got.Error, ErrMessageConfigNotSet)
				return
			}

			if name == "Dial Error" {
				assert.True(t, got.EnvironmentSet)
				assert.False(t, got.Connected)
				assert.Contains(t, got.Error, ErrMessageDial)
				return
			}

			assert.True(t, got.Connected)
			assert.True(t, got.EnvironmentSet)
			assert.NotEmpty(t, got.Order)
			assert.NotEmpty(t, got.Name)
			assert.NotEmpty(t, got.EnvironmentKeys)
		})
	}
}

func UtilTestProvider(env *environment.Env, p Provider, t *testing.T) {
	UtilTestProviderDial(env, p, t)
	UtilTestProviderInfo(env, p, t)
}
