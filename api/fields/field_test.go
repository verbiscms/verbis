package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetField() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key:  "key1",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: "test",
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "no field exists with the name: wrongval",
		},
		"Post": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1"},
			},
			key: "key2",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return([]domain.PostField{{Id: 2, Type: "text", Name: "key2", Value: "test"}}, nil)
			},
			args: []interface{}{2},
			want: "test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got, err := s.GetField(test.key, test.args...)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_GetFieldObject() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key:  "key1",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: domain.PostField{Id: 1, Type: "text", Name: "key1", OriginalValue: "test", Value: "test"},
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "no field exists with the name: wrongval",
		},
		"Post": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1"},
			},
			key: "key2",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return([]domain.PostField{{Id: 2, Type: "text", Name: "key2", OriginalValue: "test"}}, nil)
			},
			args: []interface{}{2},
			want: domain.PostField{Id: 2, Type: "text", Name: "key2", OriginalValue: "test", Value: "test"},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got, err := s.GetFieldObject(test.key, test.args...)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
