package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.PostsRepository) {
	mock := &mocks.PostsRepository{}
	return &Namespace{deps: &deps.Deps{
		Store: &models.Store{
			Posts: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {

	post := domain.Post{
		Id:     1,
		Title:  "test title",
		UserId: 1,
	}

	viewData := TplPost{
		Author:    domain.UserPart{},
		Category: &domain.Category{},
		Post:     post,
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
			},
			want: viewData,
		},
		"Format Error": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
			},
			want: nil,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.Find(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
//
//func (t *TplTestSuite) Test_GetPosts() {
//
//	post := domain.Post{Id: 1, Title: "Title"}
//	posts := []domain.Post{
//		post, post,
//	}
//
//	author := &domain.PostAuthor{}
//	category := &domain.PostCategory{}
//	viewData := []ViewPost{
//		{
//			Author:   author,
//			Category: category,
//			Post:     post,
//		},
//		{
//			Author:   author,
//			Category: category,
//			Post:     post,
//		},
//	}
//
//	categoryTest := &domain.PostCategory{
//		Name: "cat",
//	}
//	viewDataCategory := []ViewPost{
//		{
//			Author:   author,
//			Category: categoryTest,
//			Post:     post,
//		},
//		{
//			Author:   author,
//			Category: categoryTest,
//			Post:     post,
//		},
//	}
//
//	tt := map[string]struct {
//		input map[string]interface{}
//		mock  func(m *mocks.PostsRepository)
//		want  interface{}
//	}{
//		"Success": {
//			input: map[string]interface{}{"limit": 15},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
//			},
//			want: map[string]interface{}{
//				"Posts": viewData,
//				"Pagination": &vhttp.Pagination{
//					Page:  1,
//					Pages: 1,
//					Limit: 15,
//					Total: 5,
//					Next:  false,
//					Prev:  false,
//				},
//			},
//		},
//		"Failed Params": {
//			input: map[string]interface{}{"limit": "wrongval"},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
//			},
//			want: "cannot unmarshal string into Go struct field TemplateParams.limit",
//		},
//		"Not Found": {
//			input: map[string]interface{}{"limit": 15},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
//			},
//			want: map[string]interface{}(nil),
//		},
//		"Internal Error": {
//			input: map[string]interface{}{"limit": 15},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
//			},
//			want: "internal error",
//		},
//		"Format Error": {
//			input: map[string]interface{}{"limit": 15},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, fmt.Errorf("error"))
//			},
//			want: map[string]interface{}{
//				"Posts": []ViewPost(nil),
//				"Pagination": &vhttp.Pagination{
//					Page:  1,
//					Pages: 1,
//					Limit: 15,
//					Total: 5,
//					Next:  false,
//					Prev:  false,
//				},
//			},
//		},
//		"Category": {
//			input: map[string]interface{}{"limit": 15, "category": "cat"},
//			mock: func(m *mocks.PostsRepository) {
//				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 2, nil)
//				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: categoryTest}, nil)
//			},
//			want: map[string]interface{}{
//				"Posts": viewDataCategory,
//				"Pagination": &vhttp.Pagination{
//					Page:  1,
//					Pages: 1,
//					Limit: 15,
//					Total: 2,
//					Next:  false,
//					Prev:  false,
//				},
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			postsMock := mocks.PostsRepository{}
//
//			test.mock(&postsMock)
//			t.store.Posts = &postsMock
//
//			p, err := t.getPosts(test.input)
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//
//			t.EqualValues(test.want, p)
//		})
//	}
//}

