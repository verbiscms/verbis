// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_Disconnect() {
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
			"Successfully updated storage options",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Disconnect").Return(nil)
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Disconnect").Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Disconnect").Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			storageChange,
			func(m *mocks.Provider) {
				m.On("Disconnect").Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/storage", "/storage", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Disconnect(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
