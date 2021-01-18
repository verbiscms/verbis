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
		"Length of One": {
			field: domain.PostField{OriginalValue: "1", Type: "tags"},
			want:  domain.PostField{OriginalValue: "1", Type: "tags", Value: "1"},
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
