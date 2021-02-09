package recovery

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Recovery interface {
	Recover(w io.Writer, ctx *gin.Context, err interface{})
}

type Recover struct {
	deps *deps.Deps
	code int
	err  *errors.Error
}

// NotFound
//
//
func (r *Recover) NotFound() *Recover {
	r.code = http.StatusNotFound
	return r
}

// InternalServerError
//
//
func (r *Recover) InternalServerError() *Recover {
	r.code = http.StatusInternalServerError
	return r
}

// Recover
//
//
func (r *Recover) Recover(ctx *gin.Context, err interface{}) {
	r.err = getError(err)
	r.recoverWrapper(true, ctx, r.getData(ctx))
}

// RecoverTemplate
//
//
func (r *Recover) RecoverTemplate(ctx *gin.Context, err interface{}, file string, exec tpl.TemplateExecutor) {
	//path := exec.Config().GetRoot() + "/" + file + exec.Config().GetExtension()
	//_ = &FileStack{
	//	File:     path,
	//	Line:     lineNumber(err),
	//	Name:     file,
	//	Contents: fileContents(path),
	//}
}

func (r *Recover) recoverWrapper(useTheme bool, ctx *gin.Context, data interface{}) {
	path, exec, custom := r.resolver(useTheme)

	var b bytes.Buffer
	err := exec.Execute(&b, path, data)
	if err != nil {
		if custom {
			r.recoverWrapper(false, ctx, data)
			return
		}
		log.Fatal(err)
	}

	ctx.Data(r.code, "text/html", b.Bytes())
}

// getError
//
// Converts an interface{} to a Verbis internal error ready
// for processing and to return to the recovery page.
func getError(e interface{}) *errors.Error {
	switch v := e.(type) {
	case errors.Error:
		return &v
	case *errors.Error:
		return v
	case error:
		return &errors.Error{Code: errors.TEMPLATE, Message: v.Error(), Operation: "", Err: v}
	default:
		return &errors.Error{Code: errors.TEMPLATE, Message: "Internal Verbis error, please report", Operation: "", Err: fmt.Errorf("internal verbis error")}
	}
}
