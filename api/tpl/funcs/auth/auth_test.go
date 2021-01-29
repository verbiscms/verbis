package auth

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *TplTestSuite) Test_Auth() {

	tt := map[string]struct {
		want   interface{}
		cookie string
		mock   func(m *mocks.UserRepository)
	}{
		"Logged In": {
			want:   true,
			cookie: "verbis-session=token",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No Cookie": {
			want:   false,
			cookie: "",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No User": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := mocks.UserRepository{}

			test.mock(&mock)
			t.store.User = &mock
			t.gin.Request.Header.Set("Cookie", test.cookie)

			tpl := "{{ auth }}"
			t.RunT(tpl, test.want)
		})
	}
}

func (t *TplTestSuite) Test_Admin() {

	tt := map[string]struct {
		want   interface{}
		cookie string
		mock   func(m *mocks.UserRepository)
	}{
		"Is Admin": {
			want:   true,
			cookie: "verbis-session=token",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{
					UserPart: domain.UserPart{Id: 0, Role: domain.UserRole{Id: 6}},
				}, nil)
			},
		},
		"Not Admin": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{
					UserPart: domain.UserPart{Id: 0, Role: domain.UserRole{Id: 1}},
				}, nil)
			},
		},
		"No Cookie": {
			want:   false,
			cookie: "",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No User": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.UserRepository) {
				m.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := mocks.UserRepository{}

			test.mock(&mock)
			t.store.User = &mock
			t.gin.Request.Header.Set("Cookie", test.cookie)

			tpl := "{{ admin }}"
			t.RunT(tpl, test.want)
		})
	}
}
