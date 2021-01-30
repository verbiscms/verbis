package auth

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Setup(cookie string) (*Namespace, *mocks.UserRepository) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(rr)
	g.Request, _ = http.NewRequest("GET", "/get", nil)
	g.Request.Header.Set("Cookie", cookie)

	mock := &mocks.UserRepository{}
	return &Namespace{
		deps: &deps.Deps{
			Store: &models.Store{
				User: mock,
			},
		},
		ctx: g,
	}, mock
}

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
			ns, mock := Setup(test.cookie)
			test.mock(mock)
			got := ns.Auth()
			assert.Equal(t, test.want, got)
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
			ns, mock := Setup(test.cookie)
			test.mock(mock)
			got := ns.Admin()
			assert.Equal(t, test.want, got)
		})
	}
}
