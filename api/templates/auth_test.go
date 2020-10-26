package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAuth(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, nil)

	assert.True(t, f.isAuth())
}

func TestIsAuth_NoCookie(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers

	assert.False(t, f.isAuth())
}

func TestIsAuth_NoUser(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))

	assert.False(t, f.isAuth())
}

func TestIsAdmin(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	user := domain.User{
		Id:     0,
		Role:   domain.UserRole{
			Id:          6,
		},
	}

	mockUsers.On("GetByToken", "token").Return(user, nil)

	assert.True(t, f.isAdmin())
}

func TestIsAdmin_NotAdmin(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	user := domain.User{
		Id:     0,
		Role:   domain.UserRole{
			Id: 1,
		},
	}

	mockUsers.On("GetByToken", "token").Return(user, nil)

	assert.False(t, f.isAdmin())
}

func TestIsAdmin_NoCookie(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers

	assert.False(t, f.isAdmin())
}

func TestIsAdmin_NoUser(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))

	assert.False(t, f.isAdmin())
}



