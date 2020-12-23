package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Body(t *testing.T) {

	resource := "resource"

	tt := map[string]struct {
		post   domain.Post
		cookie bool
		want   string
	}{
		"ID": {
			post: domain.Post{
				Id:           123,
				Title:        "title",
				Resource:     nil,
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			cookie: false,
			want:   "page page-id-123 page-title-title page-template-template page-layout-layout",
		},
		"Resource": {
			post: domain.Post{
				Id:           1,
				Title:        "title",
				Resource:     &resource,
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			cookie: false,
			want:   "resource page-id-1 page-title-title page-template-template page-layout-layout",
		},
		"Template": {
			post: domain.Post{
				Id:           1,
				Title:        "title",
				Resource:     &resource,
				PageTemplate: "%$££@template*&",
				PageLayout:   "layout",
			},
			cookie: false,
			want:   "resource page-id-1 page-title-title page-template-template page-layout-layout",
		},
		"Logged In": {
			post: domain.Post{
				Id:           1,
				Title:        "title",
				Resource:     nil,
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			cookie: true,
			want:   "page page-id-1 page-title-title page-template-template page-layout-layout logged-in",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			f.post = &domain.PostData{
				Post: test.post,
			}

			if test.cookie {
				mockUsers := mocks.UserRepository{}
				mockUsers.On("GetByToken", "token").Return(domain.User{}, nil)
				f.store.User = &mockUsers
				f.gin.Request.Header.Set("Cookie", "verbis-session=token")
			}

			assert.Equal(t, test.want, f.body())
		})
	}
}

func Test_CSSValidString(t *testing.T) {

	tt := map[string]struct {
		input string
		want  string
	}{
		"Regex 1": {
			input: "£$verbis$£$",
			want:  "verbis",
		},
		"Regex 2": {
			input: "£@$@£$$verbis{}|%$@£%",
			want:  "verbis",
		},
		"Spaces": {
			input: "verbis cms",
			want:  "verbis-cms",
		},
		"Forward Slash": {
			input: "verbis/cms/is/the/best",
			want:  "verbis-cms-is-the-best",
		},
		"Capital Letters": {
			input: "Verbis CMS",
			want:  "verbis-cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, cssValidString(test.input))
		})
	}
}
