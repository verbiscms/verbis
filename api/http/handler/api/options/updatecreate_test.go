// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
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
		cache   func(m *cache.Store)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully created/updated options",
			optionsStruct,
			func(m *mocks.Repository) {
				m.On("Insert", dbOptions).Return(nil)
			},
			func(m *cache.Store) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
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
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				c := &cache.Store{}
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
