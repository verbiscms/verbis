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
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
	"testing"
)

func TestService_Provider(t *testing.T) {
	m := ProviderMap{domain.StorageAWS: &amazon{}}

	tt := map[string]struct {
		input    ProviderMap
		provider domain.StorageProvider
		dial     func(kind string, config stow.Config) (stow.Location, error)
		want     interface{}
	}{
		"Success": {
			m,
			domain.StorageAWS,
			dialSuccess,
			stowLocation,
		},
		"Exists Error": {
			ProviderMap{},
			domain.StorageAWS,
			dialSuccess,
			"Error connecting to storage provider",
		},
		"Dial Error": {
			m,
			domain.StorageAWS,
			dialErr,
			"Error connecting to storage provider",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			origDialler := dialler
			origProviders := Providers
			defer func() {
				dialler = origDialler
				Providers = origProviders
			}()
			dialler = test.dial
			Providers = test.input

			s := &Service{Env: &environment.Env{}}

			got, err := s.Provider(test.provider)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}

func TestService_BucketByFile(t *testing.T) {
	c := &mocks.StowContainer{}

	tt := map[string]struct {
		input domain.File
		dial  func(kind string, config stow.Config) (stow.Location, error)
		want  interface{}
	}{
		"Success": {
			domain.File{Provider: domain.StorageAWS, Bucket: "bucket"},
			func(kind string, config stow.Config) (stow.Location, error) {
				m := &mocks.StowLocation{}
				m.On("Container", "bucket").Return(c, nil)
				return m, nil
			},
			c,
		},
		"Provider Error": {
			domain.File{Provider: domain.StorageAWS, Bucket: "bucket"},
			dialErr,
			"Error connecting to storage provider",
		},
		"Container Error": {
			domain.File{Provider: domain.StorageAWS, Bucket: "bucket"},
			func(kind string, config stow.Config) (stow.Location, error) {
				m := &mocks.StowLocation{}
				m.On("Container", "bucket").Return(nil, fmt.Errorf("error"))
				return m, nil
			},
			ErrMessageInvalidBucket,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			origDialler := dialler
			defer func() {
				dialler = origDialler
			}()
			dialler = test.dial

			s := &Service{Env: &environment.Env{}}

			got, err := s.BucketByFile(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}

func TestService_Bucket(t *testing.T) {
	c := &mocks.StowContainer{}

	tt := map[string]struct {
		input domain.StorageProvider
		dial  func(kind string, config stow.Config) (stow.Location, error)
		want  interface{}
	}{
		"Success": {
			domain.StorageAWS,
			func(kind string, config stow.Config) (stow.Location, error) {
				m := &mocks.StowLocation{}
				m.On("Container", "bucket").Return(c, nil)
				return m, nil
			},
			c,
		},
		"Provider Error": {
			domain.StorageAWS,
			dialErr,
			"Error connecting to storage provider",
		},
		"Container Error": {
			domain.StorageAWS,
			func(kind string, config stow.Config) (stow.Location, error) {
				m := &mocks.StowLocation{}
				m.On("Container", "bucket").Return(nil, fmt.Errorf("error"))
				return m, nil
			},
			ErrMessageInvalidBucket,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			origDialler := dialler
			defer func() {
				dialler = origDialler
			}()
			dialler = test.dial

			s := &Service{Env: &environment.Env{}}

			got, err := s.Bucket(test.input, "bucket")
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}

func TestService_Config(t *testing.T) {
	tt := map[string]struct {
		mock func(m *options.Repository)
		want domain.StorageConfig
	}{
		"Success": {
			func(m *options.Repository) {
				m.On("Struct").Return(domain.Options{
					StorageProvider:     domain.StorageAWS,
					StorageBucket:       "bucket",
					StorageUploadRemote: true,
					StorageLocalBackup:  true,
					StorageRemoteBackup: false,
				})
			},
			domain.StorageConfig{
				Provider:     domain.StorageAWS,
				Bucket:       "bucket",
				UploadRemote: true,
				LocalBackup:  true,
				RemoteBackup: false,
			},
		},
		"Empty Provider": {
			func(m *options.Repository) {
				m.On("Struct").Return(domain.Options{
					StorageBucket:      "bucket",
					StorageLocalBackup: true,
				})
			},
			domain.StorageConfig{
				Provider:    domain.StorageLocal,
				Bucket:      "bucket",
				LocalBackup: true,
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &options.Repository{}
			test.mock(m)
			s := &Service{Env: &environment.Env{}, Options: m}
			got := s.Config()
			assert.Equal(t, test.want, got)
		})
	}
}
