package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

// Test_Header - Test verbisHeader functionality
func Test_Header(t *testing.T) {

	code := "codeinjection post"
	cannonical := "test"

	tt := map[string]struct {
		post    domain.Post
		options domain.Options
		site    domain.Site
		want    template.HTML
	}{
		"Success": {
			post: domain.Post{
				Id:             123,
				Title:          "title",
				Resource:       nil,
				PageTemplate:   "template",
				Layout:         "layout",
				CodeInjectHead: &code,
			},
			options: domain.Options{},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `codeinjection post
<meta name="robots" content="noindex">
<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"No CodeInjection for Post": {
			post: domain.Post{CodeInjectHead: nil},
			site: domain.Site{Url: "https://verbiscms.com"},
			want: `<meta name="robots" content="noindex">
<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"No CodeInjection Globally": {
			post:    domain.Post{},
			options: domain.Options{CodeInjectionHead: "test"},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `test
<meta name="robots" content="noindex">
<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"Not Public for Post": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Seo:    &domain.PostSeo{
						Public:         false,
					},
				},
			},
			options: domain.Options{},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<meta name="robots" content="noindex">
<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"Public": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Seo:    &domain.PostSeo{
						Public:         true,
					},
				},
			},
			options: domain.Options{SeoPublic: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"Canonical": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Seo:    &domain.PostSeo{
						Canonical: &cannonical,
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="test" />`,
		},
		"Meta with Post Description": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Description: "description",
					},
					Seo:   &domain.PostSeo{
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="https://verbiscms.com" />
<meta name="description" content="description">`,
		},
		"Meta with Global Description": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Description: "",
					},
					Seo:   &domain.PostSeo{
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true, MetaDescription: "description"},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="https://verbiscms.com" />
<meta name="description" content="description">`,
		},
		"Meta with Facebook Post Description": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Title:       "",
						Description: "",
						Facebook: struct {
							Title       string `json:"title,omitempty"`
							Description string `json:"description,omitempty"`
							ImageId     int    `json:"image_id,numeric,omitempty"`
						}{
							Description: "description",
						},
					},
					Seo:   &domain.PostSeo{
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true, MetaDescription: "description"},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="https://verbiscms.com" />
<meta name="description" content="description">`,
		},
		"Meta with Twitter Post Description": {
			post: domain.Post{
				SeoMeta:        domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Title:       "",
						Description: "",
						Twitter: struct {
							Title       string `json:"title,omitempty"`
							Description string `json:"description,omitempty"`
							ImageId     int    `json:"image_id,numeric,omitempty"`
						}{
							Description: "description",
						},
					},
					Seo:   &domain.PostSeo{
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true, MetaDescription: "description"},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `<link rel="canonical" href="https://verbiscms.com" />
<meta name="description" content="description">
<meta name="twitter:card" content="summary">
<meta name="twitter:description" content="">`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			mockMedia := mocks.MediaRepository{}
			mockMedia.On("GetById", 0).Return(domain.Media{}, fmt.Errorf("no image"))

			f.store.Media = &mockMedia
			f.site = &test.site
			f.post.Post = test.post
			f.options = test.options

			assert.Equal(t, test.want, f.header())
		})
	}
}
