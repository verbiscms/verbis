// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/gin-gonic/gin"
)

func (t *CacheTestSuite) TestCache_Clear() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   string
	}{
		"Success": {
			want:    nil,
			status:  200,
			message: "Successfully cleared server cache",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := &Cache{}
			t.RequestAndServe("POST", "/reset", "/reset", nil, func(ctx *gin.Context) {
				mock.Clear(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
