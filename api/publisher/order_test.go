// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

var (
//newsResource  = "news"
//testResources = map[string]domain.Resource{
//	"news": {
//		Name:         newsResource,
//		FriendlyName: "News",
//		SingularName: "News Item",
//		Slug:         "news",
//		Hidden:       false,
//	},
//}
//singlePost = domain.PostDatum{
//	post: domain.post{Slug: "news-item", Title: "News Item", Resource: &newsResource},
//}
//archivePost = domain.PostDatum{
//	post: domain.post{Slug: "news", Title: "News"},
//}
//pagePost = domain.PostDatum{
//	post: domain.post{Slug: "contact", Title: "Contact"},
//}
//customSlugPost = domain.PostDatum{
//	post: domain.post{Slug: "custom/slug", Title: "Contact"},
//}
//category = domain.Category{
//	Name: "category",
//}
//categoryPost = domain.PostDatum{
//	post: domain.post{Slug: "custom/slug", Title: "Contact"}, Category: &category,
//}
)

//func TestRennder(t *testing.T) {
//
//	tt := map[string]struct {
//		input string
//		mock  func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string)
//		post  *domain.PostDatum
//		top   *TypeOfPage
//		err   error
//	}{
//		"Single": {
//			"news-article",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", url).Return(singlePost, nil)
//			},
//			&singlePost,
//			&TypeOfPage{Single, "news"},
//			nil,
//		},
//		"Archive": {
//			"news",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", url).Return(archivePost, nil)
//			},
//			&archivePost,
//			&TypeOfPage{Archive, "news"},
//			nil,
//		},
//		"Page": {
//			"contact",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", url).Return(pagePost, nil)
//				mc.On("GetBySlug", url).Return(domain.Category{}, fmt.Errorf("err"))
//			},
//			&pagePost,
//			&TypeOfPage{Page, ""},
//			nil,
//		},
//		"Custom Slug": {
//			"custom/slug",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", "slug").Return(domain.PostDatum{}, fmt.Errorf("err")).Once()
//				m.On("GetBySlug", url).Return(customSlugPost, nil)
//			},
//			&customSlugPost,
//			&TypeOfPage{Page, ""},
//			nil,
//		},
//		"Not Found": {
//			"wrong",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", url).Return(domain.PostDatum{}, fmt.Errorf("err")).Times(2)
//			},
//			nil,
//			nil,
//			fmt.Errorf("err"),
//		},
//		"Categories": {
//			"category",
//			func(m *mocks.PostsRepository, mc *mocks.CategoryRepository, url string) {
//				m.On("GetBySlug", url).Return(categoryPost, nil)
//				mc.On("GetBySlug", url).Return(category, nil)
//			},
//			&categoryPost,
//			&TypeOfPage{PageType: "category", Data: "category"},
//			nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			m := &mocks.PostsRepository{}
//			mc := &mocks.CategoryRepository{}
//
//			test.mock(m, mc, test.input)
//
//			p := publish{
//				Deps: &deps.Deps{
//					Store: &models.Store{
//						Categories: mc,
//						Posts:      m,
//					},
//					Theme: &domain.ThemeConfig{
//						Resources: testResources,
//					},
//					Options: &domain.Options{},
//				},
//			}
//
//			post, top, err := p.resolve(test.input)
//
//			assert.Equal(t, test.post, post)
//			assert.Equal(t, test.top, top)
//			assert.Equal(t, test.err, err)
//		})
//	}
//}
