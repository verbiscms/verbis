package http

//func TestNewParams(t *testing.T) {
//	want := &Params{
//		gin: &gin.Context{},
//	}
//	got := NewParams(&gin.Context{})
//	assert.Equal(t, got, want)
//}
//
//func TestParams_Get(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	tt := map[string]struct {
//		url  string
//		want *Params
//	}{
//		"Page": {
//			url:  "page=2",
//			want: &Params{Page: 2, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Nil Page": {
//			url:  "page=wrong",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Limit All": {
//			url:  "limit=all",
//			want: &Params{Page: 1, Limit: 0, LimitAll: true, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Limit Failed": {
//			url:  "limit=wrong",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Limit Zero": {
//			url:  "limit=0",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Order": {
//			url:  "order=name,desc",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "name", OrderDirection: "desc", Filters: nil},
//		},
//		"Order One Param": {
//			url:  "order=id",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Order Comma": {
//			url:  "order=id,",
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//		"Filter": {
//			url: `&filter={"resource":[{"operator":"=", "value":"verbis"}]}`,
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: map[string][]Filter{
//				"resource": {
//					{
//						Operator: "=",
//						Value:    "verbis",
//					},
//				},
//			}},
//		},
//		"Failed Filter": {
//			url:  `&filter={"resource":[, "value":"verbis"}]}`,
//			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			rr := httptest.NewRecorder()
//			g, engine := gin.CreateTestContext(rr)
//
//			req, err := http.NewRequest("GET", "/test?"+test.url, nil)
//			assert.NoError(t, err)
//			g.Request = req
//
//			params := &Params{}
//			engine.GET("/test", func(g *gin.Context) {
//				p := NewParams(g).Get()
//				params = &p
//			})
//			engine.ServeHTTP(rr, req)
//
//			assert.Equal(t, test.want, params)
//		})
//	}
//}

//func TestGetTemplateParams(t *testing.T) {
//
//	tt := map[string]struct {
//		input  map[string]interface{}
//		params TemplateParams
//		err    string
//	}{
//		"Nil": {
//			input:  nil,
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Page": {
//			input:  map[string]interface{}{"page": 3},
//			params: TemplateParams{Params: Params{Page: 3, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Page 0": {
//			input:  map[string]interface{}{"page": 0},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Limit": {
//			input:  map[string]interface{}{"limit": 10},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 10, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Limit All": {
//			input:  map[string]interface{}{"all": true},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: true, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Category": {
//			input:  map[string]interface{}{"category": "cat"},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "all", Category: "cat"},
//			err:    "",
//		},
//		"Resource": {
//			input:  map[string]interface{}{"resource": "res"},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, Resource: "res", Category: ""},
//			err:    "",
//		},
//		"Order By": {
//			input:  map[string]interface{}{"order_by": "title"},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "title", OrderDirection: "desc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Order Direction": {
//			input:  map[string]interface{}{"order_direction": "asc"},
//			params: TemplateParams{Params: Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "asc"}, Resource: "all", Category: ""},
//			err:    "",
//		},
//		"Marshal Error": {
//			input:  map[string]interface{}{"foo": make(chan int)},
//			params: TemplateParams{},
//			err:    "Could not convert query to Template Params",
//		},
//		"Unmarshal Error": {
//			input:  map[string]interface{}{"order_direction": 123},
//			params: TemplateParams{},
//			err:    "Could not convert query to Template Params",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func(t *testing.T) {
//			p, err := GetTemplateParams(test.input)
//			assert.Equal(t, p, test.params)
//
//			if test.err != "" {
//				assert.Equal(t, test.err, errors.Message(err))
//				return
//			}
//
//			assert.Equal(t, nil, err)
//		})
//	}
//}
