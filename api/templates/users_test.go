package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	//"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"
	//"time"

	"testing"
)

func Test_GetUser(t *testing.T) {

	user := domain.User{
		UserPart: domain.UserPart{Id: 1, FirstName: "verbis"},
	}
	user.HideCredentials()
	//fmt.Printf("%+v\n", user)

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
			want: user,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.User{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			userMock := mocks.UserRepository{}

			test.mock(&userMock)
			f.store.User = &userMock

			tpl := `{{ user . }}`

			runtv(t, f, tpl, test.want, test.input)
		})
	}
}

func Test_GetUsers(t *testing.T) {

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
				"Users": users,
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
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			userMock := mocks.UserRepository{}

			test.mock(&userMock)
			f.store.User = &userMock

			c, err := f.getUsers(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.EqualValues(t, test.want, c)
		})
	}
}
