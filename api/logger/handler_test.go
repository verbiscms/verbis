// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
)

func (t *LoggerTestSuite) TestHandler() {
	tt := map[string]struct {
		err     interface{}
		message interface{}
		code    int
		want    interface{}
	}{
		"Nil": {
			nil,
			nil,
			200,
			"200 | [INFO]  | 192.0.2.1 |   GET    \"/test\"\n",
		},
		"Error": {
			&errors.Error{Code: errors.INTERNAL, Message: "message", Operation: "logger.Log", Err: fmt.Errorf("err")},
			nil,
			200,
			"200 | [INFO]  | 192.0.2.1 |   GET    \"/test\" | [code] internal [msg] message [op] logger.Log [error] err",
		},
		"Message": {
			nil,
			"message",
			200,
			"200 | [INFO]  | 192.0.2.1 |   GET    \"/test\" | [msg] message",
		},
		"400": {
			nil,
			"",
			400,
			"400 | [ERROR] | 192.0.2.1 |   GET    \"/test\"\n",
		},
		"401": {
			nil,
			"",
			401,
			"401 | [ERROR] | 192.0.2.1 |   GET    \"/test\"\n",
		},
		"404": {
			nil,
			"",
			404,
			"404 | [ERROR] | 192.0.2.1 |   GET    \"/test\"\n",
		},
		"500": {
			nil,
			"",
			500,
			"500 | [ERROR] | 192.0.2.1 |   GET    \"/test\"\n",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			buf := t.SetupHandler(func(ctx *gin.Context) {
				if test.err != nil {
					ctx.Set("verbis_error", test.err)
				}
				if test.message != nil {
					ctx.Set("verbis_message", test.message)
				}
				ctx.String(test.code, "test")
			})
			t.Contains(buf.String(), test.want)
		})
	}
}
