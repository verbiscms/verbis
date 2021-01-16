package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestValue_Post() {

	tt := map[string]struct {
		value  domain.FieldValue
		mock   func(p *mocks.PostsRepository)
		want   interface{}
	}{
		"Post": {
			value: domain.FieldValue("1"),
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(domain.Post{Title: "post"}, nil)
				p.On("Format", domain.Post{Title: "post"}).Return(domain.PostData{
					Post: domain.Post{Title: "post"},
				}, nil)
			},
			want: domain.PostData{
				Post: domain.Post{Title: "post"},
			},
		},
		"Post Error": {
			value: domain.FieldValue("1"),
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Post Format Error": {
			value: domain.FieldValue("1"),
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(domain.Post{Title: "post"}, nil)
				p.On("Format", domain.Post{Title: "post"}).Return(domain.PostData{}, fmt.Errorf("format error"))
			},
			want: "format error",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock: func(p *mocks.PostsRepository) {},
			want: `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			postMock := &mocks.PostsRepository{}

			test.mock(postMock)
			v.store.Posts = postMock

			got, err := v.post(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
