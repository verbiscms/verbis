// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	mocks "github.com/verbiscms/verbis/api/mocks/cache"
	"net/http"
)

func (t *CacheTestSuite) TestCache_Clear() {
	tt := map[string]struct {
		status  int
		message string
		mock    func(m *mocks.Store)
		want    interface{}
	}{
		"Success": {
			http.StatusOK,
			"Successfully cleared server cache",
			func(m *mocks.Store) {
				m.On("Clear", mock.Anything).Return(nil)
			},
			nil,
		},
		"Error": {
			http.StatusInternalServerError,
			"Error clearing server cache",
			func(m *mocks.Store) {
				m.On("Clear", mock.Anything).Return(fmt.Errorf("error"))
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("POST", "/reset", "/reset", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Clear(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
