// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

// VerifyEmail the users email address based on the encryption hash string passed
// Returns errors.NOTFOUND if the user was not found by the md5string email.
// Returns errors.INTERNAL if the SQL query was invalid.
//func (s *AuthStore) VerifyEmail(md5String string) error {
//const op = "AuthRepository.VerifyEmail"
//
//var userVerified = struct {
//	Id   int    `db:"id"` //nolint
//	Hash string `db:"hash"`
//}{}
//
//if err := s.DB.Get(&userVerified, "SELECT id AS id, MD5(CONCAT(id, email)) AS hash FROM users WHERE MD5(CONCAT(id, email)) = ?", md5String); err != nil {
//	return &errors.Error{Code: errors.NOTFOUND, Message: "Could not find the user for email verification", Operation: op, Err: err}
//}
//
//q := "UPDATE users SET email_verified_at = NOW() WHERE ID = ?"
//_, err := s.DB.Exec(q, userVerified.Id)
//if err != nil {
//	return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could update the user with the Id: %d", userVerified.Id), Operation: op, Err: err}
//}
//
//return nil
//}
