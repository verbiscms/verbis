package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestValue_User() {

	tt := map[string]struct {
		value  domain.FieldValue
		mock   func(u *mocks.UserRepository)
		want   interface{}
	}{
		"User": {
			value: domain.FieldValue("1"),
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 1).Return(domain.User{
					UserPart: domain.UserPart{FirstName: "user"},
				}, nil)
			},
			want: domain.UserPart{FirstName: "user"},
		},
		"User Error": {
			value: domain.FieldValue("1"),
			mock: func(u *mocks.UserRepository) {
				u.On("GetById", 1).Return(domain.User{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock: func(u *mocks.UserRepository) {},
			want: `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			userMock := &mocks.UserRepository{}

			test.mock(userMock)
			v.store.User = userMock

			got, err := v.user(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}