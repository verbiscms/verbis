// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

// Mime TODO
type mimeType string

// CanResize
//
// Returns true if the mime type is of JPG or PNG,
// determining if the image can be resized.
func (m mimeType) CanResize() bool {
	return m.IsJPG() || m.IsPNG()
}

// IsJPG
//
// Returns true if the mime type is of JPG.
func (m mimeType) IsJPG() bool {
	return m == "image/jpeg" || m == "image/jp2"
}

// IsPNG
//
// Returns true if the mime type is of PNG.
func (m mimeType) IsPNG() bool {
	return m == "image/png"
}
