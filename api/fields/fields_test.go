package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetFields() {

	tt := map[string]struct {
		fields []domain.PostField
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
			},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: 1,
		},
		"No Field": {
			fields: nil,
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "no field exists with the name: wrongval",
		},
		"Post": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
			},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return([]domain.PostField{{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: nil}}, nil)
			},
			args: []interface{}{2, true},
			want: 2,
		},
		"With Format": {
			fields: nil,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{{Id: 1, Type: "category", Name: "key1", Value: 1, Parent: nil}}, nil)
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			args: []interface{}{1, true},
			want: domain.Category{Id: 1, Name: "cat"},
		},
		"Without Format": {
			fields: nil,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{{Id: 1, Type: "category", Name: "key1", Value: 1, Parent: nil}}, nil)
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			args: []interface{}{1, false},
			want: 1,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got, err := s.GetFields(test.args...)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
