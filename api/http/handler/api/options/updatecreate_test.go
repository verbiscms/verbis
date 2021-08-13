// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/cache"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mockCache "github.com/verbiscms/verbis/api/mocks/cache"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
	"net/http"
)

func (t *OptionsTestSuite) TestOptions_UpdateCreate() {
	jsonVOptions, err := json.Marshal(optionsStruct)
	if err != nil {
		t.Error(err)
	}

	dbOptions := domain.OptionsDBMap{}
	err = json.Unmarshal(jsonVOptions, &dbOptions)
	t.NoError(err)

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		cache   func(m *mockCache.Store)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully created/updated options",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			func(m *mockCache.Store) {
				m.On("Invalidate", mock.Anything, cache.InvalidateOptions{Tags: []string{"options"}}).Return(nil)
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "site_url", Message: "Site URL is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			optionsBadValidation,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			nil,
		},
		"Validation Failed DB": {
			nil,
			http.StatusBadRequest,
			"Validation failed",
			"test",
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			nil,
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			nil,
		},
		"Invalidate Error": {
			nil,
			http.StatusInternalServerError,
			"invalidate",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			func(m *mockCache.Store) {
				m.On("Invalidate", mock.Anything, cache.InvalidateOptions{Tags: []string{"options"}}).Return(&errors.Error{Code: errors.INTERNAL, Message: "invalidate"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				c := &mockCache.Store{}
				if test.cache != nil {
					test.cache(c)
				}
				o := t.Setup(test.mock)
				o.Cache = c

				o.UpdateCreate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
