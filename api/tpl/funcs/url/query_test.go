package url

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type noStringer struct{}

func TestNamespace_Query(t *testing.T) {

	tt := map[string]struct {
		url   string
		input interface{}
		want  interface{}
	}{
		"Int Param": {
			"/?test=123",
			"test",
			"123",
		},
		"String Param": {
			"/?test=hello",
			"test",
			"hello",
		},
		"No Value": {
			"/?test=hello",
			"wrongval",
			"",
		},
		"Nasty Value": {
			"/?test=<script>alert('hacked!')</script>",
			"test",
			"&lt;script&gt;alert(&#39;hacked!&#39;)&lt;/script&gt;",
		},
		"Bad Cast": {
			"/?test=test",
			noStringer{},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			
			g, _ := gin.CreateTestContext(httptest.NewRecorder())
			g.Request, _ = http.NewRequest("GET", test.url, nil)

			ns := Namespace{
				ctx:  g,
			}

			got := ns.Query(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetPagination(t *testing.T) {

	tt := map[string]struct {
		url   string
		want  interface{}
	}{
		"Int Param": {
			"/?page=123",
			123,
		},
		"Empty": {
			"/",
			1,
		},
		"No Value": {
			"/?test=hello",
			1,
		},
		"Bad Cast": {
			`/page=wrongval`,
			1,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			g, _ := gin.CreateTestContext(httptest.NewRecorder())
			g.Request, _ = http.NewRequest("GET", test.url, nil)

			ns := Namespace{
				ctx:  g,
			}

			got := ns.Pagination()
			assert.Equal(t, test.want, got)
		})
	}
}