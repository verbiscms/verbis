// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package updates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/version/internal"
	_ "github.com/ainsleyclark/verbis/api/version/updates/v1"
	"github.com/gookit/color"
)

func Test() {
	for _, v := range internal.Registry {
		migration, err := Migrations.ReadFile(v.MigrationPath)
		if err != nil {
			color.Red.Println(err)
		}
		fmt.Println(string(migration))
	}
}
