// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Serve
//
// Serve is responsible for serving the correct data to
// the front end.
// Returns errors.NOTFOUND if the media item was not found.
func (s *Service) Serve(media domain.Media, path string, acceptWebP bool) ([]byte, domain.Mime, error) {
	const op = "Media.Serve"

	var (
		mime = media.File.Mime
		data []byte
		err  error
	)

	if acceptWebP && s.options.MediaServeWebP {
		data, _, err = s.storage.Find(path + domain.WebPExtension)
		if err != nil {
			data, _, err = s.storage.Find(path)
		} else {
			return data, "image/webp", nil
		}
	} else {
		data, _, err = s.storage.Find(path)
	}

	if err != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: "File does not exist with the path: " + path, Operation: op, Err: err}
	}

	return data, mime, nil
}
