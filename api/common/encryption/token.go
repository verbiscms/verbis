// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateUserToken generates a new user token based on
// name & email.
func GenerateUserToken(name, email string) string {
	emailHash := MD5Hash(email)
	hash := MD5Hash(name + time.Now().String() + "3de" + strconv.Itoa(rand.Intn(143-0)+0) + emailHash) //nolint
	token := strconv.Itoa(rand.Intn(143-0)+0) + hash + strconv.Itoa(rand.Intn(143-0)+0)                //nolint
	return token
}

// GenerateEmailToken generates a token based on the email
// given.
// Returns errors.INTERNAL if the bcrypt failed to generate from password.
func GenerateEmailToken(email string) (string, error) {
	const op = "Encryption.GenerateEmailToken"

	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not generate the email token with the email: %s", email), Operation: op, Err: err}
	}

	return MD5Hash(string(hash)), nil
}
