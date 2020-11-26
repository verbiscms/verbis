package controllers

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_NewCache - Test construct
func Test_NewCache(t *testing.T) {
	cache.Init()
	want := &CacheController{}
	got := newCache()
	assert.Equal(t, got, want)
}

// TestCacheController_Clear - Test Clear route
func TestCacheController_Clear(t *testing.T) {

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
			mock := &CacheController{}

			rr.RequestAndServe("POST", "/reset", "/reset", nil, func(g *gin.Context) {
				mock.Clear(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}