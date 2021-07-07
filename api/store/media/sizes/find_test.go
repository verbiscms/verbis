// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sizes

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

var (
	FindQuery = SelectStatement + "WHERE `media_sizes`.`media_id` = '" + mediaID + "'"
)

func (t *SizesTestSuite) TestStore_Find() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			sizes,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "size_key", "size_name"}).
					AddRow(sizes["hd"].Id, sizes["hd"].SizeKey, sizes["hd"].SizeName)
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
		"Internal Error": {
			database.ErrQueryMessage,
			func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnError(fmt.Errorf("error"))
			},
		},
		"Nil": {
			domain.MediaSizes{},
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"size_key", "size_name"})
				m.ExpectQuery(regexp.QuoteMeta(FindQuery)).
					WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.Find(1)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.RunT(test.want, got)
		})
	}
}
