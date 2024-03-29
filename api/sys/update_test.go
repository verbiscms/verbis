// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	sm "github.com/hashicorp/go-version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	rocket "github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/verbiscms/verbis/api/version"
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

func (t *SysTestSuite) TestSys_LatestVersion() {
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
		t.Run(name, func() {
			s := Sys{client: &rocket.Updater{Provider: test.input}}
			if test.panics {
				t.Panics(func() {
					s.LatestVersion()
				})
				return
			}
			t.Equal(test.want, s.LatestVersion())
		})
	}
}

func (t *SysTestSuite) TestSys_HasUpdate() {
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
		t.Run(name, func() {
			s := Sys{
				version: test.input,
				client:  &rocket.Updater{Provider: &mockProvider{}},
			}
			t.Equal(test.want, s.HasUpdate())
		})
	}
}

//
//func TestSys_Update(t *testing.T) {
//	logger.SetOutput(ioutil.Discard)
//
//	tt := map[string]struct {
//		mock func(m *database.Driver)
//		want interface{}
//	}{
//		"Up To Date": {
//			nil,
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
//				Driver: d,
//				client: &rocket.Updater{Provider: &mockProvider{}},
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
