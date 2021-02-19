// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getSiteMock is a helper to obtain a mock site controller
// for testing.
func getSiteMock(m models.SiteRepository) *Site {
	return &Site{
		store: &models.Store{
			Site: m,
		},
	}
}

// Test_NewSite - Test construct
func Test_NewSite(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Site{
		store:  &store,
		config: config,
	}
	got := NewSite(&store, config)
	assert.Equal(t, got, want)
}


