package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func Test_GetField(t *testing.T) {

	tt := map[string]struct {
		fields string
		tpl    string
		want   string
	}{
		"Basic": {
			fields: `{"text": "content"}`,
			tpl:    `{{ field "text" }}`,
			want:   "content",
		},
		"Empty Field": {
			fields: `{"text": "content"}`,
			tpl:    `{{ field "wrongval" }}`,
			want:   "",
		},
		"Invalid JSON": {
			fields: `{}`,
			tpl:    `{{ field "text" }}`,
			want:   "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runt(t, newTestSuite(test.fields), test.tpl, test.want)
		})
	}
}

func Test_GetFieldByPostId(t *testing.T) {

	tt := map[string]struct {
		fields string
		tpl    string
		post   domain.Post
		mock   func(m *mocks.PostsRepository, post domain.Post)
		want   string
	}{
		"Success": {
			fields: `{"text": "content"}`,
			tpl:    `{{ field "text" 1 }}`,
			post: domain.Post{
				Id: 1,
				Fields: map[string]interface{}{
					"text": "postcontent",
				},
			},
			mock: func(m *mocks.PostsRepository, post domain.Post) {
				m.On("GetById", 1).Return(post, nil)
			},
			want: "postcontent",
		},
		"No Field": {
			fields: `{"text": "content"}`,
			tpl:    `{{ field "wrongval" 1 }}`,
			post: domain.Post{
				Id: 1,
				Fields: map[string]interface{}{
					"text": "postcontent",
				},
			},
			mock: func(m *mocks.PostsRepository, post domain.Post) {
				m.On("GetById", 1).Return(post, nil)
			},
			want: "",
		},
		"No Post": {
			fields: `{"text": "content"}`,
			tpl:    `{{ field "text" 1 }}`,
			post:   domain.Post{},
			mock: func(m *mocks.PostsRepository, post domain.Post) {
				m.On("GetById", 1).Return(post, fmt.Errorf("no post"))
			},
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.fields)

			mockPosts := mocks.PostsRepository{}
			test.mock(&mockPosts, test.post)

			f.store.Posts = &mockPosts

			runt(t, f, test.tpl, test.want)
		})
	}
}

func Test_HasField(t *testing.T) {

	str := `{"text": "content"}`

	tt := map[string]struct {
		tpl   string
		input string
		want  interface{}
	}{
		"Success": {
			tpl:   `{{ hasField "text" }}`,
			input: str,
			want:  true,
		},
		"Wrong Key": {
			tpl:   `{{ hasField "wrongval" }}`,
			input: str,
			want:  false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.input)
			runt(t, f, test.tpl, test.want)
		})
	}
}

func Test_GetRepeater(t *testing.T) {

	str := `{
		"repeater":[
			{
				"text1":"content",
				"text2":"content"
			},
			{
				 "text1":"content",
				 "text2":"content"
			}
		]
	}`

	tt := map[string]struct {
		tpl   string
		input string
		want  string
	}{
		"Success": {
			tpl:   `{{ repeater "repeater" }}`,
			input: str,
			want:  "[map[text1:content text2:content] map[text1:content text2:content]]",
		},
		"Wrong Key": {
			tpl:   `{{ repeater "wrongval" }}`,
			input: str,
			want:  `[]`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.input)
			runt(t, f, test.tpl, test.want)
		})
	}
}

func Test_GetFlexible(t *testing.T) {

	str := `{
		"flexible": [
			{
				 "type": "block1",
				 "fields": {
					"text": "content",
					"text2": "content"
				 }
			},
			{
				"type": "block2",
				"fields": {
					"text": "content",
					"text1": "content",
					"text2": "content",
					"repeater": [
						{
						  "text":"content",
						  "text2":"content"
						}
					]
				}
			}
      	]
   	}`

	tt := map[string]struct {
		tpl   string
		input string
		want  string
	}{
		"Success": {
			tpl:   `{{ flexible "flexible" }}`,
			input: str,
			want:  `[map[fields:map[text:content text2:content] type:block1] map[fields:map[repeater:[map[text:content text2:content]] text:content text1:content text2:content] type:block2]]`,
		},
		"Wrong Key": {
			tpl:   `{{ flexible "wrongval" }}`,
			input: str,
			want:  `[]`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.input)
			runt(t, f, test.tpl, test.want)
		})
	}
}

func Test_GetSubField(t *testing.T) {

	str := `{
		"repeater":[
			{
				"text1":"content",
				"text2":"content"
			},
			{
				 "text1":"content",
				 "text2":"content"
			}
		]
	}`

	tt := map[string]struct {
		tpl   string
		input string
		want  string
	}{
		"Success": {
			tpl:   `{{ repeater "repeater" }}`,
			input: str,
			want:  "[map[text1:content text2:content] map[text1:content text2:content]]",
		},
		"Wrong Key": {
			tpl:   `{{ repeater "wrongval" }}`,
			input: str,
			want:  `[]`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.input)
			runt(t, f, test.tpl, test.want)
		})
	}
}

func Test_CheckField(t *testing.T) {

	postCategory := domain.PostCategory{Id:          1, Name:        "category"}
	postAuthor := domain.PostAuthor{Id:               1, FirstName:        "verbis"}

	category := domain.Category{Id: 1, Name: "category"}
	media := domain.Media{Id: 1, Url: "/media"}
	post := domain.Post{Id: 1, Title: "post title"}
	viewPost := ViewPost{Author:   &postAuthor, Category: &postCategory, Post:     post,}
	user := domain.User{Id: 1, FirstName: "verbis"}

	tt := map[string]struct {
		fields string
		tpl    string
		want   interface{}
	}{
		"Nil": {
			fields: `{
				"user": {
					"id": 1,
					"type":"wrongval"
				}
			}`,
			tpl:  `{{ field "user" }}`,
			want: "",
		},
		"Category": {
			fields: `{
				"category":[
					{
						"id": 1,
						"type":"category"
					}
				]
			}`,
			tpl:  `{{ field "category" }}`,
			want: category,
		},
		"Image": {
			fields: `{
				"image": {
					"id": 1,
					"type":"image"
				}
			}`,
			tpl:  `{{ field "image" }}`,
			want: media,
		},
		"Post": {
			fields: `{
				"post":[
					{
						"id": 1,
						"type":"post"
					}
				]
			}`,
			tpl:  `{{ field "post" }}`,
			want: viewPost,
		},
		"Users": {
			fields: `{
				"repeater":[
					{
						"users": [
							{
								"id": 1,
								"type": "user"
							},
							{
								"id": 1,
								"type": "user"
							}
						]
					}
				]
			}`,
			tpl:  `{{ repeater "repeater" }}`,
			want: []map[string]interface{}{
				{
					"users": []domain.User{
						user,
						user,
					},
				},
			},
		},
		"User": {
			fields: `{
				"user": {
					"id": 1,
					"type":"user"
				}
			}`,
			tpl:  `{{ field "user" }}`,
			want: user,
		},
		"Repeater": {
			fields: `{
				"repeater":[
					{
						"category": [
							{
								"id": 1,
								"type":"category"
							}
						],
						"image": {
							"id": 1,
							"type": "image"
						},
						"post": [
							{
								"id": 1,
								"type": "post"
							}
						],
						"user": [
							{
								"id": 1,
								"type": "user"
							}
						]
					}
				]
			}`,
			tpl:  `{{ repeater "repeater" }}`,
			want: []map[string]interface{}{
				{
					"category": category,
					"image":    media,
					"post":     viewPost,
					"user":     user,
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite(test.fields)

			var m interface{}
			err := json.Unmarshal([]byte(test.fields), &m)
			if err != nil {
				t.Error(err)
			}

			categoryMock := &mocks.CategoryRepository{}
			mediaMock := &mocks.MediaRepository{}
			postMock := &mocks.PostsRepository{}
			userMock := &mocks.UserRepository{}

			categoryMock.On("GetById", 1).Return(category, nil)

			mediaMock.On("GetById", 1).Return(media, nil)

			postMock.On("GetById", 1).Return(post, nil)
			postMock.On("Format", post).Return(domain.PostData{Post: post, Author: &postAuthor, Category: &postCategory}, nil).Once()

			userMock.On("GetById", 1).Return(user, nil).Times(10)

			f.store.Categories = categoryMock
			f.store.Media = mediaMock
			f.store.Posts = postMock
			f.store.User = userMock

			runt(t, f, test.tpl, test.want)
		})
	}
}
