package tpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	//"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	//"time"
)

func (t *TplTestSuite) Test_GetUser() {

	user := domain.User{
		UserPart: domain.UserPart{Id: 1, FirstName: "verbis"},
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.UserRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil)
			},
			want: user.HideCredentials(),
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.User{}, fmt.Errorf("error"))
			},
			want: "",
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil)
			},
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			userMock := mocks.UserRepository{}

			test.mock(&userMock)
			t.store.User = &userMock

			tpl := `{{ user . }}`
			t.RunTWithData(tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_GetUsers() {

	users := domain.Users{
		domain.User{
			UserPart: domain.UserPart{Id: 1, FirstName: "verbis"},
		},
		domain.User{
			UserPart: domain.UserPart{Id: 1, FirstName: "cms"},
		},
	}

	tt := map[string]struct {
		input map[string]interface{}
		mock  func(m *mocks.UserRepository)
		want  interface{}
	}{
		"Success": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.UserRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(users, 2, nil)
			},
			want: map[string]interface{}{
				"Users": users.HideCredentials(),
				"Pagination": &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Failed Params": {
			input: map[string]interface{}{"limit": "wrongval"},
			mock: func(m *mocks.UserRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(users, 2, nil)
			},
			want: "cannot unmarshal string into Go struct field TemplateParams.limit",
		},
		"Not Found": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.UserRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
			want: map[string]interface{}(nil),
		},
		"Internal Error": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.UserRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
			},
			want: "internal error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			userMock := mocks.UserRepository{}

			test.mock(&userMock)
			t.store.User = &userMock

			c, err := t.getUsers(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.EqualValues(test.want, c)
		})
	}
}
