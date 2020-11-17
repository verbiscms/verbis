package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func TestIsAuth(t *testing.T) {
	mockUsers := mocks.UserRepository{}
	f := newTestSuite()

	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, nil)

	tpl := "{{ isAuth }}"
	runt(t, f, tpl, true)
}

func TestIsAuth_NoCookie(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers

	tpl := "{{ isAuth }}"
	runt(t, f, tpl, false)
}

func TestIsAuth_NoUser(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))

	tpl := "{{ isAuth }}"
	runt(t, f, tpl, false)
}

func TestIsAdmin(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	user := domain.User{
		Id: 0,
		Role: domain.UserRole{
			Id: 6,
		},
	}

	mockUsers.On("GetByToken", "token").Return(user, nil)

	tpl := "{{ isAuth }}"
	runt(t, f, tpl, true)
}

func TestIsAdmin_NotAdmin(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	user := domain.User{
		Id: 0,
		Role: domain.UserRole{
			Id: 1,
		},
	}

	mockUsers.On("GetByToken", "token").Return(user, nil)

	tpl := "{{ isAdmin }}"
	runt(t, f, tpl, false)
}

func TestIsAdmin_NoCookie(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers

	tpl := "{{ isAdmin }}"
	runt(t, f, tpl, false)
}

func TestIsAdmin_NoUser(t *testing.T) {
	mockUsers := mocks.UserRepository{}

	f := newTestSuite()
	f.store.User = &mockUsers
	f.gin.Request.Header.Set("Cookie", "verbis-session=token")

	mockUsers.On("GetByToken", "token").Return(domain.User{}, fmt.Errorf("error"))

	tpl := "{{ isAdmin }}"
	runt(t, f, tpl, false)
}
