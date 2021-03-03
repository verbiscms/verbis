// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
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
		mock    func(m *mocks.OptionsRepository)
	}{
		"Success": {
			want:    nil,
			status:  200,
			message: "Successfully created/updated options",
			input:   optionsStruct,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Validation Failed": {
			want:    api.ErrorJSON{Errors: validation.Errors{{Key: "site_url", Message: "Site URL is required.", Type: "required"}}},
			status:  400,
			message: "Validation failed",
			input:   optionsBadValidation,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Validation Failed DB": {
			want:    nil,
			status:  400,
			message: "Validation failed",
			input:   "test",
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Internal Error": {
			want:    nil,
			status:  500,
			message: "internal",
			input:   optionsStruct,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/posts", "/posts", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).UpdateCreate(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
