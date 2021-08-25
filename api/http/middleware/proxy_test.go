// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"net/http"
)

func (t *MiddlewareTestSuite) Test_Proxy() {
	tt := map[string]struct {
		proxies     []domain.Proxy
		url         string
		code        int
		redirectURL string
	}{
		"Parse Error": {
			[]domain.Proxy{
				{Host: "wrong"},
			},
			"test",
			http.StatusOK,
			"/admin",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Engine.Use(Proxy(&deps.Deps{Options: &domain.Options{
				Proxies: test.proxies,
			}}))

			t.NewRequest(http.MethodGet, test.url, nil)
			t.ServeHTTP()

			fmt.Println(t.LogWriter.String())
			t.Reset()
		})
	}
}
