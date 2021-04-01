// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

// HashPassword
//
// Gets the password in byte format and generates a
// hashed password with the default cost of 10.
//
// Returns errors.INTERNAL if the bcrypt failed to
// generate from password.
func HashPassword(password string) (string, error) {
	const op = "encryption.HashPassword"
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not hash the password with the string: %s", password), Operation: op, Err: err}
	}
	return string(hashedPassword), err
}

const (
	// The amount of characters generated for random
	// passwords.
	RandomPasswordLength = 24
)

// CreatePassword
//
// Creates a random password with a character length of
// 24.
func CreatePassword() string {
	var characterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@:\\/@£$%=^&&*()_+?><")
	b := make([]rune, RandomPasswordLength)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
