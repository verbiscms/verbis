// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/version"
	sm "github.com/hashicorp/go-version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type mockErrProvider struct{}

func (m *mockErrProvider) GetLatestVersion() (string, error) {
	return "", fmt.Errorf("error")
}

func (m *mockErrProvider) Walk(walkFn provider.WalkFunc) error {
	return nil
}

func (m *mockErrProvider) Retrieve(srcPath, destPath string) error {
	return nil
}

func (m *mockErrProvider) Open() error {
	return nil
}

func (m *mockErrProvider) Close() error {
	return nil
}

type mockProvider struct{}

func (m *mockProvider) GetLatestVersion() (string, error) {
	return "v0.0.10", nil
}

func (m *mockProvider) Walk(walkFn provider.WalkFunc) error {
	return nil
}

func (m *mockProvider) Retrieve(srcPath, destPath string) error {
	return nil
}

func (m *mockProvider) Open() error {
	return nil
}

func (m *mockProvider) Close() error {
	return nil
}

func TestSys_LatestVersion(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input  provider.Provider
		panics bool
		want   interface{}
	}{
		"Success": {
			&mockProvider{},
			false,
			"v0.0.10",
		},
		"Error": {
			&mockErrProvider{},
			true,
			"Error obtaining remote version",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := Sys{updater: &updater.Updater{Provider: test.input}}
			if test.panics {
				assert.Panics(t, func() {
					s.LatestVersion()
				})
				return
			}
			assert.Equal(t, test.want, s.LatestVersion())
		})
	}
}

func TestSys_HasUpdate(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input *sm.Version
		want  interface{}
	}{
		"True": {
			version.Must("v0.0.9"),
			true,
		},
		"False": {
			version.Must("v0.0.10"),
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := Sys{
				version: test.input,
				updater: &updater.Updater{Provider: &mockProvider{}},
			}
			assert.Equal(t, test.want, s.HasUpdate())
		})
	}
}

//func TestSys_Update(t *testing.T) {
//	logger.SetOutput(ioutil.Discard)
//
//	tt := map[string]struct {
//		mock func(m *database.Driver)
//		want interface{}
//	}{
//		"Success": {
//			func(m *database.Driver) {
//
//			},
//			"Error updating Verbis with status code",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			d := &database.Driver{}
//			if test.mock != nil {
//				test.mock(d)
//			}
//			s := Sys{
//				Driver:  d,
//				updater: &updater.Updater{Provider: &mockProvider{}},
//			}
//			got, err := s.Update(false)
//			if err != nil {
//				assert.Contains(t, errors.Message(err), test.want)
//				return
//			}
//			assert.Equal(t, test.want, got)
//		})
//	}
//}
