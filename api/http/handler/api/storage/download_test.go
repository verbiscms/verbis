// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"github.com/verbiscms/verbis/api/services/storage"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_Download() {
	buf := bytes.Buffer{}

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Provider, ctx *gin.Context)
		err     bool
	}{
		"Success": {
			"test",
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Provider, ctx *gin.Context) {
				m.On("Download", ctx.Writer).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(gin.ResponseWriter)
					arg.Write([]byte("test"))
				})
			},
			false,
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Provider, ctx *gin.Context) {
				m.On("Download", mock.Anything).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			true,
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Provider, ctx *gin.Context) {
				m.On("Download", mock.Anything).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			true,
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Provider, ctx *gin.Context) {
				m.On("Download", mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer func() {
				t.Reset()
				buf.Reset()
			}()

			t.RequestAndServe("GET", "/download", "/download", nil, func(ctx *gin.Context) {
				m := &mocks.Provider{}
				test.mock(m, ctx)
				d := &deps.Deps{
					Storage: m,
				}
				New(d).Download(ctx)
			})

			t.Equal("application/octet-stream", t.Recorder.Header().Get("Content-Type"))
			t.Equal(storage.DownloadFileName, t.Recorder.Header().Get("X-Filename"))

			if test.err {
				got, _ := t.RespondData()
				t.Equal(test.message, got.Message)
				t.Equal(test.status, t.Status())
				return
			}

			t.Equal(test.want, t.Recorder.Body.String())
		})
	}
}
