// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"fmt"
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
		cache   func(m *mockCache.Cacher)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully created/updated options",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			func(m *mockCache.Cacher) {
				m.On("Set", mock.Anything, cache.OptionsKey, mock.Anything, mock.Anything).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "site_url", Message: "Site Url is required.", Type: "required"}}},
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
		"Cache Fail": {
			nil,
			http.StatusInternalServerError,
			"Error updating options cache",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			func(m *mockCache.Cacher) {
				m.On("Set", mock.Anything, cache.OptionsKey, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				c := &mockCache.Cacher{}
				if test.cache != nil {
					test.cache(c)
				}
				cache.SetDriver(c)
				t.Setup(test.mock).UpdateCreate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
