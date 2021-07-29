// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
)

func (t *ThemeTestSuite) TestTheme_Templates() {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Success": {
			"verbis",
			domain.Templates{
				{Key: "nested/template-nested", Name: "Nested/Template Nested"},
				{Key: "template-hyphen", Name: "Template Hyphen"},
				{Key: "template", Name: "Template"},
			},
		},
		"Wrong Path": {
			"wrong",
			ErrNoTemplates.Error(),
		},
		"No Layouts": {
			"empty",
			ErrNoTemplates.Error(),
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.input)
			got, err := s.Templates()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_TemplatesError() {
	o := &options.Repository{}
	o.On("GetTheme").Return("", fmt.Errorf("error"))
	theme := Theme{options: o}
	got, err := theme.Templates()
	if err == nil {
		t.Fail("expecting error")
		return
	}
	t.Nil(got)
	t.Equal("error", err.Error())
}
