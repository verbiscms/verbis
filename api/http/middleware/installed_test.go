// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"net/http"
)

func (t *MiddlewareTestSuite) Test_Installed() {
	tt := map[string]struct {
		installed   bool
		method      string
		status      int
		url         string
		redirectURL string
	}{
		"Installed": {
			true,
			http.MethodGet,
			http.StatusOK,
			"/admin",
			"",
		},
		"File": {
			false,
			http.MethodGet,
			http.StatusOK,
			"/admin/javascript.js",
			"",
		},
		"Same URL": {
			false,
			http.MethodGet,
			http.StatusOK,
			"/admin/install",
			"",
		},
		"Post Preflight": {
			false,
			http.MethodPost,
			http.StatusOK,
			app.HTTPAPIRoute + "/install/preflight",
			"",
		},
		"Post Install": {
			false,
			http.MethodPost,
			http.StatusOK,
			app.HTTPAPIRoute + "/install",
			"",
		},
		"API Call": {
			false,
			http.MethodPost,
			http.StatusBadRequest,
			app.HTTPAPIRoute + "/test",
			"",
		},
		"Redirected": {
			false,
			http.MethodGet,
			http.StatusMovedPermanently,
			"/",
			"/admin/install",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Engine.Use(Installed(&deps.Deps{
				Installed: test.installed,
			}))

			t.RequestAndServe(test.method, test.url, test.url, nil, t.DefaultHandler)

			t.Equal(test.status, t.Status())
			if test.redirectURL != "" {
				loc, err := t.Recorder.Result().Location()
				t.NoError(err)
				t.Equal(test.redirectURL, loc.Path)
			}

			t.Context.Request.Body.Close()
			t.Reset()
		})
	}
}
