// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
)

func (t *StorageTestSuite) TestMigrationInfo_Fail() {
	mi := MigrationInfo{
		Failed: 0,
		Errors: nil,
	}
	mi.fail(fileRemote, fmt.Errorf("error"))
	want := MigrationInfo{
		Failed: 1,
		Errors: []FailedMigrationFile{{Error: fmt.Errorf("error"), File: fileRemote}},
	}
	t.Equal(want, mi)
}

func (t *StorageTestSuite) TestMigrationInfo_Succeed() {
	tt := map[string]struct {
		input MigrationInfo
		want  MigrationInfo
	}{
		"Simple": {
			MigrationInfo{Total: 100, Succeeded: 0},
			MigrationInfo{Total: 100, Succeeded: 1, Progress: 1},
		},
		"Half": {
			MigrationInfo{Total: 100, Succeeded: 50},
			MigrationInfo{Total: 100, Succeeded: 51, Progress: 51},
		},
		"100": {
			MigrationInfo{Total: 100, Succeeded: 99},
			MigrationInfo{Total: 100, Succeeded: 100, Progress: 100},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			test.input.succeed()
			t.Equal(test.want, test.input)
		})
	}
}
