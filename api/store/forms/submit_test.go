// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	fields "github.com/ainsleyclark/verbis/api/mocks/store/forms/fields"
	submissions "github.com/ainsleyclark/verbis/api/mocks/store/forms/submissions"
)

func (t *FormsTestSuite) TestStore_Submit() {
	tt := map[string]struct {
		want      interface{}
		mockForms func(f *fields.Repository, s *submissions.Repository)
	}{
		"Success": {
			nil,
			func(f *fields.Repository, s *submissions.Repository) {
				s.On("Create", domain.FormSubmission{}).Return(nil)
			},
		},
		"Error": {
			fmt.Errorf("error"),
			func(f *fields.Repository, s *submissions.Repository) {
				s.On("Create", domain.FormSubmission{}).Return(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, test.mockForms)
			err := s.Submit(domain.FormSubmission{})
			t.Equal(test.want, err)
		})
	}
}
