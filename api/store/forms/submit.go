// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import "github.com/verbiscms/verbis/api/domain"

// Submit
//
// Returns nil if the submissions was stored successfully.
func (s *Store) Submit(f domain.FormSubmission) error {
	return s.submissions.Create(f)
}
