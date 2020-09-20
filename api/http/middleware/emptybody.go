package middleware

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func EmptyBody() gin.HandlerFunc {
	return func(g *gin.Context) {

		method := g.Request.Method
		contentType := g.Request.Header.Get("Content-Type")

		if contentType == "application/json" {
			if method == "POST" || method == "PUT" {
				bodyBytes, _ := ioutil.ReadAll(g.Request.Body)

				if isEmpty(g, bodyBytes) {
					controllers.AbortJSON(g, 401, "Empty JSON body", nil)
					return
				}

				if !isJSON(string(bodyBytes)) {
					controllers.AbortJSON(g, 401, "Invalid JSON", nil)
					return
				}
			}

			g.Next()
		}

		g.Next()
	}
}

// Checks if the request is empty
func isEmpty(g *gin.Context, body []byte) bool {
	_ = g.Request.Body.Close()  //  must close
	g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return len(body) == 0
}

// Checks if the request is valid json
func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}