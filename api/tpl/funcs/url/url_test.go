package url

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Setup(t *testing.T, request string, f func(ns *Namespace) string) string {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)
	engine.Use(location.Default())

	got := ""
	engine.GET("/page", func(g *gin.Context) {
		ns := &Namespace{ctx: g}
		got = f(ns)
		return
	})

	req, err := http.NewRequest("GET", request, nil)
	assert.NoError(t, err)

	engine.ServeHTTP(rr, req)

	return got
}

func TestNamespace_Base(t *testing.T) {

	tt := map[string]struct {
		request string
		want string
	}{
		"HTTP": {
			"http://verbiscms.com/page",
			"http://verbiscms.com",
		},
		"HTTPS": {
			"https://verbiscms.com/page",
			"https://verbiscms.com",
		},
		"Relative": {
			"/page",
			"http://localhost:8080",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Setup(t, test.request, func(ns *Namespace) string {
				return ns.Base()
			})
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Scheme(t *testing.T) {

	tt := map[string]struct {
		request string
		want string
	}{
		"HTTP": {
			"http://verbiscms.com/page",
			"http",
		},
		"HTTPS": {
			"https://verbiscms.com/page",
			"https",
		},
		"Relative": {
			"/page",
			"http",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Setup(t, test.request, func(ns *Namespace) string {
				return ns.Scheme()
			})
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Host(t *testing.T) {

	tt := map[string]struct {
		request string
		want string
	}{
		"Valid": {
			"http://verbiscms.com/page",
			"verbiscms.com",
		},
		"Valid 2": {
			"http://verbiscms.co.uk/page",
			"verbiscms.co.uk",
		},
		"Relative": {
			"/page",
			"localhost:8080",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Setup(t, test.request, func(ns *Namespace) string {
				return ns.Host()
			})
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Full(t *testing.T) {

	tt := map[string]struct {
		request string
		want string
	}{
		"HTTP": {
			"http://verbiscms.com/page",
			"http://verbiscms.com/page",
		},
		"HTTPS": {
			"https://verbiscms.com/page",
			"https://verbiscms.com/page",
		},
		"Relative": {
			"/page",
			"http://localhost:8080/page",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Setup(t, test.request, func(ns *Namespace) string {
				return ns.Full()
			})
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Path(t *testing.T) {

	tt := map[string]struct {
		request string
		want string
	}{
		"HTTP": {
			"http://verbiscms.com/page",
			"/page",
		},
		"HTTPS": {
			"https://verbiscms.com/page",
			"/page",
		},
		"Relative": {
			"/page",
			"/page",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Setup(t, test.request, func(ns *Namespace) string {
				return ns.Path()
			})
			assert.Equal(t, test.want, got)
		})
	}
}