// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

func (s *Service) ReGenerateWebP() (int, error) {
	const op = "Media.ReGenerateWebP"

	mm, total, err := s.repo.List(params.Params{
		LimitAll: true,
	})

	if err != nil {
		return 0, err
	}

	if total == 0 {
		return 0, &errors.Error{Code: errors.NOTFOUND, Message: "Error regenerating webp images, none found", Operation: op, Err: fmt.Errorf("no webp images to process")}
	}

	go s.generateWebP(mm)

	return total, nil
}

func (s *Service) generateWebP(items domain.MediaItems) {
	for _, m := range items {
		s.deleteWebP(m.File, false)
		//s.fileToWebP(m.File)

		for _, size := range m.Sizes {
			s.deleteWebP(size.File, false)
			//s.fileToWebP(size.File)
		}
	}
}
