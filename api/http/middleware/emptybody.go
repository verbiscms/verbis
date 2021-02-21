// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// EmptyBody
//
// Determines if the  content type is JSON and the method
// type is a post or a put, if the body is invalid
// JSON or empty, abort JSON will be called.
func EmptyBody() gin.HandlerFunc {
	return func(g *gin.Context) {
		contentType := g.Request.Header.Get("Content-Type")
		if contentType != "application/json" && contentType != "application/json; charset=utf-8" {
			g.Next()
			return
		}

		method := g.Request.Method
		if method == "POST" || method == "PUT" {
			bodyBytes, _ := ioutil.ReadAll(g.Request.Body)

			if isEmpty(g, bodyBytes) {
				api.AbortJSON(g, 401, "Empty JSON body", nil)
				return
			}

			if !isJSON(string(bodyBytes)) {
				api.AbortJSON(g, 401, "Invalid JSON", nil)
				return
			}
		}

		g.Next()
	}
}

// isEmpty
//
// Checks if the request is empty.
func isEmpty(g *gin.Context, body []byte) bool {
	_ = g.Request.Body.Close()
	g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return len(body) == 0
}

// isJSON
//
// Checks if the request is valid json.
func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
