package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/google/uuid"
)

func (t *FieldTestSuite) TestService_HandleArgs() {

	tt := map[string]struct {
		fields []domain.PostField
		args   []interface{}
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		format bool
		want   []domain.PostField
	}{
		"Default": {
			fields: []domain.PostField{{Name: "test"}},
			args:   nil,
			format: true,
			want:   []domain.PostField{{Name: "test"}},
		},
		"1 Args (Post)": {
			fields: nil,
			args:   []interface{}{1},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			format: true,
			want:   []domain.PostField{{Id: 1, Type: "text", Name: "post"}},
		},
		"1 Args (Post Error)": {
			fields: []domain.PostField{{Name: "test"}},
			args:   []interface{}{1},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(nil, fmt.Errorf("error"))
			},
			format: true,
			want:   nil,
		},
		"2 Args (Post & Format)": {
			fields: nil,
			args:   []interface{}{1, false},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			format: false,
			want:   []domain.PostField{{Id: 1, Type: "text", Name: "post"}},
		},
		"Cast to Bool Error": {
			fields: nil,
			args:   []interface{}{1, noStringer{}},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			format: true,
			want:   []domain.PostField{{Id: 1, Type: "text", Name: "post"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)
			got, format := s.handleArgs(test.args)
			t.Equal(test.format, format)
			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_GetFieldsByPost() {

	tt := map[string]struct {
		id   interface{}
		mock func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		want []domain.PostField
	}{
		"Success": {
			id: 1,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return([]domain.PostField{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			want: []domain.PostField{{Id: 1, Type: "text", Name: "post"}},
		},
		"Cast Error": {
			id:   noStringer{},
			want: nil,
		},
		"Get Error": {
			id: 1,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(nil, fmt.Errorf("error"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetMockService(nil, test.mock).getFieldsByPost(test.id))
		})
	}
}

func (t *FieldTestSuite) TestService_GetChildren() {

	id := uuid.New()

	tt := map[string]struct {
		uuid   uuid.UUID
		fields []domain.PostField
		format bool
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		want   []domain.PostField
	}{
		"Not Formatted": {
			uuid: id,
			fields: []domain.PostField{
				{Id: 1, Parent: nil},
				{Id: 2, Parent: &id},
				{Id: 3, Parent: &id},
				{Id: 4, Parent: &id},
			},
			format: false,
			want: []domain.PostField{
				{Id: 2, Parent: &id},
				{Id: 3, Parent: &id},
				{Id: 4, Parent: &id},
			},
		},
		"Formatted": {
			uuid: id,
			fields: []domain.PostField{
				{Id: 1, Parent: nil},
				{Id: 2, Parent: &id, Value: 1, Type: "category"},
			},
			format: true,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			want: []domain.PostField{
				{Id: 2, Parent: &id, Value: domain.Category{Id: 1, Name: "cat"}, Type: "category"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(nil, test.mock)
			t.Equal(test.want, s.getFieldChildren(test.uuid, test.fields, test.format))
		})
	}
}

func (t *FieldTestSuite) TestService_SortFields() {

	tt := map[string]struct {
		fields []domain.PostField
		want   []domain.PostField
	}{
		"Simple": {
			fields: []domain.PostField{
				{Id: 1, Index: 3},
				{Id: 2, Index: 2},
				{Id: 3, Index: 1},
			},
			want: []domain.PostField{
				{Id: 3, Index: 1},
				{Id: 2, Index: 2},
				{Id: 1, Index: 3},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetService(nil).sortFields(test.want))
		})
	}
}
