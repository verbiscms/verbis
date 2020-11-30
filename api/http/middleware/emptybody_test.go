package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// handler is ac helper func for the EmptyBody testing
func handler(g *gin.Context) {
	g.String(200, "verbis")
	return
}

// TestEmptyBody - Test EmptyBody handler
func TestEmptyBody(t *testing.T) {

	tt := map[string]struct {
		method        string
		header        string
		input         string
		message       string
		status        int
		returnContent string
		want          string
	}{
		"Valid": {
			want:          "verbis",
			input:         `{verbis: "cms"}`,
			method:        http.MethodDelete,
			status:        200,
			header:        "application/json",
			returnContent: "text/plain; charset=utf-8",
		},
		"Not JSON": {
			want:          "verbis",
			input:         "",
			method:        http.MethodGet,
			status:        200,
			header:        "text/plain; charset=utf-8",
			returnContent: "text/plain; charset=utf-8",
		},
		"Empty Body": {
			want:          "",
			input:         "",
			message:       "Empty JSON body",
			method:        http.MethodPost,
			status:        401,
			header:        "application/json; charset=utf-8",
			returnContent: "application/json; charset=utf-8",
		},
		"Invalid JSON": {
			want:          "",
			input:         "notjson",
			message:       "Invalid JSON",
			method:        http.MethodPost,
			status:        401,
			header:        "application/json; charset=utf-8",
			returnContent: "application/json; charset=utf-8",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			rr := httptest.NewRecorder()
			context, engine := gin.CreateTestContext(rr)
			engine.Use(EmptyBody())

			engine.GET("/test", handler)
			engine.PUT("/test", handler)
			engine.POST("/test", handler)
			engine.DELETE("/test", handler)

			context.Request, _ = http.NewRequest(test.method, "/test", bytes.NewBuffer([]byte(test.input)))
			context.Request.Header.Add("Content-Type", test.header)
			engine.ServeHTTP(rr, context.Request)

			assert.Equal(t, test.status, rr.Code)
			assert.Equal(t, test.returnContent, rr.Header().Get("content-type"))

			if test.message != "" {
				var body map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &body)
				assert.NoError(t, err)
				assert.Equal(t, test.message, body["message"])
			} else {
				assert.Equal(t, test.want, rr.Body.String())
			}
		})
	}
}

// Test_isEmpty - Test checker for empty body
func Test_isEmpty(t *testing.T) {

	tt := map[string]struct {
		want  bool
		input string
	}{
		"Empty": {
			want:  true,
			input: "",
		},
		"With Body": {
			want:  false,
			input: "{}",
		},
		"With Body JSON": {
			want:  false,
			input: `{body: "verbis"}`,
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			rr := httptest.NewRecorder()
			context, engine := gin.CreateTestContext(rr)

			var got bool
			engine.GET("/test", func(g *gin.Context) {
				bodyBytes, _ := ioutil.ReadAll(g.Request.Body)
				got = isEmpty(context, bodyBytes)
			})

			context.Request, _ = http.NewRequest("GET", "/test", bytes.NewBuffer([]byte(test.input)))
			engine.ServeHTTP(rr, context.Request)

			assert.Equal(t, test.want, got)
		})
	}
}

// Test_isJSON - Test checker for is JSON
func Test_isJSON(t *testing.T) {

	tt := map[string]struct {
		want  bool
		input string
	}{
		"Empty": {
			want:  false,
			input: "",
		},
		"With Body": {
			want:  true,
			input: "{}",
		},
		"With Body JSON": {
			want:  true,
			input: `{"body": "verbis"}`,
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			rr := httptest.NewRecorder()
			context, engine := gin.CreateTestContext(rr)

			var got bool
			engine.GET("/test", func(g *gin.Context) {
				bodyBytes, _ := ioutil.ReadAll(g.Request.Body)
				got = isJSON(string(bodyBytes))
			})

			context.Request, _ = http.NewRequest("GET", "/test", bytes.NewBuffer([]byte(test.input)))
			engine.ServeHTTP(rr, context.Request)

			assert.Equal(t, test.want, got)
		})
	}
}
