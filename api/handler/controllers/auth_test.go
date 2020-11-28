package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// getAuthMock is a helper to obtain a mock auth controller
// for testing.
func getAuthMock(m models.AuthRepository) *AuthController {
	return &AuthController{
		store: &models.Store{
			Auth: m,
		},
	}
}

// Test_NewAuth - Test construct
func Test_NewAuth(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &AuthController{
		store:  &store,
		config: config,
	}
	got := newAuth(&store, config)
	assert.Equal(t, got, want)
}

// TestAuthController_Get - Test Login route
func TestAuthController_Login(t *testing.T) {

	login := domain.Login{Email: "info@verbiscms.com", Password: "password"}
	loginBadValidation := domain.Login{Password: "password"}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   interface{}
		cookie  bool
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			want:    `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"","id":0,"instagram":null,"last_name":"","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully logged in & session started",
			input:   login,
			cookie:  true,
			mock: func(m *mocks.AuthRepository) {
				m.On("Authenticate", login.Email, login.Password).Return(domain.User{}, nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"email","message":"Email is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   loginBadValidation,
			cookie:  false,
			mock: func(m *mocks.AuthRepository) {
				m.On("Authenticate", loginBadValidation.Email, loginBadValidation.Email).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Not Authorised": {
			want:    `{}`,
			status:  401,
			message: "unauthorised",
			input:   login,
			cookie:  false,
			mock: func(m *mocks.AuthRepository) {
				m.On("Authenticate", login.Email, login.Password).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "unauthorised"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.AuthRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/login", "/login", bytes.NewBuffer(body), func(g *gin.Context) {
				getAuthMock(mock).Login(g)
			})

			rr.Run(test.want, test.status, test.message)

			if test.cookie {
				cookie := http.Cookie{
					Name:     "verbis-session",
					Expires:  time.Time{},
					MaxAge:   172800,
					Path:     "/",
					Raw:      "verbis-session=; Path=/; Max-Age=172800; HttpOnly",
					HttpOnly: true,
				}
				assert.Equal(t, rr.recorder.Result().Cookies()[0], &cookie)
			}
		})
	}
}

// TestAuthController_Logout - Test Logout route
func TestAuthController_Logout(t *testing.T) {

	token := "test"

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   string
		cookie  bool
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully logged out",
			input:   "test",
			cookie:  true,
			mock: func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   token,
			cookie:  false,
			mock: func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   token,
			cookie:  false,
			mock: func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.AuthRepository{}
			test.mock(mock)

			rr.NewRequest("POST", "/logout", nil)
			rr.gin.Request.Header.Set("token", test.input)

			getAuthMock(mock).Logout(rr.gin)

			rr.Run(test.want, test.status, test.message)

			if test.cookie {
				cookie := http.Cookie{
					Name:     "verbis-session",
					Expires:  time.Time{},
					MaxAge:   -1,
					Path:     "/",
					Raw:      "verbis-session=; Path=/; Max-Age=0; HttpOnly",
					HttpOnly: true,
				}
				assert.Equal(t, rr.recorder.Result().Cookies()[0], &cookie)
			}
		})
	}
}

// TestAuthController_ResetPassword - Test ResetPassword route
func TestAuthController_ResetPassword(t *testing.T) {

	rp := domain.ResetPassword{
		NewPassword:     "password",
		ConfirmPassword: "password",
		Token:           "token",
	}

	rpdBadValidation := domain.ResetPassword{
		NewPassword: "password",
		Token:       "token",
	}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully reset password",
			input:   rp,
			mock: func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"confirm_password","message":"Confirm Password must equal the New Password.","type":"eqfield"}]}`,
			status:  400,
			message: "Validation failed",
			input:   rpdBadValidation,
			mock: func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rpdBadValidation.Token, rpdBadValidation.NewPassword).Return(nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   rp,
			mock: func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   rp,
			mock: func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.AuthRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/reset", "/reset", bytes.NewBuffer(body), func(g *gin.Context) {
				getAuthMock(mock).ResetPassword(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestAuthController_VerifyPasswordToken - Test VerifyPasswordToken route
func TestAuthController_VerifyPasswordToken(t *testing.T) {

	token := "test"

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   string
		mock    func(m *mocks.AuthRepository)
		url     string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully verified token",
			input:   token,
			mock: func(m *mocks.AuthRepository) {
				m.On("VerifyPasswordToken", token).Return(nil)
			},
			url: "/verify/" + token,
		},
		"Not Found": {
			want:    `{}`,
			status:  404,
			message: "not found",
			input:   token,
			mock: func(m *mocks.AuthRepository) {
				m.On("VerifyPasswordToken", token).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/verify/" + token,
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.AuthRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/verify/:token", nil, func(g *gin.Context) {
				getAuthMock(mock).VerifyPasswordToken(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestAuthController_SendResetPassword - Test endResetPassword route
func TestAuthController_SendResetPassword(t *testing.T) {

	srp := domain.SendResetPassword{Email: "info@verbiscms.com"}
	srpBadvalidation := domain.SendResetPassword{}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "A fresh verification link has been sent to your email",
			input:   srp,
			mock: func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"email","message":"Email is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   srpBadvalidation,
			mock: func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srpBadvalidation.Email).Return(nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   srp,
			mock: func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   srp,
			mock: func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.AuthRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/sendreset", "/sendreset", bytes.NewBuffer(body), func(g *gin.Context) {
				getAuthMock(mock).SendResetPassword(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
