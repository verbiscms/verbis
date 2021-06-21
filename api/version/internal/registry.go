// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/logger"
	sm "github.com/hashicorp/go-version"
	"os"
	"sort"
)

type CallBackFn func() error

type UpdateRegistry []*Update

var Updates = make(UpdateRegistry, 0)

type Update struct {
	Version       string
	MajorVersion  int
	MigrationPath string
	CallBackUp    CallBackFn
	CallBackDown  CallBackFn
	Stage         Stage
}

type Stage string

const (
	Major = "major"
	Minor = "minor"
	Patch = "patch"
)

func (u Update) ToSemVer() *sm.Version {
	semver, err := sm.NewVersion(u.Version)
	if err != nil {
		logger.Panic(err.Error())
		return nil
	}
	return semver
}

// Sort UpdateRegistry is a type that implements the sort.Interface
// interface so that versions can be sorted.
func (r UpdateRegistry) Sort() {
	sort.Sort(r)
}

func (r UpdateRegistry) Len() int {
	return len(r)
}

func (r UpdateRegistry) Less(i, j int) bool {
	return r[i].ToSemVer().LessThan(r[j].ToSemVer())
}

func (r UpdateRegistry) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (u *Update) HasCallBack() bool {
	return u.CallBackUp != nil && u.CallBackDown != nil
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
// AddUpdate add's an update to the registry. the
func (r *UpdateRegistry) AddUpdate(u *Update) {
	if u.Version == "" {
		logger.Panic("No version provided for update")
	}

	if u.Stage == "" {
		logger.Panic("No stage set")
	}

	if u.MigrationPath == "" {
		logger.Panic("No migration path set")
	}

	if u.CallBackUp != nil && u.CallBackDown == nil {
		logger.Panic("CallbackDown function must be declared if CallBackUp is set")
	}

	if u.CallBackUp == nil && u.CallBackDown != nil {
		logger.Panic("CallbackUp function must be declared if CallBackDown is set")
	}

	semVer := u.ToSemVer()
	seg := semVer.Segments()

	if len(seg) != 3 {
		logger.Panic("Invalid version: " + semVer.Original())
	}

	u.MajorVersion = seg[0]
	u.MigrationPath = fmt.Sprintf("v%d%v%v", u.MajorVersion, string(os.PathSeparator), u.MigrationPath)

	*r = append(*r, u)
}
