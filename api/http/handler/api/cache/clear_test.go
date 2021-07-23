// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"net/http"
)

func (t *CacheTestSuite) TestCache_Clear() {
	tt := map[string]struct {
		status  int
		message string
		want    interface{}
	}{
		"Success": {
			http.StatusOK,
			"Successfully cleared server cache",
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := New(&deps.Deps{})
			t.RequestAndServe("POST", "/reset", "/reset", nil, func(ctx *gin.Context) {
				mock.Clear(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
