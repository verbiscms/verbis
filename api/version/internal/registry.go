// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/logger"
	sm "github.com/hashicorp/go-version"
	"sort"
)

type CallBackFn func() error

type registry []*Update

var Registry = make(registry, 0)

type Update struct {
	Version       string
	MigrationPath string
	Callback      CallBackFn
	Stage         Stage
}

type Stage string

const (
	Major = "major"
	Minor = "minor"
	Patch = "patch"
)

func (u Update) ToSemVer() *sm.Version {
	smver, err := sm.NewVersion(u.Version)
	if err != nil {
		logger.Panic(err.Error())
	}
	return smver
}

func (r registry) Sort() {
	sort.Sort(r)
}

// registry is a type that implements the sort.Interface interface
// so that versions can be sorted.
func (r registry) Len() int {
	return len(r)
}

func (r registry) Less(i, j int) bool {
	return r[i].ToSemVer().LessThan(r[j].ToSemVer())
}

func (r registry) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *registry) AddUpdate(u *Update) {
	if u.Version == "" {
		logger.Panic("No version provided for update")
	}

	// Check for panics
	_ = u.ToSemVer()

	if u.Stage == "" {
		logger.Panic("No stage set")
	}

	*r = append(*r, u)
}
