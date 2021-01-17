package tpl

import (
	"bytes"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"net/http/httptest"
)

func (t *TplTestSuite) testUrl(request string, tpl string) string {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	g, engine := gin.CreateTestContext(rr)
	engine.Use(location.Default())

	engine.GET("/page", func(context *gin.Context) {
		t.gin = context

		tt := template.Must(template.New("test").Funcs(t.GetFunctions()).Parse(tpl))

		var b bytes.Buffer
		err := tt.Execute(&b, nil)
		t.NoError(err)

		g.String(200, b.String())
		return
	})

	req, err := http.NewRequest("GET", request, nil)
	t.NoError(err)

	engine.ServeHTTP(rr, req)

	return rr.Body.String()
}

func (t *TplTestSuite)  Test_GetBaseURL() {
	want := "https://verbiscms.com"
	got := t.testUrl("https://verbiscms.com/page", "{{ baseUrl }}")
	t.Equal(want, got)
}

func (t *TplTestSuite)  Test_GetScheme_HTTP() {
	want := "http"
	got := t.testUrl("http://verbiscms.com/page", "{{ scheme }}")
	t.Equal(want, got)
}

func (t *TplTestSuite)  Test_GetScheme_HTTPS() {
	want := "https"
	got := t.testUrl("https://verbiscms.com/page", "{{ scheme }}")
	t.Equal(want, got)
}

func (t *TplTestSuite)  Test_GetHost() {
	want := "verbiscms.com"
	got := t.testUrl("https://verbiscms.com/page", "{{ host }}")
	t.Equal(want, got)
}

func (t *TplTestSuite)  Test_GetFullURL() {
	//want := "https://verbiscms.com/page"
	_ = t.testUrl("https://verbiscms.com/page", "{{ fullUrl }}")
	//assert.Equal(t, want, got)
}

func (t *TplTestSuite)  Test_GetQueryParams() {

	tt := map[string]struct {
		url   string
		data  interface{}
		input string
		want  interface{}
	}{
		"Int Param": {
			url:   "/?test=123",
			input: `{{ query "test" }}`,
			want:  "123",
		},
		"String Param": {
			url:   "/?test=hello",
			input: `{{ query "test" }}`,
			want:  "hello",
		},
		"No Value": {
			url:   "/?test=hello",
			input: `{{ query "wrongval" }}`,
			want:  "",
		},
		"Nasty Value": {
			url:   "/?test=<script>alert('hacked!')</script>",
			input: `{{ query "test" }}`,
			want:  "&lt;script&gt;alert(&#39;hacked!&#39;)&lt;/script&gt;",
		},
		"Bad Cast": {
			url: "/?test=test",
			data: map[string]interface{}{
				"Data": noStringer{},
			},
			input: `{{ query .Data }}`,
			want:  "",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			g, _ := gin.CreateTestContext(httptest.NewRecorder())
			g.Request, _ = http.NewRequest("GET", test.url, nil)
			t.gin = g

			tpl := template.Must(template.New("test").Funcs(t.GetFunctions()).Parse(test.input))

			var b bytes.Buffer
			if err := tpl.Execute(&b, test.data); err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, b.String())
		})
	}
}
