package recovery

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	deps *deps.Deps
}

func New(d *deps.Deps) *Handler {
	return &Handler{deps: d}
}

type Recovery interface {
	Recover(cfg Config) []byte
	HttpRecovery() gin.HandlerFunc
}

// Recover defines
type Recover struct {
	deps   *deps.Deps
	code   int
	err    *errors.Error
	config Config
}

// Config defines
type Config struct {
	Context *gin.Context
	Error   interface{}
	TplFile string
	TplExec tpl.TemplateExecutor
	Post    *domain.PostData
}

// New
//
// TODO: Should we be passing codes in? Or have it in the config?
func (h *Handler) New(code int) *Recover {
	return &Recover{
		deps: h.deps,
		code: code,
	}
}

// Recover
//
//
func (r *Recover) Recover(cfg Config) []byte {
	r.config = cfg
	r.err = getError(cfg.Error)
	return r.recoverWrapper(true)
}

// recoverWrapper
//
// Obtains the template executor from the resolver, this could
// be a user defined error page, or an internal Verbis page
// dependant on the pages defined in the theme. The
// error page is executed and returned as bytes.
//
// Logs errors.INTERNAL if the internal Verbis error page failed to execute.
// Sets the config error errors.TEMPLATE if the user defined error page failed to execute.
func (r *Recover) recoverWrapper(useTheme bool) []byte {
	const op = "Recovery.Recover"

	path, exec, custom := r.resolver(useTheme)

	var b bytes.Buffer
	err := exec.Execute(&b, path, r.getData())

	// Theme error template failed, use the internal error pages
	if err != nil && custom {
		r.config.TplFile = path
		r.config.TplExec = exec
		r.err = &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template with the name: " + path, Operation: op, Err: err}
		return r.recoverWrapper(false)
	}

	// Verbis error template failed, exit.
	if err != nil && !custom {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INTERNAL, Message: "Unable to execute Verbis error template", Operation: op, Err: err},
		}).Error()
		return nil
	}

	return b.Bytes()
}

