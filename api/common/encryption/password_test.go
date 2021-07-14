// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password, err := HashPassword("password")
	assert.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte("password"))
	assert.NoError(t, err)
}

func TestHashPasswordError(t *testing.T) {
	orig := DefaultCost
	defer func() {
		DefaultCost = orig
	}()
	DefaultCost = 9999999999
	_, err := HashPassword("hello")
	assert.Error(t, err)
}

func TestCreatePassword(t *testing.T) {
	got := CreatePassword()
	assert.Equal(t, RandomPasswordLength, len(got))
}
