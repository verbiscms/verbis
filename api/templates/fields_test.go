package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type FieldsTestSuite struct {
	TemplateManager *TemplateManager
}

var (
	postAuthor   = domain.PostAuthor{Id: 1, FirstName: "verbis"}
	postCategory = domain.PostCategory{Id: 1, Name: "category"}
	category     = domain.Category{Id: 1, Name: "category"}
	media        = domain.Media{Id: 1, Url: "/media"}
	post         = domain.Post{Id: 1, Title: "post title"}
	user         = domain.User{Id: 1, FirstName: "verbis"}
	viewP        = ViewPost{Author: &postAuthor, Category: &postCategory, Post: post}
)

func newFieldTestSuite(fields string) *TemplateManager {

	f := newTestSuite(fields)

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

	return f
}

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

	tt := map[string]struct {
		fields string
		input  string
		want   interface{}
	}{
		//"Success": {
		//	fields: `{
		//		"repeater":[
		//			{
		//				"text1":"content",
		//				"text2":"content"
		//			},
		//			{
		//				 "text1":"content",
		//				 "text2":"content"
		//			}
		//		]
		//	}`,
		//	input: `repeater`,
		//	want: []map[string]interface{}{
		//		{"text1": "content", "text2": "content"},
		//		{"text1": "content", "text2": "content"},
		//	},
		//},
		//"Wrong Key": {
		//	fields: `{
		//		"repeater":[
		//			{
		//				"text1":"content",
		//				"text2":"content"
		//			},
		//			{
		//				 "text1":"content",
		//				 "text2":"content"
		//			}
		//		]
		//	}`,
		//	input: `wrongval`,
		//	want:  make([]map[string]interface{}, 0),
		//},
		//"With Types": {
		//	fields: `{
		//		"repeater":[
		//			{
		//				"category": [
		//					{
		//						"id": 1,
		//						"type":"category"
		//					}
		//				],
		//				"image": {
		//					"id": 1,
		//					"type": "image"
		//				},
		//				"post": [
		//					{
		//						"id": 1,
		//						"type": "post"
		//					}
		//				],
		//				"user": [
		//					{
		//						"id": 1,
		//						"type": "user"
		//					}
		//				]
		//			}
		//		]
		//	}`,
		//	input: `repeater`,
		//	want: []map[string]interface{}{
		//		{
		//			"category": &category,
		//			"image":    &media,
		//			"post":     &viewP,
		//			"user":     &user,
		//		},
		//	},
		//},
		//"Nested Slices": {
		//	fields: `{
		//		"repeater":[
		//			{
		//				"users": [
		//					{
		//						"id": 1,
		//						"type": "user"
		//					},
		//					{
		//						"id": 1,
		//						"type": "category"
		//					}
		//				]
		//			}
		//		]
		//	}`,
		//	input: `repeater`,
		//	want: []map[string]interface{}{
		//		{
		//			"users": []interface{}{&user, &category},
		//		},
		//	},
		//},
		"Nested Slices 2": {
			fields: `{
				"repeater":[
					{
						"ext": "text",
						"image": {
							"id": 1,
							"type": "image"
						},
						"users": [
							{
								"id": 1,
								"type": "user"
							},
							{
								"id": 1,
								"type": "user"
							}
						],
						"nested":[
							{
								"nestedtext": "text",
								"nestedimage": {
									"id": 1,
									"type": "image"
								},
								"nestedusers": [
									{
										"id": 1,
										"type": "user"
									}
								]
							}
						]
					}
				]
			}`,
			input: `repeater`,
			want: []map[string]interface{}{
				{
					"users": []interface{}{&user},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {

			f := newFieldTestSuite(test.fields)

			layout, err := ioutil.ReadFile("/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata/fields/layout.json")
			assert.NoError(t, err)

			var fields domain.FieldGroup
			err = json.Unmarshal(layout, &fields)
			assert.NoError(t, err)

			fg := make([]domain.FieldGroup, 1)
			fg[0] = fields
			f.post.Layout = &fg


			repeater, err := f.getRepeater(test.input)

			color.Yellow.Println(repeater)

			tem, err := execute(f, `{{ $test := repeater "repeater" }}
				{{ $nest := index $test 0 }}
				{{ repeater $nest.nested }}`, nil)
			fmt.Println(err)
			fmt.Println(tem)
		})
	}
}

func Test_GetFlexible(t *testing.T) {

//fields: `{
//				"flexible": [
//					{
//						 "block": "block1",
//						 "fields": {
//							"text": "content",
//							"text2": "content"
//						 }
//					},
//					{
//						"block": "block2",
//						"fields": {
//							"text": "content",
//							"text1": "content"
//						}
//					}
//				]
//			}`,

	tt := map[string]struct {
		fields string
		input  string
		want   interface{}
	}{
		"Simple": {
			fields: `{
				"flexible": [
					{
						 "block": "block",
						 "fields": {
							"text": "content",
							"text2": "content"
						 }
					}
				]
			}`,
			input: `flexible`,
			want: []map[string]interface{}{
				{
					"block": "block",
					"fields": map[string]interface{}{"text": "content", "text2": "content"},
				},
			},
		},
		"Category": {
			fields: `{
				"flexible": [
					{
						 "block": "block",
						 "fields": {
							"text": "content",
							"category": [
								{
									"id": 1,
									"type":"category"
								}
							]
						 }
					}
				]
			}`,
			input: `flexible`,
			want: []map[string]interface{}{
				{
					"block": "block",
					"fields": map[string]interface{}{"category": &category, "text2": "content"},
				},
			},
		},

	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, newFieldTestSuite(test.fields).getFlexible(test.input))
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

	tt := map[string]struct {
		fields string
		want   interface{}
	}{
		"Nil": {
			fields: `{
				"user": {
					"id": 1,
					"type":"wrongval"
				}
			}`,
			want: nil,
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
			want: &category,
		},
		"Image": {
			fields: `{
				"image": {
					"id": 1,
					"type":"image"
				}
			}`,
			want: &media,
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
			want: &viewP,
		},
		"User": {
			fields: `{
				"user": {
					"id": 1,
					"type":"user"
				}
			}`,
			want: &user,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newFieldTestSuite("{}")
			var d interface{}
			err := json.Unmarshal([]byte(test.fields), &d)
			assert.NoError(t, err)
			assert.Equal(t, test.want, f.checkFieldType(d))
		})
	}
}
