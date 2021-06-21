// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package updates

import (
	"database/sql"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/ainsleyclark/verbis/api/version/internal"
	_ "github.com/ainsleyclark/verbis/api/version/updates/v0"
)

type updater struct {
	tx        *sql.Tx
	downFuncs []internal.CallBackFn
}

func New(db database.Driver) (*updater, error) {
	tx, err := db.DB().Begin()
	if err != nil {
		return nil, err
	}
	return &updater{
		tx:        tx,
		downFuncs: nil,
	}, nil
}

func (u *updater) Run() error {
	internal.Updates.Sort()

	for _, update := range internal.Updates {
		shouldRun := version.SemVer.LessThanOrEqual(update.ToSemVer())
		if !shouldRun {
			continue
		}

		err := u.process(update)
		if err != nil {
			rollBackErr := u.rollBack()
			if rollBackErr != nil {
				// In a dirty state
				return rollBackErr
			}
			return err
		}

		u.downFuncs = append(u.downFuncs, update.CallBackDown)
	}

	err := u.tx.Commit()
	if err != nil {
		return err
	}

	u.downFuncs = nil

	return nil
}

func (u *updater) rollBack() error {
	err := u.tx.Rollback()
	if err != nil {
		return err
	}
	for _, fn := range u.downFuncs {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *updater) process(update *internal.Update) error {
	migration, err := Migrations.ReadFile(update.MigrationPath)
	if err != nil {
		return err
	}

	_, err = u.tx.Exec(string(migration))
	if err != nil {
		return err
	}

	if update.HasCallBack() {
		err := update.CallBackUp()
		if err != nil {
			return err
		}
	}

	return nil
}
