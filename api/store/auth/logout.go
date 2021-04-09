// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

// Logout
//
// Logout checks to see if see if the the token is valid & then
// proceeds to create a new token and returns the user ID.
// Returns errors.NOTFOUND if the user was not found by the given token.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *Store) Logout(token string) (int, error) {
	const op = "AuthStore.Logout"

	user, err := s.userStore.FindByToken(token)
	if err != nil {
		return -1, err
	}

	newToken := s.generateTokeFunc(user.FirstName+user.LastName, user.Email)

	err = s.userStore.UpdateToken(newToken)
	if err != nil {
		return -1, err
	}

	return user.Id, nil
}
