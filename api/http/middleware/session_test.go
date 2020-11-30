package middleware

import (
	"encoding/json"
	"fmt"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRedirects - Test redirects from options are working correctly
// with correct code & location.
func Test_SessionCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Session Expired", func(t *testing.T) {
		token := "tokenVal"
		mock := mocks.UserRepository{}
		mock.On("CheckSession", token).Return(fmt.Errorf("error"))

		rr := httptest.NewRecorder()
		g, engine := gin.CreateTestContext(rr)
		engine.Use(SessionCheck(&mock))

		engine.GET("/test", func(context *gin.Context) {
			g.String(200, "verbis")
		})

		g.Request, _ = http.NewRequest("GET", "/test", nil)
		g.Request.Header.Set("token", token)
		engine.ServeHTTP(rr, g.Request)

		var body map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &body)
		assert.NoError(t, err)

		data, err := json.Marshal(body["data"].(map[string]interface{}))
		assert.NoError(t, err)

		assert.JSONEq(t, `{"errors":{"session":"expired"}}`, string(data))
		assert.Equal(t, "Session expired, please login again", body["message"])
		assert.Equal(t, 401, rr.Code)
		assert.Equal(t, &http.Cookie{
			Name:     "verbis-session",
			Value:    "",
			Path:     "/",
			Raw:      "verbis-session=; Path=/; Max-Age=0; HttpOnly",
			Domain:   "",
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
		}, rr.Result().Cookies()[0])
	})

	t.Run("Next", func(t *testing.T) {
		token := "tokenVal"
		mock := mocks.UserRepository{}
		mock.On("CheckSession", token).Return(nil)

		rr := httptest.NewRecorder()
		g, engine := gin.CreateTestContext(rr)
		engine.Use(SessionCheck(&mock))

		engine.GET("/test", func(context *gin.Context) {
			g.String(200, "verbis")
		})

		g.Request, _ = http.NewRequest("GET", "/test", nil)
		g.Request.Header.Set("token", token)
		engine.ServeHTTP(rr, g.Request)

		assert.Equal(t, "verbis", rr.Body.String())
		assert.Equal(t, 200, rr.Code)
	})
}
