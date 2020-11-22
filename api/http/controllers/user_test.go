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
	"github.com/magiconair/properties/assert"
	"testing"
)

// getUserMock is a helper to obtain a mock user controller
// for testing
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

// TestUserController_Get - Test get all users endpoint
func TestUserController_Get(t *testing.T) {

	users := domain.Users{
		{Id: 123, FirstName: "Verbis", LastName: "CMS"},
		{Id: 124, FirstName: "Verbis", LastName: "CMS"},
	}

	pagination := http.Params{
		Page:           1,
		Limit:          15,
		OrderBy:        "id",
		OrderDirection: "asc",
		Filters:        nil,
	}

	// Test success
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Get", pagination).Return(users, 1, nil)

		test.NewRequest("GET", "/users", nil)
		getUserMock(&userMock).Get(test.gin)

		test.RunSuccess(users)
		assert.Equal(test.testing, test.Message(), "Successfully obtained users")
	})

	// Test errors.NOTFOUND
	t.Run("Not Found", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND})

		test.NewRequest("GET", "/users", nil)
		getUserMock(&userMock).Get(test.gin)

		assert.Equal(t, 200, test.recorder.Code)
	})

	// Test errors.CONFLICT
	t.Run("Conflict", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT})

		test.NewRequest("GET", "/users", nil)
		getUserMock(&userMock).Get(test.gin)

		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INVALID
	t.Run("Invalid", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID})

		test.NewRequest("GET", "/users", nil)
		getUserMock(&userMock).Get(test.gin)

		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INTERNAL
	t.Run("Internal Server Error", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL})

		test.NewRequest("GET", "/users", nil)
		getUserMock(&userMock).Get(test.gin)

		test.RunInternalError()
	})
}

// TestUserController_GetById - Test get by ID endpoint
func TestUserController_GetById(t *testing.T) {

	// Test success
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		user := domain.User{Id: 123, FirstName: "Verbis", LastName: "CMS"}
		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(user, nil)

		test.RequestAndServe("GET", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).GetById(g)
		})

		test.RunSuccess(&user)
		assert.Equal(test.testing, test.Message(), "Successfully obtained user with ID: 123")
	})

	// Test errors.INTERNAL
	t.Run("Internal Server Error", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))

		test.RequestAndServe("GET", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).GetById(g)
		})

		test.RunInternalError()
	})

	// Test invalid ID passed
	t.Run("Invalid ID", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))

		test.RequestAndServe("GET", "/users/wrongid", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).GetById(g)
		})

		test.RunParamError()
		assert.Equal(test.testing, test.Message(), "Pass a valid number to obtain the user by ID")
	})
}

// TestUserController_GetRoles - Test get roles endpoint
func TestUserController_GetRoles(t *testing.T) {

	roles := []domain.UserRole{
		{Id: 1, Name: "Banned", Description: "Banned Role"},
		{Id: 2, Name: "Administrator", Description: "Administrator Role"},
	}

	// Test success
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetRoles").Return(roles, nil)

		test.NewRequest("GET", "/roles", nil)
		getUserMock(&userMock).GetRoles(test.gin)

		test.RunSuccess(&roles)
		assert.Equal(test.testing, test.Message(), "Successfully obtained user roles")
	})

	// Test errors.INTERNAL
	t.Run("Fail", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetRoles").Return(nil, fmt.Errorf("err"))

		test.NewRequest("GET", "/roles", nil)

		getUserMock(&userMock).GetRoles(test.gin)

		test.RunInternalError()
	})
}

// TestUserController_Create - Test create endpoint
func TestUserController_Create(t *testing.T) {

	userCreate := &domain.UserCreate{
		User: domain.User{
			FirstName: "Verbis",
			LastName:  "CMS",
			Email:     "verbis@verbiscms.com",
			Role: domain.UserRole{
				Id: 1,
			},
		},
		Password:        "password",
		ConfirmPassword: "password",
	}

	user := domain.User{
		Id:        1,
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

	// Test success
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Create", userCreate).Return(user, nil)

		userString, err := json.Marshal(userCreate)
		if err != nil {
			t.Fatal(err)
		}

		test.NewRequest("POST", "/users", bytes.NewBuffer(userString))
		getUserMock(&userMock).Create(test.gin)

		test.RunSuccess(&user)
		assert.Equal(test.testing, test.Message(), "Successfully created user with ID: 1")
	})

	// Test bad validation
	t.Run("Validation Failed", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Create", userBadValidation).Return(nil, fmt.Errorf("error"))

		userString, err := json.Marshal(userBadValidation)
		if err != nil {
			t.Fatal(err)
		}

		test.NewRequest("POST", "/users", bytes.NewBuffer(userString))
		getUserMock(&userMock).Create(test.gin)

		test.RunValidationError()
		assert.Equal(test.testing, test.Message(), "Validation failed")
	})

	// Test errors.INVALID
	t.Run("Internal Server Error", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL})

		userString, err := json.Marshal(userCreate)
		if err != nil {
			t.Fatal(err)
		}

		test.NewRequest("POST", "/users", bytes.NewBuffer(userString))
		getUserMock(&userMock).Create(test.gin)
		test.RunInternalError()
	})

	// Test errors.CONFLICT
	t.Run("Conflict", func(t *testing.T) {
		test := newResponseRecorder(t)
		userMock := modelMocks.UserRepository{}

		userMock.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.CONFLICT})

		userString, err := json.Marshal(userCreate)
		if err != nil {
			t.Fatal(err)
		}

		test.NewRequest("POST", "/users", bytes.NewBuffer(userString))

		getUserMock(&userMock).Create(test.gin)
		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INVALID
	t.Run("Invalid", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Create", userCreate).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "error", Err: fmt.Errorf("err")})

		userString, err := json.Marshal(userCreate)
		if err != nil {
			t.Fatal(err)
		}

		test.NewRequest("POST", "/users", bytes.NewBuffer(userString))
		getUserMock(&userMock).Create(test.gin)

		assert.Equal(t, 400, test.recorder.Code)
	})
}

// TestUserController_Update - Test update endpoint
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

	// Test success
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Update", &user).Return(user, nil)
		userString, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		test.RequestAndServe("PUT", "/users/123", "/users/:id", bytes.NewBuffer(userString), func(g *gin.Context) {
			getUserMock(&userMock).Update(g)
		})

		test.RunSuccess(&user)
		assert.Equal(test.testing, test.Message(), "Successfully updated user with ID: 123")
	})

	// Test bad validation
	t.Run("Validation Failed", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Update", userBadValidation).Return(nil, fmt.Errorf("error"))
		userString, err := json.Marshal(userBadValidation)
		if err != nil {
			t.Fatal(err)
		}

		test.RequestAndServe("PUT", "/users/123", "/users/:id", bytes.NewBuffer(userString), func(g *gin.Context) {
			getUserMock(&userMock).Update(g)
		})

		test.RunValidationError()
		assert.Equal(test.testing, test.Message(), "Validation failed")
	})

	// Test invalid ID passed
	t.Run("Invalid ID", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Update", 123).Return(domain.User{}, fmt.Errorf("error"))
		userString, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		test.RequestAndServe("PUT", "/users/wrongid", "/users/:id", bytes.NewBuffer(userString), func(g *gin.Context) {
			getUserMock(&userMock).Update(g)
		})

		test.RunParamError()
		assert.Equal(test.testing, test.Message(), "A valid ID is required to update the user")
	})

	// Test errors.NOTFOUND
	t.Run("Not Found", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Update", userBadValidation).Return(nil, errors.Error{Code:      errors.NOTFOUND})
		userString, err := json.Marshal(userBadValidation)
		if err != nil {
			t.Fatal(err)
		}

		test.RequestAndServe("PUT", "/users/123", "/users/:id", bytes.NewBuffer(userString), func(g *gin.Context) {
			getUserMock(&userMock).Update(g)
		})

		getUserMock(&userMock).Create(test.gin)
		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INVALID
	t.Run("Internal Server Error", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Update", &user).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL})

		userString, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		test.RequestAndServe("PUT", "/users/123", "/users/:id", bytes.NewBuffer(userString), func(g *gin.Context) {
			getUserMock(&userMock).Update(g)
		})

		test.RunInternalError()
	})
}

// TestUserController_Delete - Test delete endpoint
func TestUserController_Delete(t *testing.T) {

	// Test invalid ID passed
	t.Run("Success", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Delete", 123).Return(nil)

		test.RequestAndServe("DELETE", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).Delete(g)
		})

		test.RunSuccess(nil)
	})

	// Test invalid ID passed
	t.Run("Invalid ID", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Delete", 0).Return(fmt.Errorf("error"))

		test.RequestAndServe("DELETE", "/users/wrongval", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).Delete(g)
		})

		test.RunParamError()
		assert.Equal(test.testing, test.Message(), "A valid ID is required to delete a user")
	})

	// Test errors.NOTFOUND
	t.Run("Not Found", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND})

		test.RequestAndServe("DELETE", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).Delete(g)
		})

		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INVALID
	t.Run("Invalid", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Delete", 123).Return(&errors.Error{Code: errors.INVALID})

		test.RequestAndServe("DELETE", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).Delete(g)
		})

		assert.Equal(t, 400, test.recorder.Code)
	})

	// Test errors.INVALID
	t.Run("Internal Server Error", func(t *testing.T) {
		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL})

		test.RequestAndServe("DELETE", "/users/123", "/users/:id", nil, func(g *gin.Context) {
			getUserMock(&userMock).Delete(g)
		})

		test.RunInternalError()
	})
}