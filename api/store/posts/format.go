// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import "github.com/ainsleyclark/verbis/api/domain"

func (s *Store) format(raw []postsRaw, layout bool) domain.PostData {
	var posts = make(domain.PostData, 0)

	for _, v := range raw {
		if !s.find(posts, v.Id) {
			var category domain.Category
			if v.Category.Id != 0 {
				category = v.Category
			}

			p := domain.PostDatum{
				Post:     v.Post,
				Author:   v.Author.HideCredentials(),
				Category: &category,
				Fields:   make(domain.PostFields, 0),
			}

			//if layout {
			//p.Layout = s.fieldsModel.Layout(p)
			//}

			posts = append(posts, p)
		}

		if v.Field.UUID != nil {
			field := domain.PostField{
				Id:            v.Field.Id,
				PostId:        v.Field.PostId,
				UUID:          *v.Field.UUID,
				Type:          v.Field.Type,
				Name:          v.Field.Name,
				Key:           v.Field.Key,
				Value:         nil,
				OriginalValue: domain.FieldValue(v.Field.OriginalValue),
			}

			for fi, fv := range posts {
				if fv.Id == v.Id {
					posts[fi].Fields = append(posts[fi].Fields, field)
				}
			}
		}
	}

	return posts
}

func (s *Store) find(posts domain.PostData, id int) bool {
	for _, v := range posts {
		if v.Id == id {
			return true
		}
	}
	return false
}
