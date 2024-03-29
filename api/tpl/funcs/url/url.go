// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package url

import (
	"github.com/gin-contrib/location"
)

// Base
//
// Returns the current base url.
//
// Example: {{ baseUrl }}
// Returns: `http://verbiscms.com` (for example)
func (ns *Namespace) Base() string {
	return location.Get(ns.ctx).String()
}

// Scheme
//
// Returns the scheme of the current url
// `http` or `https`
//
// Example: {{ scheme }}
// Returns: `http` or `https` (for example)
func (ns *Namespace) Scheme() string {
	return location.Get(ns.ctx).Scheme
}

// Host
//
// Returns the host of the current url
//
// Example: {{ host }}
// Returns: `verbiscms.com` (for example)
func (ns *Namespace) Host() string {
	return location.Get(ns.ctx).Host
}

// Full
//
// Returns the current full url
//
// Example: {{ fullUrl }}
// Returns: `http://verbiscms.com/page` (for example)
func (ns *Namespace) Full() string {
	return location.Get(ns.ctx).String() + ns.ctx.Request.URL.Path
}

// Path
//
// Returns the path of the current url
//
// Example: {{ path }}
// Returns: `/page` (for example)
func (ns *Namespace) Path() string {
	return ns.ctx.Request.URL.Path
}
