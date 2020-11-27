package api

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_NewCache - Test construct
func Test_NewCache(t *testing.T) {
	cache.Init()
	want := &Cache{}
	got := NewCache()
	assert.Equal(t, got, want)
}

// TestCache_Clear - Test Clear route
func TestCache_Clear(t *testing.T) {

	cache.Init()

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully cleared server cache",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &Cache{}

			rr.RequestAndServe("POST", "/reset", "/reset", nil, func(g *gin.Context) {
				mock.Clear(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
