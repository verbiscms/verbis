// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeaders_Cache(t *testing.T) {

	tt := map[string]struct {
		url     string
		want    string
		options domain.Options
	}{
		"Max Age": {
			want: "max-age=1000, public",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "max-age",
				CacheFrontendExtension: []string{"jpg"},
				CacheFrontendSeconds:   1000,
			},
		},
		"Max Stale": {
			want: "max-stale=1000, public",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "max-stale",
				CacheFrontendExtension: []string{"jpg"},
				CacheFrontendSeconds:   1000,
			},
		},
		"Min Fresh": {
			want: "min-fresh=1000, public",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "min-fresh",
				CacheFrontendExtension: []string{"jpg"},
				CacheFrontendSeconds:   1000,
			},
		},
		"No Cache": {
			want: "no-cache",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "no-cache",
				CacheFrontendExtension: []string{"jpg"},
			},
		},
		"No Store": {
			want: "no-store",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "no-store",
				CacheFrontendExtension: []string{"jpg"},
			},
		},
		"No Transform": {
			want: "no-transform",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "no-transform",
				CacheFrontendExtension: []string{"jpg"},
			},
		},
		"Only If Cached": {
			want: "only-if-cached",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "only-if-cached",
				CacheFrontendExtension: []string{"jpg"},
			},
		},
		"Cache Disabled": {
			want: "",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend: false,
			},
		},
		"Wrong Request": {
			want: "max-age=1000, public",
			url:  "/assets/images/test.jpg",
			options: domain.Options{
				CacheFrontend:          true,
				CacheFrontendRequest:   "wrongval",
				CacheFrontendExtension: []string{"jpg"},
				CacheFrontendSeconds:   1000,
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			r := gin.Default()

			h := newHeaders(test.options)

			server := httptest.NewServer(r)
			defer server.Close()
			url := "http://" + server.Listener.Addr().String() + test.url

			r.GET("/assets/*any", func(context *gin.Context) {
				h.Cache(context)
			})

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			assert.NoError(t, err)

			get, err := client.Do(req)
			assert.NoError(t, err)

			o := get.Header.Get("Cache-Control")
			assert.Equal(t, test.want, o)
		})
	}
}
