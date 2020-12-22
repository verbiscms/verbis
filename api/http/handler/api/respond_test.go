package api

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_calculateRequestTime(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//tt := map[string]struct {
	//	requestTime time.Time
	//	want        Meta
	//}{
	//	"Success": {
	//		requestTime: time.Now().Add(-10 * time.Second),
	//		want: Meta{
	//			RequestTime:  time.Now().Add(-10 * time.Second).UTC().String(),
	//			ResponseTime: time.Now().UTC().String(),
	//			LatencyTime:  "10s",
	//			Pagination:   nil,
	//		},
	//	},
	//}

	//for name, test := range tt {
	//t.Run(name, func(t *testing.T) {
	//	rr := httptest.NewRecorder()
	//	g, engine := gin.CreateTestContext(rr)
	//
	//	req, err := http.NewRequest("GET", "/test", nil)
	//	assert.NoError(t, err)
	//
	//	g.Request = req
	//	g.Set("request_time", test.requestTime)
	//
	//	meta := Meta{}
	//	engine.GET("/test", func(gin *gin.Context) {
	//		m := calculateRequestTime(g)
	//		meta = m
	//	})
	//	engine.ServeHTTP(rr, req)
	//
	//	time1, err := time.Parse(time.RFC850, meta.RequestTime)
	//	assert.NoError(t, err)
	//
	//	assert.WithinDuration(t, time1, test.requestTime, time.Millisecond*3)
	//	//assert.Equal(t, test.want, meta)
	//})
	//}
}
