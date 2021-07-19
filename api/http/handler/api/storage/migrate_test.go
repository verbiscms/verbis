// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	validation "github.com/ainsleyclark/verbis/api/common/vaidation"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

var migrate = migration{
	From: storageChange,
	To:   storageChange,
}

var migrateBadValidation = migration{
	From: storageChangeBadValidation,
	To:   storageChange,
}

func (t *StorageTestSuite) TestStorage_Migrate() {
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
			"Successfully started migration, processing 5 files",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", migrate.From, migrate.To).Return(5, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "from_provider", Message: "From Provider is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			migrateBadValidation,
			nil,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", migrate.From, migrate.To).Return(0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", migrate.From, migrate.To).Return(0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", migrate.From, migrate.To).Return(0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/storage/migrate", "/storage/migrate", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Migrate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
