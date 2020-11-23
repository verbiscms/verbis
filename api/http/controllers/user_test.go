package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	modelMocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getUserMock is a helper to obtain a mock user controller
// for testing.
func getUserMock(userMock models.UserRepository) *UserController {
	return &UserController{
		store: &models.Store{
			User: userMock,
		},
	}
}

// Test_NewUser - Test construct
func Test_NewUser(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &UserController{
		store:  &store,
		config: config,
	}
	got := newUser(&store, config)
	assert.Equal(t, got, want)
}

// TestUserController_Get - Test Get route
func TestUserController_Get(t *testing.T) {

	users := domain.Users{
		{Id: 123, FirstName: "Verbis", LastName: "CMS"},
		{Id: 124, FirstName: "Verbis", LastName: "CMS"},
	}
	pagination := http.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "asc", Filters: nil}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *modelMocks.UserRepository)
	}{
		{
			name:       "Success",
			want:       `[{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":124,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}]`,
			status:     200,
			message:    "Successfully obtained users",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Get", pagination).Return(users, 1, nil)
			},
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     200,
			message:    "no users found",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
		},
		{
			name:       "Conflict",
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		{
			name:       "Invalid",
			want:       `{}`,
			status:     400,
			message:    "invalid",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/users", "/users", nil, func(g *gin.Context) {
				getUserMock(mock).Get(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_GetById - Test GetByID route
func TestUserController_GetById(t *testing.T) {

	user := domain.User{Id: 123, FirstName: "Verbis", LastName: "CMS"}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *modelMocks.UserRepository)
		url string
	}{
		{
			name:       "Success",
			want:       `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully obtained user with ID: 123",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetById", 123).Return(user, nil)
			},
			url: "/users/123",
		},
		{
			name:       "Invalid ID",
			want:       `{}`,
			status:     400,
			message:    "Pass a valid number to obtain the user by ID",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/wrongid",
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     200,
			message:    "no users found",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
			url: "/users/123",
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/users/:id", nil, func(g *gin.Context) {
				getUserMock(mock).GetById(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_GetRoles - Test GetRoles route
func TestUserController_GetRoles(t *testing.T) {

	roles := []domain.UserRole{
		{Id: 1, Name: "Banned", Description: "Banned Role"},
		{Id: 2, Name: "Administrator", Description: "Administrator Role"},
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *modelMocks.UserRepository)
	}{
		{
			name:       "Success",
			want:       `[{"description":"Banned Role","id":1,"name":"Banned"},{"description":"Administrator Role","id":2,"name":"Administrator"}]`,
			status:     200,
			message:    "Successfully obtained user roles",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetRoles").Return(roles, nil)
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetRoles").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/roles", "/roles", nil, func(g *gin.Context) {
				getUserMock(mock).GetRoles(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_Create - Test Create route
func TestUserController_Create(t *testing.T) {

	userCreate := &domain.UserCreate{
		User: domain.User{
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.UserRole{
				Id: 123,
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}

	user := domain.User{
		Id:        123,
		FirstName: "Verbis",
		LastName:  "CMS",
		Email:     "verbis@verbiscms.com",
	}

	userBadValidation := &domain.UserCreate{
		User: domain.User{
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
		},
		Password:        "password",
		ConfirmPassword: "password",
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		input interface{}
		mock func(u *modelMocks.UserRepository)
	}{
		{
			name:       "Success",
			want:       `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"verbis@verbiscms.com","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully created user with ID: 123",
			input: userCreate,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Create", userCreate).Return(user, nil)
			},
		},
		{
			name:       "Validation Failed",
			want:       `{"errors":[{"key":"role_id","message":"User Role Id is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: userBadValidation,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Create", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		{
			name:       "Invalid",
			want:       `{}`,
			status:     400,
			message:    "invalid",
			input: userCreate,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		{
			name:       "Conflict",
			want:       `{}`,
			status:     400,
			message:    "conflict",
			input: userCreate,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: userCreate,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/users", "/users", bytes.NewBuffer(body), func(g *gin.Context) {
				getUserMock(mock).Create(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_Update - Test Update route
func TestUserController_Update(t *testing.T) {

	user := domain.User{
		Id:        123,
		FirstName: "Verbis",
		LastName:  "CMS",
		Email:     "verbis@verbiscms.com",
		Role: domain.UserRole{
			Id: 1,
		},
	}

	userBadValidation := &domain.User{
		FirstName: "Verbis",
		LastName:  "CMS",
		Email:     "verbis@verbiscms.com",
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		input interface{}
		mock func(u *modelMocks.UserRepository)
		url string
	}{
		{
			name:       "Success",
			want:       `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"verbis@verbiscms.com","email_verified_at":null,"facebook":null,"first_name":"Verbis","id":123,"instagram":null,"last_name":"CMS","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":1,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully updated user with ID: 123",
			input: user,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Update", &user).Return(user, nil)
			},
			url: "/users/123",
		},
		{
			name:       "Validation Failed",
			want:       `{"errors":[{"key":"role_id","message":"Role Id is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: userBadValidation,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/123",
		},
		{
			name:       "Invalid ID",
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to update the user",
			input: user,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			url: "/users/wrongid",
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     400,
			message:    "not found",
			input: user,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Update", &user).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/users/123",
		},
		{
			name:       "Internal",
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: user,
			mock: func(u *modelMocks.UserRepository) {
				u.On("Update", &user).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},

	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/users/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getUserMock(mock).Update(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_Delete - Test Delete route
func TestUserController_Delete(t *testing.T) {

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *modelMocks.UserRepository)
		url string
	}{
		{
			name:       "Success",
			want:       `null`,
			status:     200,
			message:    "Successfully deleted user with ID: 123",
			mock: func(u *modelMocks.UserRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/users/123",
		},
		{
			name:       "Invalid ID",
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to delete a user",
			mock:  func(u *modelMocks.UserRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/users/wrongid",
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     400,
			message:    "not found",
			mock:  func(u *modelMocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/users/123",
		},
		{
			name:       "Conflict",
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock:  func(u *modelMocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/users/123",
		},
		{
			name:       "Internal",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock:  func(u *modelMocks.UserRepository) {
				u.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/users/123",
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/users/:id", nil, func(g *gin.Context) {
				getUserMock(mock).Delete(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestUserController_ResetPassword - Test Reset Password route
func TestUserController_ResetPassword(t *testing.T) {

	reset := &domain.UserPasswordReset{
		DBPassword: "password",
		CurrentPassword: "password",
		NewPassword:     "verbiscms",
		ConfirmPassword: "verbiscms",
	}

	//resetBadValidation := &domain.UserPasswordReset{
	//	CurrentPassword: "password",
	//	NewPassword:     "verbis",
	//	ConfirmPassword: "verbisc",
	//}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		input interface{}
		mock func(u *modelMocks.UserRepository)
		url string
	}{
		{
			name:       "Success",
			want:       `{}`,
			status:     200,
			message:    "Successfully updated password for the user with ID: 123",
			mock: func(u *modelMocks.UserRepository) {
				u.On("GetById", 123).Return(domain.User{}, nil)
				u.On("ResetPassword", reset).Return(nil)
			},
			input: reset,
			url: "/users/reset/123",
		},
		//{
		//	name:       "Invalid ID",
		//	want:       `{}`,
		//	status:     400,
		//	message:    "A valid ID is required to update a user's password",
		//	mock:  func(u *modelMocks.UserRepository) {
		//		u.On("ResetPassword", 123).Return(nil)
		//		u.On("GetById", )
		//	},
		//	input: reset,
		//	url: "/users/wrongid",
		//},
		//{
		//	name:       "Not Found",
		//	want:       `{}`,
		//	status:     400,
		//	message:    "not found",
		//	mock:  func(u *modelMocks.UserRepository) {
		//		u.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
		//	},
		//	url: "/users/123",
		//},
		//{
		//	name:       "Conflict",
		//	want:       `{}`,
		//	status:     400,
		//	message:    "conflict",
		//	mock:  func(u *modelMocks.UserRepository) {
		//		u.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
		//	},
		//	url: "/users/123",
		//},
		//{
		//	name:       "Internal",
		//	want:       `{}`,
		//	status:     500,
		//	message:    "internal",
		//	mock:  func(u *modelMocks.UserRepository) {
		//		u.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
		//	},
		//	url: "/users/123",
		//},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &modelMocks.UserRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("DELETE", test.url, "/users/reset/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getUserMock(mock).ResetPassword(g)
			})

			//assert.JSONEq(t, test.want, rr.Data())
			//assert.Equal(t, test.status, rr.recorder.Code)
			//assert.Equal(t, test.message, rr.Message())
			//assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}
