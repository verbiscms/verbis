package controllers

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"testing"
)

func TestUserController_Get(t *testing.T) {

	test := newResponseRecorder(t)

	t.Run("Success", func(t *testing.T) {

		users := []domain.User{
			{
				Id:        123,
				FirstName: "Verbis",
				LastName:  "CMS",
			},
		}
		userMock := mocks.UserRepository{}
		userMock.On("Get", http.Params{}).Return(&users, 1, nil)

		userController := UserController{
			store:  &models.Store{
				User:       &userMock,
			},
		}

 		userController.Get(test.gin)

		test.runSuccess(users)
	})
}
