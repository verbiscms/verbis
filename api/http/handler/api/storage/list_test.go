// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_ListBuckets() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Provider)
		url     string
	}{
		"Success": {
			buckets,
			http.StatusOK,
			"Successfully obtained buckets",
			func(m *mocks.Provider) {
				m.On("ListBuckets", domain.StorageAWS).Return(buckets, nil)
			},
			"/buckets/aws",
		},
		"Local": {
			nil,
			http.StatusForbidden,
			"Obtaining local buckets are forbidden",
			nil,
			"/buckets/local",
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Provider) {
				m.On("ListBuckets", domain.StorageAWS).Return(nil, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			"/buckets/aws",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Provider) {
				m.On("ListBuckets", domain.StorageAWS).Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/buckets/aws",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("GET", test.url, "/buckets/:name", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).ListBuckets(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
