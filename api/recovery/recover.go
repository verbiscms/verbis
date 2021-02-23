// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/recovery/trace"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
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
	deps    *deps.Deps
	err     *errors.Error
	config  Config
	resolve Resolver
	data    TplData
	tracer  trace.Tracer
}

type Resolver func(custom bool) (string, tpl.TemplateExecutor, bool)

type TplData func() *Data

// Config defines
type Config struct {
	Code    int
	Context *gin.Context
	Error   interface{}
	TplFile string
	TplExec tpl.TemplateExecutor
	Post    *domain.PostData
}

func (h *Handler) newRecover(cfg Config) *Recover {
	r := &Recover{
		deps:   h.deps,
		err:    errors.ToError(cfg.Error),
		config: cfg,
		tracer: trace.New(),
	}
	r.resolve = r.resolveErrorPage
	r.data = r.getData
	return r
}

// Recover
//
//
func (h *Handler) Recover(cfg Config) []byte {
	r := h.newRecover(cfg)
	var tpl []byte
	custom := true

	r.recoverWrapper(custom, func(b []byte, err *errors.Error) {
		if err != nil {
			custom = false
		}
		tpl = b
	})

	//
	if !custom {
		r.recoverWrapper(custom, func(b []byte, err *errors.Error) {
			if err != nil {
				fmt.Println(err)
			}
			tpl = b
		})
	}

	return tpl
}

// HttpRecovery
//
//
func (h *Handler) HttpRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// If the connection is dead, we can't write a status to it.
				// Otherwise we will send the recover data back.
				if brokenPipe {
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				} else {
					b := h.Recover(Config{
						Context: ctx,
						Error:   err,
					})
					ctx.Data(500, "text/html", b)
					return
				}
			}
		}()
		ctx.Next()
	}
}

// recoverWrapper
//
// Obtains the template executor from the resolver, this could
// be a user defined error page, or an internal Verbis page
// dependant on the pages defined in the theme. The
// error page is executed and returned as bytes.
//
// Logs errors.INTERNAL if the internal Verbis error page
// failed to execute.
// Sets the config error errors.TEMPLATE if the user defined
// error page failed to execute.
func (r *Recover) recoverWrapper(useTheme bool, fn func(b []byte, err *errors.Error)) {
	const op = "Recovery.Recover"

	path, exec, custom := r.resolve(useTheme)

	var b bytes.Buffer
	_, err := exec.Execute(&b, path, r.data())

	// Theme error template failed, use the internal error pages
	if err != nil && custom {
		r.config.TplFile = path
		r.config.TplExec = exec
		fn(nil, &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template with the name: " + path, Operation: op, Err: err})
	}

	// Verbis error template failed, exit
	if err != nil && !custom {
		fn(nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to execute Verbis error template", Operation: op, Err: err})
	}

	fn(b.Bytes(), nil)
}
