package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func Test_Auth(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			mock := mocks.UserRepository{}
			f := newTestSuite()

			test.mock(&mock)
			f.store.User = &mock
			f.gin.Request.Header.Set("Cookie", test.cookie)

			tpl := "{{ auth }}"
			runt(t, f, tpl, test.want)
		})
	}
}

func Test_Admin(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			mock := mocks.UserRepository{}
			f := newTestSuite()

			test.mock(&mock)
			f.store.User = &mock
			f.gin.Request.Header.Set("Cookie", test.cookie)

			tpl := "{{ admin }}"
			runt(t, f, tpl, test.want)
		})
	}
}
