package fields

//func (t *FieldTestSuite) TestService_GetFlexible() {
//
//	l1 := "layout1"
//	l2 := "layout2"
//	l3 := "layout3"
//
//	tt := map[string]struct {
//		fields []domain.PostField
//		key    string
//		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
//		args   []interface{}
//		want   interface{}
//	}{
//		"Simple": {
//			fields: []domain.PostField{
//				{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1", Parent: nil, Layout: nil, Index: nil},
//				{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &l1, Index: nil},
//				{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &l1, Index: nil},
//			},
//			key:  "flex",
//			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
//			args: nil,
//			want: Flexible{
//				{
//					Name: "layout1",
//					SubFields: SubFields{
//						{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &l1, Index: nil},
//						{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &l1, Index: nil},
//					},
//				},
//			},
//		},
//		"No Field": {
//			fields: nil,
//			key:    "wrongval",
//			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
//			args:   nil,
//			want:   "no field exists with the name: wrongval",
//		},
//		"Multiple Layouts": {
//			fields: []domain.PostField{
//				{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout2,layout3", Parent: nil, Layout: nil, Index: nil},
//				{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &l1, Index: nil},
//				{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &l2, Index: nil},
//				{Id: 4, Type: "text", Name: "text 3", Value: "text", Layout: &l3, Index: nil},
//			},
//			key:  "flex",
//			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
//			args: nil,
//			want: Flexible{
//				{
//					Name:      "layout1",
//					SubFields: SubFields{{Id: 2, Type: "text", Name: "text 1", Value: "text", Layout: &l1, Index: nil}},
//				},
//				{
//					Name:      "layout2",
//					SubFields: SubFields{{Id: 3, Type: "text", Name: "text 2", Value: "text", Layout: &l2, Index: nil}},
//				},
//				{
//					Name:      "layout3",
//					SubFields: SubFields{{Id: 4, Type: "text", Name: "text 3", Value: "text", Layout: &l3, Index: nil}},
//				},
//			},
//		},
//		"Format": {
//			fields: nil,
//			key:    "flex",
//			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
//				f.On("GetByPost", 1).Return([]domain.PostField{
//					{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1", Parent: nil, Layout: nil, Index: nil},
//					{Id: 2, Type: "category", Name: "text 1", Value: 1, Layout: &l1, Index: nil},
//				}, nil)
//				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
//			},
//			args: []interface{}{1, false},
//			want: Flexible{
//				{
//					Name:      "layout1",
//					SubFields: SubFields{{Id: 2, Type: "category", Name: "text 1", Value: 1, Layout: &l1, Index: nil}},
//				},
//			},
//		},
//		"Without Format": {
//			fields: nil,
//			key:    "flex",
//			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
//				f.On("GetByPost", 1).Return([]domain.PostField{
//					{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1", Parent: nil, Layout: nil, Index: nil},
//					{Id: 2, Type: "category", Name: "text 1", Value: 1, Layout: &l1, Index: nil},
//				}, nil)
//				c.On("GetById", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
//			},
//			args: []interface{}{1, true},
//			want: Flexible{
//				{
//					Name:      "layout1",
//					SubFields: SubFields{{Id: 2, Type: "category", Name: "text 1", Value: 1, Layout: &l1, Index: nil}},
//				},
//			},
//		},
//		"Invalid Type": {
//			fields: []domain.PostField{{Id: 1, Type: "text", Name: "flex", Value: 1, Parent: nil}},
//			key:    "flex",
//			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
//			want:   "field with the name: flex, is not flexible content",
//		},
//		//"Bad Cast to Layouts": {
//		//	fields: []domain.PostField{{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "333", Parent: nil}},
//		//	key:    "flex",
//		//	mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
//		//	want:   "unable to cast fields.noStringer{} of type fields.noStringer",
//		//},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.GetMockService(test.fields, test.mock)
//
//			got, err := s.GetFlexible(test.key, test.args...)
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//
//			t.Equal(test.want, got)
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestFlexible_HasRows() {
//
//	tt := map[string]struct {
//		flexible Flexible
//		want     interface{}
//	}{
//		"With Rows": {
//			flexible: Flexible{
//				{Name: "layout", SubFields: SubFields{domain.PostField{Id: 1, Name: "test"}}},
//			},
//			want: true,
//		},
//		"Without Rows": {
//			flexible: Flexible{},
//			want:     false,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.flexible.HasRows())
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestSubFields_SubField() {
//
//	subfield := SubFields{
//		{Id: 1, Name: "test1", Value: 1},
//		{Id: 2, Name: "test2", Value: 2},
//		{Id: 3, Name: "test3", Value: 3},
//	}
//
//	tt := map[string]struct {
//		key  string
//		want interface{}
//	}{
//		"Found": {
//			key:  "test1",
//			want: 1,
//		},
//		"Not Found": {
//			key:  "wrongval",
//			want: nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, subfield.SubField(test.key))
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestSubFields_First() {
//
//	tt := map[string]struct {
//		subfields SubFields
//		want      interface{}
//	}{
//		"Found": {
//			subfields: SubFields{
//				{Id: 1, Name: "test1", Value: 1},
//				{Id: 2, Name: "test2", Value: 2},
//				{Id: 3, Name: "test3", Value: 3},
//			},
//			want: domain.PostField{Id: 1, Name: "test1", Value: 1},
//		},
//		"Not Found": {
//			subfields: SubFields{},
//			want:      nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.subfields.First())
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestSubFields_Last() {
//
//	tt := map[string]struct {
//		subfields SubFields
//		want      interface{}
//	}{
//		"Found": {
//			subfields: SubFields{
//				{Id: 1, Name: "test1", Value: 1},
//				{Id: 2, Name: "test2", Value: 2},
//				{Id: 3, Name: "test3", Value: 3},
//			},
//			want: domain.PostField{Id: 3, Name: "test3", Value: 3},
//		},
//		"Not Found": {
//			subfields: SubFields{},
//			want:      nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.subfields.Last())
//		})
//	}
//}
