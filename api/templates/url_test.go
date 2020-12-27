package templates

import (
	"bytes"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testUrl(t *testing.T, request string, tpl string) string {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	g, engine := gin.CreateTestContext(rr)
	engine.Use(location.Default())

	engine.GET("/page", func(context *gin.Context) {
		f := newTestSuite()
		f.gin = context

		tt := template.Must(template.New("test").Funcs(f.GetFunctions()).Parse(tpl))

		var b bytes.Buffer
		err := tt.Execute(&b, nil)
		assert.NoError(t, err)

		g.String(200, b.String())
		return
	})

	req, err := http.NewRequest("GET", request, nil)
	assert.NoError(t, err)

	engine.ServeHTTP(rr, req)

	return rr.Body.String()
}

func Test_GetBaseURL(t *testing.T) {
	want := "https://verbiscms.com"
	got := testUrl(t, "https://verbiscms.com/page", "{{ baseUrl }}")
	assert.Equal(t, want, got)
}

func Test_GetScheme_HTTP(t *testing.T) {
	want := "http"
	got := testUrl(t, "http://verbiscms.com/page", "{{ scheme }}")
	assert.Equal(t, want, got)
}

func Test_GetScheme_HTTPS(t *testing.T) {
	want := "https"
	got := testUrl(t, "https://verbiscms.com/page", "{{ scheme }}")
	assert.Equal(t, want, got)
}

func Test_GetHost(t *testing.T) {
	want := "verbiscms.com"
	got := testUrl(t, "https://verbiscms.com/page", "{{ host }}")
	assert.Equal(t, want, got)
}

func Test_GetFullURL(t *testing.T) {
	//want := "https://verbiscms.com/page"
	_ = testUrl(t, "https://verbiscms.com/page", "{{ fullUrl }}")
	//assert.Equal(t, want, got)
}

func Test_GetQueryParams(t *testing.T) {

	tt := map[string]struct {
		url string
		data interface{}
		input  string
		want  interface{}
	}{
		"Int Param": {
			url: "/?test=123",
			input: `{{ query "test" }}`,
			want: "123",
		},
		"String Param": {
			url: "/?test=hello",
			input: `{{ query "test" }}`,
			want: "hello",
		},
		"No Value": {
			url: "/?test=hello",
			input: `{{ query "wrongval" }}`,
			want: "",
		},
		"Nasty Value": {
			url: "/?test=<script>alert('hacked!')</script>",
			input: `{{ query "test" }}`,
			want: "&lt;script&gt;alert(&#39;hacked!&#39;)&lt;/script&gt;",
		},
		"Bad Cast": {
			url: "/?test=test",
			data: map[string]interface{}{
				"Data": noStringer{},
			},
			input: `{{ query .Data }}`,
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			gin.SetMode(gin.TestMode)
			g, _ := gin.CreateTestContext(httptest.NewRecorder())
			g.Request, _ = http.NewRequest("GET", test.url, nil)
			f.gin = g

			tt := template.Must(template.New("test").Funcs(f.GetFunctions()).Parse(test.input))

			var b bytes.Buffer
			if err := tt.Execute(&b, test.data); err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, b.String())
		})
	}
}