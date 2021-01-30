package url

import (
	"github.com/gin-contrib/location"
)

// Base
//
// Returns the current base URL.
//
// Example: {{ baseUrl }}
// Returns: `http://verbiscms.com` (for example)
func (ns *Namespace) Base() string {
	return location.Get(ns.ctx).String()
}

// Scheme
//
// Returns the scheme of the current URL
// `http` or `https`
//
// Example: {{ scheme }}
// Returns: `http` or `https` (for example)
func (ns *Namespace) Scheme() string {
	return location.Get(ns.ctx).Scheme
}

// Host
//
// Returns the host of the current URL
//
// Example: {{ host }}
// Returns: `verbiscms.com` (for example)
func (ns *Namespace) Host() string {
	return location.Get(ns.ctx).Host
}

// Full
//
// Returns the current full URL
//
// Example: {{ fullUrl }}
// Returns: `http://verbiscms.com/page` (for example)
func (ns *Namespace) Full() string {
	return location.Get(ns.ctx).String() + ns.ctx.Request.URL.Path
}

// Path
//
// Returns the path of the current URL
//
// Example: {{ path }}
// Returns: `/page` (for example)
func (ns *Namespace) Path() string {
	return ns.ctx.Request.URL.Path
}
