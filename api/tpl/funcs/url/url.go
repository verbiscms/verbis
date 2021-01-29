package url

import (
	"github.com/gin-contrib/location"
)

/*
	TODO:
		- Coverage
		- Comments
  		- Docs
		- Teardown in test setup
		-
 */

func (ns *Namespace) Base() string {
	return location.Get(ns.ctx).String()
}

func (ns *Namespace) Scheme() string {
	return location.Get(ns.ctx).Scheme
}

func (ns *Namespace) Host() string {
	return location.Get(ns.ctx).Host
}

func (ns *Namespace) Full() string {
	return location.Get(ns.ctx).String() + ns.ctx.Request.URL.Path
}

func (ns *Namespace) Path() string {
	return ns.ctx.Request.URL.Path
}
