// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/fields/converter"
)

func (t *FieldsTestSuite) TestStore_Layout() {
	s := t.Setup()

	post := domain.PostDatum{}
	mock := &mocks.Finder{}
	mock.On("Layout", post, s.Options.CacheServerFields).Return(groups)
	s.finder = mock

	got := s.Layout(domain.PostDatum{})
	t.Equal(groups, got)
}
