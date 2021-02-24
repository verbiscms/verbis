// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	tplMocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/safe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"html/template"
	"io"
	"testing"
)

func Setup(opts domain.Options, site domain.Site, post domain.Post, tpl tpl.TemplateHandler) (*Namespace, *mocks.MediaRepository) {

	mock := &mocks.MediaRepository{}
	d := &deps.Deps{
		Store:   &models.Store{Media: mock},
		Site:    site,
		Options: &opts,
	}

	d.SetTmpl(tpl)

	ns := Namespace{
		deps: d,
		post: &domain.PostDatum{Post: post},
		funcs: template.FuncMap{
			"safeHTML": safe.New(d).HTML,
		},
	}

	return &ns, mock
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
			ns, _ := Setup(test.options, domain.Site{}, post, nil)
			got := ns.MetaTitle()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace(t *testing.T) {

	tt := map[string]struct {
		mock func(th *tplMocks.TemplateHandler, te *tplMocks.TemplateExecutor)
		fn   func(ns *Namespace) interface{}
		want interface{}
	}{
		"Header": {
			func(th *tplMocks.TemplateHandler, te *tplMocks.TemplateExecutor) {
				th.On("Prepare", tpl.Config{Root: DevEmbeddedPath, Extension: EmbeddedExtension}).Return(te)
				buf := bytes.Buffer{}
				te.On("Execute", &buf, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return("", nil)
			},
			func(ns *Namespace) interface{} {
				return ns.Header()
			},
			template.HTML("test\ntest\ntest\n"),
		},
		"Header Error": {
			func(th *tplMocks.TemplateHandler, te *tplMocks.TemplateExecutor) {
				th.On("Prepare", tpl.Config{Root: DevEmbeddedPath, Extension: EmbeddedExtension}).Return(te)
				te.On("Execute", &bytes.Buffer{}, mock.Anything, mock.Anything).Return("", fmt.Errorf("error"))
			},
			func(ns *Namespace) interface{} {
				return ns.Header()

			},
			template.HTML(""),
		},
		"Footer": {
			func(th *tplMocks.TemplateHandler, te *tplMocks.TemplateExecutor) {
				th.On("Prepare", tpl.Config{Root: DevEmbeddedPath, Extension: EmbeddedExtension}).Return(te)
				buf := bytes.Buffer{}
				te.On("Execute", &buf, "footer", mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return("", nil)
			},
			func(ns *Namespace) interface{} {
				return ns.Footer()
			},
			template.HTML("test\n"),
		},
		"Footer Error": {
			func(th *tplMocks.TemplateHandler, te *tplMocks.TemplateExecutor) {
				th.On("Prepare", tpl.Config{Root: DevEmbeddedPath, Extension: EmbeddedExtension}).Return(te)
				te.On("Execute", &bytes.Buffer{}, "footer", mock.Anything).Return("", fmt.Errorf("error"))
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
			tplExecutor := &tplMocks.TemplateExecutor{}

			test.mock(tplMock, tplExecutor)
			ns, _ := Setup(domain.Options{}, domain.Site{}, domain.Post{}, tplMock)

			got := test.fn(ns)

			assert.Equal(t, test.want, got)
			tplMock.AssertExpectations(t)
			tplExecutor.AssertExpectations(t)
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
				m.On("GetById", 1).Return(media, nil)
			},
			media.Url,
		},
		"Error": {
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup(domain.Options{}, domain.Site{}, domain.Post{}, nil)
			test.mock(mock)
			tm := TemplateMeta{deps: ns.deps}
			got := tm.GetImage(1)
			assert.Equal(t, test.want, got)
		})
	}
}
