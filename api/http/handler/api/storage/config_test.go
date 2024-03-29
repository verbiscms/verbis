// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	mocks "github.com/verbiscms/verbis/api/mocks/services/storage"
	"net/http"
)

func (t *StorageTestSuite) TestStorage_Config() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Provider)
	}{
		"Success": {
			storageConfig,
			http.StatusOK,
			"Successfully obtained configuration",
			func(m *mocks.Provider) {
				m.On("Info", mock.Anything).Return(storageConfig, nil)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/storage/config", "/storage/config", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Config(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
