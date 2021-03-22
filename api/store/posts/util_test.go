// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	//"github.com/ainsleyclark/verbis/api/errors"
	"regexp"
)

func (t *PostsTestSuite) TestStore_Validate() {
	tt := map[string]struct {
		want interface{}
		mock func(m sqlmock.Sqlmock)
	}{
		"Success": {
			postDatum,
			func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "slug", "title"}).
					AddRow(post.Id, post.Slug, post.Title)
				m.ExpectQuery(regexp.QuoteMeta(selectStmt(FindBySlugQuery))).WillReturnRows(rows)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)

			//r := "resource"
			c := 1
			p := domain.PostCreate{
				Post: domain.Post{
					Slug: "slug",
					//Resource: &r,
				},
				Category: &c,
				Fields:   nil,
			}

			s.validate(p)
			fmt.Println(test.want)
			//if err != nil {
			//	t.Contains(errors.Message(err), test.want)
			//	return
			//}
			//t.RunT(test.want, err)
		})
	}
}
