// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"net/http"
)

var migrate = migration{
	From:   storageChange,
	To:     storageChange,
	Delete: false,
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
				m.On("Migrate", mock.Anything, migrate.From, migrate.To, migrate.Delete).
					Return(5, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "provider", Message: "Provider is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			migrateBadValidation,
			nil,
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", mock.Anything, migrate.From, migrate.To, migrate.Delete).
					Return(0, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", mock.Anything, migrate.From, migrate.To, migrate.Delete).
					Return(0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", mock.Anything, migrate.From, migrate.To, migrate.Delete).
					Return(0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			migrate,
			func(m *mocks.Provider) {
				m.On("Migrate", mock.Anything, migrate.From, migrate.To, migrate.Delete).
					Return(0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
