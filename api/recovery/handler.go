package recovery

import (
	"github.com/ainsleyclark/verbis/api/deps"
)

const (
	Prefix = "error"
)

type Handler struct {
	deps *deps.Deps
}

func New(d *deps.Deps) *Handler {
	return &Handler{deps: d}
}

func (h *Handler) New() *Recover {
	return &Recover{deps: h.deps, code: 0}
}
