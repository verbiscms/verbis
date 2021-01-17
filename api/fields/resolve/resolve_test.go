package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
)

func (t *ResolverTestSuite) Test_Field() {
	store := &models.Store{}
	field := domain.PostField{Id: 1, Type: "text", OriginalValue: "test"}

	got := Field(field, store)

	t.Equal(domain.PostField{Id: 1, Type: "text", OriginalValue: "test", Value: "test"}, got)
}

func (t *ResolverTestSuite) TestValue_Resolve() {

	tt := map[string]struct {
		field domain.PostField
		want  domain.PostField
	}{
		"Empty": {
			field: domain.PostField{OriginalValue: ""},
			want:  domain.PostField{OriginalValue: "", Value: ""},
		},
		"Not Iterable": {
			field: domain.PostField{OriginalValue: "999", Type: "number"},
			want:  domain.PostField{OriginalValue: "999", Type: "number", Value: int64(999)},
		},
		"Iterable": {
			field: domain.PostField{OriginalValue: "1,2,3,4,5", Type: "tags"},
			want:  domain.PostField{OriginalValue: "1,2,3,4,5", Type: "tags", Value: []interface{}{"1", "2", "3", "4", "5"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetValue().resolve(test.field))
		})
	}
}

func (t *ResolverTestSuite) TestValue_Execute() {

	tt := map[string]struct {
		value string
		typ   string
		want  interface{}
	}{
		"Not found": {
			value: "test",
			typ:   "wrongval",
			want:  "test",
		},
		"Found": {
			value: "999",
			typ:   "number",
			want:  int64(999),
		},
		"Error": {
			value: "wrongval",
			typ:   "number",
			want:  nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetValue().execute(test.value, test.typ))
		})
	}
}

//func (t *ResolverTestSuite) TestService_ResolveField() {
//
//	tt := map[string]struct {
//		field  domain.PostField
//		mock   func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)
//		want   interface{}
//		hasErr bool
//	}{
//		"Category Array": {
//			field: domain.PostField{
//				Id:            1,
//				Type:          "category",
//				OriginalValue: "1,2,3",
//			},
//			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
//				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil).Once()
//				c.On("GetById", 2).Return(domain.Category{Id: 2, Name: "cat"}, nil).Once()
//				c.On("GetById", 3).Return(domain.Category{Id: 3, Name: "cat"}, nil).Once()
//			},
//			want: domain.PostField{
//				Id:            1,
//				Type:          "category",
//				OriginalValue: "1,2,3",
//				Value: []interface{}{
//					domain.Category{Id: 1, Name: "cat"},
//					domain.Category{Id: 2, Name: "cat"},
//					domain.Category{Id: 3, Name: "cat"},
//				},
//			},
//			hasErr: false,
//		},
//		"Media Array": {
//			field: domain.PostField{
//				Id:            1,
//				Type:          "image",
//				OriginalValue: "1,2,3",
//			},
//			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
//				m.On("GetById", 1).Return(domain.Media{Id: 1, Url: "image"}, nil).Once()
//				m.On("GetById", 2).Return(domain.Media{Id: 2, Url: "image"}, nil).Once()
//				m.On("GetById", 3).Return(domain.Media{Id: 3, Url: "image"}, nil).Once()
//			},
//			want: domain.PostField{
//				Id:            1,
//				Type:          "image",
//				OriginalValue: "1,2,3",
//				Value: []interface{}{
//					domain.Media{Id: 1, Url: "image"},
//					domain.Media{Id: 2, Url: "image"},
//					domain.Media{Id: 3, Url: "image"},
//				},
//			},
//			hasErr: false,
//		},
//		"Post Array": {
//			field: domain.PostField{
//				Id:            1,
//				Type:          "post",
//				OriginalValue: "1,2,3",
//			},
//			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
//				p.On("GetById", 1).Return(domain.Post{Id: 1, Title: "post"}, nil).Once()
//				p.On("Format", domain.Post{Id: 1, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 1, Title: "post"}}, nil).Once()
//				p.On("GetById", 2).Return(domain.Post{Id: 2, Title: "post"}, nil).Once()
//				p.On("Format", domain.Post{Id: 2, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 2, Title: "post"}}, nil).Once()
//				p.On("GetById", 3).Return(domain.Post{Id: 3, Title: "post"}, nil).Once()
//				p.On("Format", domain.Post{Id: 3, Title: "post"}).Return(domain.PostData{Post: domain.Post{Id: 3, Title: "post"}}, nil).Once()
//			},
//			want: domain.PostField{
//				Id:            1,
//				Type:          "post",
//				OriginalValue: "1,2,3",
//				Value: []interface{}{
//					domain.PostData{Post: domain.Post{Id: 1, Title: "post"}},
//					domain.PostData{Post: domain.Post{Id: 2, Title: "post"}},
//					domain.PostData{Post: domain.Post{Id: 3, Title: "post"}},
//				},
//			},
//			hasErr: false,
//		},
//		"User Array": {
//			field: domain.PostField{
//				Id:            1,
//				Type:          "user",
//				OriginalValue: "1,2,3",
//			},
//			mock: func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository) {
//				u.On("GetById", 1).Return(domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "user"}}, nil).Once()
//				u.On("GetById", 2).Return(domain.User{UserPart: domain.UserPart{Id: 2, FirstName: "user"}}, nil).Once()
//				u.On("GetById", 3).Return(domain.User{UserPart: domain.UserPart{Id: 3, FirstName: "user"}}, nil).Once()
//			},
//			want: domain.PostField{
//				Id:            1,
//				Type:          "user",
//				OriginalValue: "1,2,3",
//				Value: []interface{}{
//					domain.UserPart{Id: 1, FirstName: "user"},
//					domain.UserPart{Id: 2, FirstName: "user"},
//					domain.UserPart{Id: 3, FirstName: "user"},
//				},
//			},
//			hasErr: false,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.GetTypeMockService(test.mock)
//
//			got := s.resolveField(test.field)
//			if test.hasErr {
//				t.Nil(t, got.Value)
//				return
//			}
//
//			t.Equal(test.want, got)
//		})
//	}
//}
