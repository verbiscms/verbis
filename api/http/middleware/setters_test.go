// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

//
//func (t *MiddlewareTestSuite) Test_Setters() {
//	opts := &domain.Options{
//		SiteTitle: "middleware",
//	}
//
//	tt := map[string]struct {
//		url    string
//		method string
//		mock   func(m *options.Repository, c *cache.Store, th *theme.Service)
//	}{
//		"Method Post": {
//			"/news",
//			http.MethodPost,
//			nil,
//		},
//		"File": {
//			"/news/sports.jpg",
//			http.MethodGet,
//			nil,
//		},
//		"From Cache": {
//			"/news",
//			http.MethodGet,
//			func(m *options.Repository, c *cache.Store, th *theme.Service) {
//				c.On("Get").Return(mock.Anything, mock.Anything).Return(opts, nil)
//			},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			var (
//				o  = &options.Repository{}
//				c  = &cache.Store{}
//				th = &theme.Service{}
//			)
//			if test.mock != nil {
//				test.mock(o, c, th)
//			}
//			d := &deps.Deps{
//				Store: &store.Repository{Options: o},
//				Cache: c,
//			}
//			t.Engine.Use(Setters(d))
//			t.RequestAndServe(test.method, test.url, test.url, nil, t.DefaultHandler)
//			t.Context.Request.Body.Close()
//			t.Reset()
//		})
//	}
//}
//
//func (t *MiddlewareTestSuite) Test_Setters_Options() {
//	opts := &domain.Options{
//		SiteTitle: "middleware",
//	}
//
//	tt := map[string]struct {
//		mock   func(m *options.Repository, c *mockCache.Cacher)
//		panics bool
//		want   interface{}
//	}{
//		"From Cache": {
//			func(m *options.Repository, c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.OptionsKey).Return(*opts, nil)
//			},
//			false,
//			opts,
//		},
//		"From DB": {
//			func(m *options.Repository, c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.OptionsKey).Return(nil, fmt.Errorf("error"))
//				m.On("Struct").Return(*opts)
//				c.On("Set", mock.Anything, cache.OptionsKey, *opts, mock.Anything).Return(nil)
//			},
//			false,
//			opts,
//		},
//		"Set Error": {
//			func(m *options.Repository, c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.OptionsKey).Return(nil, fmt.Errorf("error"))
//				m.On("Struct").Return(*opts)
//				c.On("Set", mock.Anything, cache.OptionsKey, *opts, mock.Anything).Return(fmt.Errorf("error"))
//			},
//			true,
//			opts,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			logger.SetOutput(ioutil.Discard)
//			defer t.Reset()
//			var (
//				o = &options.Repository{}
//				c = &mockCache.Cacher{}
//			)
//			if test.mock != nil {
//				test.mock(o, c)
//			}
//			d := &deps.Deps{
//				Store:   &store.Repository{Options: o},
//				Options: &domain.Options{},
//			}
//			cache.SetDriver(c)
//			t.RequestAndServe(http.MethodGet, "/set", "/set", nil, func(ctx *gin.Context) {
//				if test.panics {
//					t.Panics(func() {
//						setOptions(d, ctx)
//					})
//					return
//				}
//				setOptions(d, ctx)
//				t.Equal(test.want, d.Options)
//			})
//		})
//	}
//}
//
//func (t *MiddlewareTestSuite) Test_Setters_Theme() {
//	theme := &domain.ThemeConfig{
//		Theme: domain.Theme{
//			Title: "middleware",
//		},
//	}
//
//	tt := map[string]struct {
//		mock         func(mc *mockCache.Cacher)
//		themeFetcher func(path string) *domain.ThemeConfig
//		panics       bool
//		want         interface{}
//	}{
//		"From Cache": {
//			func(c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.ThemeConfigKey).Return(*theme, nil)
//			},
//			func(path string) *domain.ThemeConfig {
//				return theme
//			},
//			false,
//			theme,
//		},
//		"From Function": {
//			func(c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.ThemeConfigKey).Return(nil, fmt.Errorf("error"))
//				c.On("Set", mock.Anything, cache.ThemeConfigKey, *theme, mock.Anything).Return(nil)
//			},
//			func(path string) *domain.ThemeConfig {
//				return theme
//			},
//			false,
//			theme,
//		},
//		"Set Error": {
//			func(c *mockCache.Cacher) {
//				c.On("Get", mock.Anything, cache.ThemeConfigKey).Return(nil, fmt.Errorf("error"))
//				c.On("Set", mock.Anything, cache.ThemeConfigKey, *theme, mock.Anything).Return(fmt.Errorf("error"))
//			},
//			func(path string) *domain.ThemeConfig {
//				return theme
//			},
//			true,
//			theme,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			logger.SetOutput(ioutil.Discard)
//			orig := themeFetcher
//			defer func() {
//				t.Reset()
//				themeFetcher = orig
//			}()
//
//			themeFetcher = test.themeFetcher
//
//			var (
//				c = &mockCache.Cacher{}
//			)
//			if test.mock != nil {
//				test.mock(c)
//			}
//			d := &deps.Deps{
//				Config:  &domain.ThemeConfig{},
//				Options: &domain.Options{},
//			}
//			cache.SetDriver(c)
//			t.RequestAndServe(http.MethodGet, "/set", "/set", nil, func(ctx *gin.Context) {
//				if test.panics {
//					t.Panics(func() {
//						setTheme(d, ctx)
//					})
//					return
//				}
//				setTheme(d, ctx)
//				t.Equal(test.want, d.Config)
//			})
//		})
//	}
//}
