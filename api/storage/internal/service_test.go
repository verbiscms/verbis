// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/storage/mocks"
	options "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/assert"
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
		mock     func(m *options.Repository)
		error    string
		provider domain.StorageProvider
		bucket   string
	}{
		"Success": {
			func(m *options.Repository) {
				m.On("Find", "storage_provider").Return("\"aws\"", nil)
				m.On("Find", "storage_bucket").Return("\"bucket\"", nil)
			},
			"",
			domain.StorageAWS,
			"bucket",
		},
		"Empty Provider": {
			func(m *options.Repository) {
				m.On("Find", "storage_provider").Return("", nil)
				m.On("Find", "storage_bucket").Return("\"bucket\"", nil)
			},
			"",
			domain.StorageLocal,
			"bucket",
		},
		"Provider Error": {
			func(m *options.Repository) {
				m.On("Find", "storage_provider").Return(nil, fmt.Errorf("error"))
			},
			"error",
			"",
			"",
		},
		"Bucket Error": {
			func(m *options.Repository) {
				m.On("Find", "storage_provider").Return("amazon", nil)
				m.On("Find", "storage_bucket").Return(nil, fmt.Errorf("error"))
			},
			"error",
			"",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &options.Repository{}
			test.mock(m)
			s := &Service{Env: &environment.Env{}, Options: m}

			provider, bucket, err := s.Config()
			if err != nil {
				assert.Contains(t, errors.Message(err), test.error)
				return
			}

			assert.Equal(t, test.provider, provider)
			assert.Equal(t, test.bucket, bucket)
		})
	}
}
