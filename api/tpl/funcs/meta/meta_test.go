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
	siteMocks "github.com/ainsleyclark/verbis/api/mocks/services/site"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/media"
	tplMocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/safe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"html/template"
	"io"
	"io/ioutil"
	"testing"
)

func Setup(opts domain.Options, post domain.Post) (*Namespace, *mocks.Repository) {
	mockSite := &siteMocks.Repository{}
	mockSite.On("Global").Return(domain.Site{})

	m := &mocks.Repository{}
	d := &deps.Deps{
		Store:   &store.Repository{Media: m},
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
		"With options": {
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
			ns, _ := Setup(test.options, post)
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
			template.HTML("test\ntest\ntest\ntest"),
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
		"Execute Error": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			func(ns *Namespace) interface{} {
				return ns.executeTemplates(&TemplateMeta{}, []string{"wrong"})
			},
			"",
		},
		"Regex Error": {
			func(th *tplMocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("\n\n\n\n\n"))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) interface{} {
				return ns.Header()
			},
			template.HTML(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if name == "Regex Error" {
				orig := newLineRegex
				defer func() {
					newLineRegex = orig
				}()
				newLineRegex = "[)"
			}

			tplMock := &tplMocks.TemplateHandler{}
			ns, _ := Setup(domain.Options{}, domain.Post{})
			ns.deps.SetTmpl(tplMock)

			test.mock(tplMock)
			got := test.fn(ns)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestTemplateMeta_GetImage(t *testing.T) {
	media := domain.Media{Id: 1, File: domain.File{Url: "testurl"}}

	tt := map[string]struct {
		mock func(m *mocks.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			media.File.Url,
		},
		"Error": {
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, m := Setup(domain.Options{}, domain.Post{})
			test.mock(m)
			tm := TemplateMeta{deps: ns.deps}
			got := tm.GetImage(1)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTemplateMeta_IsHomepage(t *testing.T) {
	tt := map[string]struct {
		input domain.Post
		want  bool
	}{
		"Success": {
			domain.Post{Id: 1},
			true,
		},
		"Error": {
			domain.Post{Id: 2},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tm := TemplateMeta{
				deps: &deps.Deps{
					Options: &domain.Options{Homepage: 1},
				},
				Post: &domain.PostDatum{Post: test.input},
			}
			got := tm.IsHomepage()
			assert.Equal(t, test.want, got)
		})
	}
}
