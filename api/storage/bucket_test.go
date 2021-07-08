// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/files"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
)

func (t *StorageTestSuite) TestBucket_Upload() {
	tt := map[string]struct {
		input domain.Upload
		want interface{}
	}{
		"Resource": {
			domain.Upload{}
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			fmt.Println(test)
		})
	}
}
