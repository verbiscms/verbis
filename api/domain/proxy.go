// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type (
	// Proxy defines the configuration for a singular reverse proxy.
	Proxy struct {
		// Name is an identifier for the proxy used as a helper.
		Name string `json:"name"`
		// Path defines the original URL of the proxy, this will be
		// targeted and compared
		Path string `json:"path"`
		// Host defines the target URL of the reverse proxy.
		Host string `json:"host"`
		// Rewrite defines URL path rewrite rules. The values captured in
		// asterisk can be retrieved by index e.g. $1, $2 and so on.
		// Examples:
		// "/old":              "/new",
		// "/api/*":            "/$1",
		// "/js/*":             "/public/javascripts/$1",
		// "/users/*/orders/*": "/user/$1/order/$2",
		Rewrite map[string]string `json:"rewrite,omitempty"`
		// RegexRewrite defines rewrite rules using regexp.Rexexp
		// with captures. Every capture group in the values can
		// be retrieved by index e.g. $1, $2 and so on.
		// Example:
		// "^/old/[0.9]+/":     "/new",
		// "^/api/.+?/(.*)":    "/v2/$1",
		RegexRewrite map[string]string `json:"rewrite_regex,omitempty"`
	}
)
