package meta

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
	"time"
)

func Setup(opts domain.Options, site domain.Site, post domain.Post) (*Namespace, *mocks.MediaRepository) {
	mock := &mocks.MediaRepository{}
	ns := Namespace{
		deps: &deps.Deps{
			Store: &models.Store{
				Media: mock,
			},
			Site:    site,
			Options: opts,
		},
		post: &domain.PostData{
			Post: post,
		},
	}
	return &ns, mock
}

func TestNamespace_Header(t *testing.T) {

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
				SeoMeta: domain.PostOptions{
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
				SeoMeta: domain.PostOptions{
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
				SeoMeta: domain.PostOptions{
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
		"Canonical With Slash": {
			post: domain.Post{
				SeoMeta: domain.PostOptions{
					Seo: &domain.PostSeo{
						Canonical: &cannonical,
						Public:    true,
					},
				},
			},
			options: domain.Options{SeoPublic: true, SeoEnforceSlash: true},
			site:    domain.Site{Url: "https://verbiscms.com"},
			want:    `<link rel="canonical" href="test/" />`,
		},
		"Meta with Post Description": {
			post: domain.Post{
				SeoMeta: domain.PostOptions{
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
				SeoMeta: domain.PostOptions{
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
				SeoMeta: domain.PostOptions{
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
				SeoMeta: domain.PostOptions{
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
			ns, mock := Setup(test.options, test.site, test.post)
			mock.On("GetById", 0).Return(domain.Media{}, fmt.Errorf("no image"))
			got := ns.Header()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_WriteMeta(t *testing.T) {

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
			post := domain.Post{}
			if test.publishedAt != nil {
				post.PublishedAt = test.publishedAt
			}

			ns, mock := Setup(domain.Options{}, domain.Site{}, post)
			mock.On("GetById", 0).Return(domain.Media{}, fmt.Errorf("no image"))

			var b bytes.Buffer
			ns.writeMeta(&b, test.description)
			assert.Equal(t, test.want, b.String())
		})
	}
}

func TestNamespace_WriteFacebook(t *testing.T) {

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
			ns, mock := Setup(test.options, domain.Site{}, domain.Post{})
			test.mock(mock)

			var b bytes.Buffer
			ns.writeFacebook(&b, test.title, test.description, 1)
			assert.Equal(t, test.want, b.String())
		})
	}
}

func TestNamespace_WriteTwitter(t *testing.T) {

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
			ns, mock := Setup(domain.Options{}, domain.Site{}, domain.Post{})
			test.mock(mock)

			var b bytes.Buffer
			ns.writeTwitter(&b, test.title, test.description, 1)
			assert.Equal(t, test.want, b.String())
		})
	}
}

func TestNamespace_MetaTitle(t *testing.T) {

	tt := map[string]struct {
		meta    domain.PostOptions
		options domain.Options
		want    string
	}{
		"Nil Meta": {
			meta:    domain.PostOptions{Meta: nil},
			options: domain.Options{},
			want:    "",
		},
		"With Post Title": {
			meta:    domain.PostOptions{Meta: &domain.PostMeta{Title: "post-title-verbis"}},
			options: domain.Options{},
			want:    "post-title-verbis",
		},
		"With Options": {
			meta:    domain.PostOptions{Meta: &domain.PostMeta{Title: ""}},
			options: domain.Options{MetaTitle: "post-title-verbis"},
			want:    "post-title-verbis",
		},
		"None": {
			meta:    domain.PostOptions{Meta: &domain.PostMeta{Title: ""}},
			options: domain.Options{MetaTitle: ""},
			want:    "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			post := domain.Post{
				SeoMeta: test.meta,
			}
			ns, _ := Setup(test.options, domain.Site{}, post)
			got := ns.MetaTitle()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Footer(t *testing.T) {

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
			ns, _ := Setup(test.options, domain.Site{}, test.post)
			got := ns.Footer()
			assert.Equal(t, test.want, got)
		})
	}
}
