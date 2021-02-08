package errors

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
)

// If theres an error executing the template
// it will return the bytes of the error file
// with all the stack info etc
// Its up to the caller to log or do whatever it wants
// to do with the file contents
type Handler struct {
	deps *deps.Deps
}

type Recovery interface {
	Recover(w io.Writer, ctx *gin.Context, err interface{})
}

type Recover struct {
	deps *deps.Deps
	tpl  *resolve
}

func New(d *deps.Deps) *Handler {
	return &Handler{deps: d}
}

func (r *Recover) Recover(w io.Writer, ctx *gin.Context, err interface{}) {
	var b bytes.Buffer

	ok := r.tpl.Exec.Execute(&b, r.tpl.Path, r.GetData(ctx, err))
	if ok != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return
	}

	_, ok = w.Write(b.Bytes())
	if ok != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		return
	}
}


func getError(e interface{}) *errors.Error {
	var err *errors.Error

	switch v := e.(type) {
	case errors.Error:
		err = &v
	case *errors.Error:
		err = v
	case error:
		err = &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template", Operation: "TemplateEngine.Execute", Err: v}
	default:
		err = &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute template", Operation: "TemplateEngine.Execute", Err: fmt.Errorf("templte engine: unable to execute template")}
	}

	return err
}
