// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

//func (t *PostsTestSuite) TestStore_PostType() {
//	tt := map[string]struct {
//		input     domain.PostDatum
//		resources domain.Resources
//		home      int
//		want      domain.PostType
//	}{
//		"Homepage": {
//			domain.PostDatum{Post: domain.Post{Id: 1}},
//			nil,
//			1,
//			domain.PostType{
//				PageType: domain.HomeType,
//			},
//		},
//		"Archive Set": {
//			domain.PostDatum{Post: domain.Post{Id: 1, IsArchive: true}},
//			domain.Resources{"news": domain.Resource{Name: "News"}},
//			999,
//			domain.PostType{
//				PageType: domain.ArchiveType,
//				// TODO, this should resolve!
//				Data: domain.Resource{},
//			},
//		},
//		"Archive Slug": {
//			domain.PostDatum{Post: domain.Post{Id: 1, IsArchive: true, Slug: "news"}},
//			domain.Resources{"news": domain.Resource{Name: "News"}},
//			999,
//			domain.PostType{
//				PageType: domain.ArchiveType,
//				Data:     domain.Resource{Name: "News"},
//			},
//		},
//		"Resource": {
//			domain.PostDatum{Post: domain.Post{Id: 1, Resource: "news"}},
//			domain.Resources{"news": domain.Resource{Name: "News"}},
//			999,
//			domain.PostType{
//				PageType: domain.SingleType,
//				Data:     domain.Resource{Name: "News"},
//			},
//		},
//		"Single": {
//			domain.PostDatum{Post: domain.Post{Id: 1}},
//			nil,
//			999,
//			domain.PostType{
//				PageType: domain.PageType,
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(nil)
//			s.Theme.Resources = test.resources
//			s.Options.Homepage = test.home
//			got := s.postType(&test.input)
//			t.Equal(test.want, got)
//		})
//	}
//}
