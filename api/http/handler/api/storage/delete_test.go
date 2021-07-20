// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	validation "github.com/ainsleyclark/verbis/api/common/vaidation"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_DeleteBucket() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Provider)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully deleted bucket: " + storageChange.Bucket,
			storageChange,
			func(m *mocks.Provider) {
				m.On("DeleteBucket", storageChange.Provider, storageChange.Bucket).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "provider", Message: "Provider is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			storageChangeBadValidation,
			nil,
		},
		"Local": {
			nil,
			http.StatusBadRequest,
			"Local bucket cannot be deleted",
			domain.StorageChange{Provider: domain.StorageLocal},
			nil,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			storageChange,
			func(m *mocks.Provider) {
				m.On("DeleteBucket", storageChange.Provider, storageChange.Bucket).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			storageChange,
			func(m *mocks.Provider) {
				m.On("DeleteBucket", storageChange.Provider, storageChange.Bucket).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			storageChange,
			func(m *mocks.Provider) {
				m.On("DeleteBucket", storageChange.Provider, storageChange.Bucket).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, "/storage", "/storage", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).DeleteBucket(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
