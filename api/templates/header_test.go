package templates

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
	"time"
)

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
				Id:                123,
				Title:             "title",
				Resource:          nil,
				PageTemplate:      "template",
				PageLayout:        "layout",
				CodeInjectionHead: &code,
			},
			options: domain.Options{},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want: `codeinjection post
<meta name="robots" content="noindex">
<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"No CodeInjection for Post": {
			post: domain.Post{CodeInjectionHead: nil},
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
				SeoMeta: domain.PostSeoMeta{
					Seo: &domain.PostSeo{
						Public: false,
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
				SeoMeta: domain.PostSeoMeta{
					Seo: &domain.PostSeo{
						Public: true,
					},
				},
			},
			options: domain.Options{SeoPublic: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want:    `<link rel="canonical" href="https://verbiscms.com" />`,
		},
		"Canonical": {
			post: domain.Post{
				SeoMeta: domain.PostSeoMeta{
					Seo: &domain.PostSeo{
						Canonical: &cannonical,
						Public:    true,
					},
				},
			},
			options: domain.Options{SeoPublic: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want:    `<link rel="canonical" href="test" />`,
		},
		"Meta with Post Description": {
			post: domain.Post{
				SeoMeta: domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Description: "description",
					},
					Seo: &domain.PostSeo{
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
				SeoMeta: domain.PostSeoMeta{
					Meta: &domain.PostMeta{
						Description: "",
					},
					Seo: &domain.PostSeo{
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
				SeoMeta: domain.PostSeoMeta{
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
					Seo: &domain.PostSeo{
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
				SeoMeta: domain.PostSeoMeta{
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
					Seo: &domain.PostSeo{
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

func Test_WriteMeta(t *testing.T) {

	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "22 May 90 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}

	tt := map[string]struct {
		description string
		publishedAt *time.Time
		want        string
	}{
		"With Description": {
			publishedAt: nil,
			description: "verbis",
			want:        `<meta name="description" content="verbis">`,
		},
		"Without Description": {
			publishedAt: nil,
			description: "",
			want:        ``,
		},
		"With Time": {
			publishedAt: &tm,
			want:        `<meta property="article:modified_time" content="1990-05-22 21:39:39 +0100 BST" />`,
		},
		"Nothing": {
			publishedAt: nil,
			want:        ``,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			if test.publishedAt != nil {
				f.post.PublishedAt = test.publishedAt
			}

			var b bytes.Buffer
			f.writeMeta(&b, test.description)

			assert.Equal(t, test.want, b.String())
		})
	}
}

func Test_WriteFacebook(t *testing.T) {

	media := domain.Media{
		Id:  1,
		Url: "/media/url",
	}
	opts := domain.Options{
		SiteTitle:     "verbis",
		GeneralLocale: "en-gb",
	}

	tt := map[string]struct {
		title       string
		description string
		mock        func(m *mocks.MediaRepository)
		options     domain.Options
		want        string
	}{
		"With Title & Description": {
			title:       "cms",
			description: "verbis",
			options:     opts,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: `<meta property="og:type" content="website"><meta property="og:site_name" content="verbis"><meta property="og:locale" content="en-gb"><meta property="og:title" content="cms"><meta property="og:description" content="verbis"><meta property="og:image" content="/media/url">`,
		},
		"Without Title & Description": {
			title:       "",
			description: "",
			options:     opts,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("err"))
			},
			want: ``,
		},
		"No Image": {
			title:       "cms",
			description: "verbis",
			options:     opts,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("err"))
			},
			want: `<meta property="og:type" content="website"><meta property="og:site_name" content="verbis"><meta property="og:locale" content="en-gb"><meta property="og:title" content="cms"><meta property="og:description" content="verbis">`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			mock := mocks.MediaRepository{}
			f := newTestSuite()

			test.mock(&mock)
			f.store.Media = &mock
			f.options = test.options

			var b bytes.Buffer
			f.writeFacebook(&b, test.title, test.description, 1)

			assert.Equal(t, test.want, b.String())
		})
	}
}

func Test_WriteTwitter(t *testing.T) {

	media := domain.Media{
		Id:  1,
		Url: "/media/url",
	}

	tt := map[string]struct {
		title       string
		description string
		mock        func(m *mocks.MediaRepository)
		want        string
	}{
		"With Title & Description": {
			title:       "cms",
			description: "verbis",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: `<meta name="twitter:card" content="summary"><meta name="twitter:title" content="cms"><meta name="twitter:description" content="cms"><meta name="twitter:image" content="/media/url">`,
		},
		"Without Title & Description": {
			title:       "",
			description: "",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("err"))
			},
			want: ``,
		},
		"No Image": {
			title:       "cms",
			description: "verbis",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("err"))
			},
			want: `<meta name="twitter:card" content="summary"><meta name="twitter:title" content="cms"><meta name="twitter:description" content="cms">`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			mock := mocks.MediaRepository{}
			f := newTestSuite()

			test.mock(&mock)
			f.store.Media = &mock

			var b bytes.Buffer
			f.writeTwitter(&b, test.title, test.description, 1)

			assert.Equal(t, test.want, b.String())
		})
	}
}

func Test_MetaTitle(t *testing.T) {

	tt := map[string]struct {
		meta    domain.PostMeta
		options domain.Options
		want    string
	}{
		"With Post Title": {
			meta:    domain.PostMeta{Title: "post-title-verbis"},
			options: domain.Options{},
			want:    "post-title-verbis",
		},
		"With Options": {
			meta:    domain.PostMeta{Title: ""},
			options: domain.Options{MetaTitle: "post-title-verbis"},
			want:    "post-title-verbis",
		},
		"None": {
			meta:    domain.PostMeta{Title: ""},
			options: domain.Options{MetaTitle: ""},
			want:    "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			f.post.Post = domain.Post{
				SeoMeta: domain.PostSeoMeta{
					Meta: &test.meta,
				},
			}
			f.options = test.options

			assert.Equal(t, test.want, f.metaTitle())
		})
	}
}
