// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database/updates"
	"github.com/ainsleyclark/verbis/api/logger"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
)

type Migrator interface {
	Migrate(version *sm.Version) error
}

type migrate struct {
	down   []CallBackFn
	Driver string
	DB     *sqlx.DB
}

func NewMigrator(driver string, db *sqlx.DB) (Migrator, error) {
	if driver != MySQLDriver && driver != PostgresDriver {
		return nil, fmt.Errorf("wrong driver")
	}
	return &migrate{
		down:   make([]CallBackFn, 0),
		Driver: driver,
		DB:     db,
	}, nil
}

func (r *migrate) Migrate(version *sm.Version) (err error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Panic(err)
			}
			for _, fn := range r.down {
				err := fn()
				if err != nil {
					logger.Panic(err)
				}
			}
			return
		}
	}()

	migrations.Sort()
	for _, migration := range migrations {
		if version.GreaterThanOrEqual(migration.SemVer) {
			continue
		}

		err = r.process(migration, tx)
		if err != nil {
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
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
		migration, err := updates.Static.ReadFile(path)
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
