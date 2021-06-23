// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/version"
	sm "github.com/hashicorp/go-version"
	"sort"
)

// Migration represents a singular migration for a single
// version.
type Migration struct {
	// The main version of the migration such as "v0.0.1"
	Version string
	// CallBackUp is a function called when the migration
	// is going up, this can be useful when manipulating
	// files and directories for the current version.
	CallBackUp CallBackFn
	// CallBackUp is a function called when the migration
	// is going Down, this is only called if an update
	// failed. And must be provided if CallBackUp is
	// defined.
	CallBackDown CallBackFn
	// Stage defines the release stage of the migration such as
	// Major, Minor or Patch,
	Stage version.Stage
	// The path of the MySQL file.
	SQLPath string
	// The path of the Postgres SQL file.
	PostgresPath string
	// Parsed SemVer of the Version
	SemVer *sm.Version
}

// CallBackFn is the function type when migrations are
// running up or Down.
type CallBackFn func() error

// MigrationRegistry contains a slice of pointers to each
// migration.
type MigrationRegistry []*Migration

// migrations is the in memory registry store of
// migrations.
var migrations = make(MigrationRegistry, 0)

var (
	// ErrCallBackMismatch is returned by AddMigration when
	// there has been a mismatch in the amount of callbacks
	// passed. Each migration should have two callbacks,
	// one up and one Down, or none at all.
	ErrCallBackMismatch = errors.New("both CallBackUp and CallBackDown must be set")
)

// AddMigration adds a migration to the update registry
// which will be called when Update() is run. The
// version and Stage must be attached to the
// migration.
func AddMigration(m *Migration) error {
	if m.Version == "" {
		return errors.New("no version provided for update")
	}

	if m.Stage == "" {
		return errors.New("no stage set")
	}

	if m.CallBackUp != nil && m.CallBackDown == nil {
		return ErrCallBackMismatch
	}

	if m.CallBackUp == nil && m.CallBackDown != nil {
		return ErrCallBackMismatch
	}

	if m.SemVer == nil {
		ver, err := sm.NewVersion(m.Version)
		if err != nil {
			return errors.New("malformed version: " + m.Version)
		}
		m.SemVer = ver
	}

	for _, migration := range migrations {
		if migration.SemVer.Equal(m.SemVer) {
			return errors.New("duplicate version")
		}
	}

	migrations = append(migrations, m)

	return nil
}

// hasCallBack returns true if CallBackUp and CallBackDown
// are both defined.
func (m *Migration) hasCallBack() bool {
	return m.CallBackUp != nil && m.CallBackDown != nil
}

// Sort MigrationRegistry is a type that implements the
// sort.Interface interface so that versions can be
// sorted.
func (m MigrationRegistry) Sort() {
	sort.Sort(m)
}

func (m MigrationRegistry) Len() int {
	return len(m)
}

func (m MigrationRegistry) Less(i, j int) bool {
	return m[i].SemVer.LessThan(m[j].SemVer)
}

func (m MigrationRegistry) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
