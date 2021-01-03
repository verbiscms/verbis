package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetLayout() {

	tt := map[string]struct {
		id     interface{}
		name   string
		layout []domain.FieldGroup
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			id:   1,
			name: "key3",
			layout: []domain.FieldGroup{
				{
					Title:  "test1",
					Fields: &[]domain.Field{{Name: "key1"}, {Name: "key2"}},
				},
				{
					Title:  "test2",
					Fields: &[]domain.Field{{Name: "key3"}, {Name: "key4"}},
				},
			},
			args: nil,
			want: domain.Field{Name: "key3"},
		},
		"Error": {
			id:     1,
			name:   "key3",
			layout: nil,
			args:   nil,
			want:   "no groups exist",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(nil)
			s.layout = test.layout

			got, err := s.GetLayout(test.name, test.args...)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

//func (t *FieldTestSuite) TestService_GetLayouts() {
//
//	tt := map[string]struct {
//		id   interface{}
//		name string
//		layout []domain.FieldGroup
//		args []interface{}
//		want interface{}
//	}{
//		"Success": {
//			id: 1,
//			name: "key3",
//			layout: []domain.FieldGroup{
//				{
//					Title:     "test1",
//					Fields:    &[]domain.Field{{Name: "key1"},{Name: "key2"}},
//				},
//				{
//					Title:     "test2",
//					Fields:    &[]domain.Field{{Name: "key3"},{Name: "key4"}},
//				},
//			},
//			args: nil,
//			want: domain.Field{Name: "key3"},
//		},
//		"Error": {
//			id: 1,
//			name: "key3",
//			layout: nil,
//			args: nil,
//			want: "no groups exist",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.GetService(nil)
//			s.layout = test.layout
//
//			got, err := s.GetLayout(test.name, test.args...)
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//
//			t.Equal(test.want, got)
//		})
//	}
//}

func (t *FieldTestSuite) TestService_GetLayoutsByPost() {

	post := domain.Post{Id: 1, Title: "post"}
	fg := &[]domain.FieldGroup{{Title: "test"}}

	tt := map[string]struct {
		id   interface{}
		mock func(p *mocks.PostsRepository)
		want []domain.FieldGroup
	}{
		"Success": {
			id: 1,
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(post, nil)
				p.On("Format", post).Return(domain.PostData{
					Post:   domain.Post{Id: 1, Title: "post"},
					Layout: fg,
				}, nil)
			},
			want: *fg,
		},
		"Cast Error": {
			id:   noStringer{},
			want: nil,
		},
		"Not Found": {
			id: 1,
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Format Error": {
			id: 1,
			mock: func(p *mocks.PostsRepository) {
				p.On("GetById", 1).Return(post, nil)
				p.On("Format", post).Return(domain.PostData{}, fmt.Errorf("error"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetPostsMockService(nil, test.mock).getLayoutByPost(test.id))
		})
	}
}
