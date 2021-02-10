// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
)

func (t *RecoverTestSuite) TestHandler_New() {
	d := &deps.Deps{}
	h := &Handler{d}
	t.Equal(h, New(d))
}

func (t *RecoverTestSuite) TestHandler_HttpRecovery() {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)
	engine.Use(location.Default())

	handlerMock := &mocks.TemplateHandler{}
	templateMock := &mocks.TemplateExecutor{}
	handlerMock.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(templateMock)
	templateMock.On("Exists", "error-500").Return(true)
	templateMock.On("Execute", &bytes.Buffer{}, "error-500", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(io.Writer)
		_, err := arg.Write([]byte("test"))
		t.NoError(err)
	}).Return(nil)

	d.SetTmpl(handlerMock)
	engine.Use(New(d).HttpRecovery())

	engine.GET("/test", func(ctx *gin.Context) {
		panic(&errors.Error{Message: "test"})
	})

	request, err := http.NewRequest("GET", "/test", &bytes.Buffer{})
	t.NoError(err)

	engine.ServeHTTP(rr, request)

	t.Equal("test", rr.Body.String())
	t.Equal(500, rr.Code)
}

func (t *RecoverTestSuite) TestHandler_HttpRecovery_Panics() {

	expectMsgs := map[syscall.Errno]string{
		syscall.EPIPE:      "broken pipe",
		syscall.ECONNRESET: "connection reset by peer",
	}

	for errno, expectMsg := range expectMsgs {
		t.Run(expectMsg, func() {
			t.NotPanics(func() {
				gin.SetMode(gin.TestMode)

				rr := httptest.NewRecorder()
				_, engine := gin.CreateTestContext(rr)
				engine.Use(New(d).HttpRecovery())

				engine.GET("/test", func(ctx *gin.Context) {
					ctx.Header("X-Test", "Value")
					ctx.Status(204)
					e := &net.OpError{Err: &os.SyscallError{Err: errno}}
					panic(e)
				})

				request, err := http.NewRequest("GET", "/test", nil)
				t.NoError(err)

				engine.ServeHTTP(rr, request)

				t.Equal("", rr.Body.String())
			})
		})
	}
}

func (t *RecoverTestSuite) TestRecover_RecoverWrapper() {

	var data = func() *Data {
		return &Data{}
	}

	var nilBytes []byte

	tt := map[string]struct {
		input    bool
		resolver resolver
		want     []byte
	}{
		"Theme Error Page": {
			true,
			func(custom bool) (string, tpl.TemplateExecutor, bool) {
				m := mocks.TemplateExecutor{}
				m.On("Execute", &bytes.Buffer{}, "root", mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					t.NoError(err)
				}).Return(nil)
				return "root", &m, true
			},
			[]byte("test"),
		},
		//"Error Executing Theme Error Page": {
		//	true,
		//	func(custom bool) (string, tpl.TemplateExecutor, bool) {
		//		m := mocks.TemplateExecutor{}
		//		m.On("Execute", &bytes.Buffer{}, "root", data()).Return(fmt.Errorf("error")).Once()
		//		m.On("Execute", &bytes.Buffer{}, "root", data()).Run(func(args mock.Arguments) {
		//			arg := args.Get(0).(io.Writer)
		//			_, err := arg.Write([]byte("test"))
		//			t.NoError(err)
		//		}).Return(fmt.Errorf("error")).Once()
		//		return "root", &m, true
		//	},
		//	[]byte("test"),
		//},
		"Verbis Error": {
			false,
			func(custom bool) (string, tpl.TemplateExecutor, bool) {
				m := mocks.TemplateExecutor{}
				m.On("Execute", &bytes.Buffer{}, "root", data()).Return(fmt.Errorf("error"))
				return "root", &m, false
			},
			nilBytes,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			r := &Recover{
				deps:     d,
				config:   Config{},
				resolver: test.resolver,
				data:     data,
			}
			got := r.recoverWrapper(test.input)
			t.Equal(test.want, got)
		})
	}
}
