// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

const (
	UpdatePasswordQuery      = "UPDATE users SET password = ? WHERE email = ?"
	DeletePasswordResetQuery = "DELETE FROM password_resets WHERE token = ?"
	InsertPasswordQuery = "INSERT INTO password_resets (email, token, created_at) VALUES (?, ?, NOW())"
)
