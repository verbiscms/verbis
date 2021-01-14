package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_ResolveField() {

	tt := map[string]struct {
		field  domain.PostField
		mock   func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)
		want   interface{}
		hasErr bool
	}{
		"Nil Value": {
			field: domain.PostField{
				Id:            1,
				Type:          "text",
				OriginalValue: "",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
			},
			want: domain.PostField{
				Id:    1,
				Type:  "text",
				Value: nil,
			},
			hasErr: false,
		},
		"Category Array": {
			field: domain.PostField{
				Id:            1,
				Type:          "category",
				OriginalValue: "1,2,3",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil).Once()
				c.On("GetById", 2).Return(domain.Category{Id: 2, Name: "cat"}, nil).Once()
				c.On("GetById", 3).Return(domain.Category{Id: 3, Name: "cat"}, nil).Once()
			},
			want: domain.PostField{
				Id:            1,
				Type:          "category",
				OriginalValue: "1,2,3",
				Value: []interface{}{
					domain.Category{Id: 1, Name: "cat"},
					domain.Category{Id: 2, Name: "cat"},
					domain.Category{Id: 3, Name: "cat"},
				},
			},
			hasErr: false,
		},
		"Media Array": {
			field: domain.PostField{
				Id:            1,
				Type:          "image",
				OriginalValue: "1,2,3",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.Media{Id: 1, Url: "image"}, nil).Once()
				m.On("GetById", 2).Return(domain.Media{Id: 2, Url: "image"}, nil).Once()
				m.On("GetById", 3).Return(domain.Media{Id: 3, Url: "image"}, nil).Once()
			},
			want: domain.PostField{
				Id:            1,
				Type:          "image",
				OriginalValue: "1,2,3",
				Value: []interface{}{
					domain.Media{Id: 1, Url: "image"},
					domain.Media{Id: 2, Url: "image"},
					domain.Media{Id: 3, Url: "image"},
				},
			},
			hasErr: false,
		},
		"Post Array": {
			field: domain.PostField{
				Id:            1,
				Type:          "post",
				OriginalValue: "1,2,3",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				p.On("GetById", 1).Return(domain.Post{Id: 1, Title: "post"}, nil).Once()
				p.On("Format", domain.Post{Id: 1, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 1, Title: "post"}}, nil).Once()
				p.On("GetById", 2).Return(domain.Post{Id: 2, Title: "post"}, nil).Once()
				p.On("Format", domain.Post{Id: 2, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 2, Title: "post"}}, nil).Once()
				p.On("GetById", 3).Return(domain.Post{Id: 3, Title: "post"}, nil).Once()
				p.On("Format", domain.Post{Id: 3, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 3, Title: "post"}}, nil).Once()
			},
			want: domain.PostField{
				Id:            1,
				Type:          "post",
				OriginalValue: "1,2,3",
				Value: []interface{}{
					domain.PostData{Post: domain.Post{Id: 1, Title: "post"}},
					domain.PostData{Post: domain.Post{Id: 2, Title: "post"}},
					domain.PostData{Post: domain.Post{Id: 3, Title: "post"}},
				},
			},
			hasErr: false,
		},
		"User Array": {
			field: domain.PostField{
				Id:            1,
				Type:          "user",
				OriginalValue: "1,2,3",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				u.On("GetById", 1).Return(domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "user"}}, nil).Once()
				u.On("GetById", 2).Return(domain.User{UserPart: domain.UserPart{Id: 2, FirstName: "user"}}, nil).Once()
				u.On("GetById", 3).Return(domain.User{UserPart: domain.UserPart{Id: 3, FirstName: "user"}}, nil).Once()
			},
			want: domain.PostField{
				Id:            1,
				Type:          "user",
				OriginalValue: "1,2,3",
				Value: []interface{}{
					domain.UserPart{Id: 1, FirstName: "user"},
					domain.UserPart{Id: 2, FirstName: "user"},
					domain.UserPart{Id: 3, FirstName: "user"},
				},
			},
			hasErr: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetTypeMockService(test.mock)

			got := s.resolveField(test.field)
			if test.hasErr {
				t.Nil(t, got.Value)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_ResolveValue() {

	tt := map[string]struct {
		field  domain.PostField
		mock   func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)
		want   interface{}
		hasErr bool
	}{
		"Category": {
			field: domain.PostField{
				Id:            1,
				Type:          "category",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				c.On("GetById", 1).Return(domain.Category{Name: "cat"}, nil)
			},
			want: domain.PostField{
				Id:            1,
				Type:          "category",
				Value:         domain.Category{Name: "cat"},
				OriginalValue: "1",
			},
			hasErr: false,
		},
		"Category Error": {
			field: domain.PostField{
				Id:            1,
				Type:          "category",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				c.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("not found"))
			},
			hasErr: true,
		},
		"Image": {
			field: domain.PostField{
				Id:            1,
				Type:          "image",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.Media{Url: "image"}, nil)
			},
			want: domain.PostField{
				Id:            1,
				Type:          "image",
				Value:         domain.Media{Url: "image"},
				OriginalValue: "1",
			},
			hasErr: false,
		},
		"Image Error": {
			field: domain.PostField{
				Id:            1,
				Type:          "image",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("not found"))
			},
			hasErr: true,
		},
		"Post": {
			field: domain.PostField{
				Id:            1,
				Type:          "post",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				p.On("GetById", 1).Return(domain.Post{Title: "post"}, nil)
				p.On("Format", domain.Post{Title: "post"}).Return(domain.PostData{
					Post: domain.Post{Title: "post"},
				}, nil)
			},
			want: domain.PostField{
				Id:   1,
				Type: "post",
				Value: domain.PostData{
					Post: domain.Post{Title: "post"},
				},
				OriginalValue: "1",
			},
			hasErr: false,
		},
		"Post Error": {
			field: domain.PostField{
				Id:            1,
				Type:          "post",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				p.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("not found"))
			},
			hasErr: true,
		},
		"Post Format Error": {
			field: domain.PostField{
				Id:            1,
				Type:          "post",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				p.On("GetById", 1).Return(domain.Post{Title: "post"}, nil)
				p.On("Format", domain.Post{Title: "post"}).Return(domain.PostData{}, fmt.Errorf("format error"))
			},
			hasErr: true,
		},
		"User": {
			field: domain.PostField{
				Id:            1,
				Type:          "user",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				u.On("GetById", 1).Return(domain.User{
					UserPart: domain.UserPart{FirstName: "user"},
				}, nil)
			},
			want: domain.PostField{
				Id:            1,
				Type:          "user",
				Value:         domain.UserPart{FirstName: "user"},
				OriginalValue: "1",
			},
			hasErr: false,
		},
		"User Error": {
			field: domain.PostField{
				Id:            1,
				Type:          "user",
				OriginalValue: "1",
			},
			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
				u.On("GetById", 1).Return(domain.User{}, fmt.Errorf("not found"))
			},
			hasErr: true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetTypeMockService(test.mock)

			got := s.resolveField(test.field)
			if test.hasErr {
				t.Nil(got.Value)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
