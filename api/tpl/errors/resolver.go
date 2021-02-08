package errors

import (
	"github.com/ainsleyclark/verbis/api/tpl"
)

const (
	Prefix              = "error"
	NotFound            = "404"
	InternalServerError = "500"
)

type resolve struct {
	Path string
	Exec tpl.TemplateExecutor
}

// NotFound
//
//
func (r *Handler) NotFound() *Recover {
	return &Recover{
		deps: r.deps,
		tpl:  r.resolver(NotFound),
	}
}

// InternalServerError
//
//
func (r *Handler) InternalServerError() *Recover {
	return &Recover{
		deps: r.deps,
		tpl:  r.resolver(InternalServerError),
	}
}

// resolver
//
//
func (r *Handler) resolver(code string) *resolve {
	path := ""
	e := r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Theme + "/" + r.deps.Theme.TemplateDir,
		Extension: r.deps.Theme.FileExtension,
	})

	// Look for `errors-404.cms` for example
	path = Prefix + "-" + code
	if e.Exists(path) {
		return &resolve{Path: path, Exec: e}
	}

	// Look for `error.cms` for example
	path = Prefix
	if e.Exists(Prefix) {
		return &resolve{Path: path, Exec: e}
	}

	// Return native error page
	return &resolve{Path: "templates/error", Exec: r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Web,
		Extension: ".html",
		Master:    "layouts/main",
	})}
}
