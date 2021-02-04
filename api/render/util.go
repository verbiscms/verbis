// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import "github.com/ainsleyclark/verbis/api/cache"

// GetCachedAsset checks to see if there is a cached version of the file
// and mimetypes, returns nil for both if nothing was found.
func (r *Render) getCachedAsset(url string) (*[]byte, *string) {

	if r.Options.CacheServerAssets {
		return nil, nil
	}

	file, foundFile := cache.Store.Get(url)
	mimeType, foundMime := cache.Store.Get(url + "mimetype")

	if foundFile && foundMime {
		file := file.(*[]byte)
		mimeType := mimeType.(*string)
		return file, mimeType
	}

	return nil, nil
}
