// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"database/sql"
	"embed"
	"fmt"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
	"github.com/verbiscms/verbis/api/database/updates"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Migrator is the core updater for Verbis, it takes in a
// version and migrates the migrations up to the major
// version. For example, if the current version is
/// 1.0.2, the migrator will run up to version
// 1.9.9 and no more, to allow for schema
// changes.
type Migrator interface {
	Migrate(version *sm.Version) error
}

// migrate is the implementation of the Migrator which has
// a database, driver, embed file system and callback
// down functions attached ti it.
type migrate struct {
	Driver string
	DB     *sqlx.DB
	Embed  embed.FS
	Down   []CallBackFn
}

// NewMigrator creates a migrator for updating and
// installing Verbis. Returns an error on driver
// mismatch.
func NewMigrator(driver string, db *sqlx.DB) (Migrator, error) {
	if driver != MySQLDriver && driver != PostgresDriver {
		return nil, fmt.Errorf("wrong driver")
	}
	return &migrate{
		Embed:  updates.Static,
		Driver: driver,
		DB:     db,
		Down:   make([]CallBackFn, 0),
	}, nil
}

// Migrate migrates Verbis up to the most recent version
// but no more than a version.Major.
// The function panics if the migration could not be
// rolled back if an error occurred.
func (r *migrate) Migrate(version *sm.Version) (err error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				logger.Panic(rollBackErr)
			}
			for _, fn := range r.Down {
				downErr := fn()
				if downErr != nil {
					logger.Panic(downErr)
				}
			}
			r.Down = nil
			return
		}
		r.Down = nil
	}()

	migrations.Sort()
	for _, migration := range migrations {
		var canMigrate bool

		canMigrate, err = r.canMigrate(version, migration.SemVer)
		if err != nil {
			return
		}
		if !canMigrate {
			continue
		}

		err = r.process(migration, tx)
		if err != nil {
			return
		}

		r.Down = append(r.Down, migration.CallBackDown)
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return //nolint
}

// canMigrate checks to see if the migration can run by
// comparing against the currently running version.
// If the base is greater than or equal to, or
// the migration is not inside a major version
// e.g (v1 - v2) the function will return false.
func (r *migrate) canMigrate(base, migrate *sm.Version) (bool, error) {
	if base.GreaterThan(migrate) {
		return false, nil
	}

	seg := base.Segments()
	if len(seg) < 1 {
		return false, errors.New("error parsing version, invalid length")
	}

	constraint, err := sm.NewConstraint(fmt.Sprintf(">= %d.0, < %d.0", seg[0], seg[0]+1))
	if err != nil {
		return false, err
	}

	if !constraint.Check(migrate) {
		return false, nil
	}

	return true, nil
}

// process reads the migration and executes the migration
// if there is one. Calls the callback function if there
// is one set.
func (r *migrate) process(m *Migration, tx *sql.Tx) error {
	path := m.SQLPath
	if r.Driver == PostgresDriver {
		path = m.PostgresPath
	}

	if path != "" {
		migration, err := r.Embed.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = tx.Exec(string(migration))
		if err != nil {
			return err
		}
	}

	if m.hasCallBack() {
		err := m.CallBackUp()
		if err != nil {
			return err
		}
	}

	return nil
}
