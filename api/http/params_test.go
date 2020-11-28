package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewPagination - Test construct
func TestNewParams(t *testing.T) {
	want := &Params{
		gin: &gin.Context{},
	}
	got := NewParams(&gin.Context{})
	assert.Equal(t, got, want)
}

// TestParams_Get - Test get Params
func TestParams_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tt := map[string]struct {
		url  string
		want *Params
	}{
		"Page": {
			url:  "page=2",
			want: &Params{Page: 2, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Nil Page": {
			url:  "page=wrong",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Limit All": {
			url:  "limit=all",
			want: &Params{Page: 1, Limit: 0, LimitAll: true, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Limit Failed": {
			url:  "limit=wrong",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Limit Zero": {
			url:  "limit=0",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Order": {
			url:  "order=name,desc",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "name", OrderDirection: "desc", Filters: nil},
		},
		"Order One Param": {
			url:  "order=id",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Order Comma": {
			url:  "order=id,",
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
		"Filter": {
			url: `&filter={"resource":[{"operator":"=", "value":"verbis"}]}`,
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: map[string][]Filter{
				"resource": {
					{
						Operator: "=",
						Value:    "verbis",
					},
				},
			}},
		},
		"Failed Filter": {
			url:  `&filter={"resource":[, "value":"verbis"}]}`,
			want: &Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "id", OrderDirection: "ASC", Filters: nil},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			g, engine := gin.CreateTestContext(rr)

			req, err := http.NewRequest("GET", "/test?"+test.url, nil)
			assert.NoError(t, err)
			g.Request = req

			params := &Params{}
			engine.GET("/test", func(g *gin.Context) {
				p := NewParams(g).Get()
				params = &p
			})
			engine.ServeHTTP(rr, req)

			assert.Equal(t, test.want, params)
		})
	}
}
