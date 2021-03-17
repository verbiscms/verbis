// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	tplMocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	siteMocks "github.com/ainsleyclark/verbis/api/mocks/verbis/site"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/safe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"html/template"
	"io"
	"io/ioutil"
	"testing"
)

func Setup(opts domain.Options, site domain.Site, post domain.Post) (*Namespace, *mocks.MediaRepository) {
	mockSite := &siteMocks.Repository{}
	mockSite.On("Global").Return(domain.Site{})

	m := &mocks.MediaRepository{}
	d := &deps.Deps{
		Store:   &models.Store{Media: m},
		Site:    mockSite,
		Options: &opts,
	}

	logger.SetOutput(ioutil.Discard)

	ns := Namespace{
		deps: d,
		post: &domain.PostDatum{Post: post},
		funcs: template.FuncMap{
			"safeHTML": safe.New(d).HTML,
		},
	}

	return &ns, m
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
		"With post Title": {
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

func TestNamespace(t *testing.T) {
	tt := map[string]struct {
		mock func(th *tplMocks.TemplateHandler)
		fn   func(ns *Namespace) interface{}
		want interface{}
	}{
		"Header": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) interface{} {
				return ns.Header()
			},
			template.HTML("test\ntest\ntest"),
		},
		"Header Error": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			func(ns *Namespace) interface{} {
				return ns.Header()
			},
			template.HTML(""),
		},
		"Footer": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) interface{} {
				return ns.Footer()
			},
			template.HTML("test"),
		},
		"Footer Error": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			func(ns *Namespace) interface{} {
				return ns.Footer()
			},
			template.HTML(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tplMock := &tplMocks.TemplateHandler{}
			ns, _ := Setup(domain.Options{}, domain.Site{}, domain.Post{})
			ns.deps.SetTmpl(tplMock)

			test.mock(tplMock)
			got := test.fn(ns)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestTemplateMeta_GetImage(t *testing.T) {
	media := domain.Media{Id: 1, Url: "testurl"}

	tt := map[string]struct {
		mock func(m *mocks.MediaRepository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.MediaRepository) {
				m.On("GetByID", 1).Return(media, nil)
			},
			media.Url,
		},
		"Error": {
			func(m *mocks.MediaRepository) {
				m.On("GetByID", 1).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, m := Setup(domain.Options{}, domain.Site{}, domain.Post{})
			test.mock(m)
			tm := TemplateMeta{deps: ns.deps}
			got := tm.GetImage(1)
			assert.Equal(t, test.want, got)
		})
	}
}
