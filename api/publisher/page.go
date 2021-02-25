// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
)

func (r *publish) Page(ctx *gin.Context) ([]byte, error) {
	const op = "publish.GetPage"

	api.ServeChan <- 1
	defer func() {
		<-api.ServeChan
	}()

	url, hasRedirected := r.handleTrailingSlash(ctx)
	if hasRedirected {
		return nil, nil
	}

	post, typ, err := r.resolve(url)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No page found with the url: %s", url), Operation: op, Err: err}
	}

	p := &page{
		Deps:     r.Deps,
		Context:  ctx,
		Post:     post,
		Url:      url,
		CacheKey: cache.GetPostKey(post.Id),
		Type:     typ,
	}

	err = p.CheckSession()
	if err != nil {
		return nil, err
	}

	err = p.IsResourcePublic()
	if err != nil {
		return nil, err
	}

	c, ok := p.GetCached()
	if ok {
		return c, nil
	}

	return p.Execute()
}
