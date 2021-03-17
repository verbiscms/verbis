// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/gin-gonic/gin"
)

func (r *publish) Page(ctx *gin.Context) ([]byte, error) {
	api.ServeChan <- 1
	defer func() {
		<-api.ServeChan
	}()

	pager, redirected, err := newPage(r.Deps, ctx)
	if redirected {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	err = pager.CheckSession()
	if err != nil {
		return nil, err
	}

	err = pager.IsResourcePublic()
	if err != nil {
		return nil, err
	}

	c, ok := pager.GetCached()
	if ok {
		return c, nil
	}

	return pager.Execute()
}
