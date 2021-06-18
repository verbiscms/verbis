// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type mockPatcherSuccess struct{}

func (m *mockPatcherSuccess) HasUpdate() (bool, error) {
	return true, nil
}

func (m *mockPatcherSuccess) LatestVersion() (string, error) {
	return "0.0.1", nil
}

func (m *mockPatcherSuccess) Update(archive string) (updater.Status, error) {
	return 1, nil
}

type mockPatcherError struct{}

func (m *mockPatcherError) HasUpdate() (bool, error) {
	return false, fmt.Errorf("error")
}

func (m *mockPatcherError) LatestVersion() (string, error) {
	return "", fmt.Errorf("error")
}

func (m *mockPatcherError) Update(archive string) (updater.Status, error) {
	return 1, fmt.Errorf("error")
}

func TestSys_LatestVersion(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		patcher updater.Patcher
		panics  bool
		want    interface{}
	}{
		"Success": {
			&mockPatcherSuccess{},
			false,
			"0.0.1",
		},
		"Error": {
			&mockPatcherError{},
			true,
			"Error obtaining remote version",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := Sys{updater: test.patcher}
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
		patcher updater.Patcher
		panics  bool
		want    interface{}
	}{
		"Success": {
			&mockPatcherSuccess{},
			false,
			true,
		},
		"Error": {
			&mockPatcherError{},
			true,
			"Error obtaining remote version",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := Sys{updater: test.patcher}
			if test.panics {
				assert.Panics(t, func() {
					s.HasUpdate()
				})
				return
			}
			assert.Equal(t, test.want, s.HasUpdate())
		})
	}
}

func TestSys_Update(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		patcher func() updater.Patcher
		want    interface{}
	}{
		"Success": {
			func() updater.Patcher { return &mockPatcherSuccess{} },
			"0.0.1",
		},
		//"Error": {
		//	func() updater.Patcher {},
		//	"0.0.1",
		//},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := Sys{updater: test.patcher()}
			got, err := s.Update()
			if err != nil {
				assert.Contains(t, errors.Message(err), err)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
