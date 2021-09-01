// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Proxy creates a reverse proxy as defined in the options.
// The URL's will be parsed and compared. A new
// httputil.ReverseProxy will be created to
// handle traffic. URL's will be rewritten
// if there are regex and rewrite rules
// attached to the proxy.
// Capture tokens, rewrite rules regex are replicated from Echo as below.
// https://github.com/labstack/echo/blob/master/middleware/middleware.go
func Proxy(d *deps.Deps) gin.HandlerFunc {
	const op = "Middleware.Proxy"

	return func(ctx *gin.Context) {
		for _, config := range d.Options.Proxies {
			target, err := url.Parse(config.Host)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error parsing proxy configuration", Operation: op, Err: err}).Error()
				continue
			}

			if !strings.Contains(ctx.Request.URL.Path, config.Path) {
				continue
			}

			regexRewrites := make(map[*regexp.Regexp]string)
			if config.RegexRewrite == nil {
				regexRewrites = make(map[*regexp.Regexp]string)
			}

			for k, v := range config.RegexRewrite {
				regexRewrites[regexp.MustCompile(k)] = v
			}

			if config.Rewrite != nil {
				for k, v := range rewriteRulesRegex(config.Rewrite) {
					regexRewrites[k] = v
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

					err := rewriteURL(regexRewrites, req)
					if err != nil {
						logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error rewriting proxy URL", Operation: op, Err: err}).Error()
						ctx.Next()
						return
					}
				},
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}

			proxy.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		}

		ctx.Next()
	}
}

func captureTokens(pattern *regexp.Regexp, input string) *strings.Replacer {
	groups := pattern.FindAllStringSubmatch(input, -1)
	if groups == nil {
		return nil
	}
	values := groups[0][1:]
	replace := make([]string, 2*len(values)) //nolint
	for i, v := range values {
		j := 2 * i
		replace[j] = "$" + strconv.Itoa(i+1)
		replace[j+1] = v
	}
	return strings.NewReplacer(replace...)
}

func rewriteRulesRegex(rewrite map[string]string) map[*regexp.Regexp]string {
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

	// Depending on how the HTTP request is sent RequestURI could
	// contain Scheme://Host/path or be just /path. We only
	// want to use path part for rewriting and therefore
	// trim prefix if it exists.
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
		replacer := captureTokens(k, rawURI)
		if replacer != nil {
			uri, err := req.URL.Parse(replacer.Replace(v))
			if err != nil {
				return err
			}
			req.URL = uri
			return nil // Rewrite only once
		}
	}

	return nil
}
