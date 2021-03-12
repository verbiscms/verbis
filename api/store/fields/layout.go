// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import "github.com/ainsleyclark/verbis/api/domain"

// Layout
//
// Loops over all of the locations within the config json
// file that is defined. Produces an array of field
// groups that can be returned for the post.
func (s *Store) Layout(post domain.PostDatum) domain.FieldGroups {
	return s.finder.Layout(post, s.Options.CacheServerFields)
}
