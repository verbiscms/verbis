// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getFieldsMock is a helper to obtain a mock fields controller
// for testing.
func getFieldsMock(m models.FieldsRepository) *Fields {
	return &Fields{
		store: &models.Store{
			Fields: m,
		},
	}
}

// Test_NewAuth - Test construct
func Test_NewFields(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Fields{
		store:  &store,
		config: config,
	}
	got := NewFields(&store, config)
	assert.Equal(t, got, want)
}

// TestAuthController_Get - Test Get route
func TestFieldController_Get(t *testing.T) {

}
