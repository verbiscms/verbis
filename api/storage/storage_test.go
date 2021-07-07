// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
)

//func TestStorage_Delete() {
//	id := 1
//
//	tt := map[string]struct {
//		input *int
//		want  bool
//	}{
//		"Resource": {
//			&id,
//			true,
//		},
//		"No Resource": {
//			nil,
//			false,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			p := Category{
//				ParentId: test.input,
//			}
//			got := p.HasParent()
//			assert.Equal(t, test.want, got)
//		})
//	}
//}
