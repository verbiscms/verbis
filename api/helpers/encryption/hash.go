// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

// GenerateRandomHash generates a unique random md5 hash
// Returns errors.INTERNAL if the hash failed to generate.
func GenerateRandomHash() (string, error) {
	const op = "encryption.GenerateRandomHash"
	hash, err := bcrypt.GenerateFromPassword([]byte(newSHA1Hash(36)), bcrypt.DefaultCost) //nolint
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Could not generate a random hash", Operation: op, Err: err}
	}
	hasher := md5.New()
	hasher.Write(hash) //nolint
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// newSHA1Hash generates a new SHA1 hash based on
// a random number of characters.
func newSHA1Hash(n ...int) string {
	noRandomCharacters := 32

	if len(n) > 0 {
		noRandomCharacters = n[0]
	}

	randString := RandomString(int64(noRandomCharacters), true)

	hash := sha1.New()
	hash.Write([]byte(randString)) //nolint
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// randomString generates a random string of n length
func RandomString(n int64, numeric bool) string {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if !numeric {
		characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
