package templates

import (
	"fmt"
	"testing"
)

// Test_Dict - Test dict function
func Test_Dict(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want string
		err  error
	}{
		"Valid": {
			tmpl: `{{ dict "test" 123 }}`,
			want: "map[test:123]",
			err: fmt.Errorf(""),
		},
		"Odd Value": {
			tmpl: `{{ dict "test" }}`,
			want: "",
			err: fmt.Errorf("Invalid dict call"),
		},
		"Not a String": {
			tmpl: `{{ dict 2 2 }}`,
			want: "",
			err: fmt.Errorf("Dict keys must be strings"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			if test.err.Error() != "" {
				rune(t, f, test.tmpl, test.want, test.err)
				return
			}
			runt(t, f, test.tmpl, test.want)
		})
	}
}
