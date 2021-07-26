// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_Session() {
	t.RequestAndServe(http.MethodPost, "/session", "/session", nil, func(ctx *gin.Context) {
		t.Setup(nil).CheckSession(ctx)
		t.RunT(nil, http.StatusOK, "Session valid")
	})
}
