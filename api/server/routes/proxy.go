// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	serpSpeed  = "https://35.214.23.223:5000/tools/serp-speed/"
	rereDirect = "http://35.246.93.18"
)

type ProxyConfig struct {
	// Path defines the original URL of the
	Path string
	// Host defines
	Host string
	// Rewrite defines URL path rewrite rules. The values captured in asterisk can be
	// retrieved by index e.g. $1, $2 and so on.
	// Examples:
	// "/old":              "/new",
	// "/api/*":            "/$1",
	// "/js/*":             "/public/javascripts/$1",
	// "/users/*/orders/*": "/user/$1/order/$2",
	Rewrite map[string]string
	// RegexRewrite defines rewrite rules using regexp.Rexexp with captures
	// Every capture group in the values can be retrieved by index e.g. $1, $2 and so on.
	// Example:
	// "^/old/[0.9]+/":     "/new",
	// "^/api/.+?/(.*)":    "/v2/$1",
	RegexRewrite map[*regexp.Regexp]string
}

var configuration = []ProxyConfig{
	{
		Path:    "/serp-speed",
		Host:    "https://35.214.23.223:5000",
		Rewrite: map[string]string{
			//"/tools/serp-speed/*": "$1",
		},
	},
}

func Proxy(ctx *gin.Context) {
	for _, config := range configuration {
		target, err := url.Parse(config.Host)
		if err != nil {
			ctx.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
			return
		}

		if config.Rewrite != nil {
			if config.RegexRewrite == nil {
				config.RegexRewrite = make(map[*regexp.Regexp]string)
			}
			for k, v := range rewriteRulesRegex(config.Rewrite) {
				config.RegexRewrite[k] = v
			}
		}

		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.Header = ctx.Request.Header
				req.Header.Add("X-Forwarded-Host", req.Host)
				req.Header.Add("X-Origin-Host", target.Host)
				req.Host = target.Host
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host

				if err := rewriteURL(config.RegexRewrite, req); err != nil {
					fmt.Println(err)
				}
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}

		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func captureTokens(pattern *regexp.Regexp, input string) *strings.Replacer {
	groups := pattern.FindAllStringSubmatch(input, -1)
	if groups == nil {
		return nil
	}
	values := groups[0][1:]
	replace := make([]string, 2*len(values))
	for i, v := range values {
		j := 2 * i
		replace[j] = "$" + strconv.Itoa(i+1)
		replace[j+1] = v
	}
	return strings.NewReplacer(replace...)
}

func rewriteRulesRegex(rewrite map[string]string) map[*regexp.Regexp]string {
	// Initialize
	rulesRegex := map[*regexp.Regexp]string{}
	for k, v := range rewrite {
		k = regexp.QuoteMeta(k)
		k = strings.Replace(k, `\*`, "(.*?)", -1)
		if strings.HasPrefix(k, `\^`) {
			k = strings.Replace(k, `\^`, "^", -1)
		}
		k = k + "$"
		rulesRegex[regexp.MustCompile(k)] = v
	}
	return rulesRegex
}

func rewriteURL(rewriteRegex map[*regexp.Regexp]string, req *http.Request) error {
	if len(rewriteRegex) == 0 {
		return nil
	}

	// Depending how HTTP request is sent RequestURI could contain Scheme://Host/path or be just /path.
	// We only want to use path part for rewriting and therefore trim prefix if it exists
	rawURI := req.RequestURI
	if rawURI != "" && rawURI[0] != '/' {
		prefix := ""
		if req.URL.Scheme != "" {
			prefix = req.URL.Scheme + "://"
		}
		if req.URL.Host != "" {
			prefix += req.URL.Host // host or host:port
		}
		if prefix != "" {
			rawURI = strings.TrimPrefix(rawURI, prefix)
		}
	}

	for k, v := range rewriteRegex {
		if replacer := captureTokens(k, rawURI); replacer != nil {
			url, err := req.URL.Parse(replacer.Replace(v))
			if err != nil {
				return err
			}
			req.URL = url

			return nil // rewrite only once
		}
	}

	return nil
}
