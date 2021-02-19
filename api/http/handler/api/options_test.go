// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
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

// getPostsMock is a helper to obtain a mock options controller
// for testing.
func getOptionsMock(m models.OptionsRepository) *Options {
	return &Options{
		store: &models.Store{
			Options: m,
		},
	}
}

// Test_NewOptions - Test construct
func Test_NewOptions(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Options{
		store:  &store,
		config: config,
	}
	got := NewwOptions(&store, config)
	assert.Equal(t, got, want)
}






