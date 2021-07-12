// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_ListBuckets() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Provider)
	}{
		"Success": {
			buckets,
			http.StatusOK,
			"Successfully obtained buckets",
			func(m *mocks.Provider) {
				m.On("ListBuckets").Return(buckets, nil)
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Provider) {
				m.On("ListBuckets").Return(nil, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Provider) {
				m.On("ListBuckets").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("GET", "/buckets", "/buckets", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).ListBuckets(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
