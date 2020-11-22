package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http"

	//"github.com/ainsleyclark/verbis/api/http"
	modelMocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	gohttp "net/http"
	"testing"
)

func TestUserController_Get(t *testing.T) {

	test := newResponseRecorder(t)

	t.Run("Success", func(t *testing.T) {

		users := domain.Users{
			{Id: 123, FirstName: "Verbis", LastName: "CMS"},
			{Id: 124, FirstName: "Verbis", LastName: "CMS"},
		}
		userMock := modelMocks.UserRepository{}
		userMock.On("Get", http.Params{
			Page:           1,
			Limit:          15,
			OrderBy:        "id",
			OrderDirection: "asc",
			Filters:        nil,
		}).Return(users, 1, nil)

		userController := UserController{
			store: &models.Store{
				User: &userMock,
			},
		}

		req, err := gohttp.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		test.gin.Request = req

		userController.Get(test.gin)

		test.runSuccess(users)
		assert.Equal(test.testing, test.GetMessage(), "Successfully obtained users")
	})
}

func TestUserController_GetById(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		test := newResponseRecorder(t)

		user := domain.User{
			Id:        123,
			FirstName: "Verbis",
			LastName:  "CMS",
		}
		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(user, nil)

		userController := UserController{
			store: &models.Store{
				User: &userMock,
			},
		}

		test.engine.GET("/users/:id", func(g *gin.Context) {
			userController.GetById(g)
		})

		req, _ := gohttp.NewRequest("GET", "/users/123", nil)
		test.engine.ServeHTTP(test.recorder, req)

		test.runSuccess(&user)
		assert.Equal(test.testing, test.GetMessage(), "Successfully obtained user with ID: 123")
	})

	t.Run("Internal Server Error", func(t *testing.T) {

		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))
		userController := UserController{
			store: &models.Store{
				User: &userMock,
			},
		}

		test.engine.GET("/users/:id", func(g *gin.Context) {
			userController.GetById(g)
		})

		req, _ := gohttp.NewRequest("GET", "/users/123", nil)
		test.engine.ServeHTTP(test.recorder, req)

		test.runInternalError()
	})

	t.Run("Invalid ID", func(t *testing.T) {

		test := newResponseRecorder(t)

		userMock := modelMocks.UserRepository{}
		userMock.On("GetById", 123).Return(domain.User{}, fmt.Errorf("error"))
		userController := UserController{
			store: &models.Store{
				User: &userMock,
			},
		}

		test.engine.GET("/users/:id", func(g *gin.Context) {
			userController.GetById(g)
		})

		req, _ := gohttp.NewRequest("GET", "/users/wrongid", nil)
		test.engine.ServeHTTP(test.recorder, req)

		test.runValidationError()
		assert.Equal(test.testing, test.GetMessage(), "Pass a valid number to obtain the user by ID")
	})
}
