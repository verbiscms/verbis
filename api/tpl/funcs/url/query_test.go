package url

//func TestNamespace_Query(t *testing.T) {
//
//	tt := map[string]struct {
//		url   string
//		data  interface{}
//		input string
//		want  interface{}
//	}{
//		"Int Param": {
//			url:   "/?test=123",
//			input: `{{ query "test" }}`,
//			want:  "123",
//		},
//		"String Param": {
//			url:   "/?test=hello",
//			input: `{{ query "test" }}`,
//			want:  "hello",
//		},
//		"No Value": {
//			url:   "/?test=hello",
//			input: `{{ query "wrongval" }}`,
//			want:  "",
//		},
//		"Nasty Value": {
//			url:   "/?test=<script>alert('hacked!')</script>",
//			input: `{{ query "test" }}`,
//			want:  "&lt;script&gt;alert(&#39;hacked!&#39;)&lt;/script&gt;",
//		},
//		"Bad Cast": {
//			url: "/?test=test",
//			data: map[string]interface{}{
//				"Data": noStringer{},
//			},
//			input: `{{ query .Data }}`,
//			want:  "",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			g, _ := gin.CreateTestContext(httptest.NewRecorder())
//			g.Request, _ = http.NewRequest("GET", test.url, nil)
//			t.gin = g
//
//			tpl := template.Must(template.New("test").Funcs(t.GetFunctions()).Parse(test.input))
//
//			var b bytes.Buffer
//			if err := tpl.Execute(&b, test.data); err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//
//			t.Equal(test.want, b.String())
//		})
//	}
//}

//func (t *TplTestSuite) TestGetPagination() {
//	g, _ := gin.CreateTestContext(httptest.NewRecorder())
//	g.Request, _ = http.NewRequest("GET", "/get?page=123", nil)
//	t.gin = g
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, 123)
//}
//
//func (t *TplTestSuite) TestGetPagination_NoPage() {
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, 1)
//}
//
//func (t *TplTestSuite) TestGetPagination_ConvertString() {
//	g, _ := gin.CreateTestContext(httptest.NewRecorder())
//	g.Request, _ = http.NewRequest("GET", "/get?page=wrongval", nil)
//	t.gin = g
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, "1")
//}
