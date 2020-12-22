package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func Test_Footer(t *testing.T) {

	code := `codeinjection `

	tt := map[string]struct {
		post    domain.Post
		options domain.Options
		want    template.HTML
	}{
		"CodeInjection Options & Post": {
			post:    domain.Post{CodeInjectionFoot: &code},
			options: domain.Options{CodeInjectionFoot: code},
			want:    "codeinjection codeinjection ",
		},
		"CodeInjection Post": {
			post:    domain.Post{CodeInjectionFoot: &code},
			options: domain.Options{},
			want:    "codeinjection ",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			f.post.Post = test.post
			f.options = test.options

			assert.Equal(t, test.want, f.footer())
		})
	}
}
