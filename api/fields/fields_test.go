package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/google/uuid"
)

func (t *FieldTestSuite) TestService_GetFields() {

	uniq := uuid.New()
	layout := "layout"

	tt := map[string]struct {
		fields []domain.PostField
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
	}{
		"Simple": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: nil},
				{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: nil},
			},
			args: nil,
			want: Fields{
				"key1": 1,
				"key2": 2,
				"key3": 3,
			},
		},
		"Repeater": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "repeater", UUID: uniq, Name: "key2", Value: nil, Parent: nil},
				{Id: 3, Type: "text", Name: "key3", Value: 2, Parent: &uniq},
				{Id: 4, Type: "text", Name: "key4", Value: 3, Parent: &uniq},
				{Id: 5, Type: "text", Name: "key5", Value: 4, Parent: &uniq},
			},
			args: nil,
			want: Fields{
				"key1": 1,
				"key2": []interface{}{
					domain.PostField{Id: 3, Type: "text", Name: "key3", Value: 2, Parent: &uniq},
					domain.PostField{Id: 4, Type: "text", Name: "key4", Value: 3, Parent: &uniq},
					domain.PostField{Id: 5, Type: "text", Name: "key5", Value: 4, Parent: &uniq},
				},
			},
		},
		"Flexible Content": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "flexible", Name: "key2", Value: []string{"layout"}, Parent: nil, Layout: nil, Index: 0},
				{Id: 3, Type: "text", Name: "text 1", Value: "text", Layout: &layout, Index: 0},
				{Id: 4, Type: "text", Name: "text 2", Value: "text", Layout: &layout, Index: 0},
			},
			args: nil,
			want: Fields{
				"key1": 1,
				"key2": []interface{}{
					Layout{
						Name:      "layout",
						SubFields: SubFields{
							domain.PostField{Id: 3, Type: "text", Name: "text 1", Value: "text", Layout: &layout, Index: 0},
							domain.PostField{Id: 4, Type: "text", Name: "text 2", Value: "text", Layout: &layout, Index: 0},
						},
					},
				},
			},
		},
		"Format": {
			fields: nil,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "category", Name: "key1", Value: 1, Parent: nil},
				}, nil)
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			args: []interface{}{1, true},
			want: Fields{
				"key1": domain.Category{Id: 1, Name: "cat"},
			},
		},
		"Without Format": {
			fields: nil,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "category", Name: "key1", Value: 1, Parent: nil},
				}, nil)
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			args: []interface{}{1, false},
			want: Fields{
				"key1": 1,
			},
		},
		"Category Array": {
			fields: []domain.PostField{
				{Id: 1, Type: "category", Name: "key1", Value: []int{1, 2}, Parent: nil},
			},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil).Once()
				c.On("GetById", 2).Return(domain.Category{Id: 2, Name: "cat"}, nil).Once()
			},
			args: nil,
			want: Fields{
				"key1": []interface{}{
					domain.Category{Id: 1, Name: "cat"},
					domain.Category{Id: 2, Name: "cat"},
				},
			},
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

func (t *FieldTestSuite) TestService_Mapper() {

	uniq := uuid.New()
	layout := "layout"

	tt := map[string]struct {
		fields []domain.PostField
		want    interface{}
	}{
		"Simple": {
			fields: []domain.PostField{{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil}},
			want: domain.PostField{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: nil},
		},
		"Parent": {
			fields: []domain.PostField{{Id: 1, Type: "text", Name: "key1", Value: 1, Parent: &uniq}},
			want: nil,
		},
		"Layout": {
			fields: []domain.PostField{{Id: 1, Type: "text", Name: "key1", Value: 1, Layout: &layout}},
			want: nil,
		},
		"Repeater": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", UUID: uniq, Name: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq},
				{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq},
				{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq},
			},
			want: domain.PostField{
				Id: 1, Type: "repeater", UUID: uniq, Name: "key1", Parent: nil,
				Value:  Repeater{
					{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq},
					{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq},
					{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq},
				},
			},
		},
		"Flexible": {
			fields: []domain.PostField{
				{Id: 1, Type: "flexible", Name: "key1", Value: []string{"layout"}, Parent: nil, Layout: nil, Index: 0},
				{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &layout, Index: 0},
				{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &layout, Index: 0},
			},
			want: domain.PostField{
				Id: 1, Type: "flexible", Name: "key1", Parent: nil, Layout: nil, Index: 0,
				Value:  Flexible{
					Layout{
						Name:      "layout",
						SubFields: SubFields{
							domain.PostField{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &layout, Index: 0},
							domain.PostField{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &layout, Index: 0},
						},
					},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			s.mapper(test.fields, func(field domain.PostField) {
				t.Equal(test.want, field)
			})
		})
	}
}

