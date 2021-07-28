// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	events "github.com/verbiscms/verbis/api/mocks/events"
	mocks "github.com/verbiscms/verbis/api/mocks/store/auth"
	users "github.com/verbiscms/verbis/api/mocks/store/users"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_SendResetPassword() {
	var (
		srp = SendResetPassword{
			Email: "info@verbiscms.com",
		}
		srpBadValidation = SendResetPassword{}
		dispatchSuccess  = func(m *events.Dispatcher) {
			m.On("Dispatch", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		}
		hashSuccess = func(email string) (string, error) {
			return "token", nil
		}
	)

	tt := map[string]struct {
		want       interface{}
		status     int
		message    string
		input      interface{}
		dispatcher func(m *events.Dispatcher)
		mock       func(m *mocks.Repository, c *cache.Store, u *users.Repository)
		tokenFunc  func(email string) (string, error)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"A fresh verification link has been sent to your email",
			srp,
			dispatchSuccess,
			func(m *mocks.Repository, c *cache.Store, u *users.Repository) {
				u.On("FindByEmail", srp.Email).Return(user, nil)
				c.On("Set", mock.Anything, "token", user, mock.Anything).Return(nil)
			},
			hashSuccess,
		},
		"Validation Failed": {
			`{"errors":[{"key":"email","message":"Email is required.","type":"required"}]}`,
			http.StatusBadRequest,
			"Validation failed",
			srpBadValidation,
			dispatchSuccess,
			nil,
			hashSuccess,
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"No user found with email: " + srp.Email,
			srp,
			dispatchSuccess,
			func(m *mocks.Repository, c *cache.Store, u *users.Repository) {
				u.On("FindByEmail", srp.Email).Return(domain.User{}, fmt.Errorf("error"))
			},
			hashSuccess,
		},
		"Token Error": {
			nil,
			http.StatusInternalServerError,
			"Error generating user token",
			srp,
			dispatchSuccess,
			func(m *mocks.Repository, c *cache.Store, u *users.Repository) {
				u.On("FindByEmail", srp.Email).Return(user, nil)
			},
			func(email string) (string, error) {
				return "", fmt.Errorf("error")
			},
		},
		"Dispatch Error": {
			nil,
			http.StatusInternalServerError,
			"dispatch",
			srp,
			func(m *events.Dispatcher) {
				m.On("Dispatch", mock.Anything, mock.Anything, mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "dispatch"})
			},
			func(m *mocks.Repository, c *cache.Store, u *users.Repository) {
				u.On("FindByEmail", srp.Email).Return(user, nil)
				c.On("Set", mock.Anything, "token", user, mock.Anything).Return(nil)
			},
			hashSuccess,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/sendreset", "/sendreset", test.input, func(ctx *gin.Context) {
				a := t.SetupDispatcher(test.mock, test.dispatcher)
				a.generateTokenFunc = test.tokenFunc
				a.SendResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
