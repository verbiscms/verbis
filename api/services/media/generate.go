// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import "github.com/ainsleyclark/verbis/api/common/params"

func (s *Service) Generate(webp bool, sizes bool) (int, error) {

	mm, total, err := s.repo.List(params.Params{
		LimitAll: true,
	})

	if err != nil {
		return 0, err
	}

	for _, m := range mm {
		s.fileToWebP(m.File)
	}

	return total, nil

}
