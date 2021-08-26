// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
)

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func (t *MiddlewareTestSuite) SetupProxy(proxy domain.Proxy) (string, int, func()) {
	receivedRequestURI := make(chan string, 1)

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedRequestURI <- r.RequestURI
	}))

	serverURL, _ := url.Parse(upstream.URL)
	targetURL, _ := serverURL.Parse(proxy.Path)

	t.Engine.Use(Proxy(&deps.Deps{Options: &domain.Options{
		Proxies: []domain.Proxy{
			{Path: targetURL.Path, Host: serverURL.String(), Rewrite: proxy.Rewrite, RegexRewrite: proxy.RegexRewrite},
		},
	}}))

	rec := CreateTestResponseRecorder()
	req, err := http.NewRequest(http.MethodGet, proxy.Path, nil)
	req.RequestURI = proxy.Path

	t.NoError(err)
	t.Engine.ServeHTTP(rec, req)
	actualRequestURI := <-receivedRequestURI

	return actualRequestURI, t.Recorder.Code, func() {
		upstream.Close()
	}
}

func (t *MiddlewareTestSuite) TestProxy() {
	tt := map[string]struct {
		proxy  domain.Proxy
		path   string
		status int
		want   interface{}
	}{
		"Parse Error": {
			domain.Proxy{Host: "mysql://user:wrong", Path: "/path"},
			"/path",
			http.StatusOK,
			"verbis",
		},
		"Skip Path": {
			domain.Proxy{Host: "https://github.com", Path: "/test"},
			"/wrong",
			http.StatusOK,
			"verbis",
		},
		"Simple": {
			domain.Proxy{Host: "https://github.com", Path: "/test"},
			"/test",
			http.StatusOK,
			"github",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Engine.Use(Proxy(&deps.Deps{Options: &domain.Options{
				Proxies: []domain.Proxy{test.proxy},
			}}))

			t.Engine.GET(test.path, t.DefaultHandler)
			rec := CreateTestResponseRecorder()
			req, err := http.NewRequest(http.MethodGet, test.path, nil)
			t.NoError(err)
			t.Engine.ServeHTTP(rec, req)

			t.Equal(rec.Code, test.status)
			t.Contains(rec.ResponseRecorder.Body.String(), test.want)
			t.Reset()
		})
	}
}

func (t *MiddlewareTestSuite) TestProxy_Rewrite() {
	rewrites := map[string]string{
		"/old":              "/new",
		"/path":             "/new-path",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}

	tt := map[string]struct {
		path   string
		uri    string
		status int
	}{
		"Simple": {
			"/api/users",
			"/users",
			http.StatusOK,
		},
		"File": {
			"/js/main.js",
			"/public/javascripts/main.js",
			http.StatusOK,
		},
		"Absolute": {
			"/old",
			"/new",
			http.StatusOK,
		},
		"Multiple": {
			"/users/ainsley/orders/1",
			"/user/ainsley/order/1",
			http.StatusOK,
		},
		"Space": {
			"/api/new users",
			"/new%20users",
			http.StatusOK,
		},
		"Host": {
			"http://localhost:5000/path",
			"/new-path",
			http.StatusOK,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			proxy := domain.Proxy{Path: test.path, Rewrite: rewrites}
			uri, code, teardown := t.SetupProxy(proxy)
			defer teardown()
			t.Equal(test.status, code)
			t.Equal(test.uri, uri)
		})
	}
}

func (t *MiddlewareTestSuite) TestProxy_RewriteRegex() {
	rewrites := map[string]string{
		"^/a/*":     "/v1/$1",
		"^/b/*/c/*": "/v2/$2/$1",
		"^/c/*/*":   "/v3/$2",
	}
	rewritesRegex := map[*regexp.Regexp]string{
		regexp.MustCompile("^/x/.+?/(.*)"):   "/v4/$1",
		regexp.MustCompile("^/y/(.+?)/(.*)"): "/v5/$2/$1",
	}

	tt := map[string]struct {
		path   string
		uri    string
		status int
	}{
		"Unmatched": {
			"/unmatched",
			"/unmatched",
			http.StatusOK,
		},
		"Simple": {
			"/a/test",
			"/v1/test",
			http.StatusOK,
		},
		"Multiple": {
			"/b/foo/c/bar/baz",
			"/v2/bar/baz/foo",
			http.StatusOK,
		},
		"Multiple 2": {
			"/y/foo/bar",
			"/v5/bar/foo",
			http.StatusOK,
		},
		"Ignored": {
			"/c/ignore/test",
			"/v3/test",
			http.StatusOK,
		},
		"Ignore 2": {
			"/c/ignore1/test/this",
			"/v3/test/this",
			http.StatusOK,
		},
		"Ignore 3": {
			"/x/ignore/test",
			"/v4/test",
			http.StatusOK,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			proxy := domain.Proxy{Path: test.path, Rewrite: rewrites, RegexRewrite: rewritesRegex}
			uri, code, teardown := t.SetupProxy(proxy)
			defer teardown()
			t.Equal(test.status, code)
			t.Equal(test.uri, uri)
		})
	}
}
