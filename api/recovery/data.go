package recovery

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/gin-contrib/location"
	"io/ioutil"
	"net/http"
)

type (
	// Data represents the main struct for sending back data to the
	// template for recovery.
	Data struct {
		Error    Error
		Template *FileStack
		Request  Request
		Post     *domain.PostData
		Stack    []*FileStack
		Debug    bool
	}
	// Error represents a errors.Error in friendly form (strings) to
	// for the recovery template.
	Error struct {
		Code      string
		Message   string
		Operation string
		Err       string
	}
	// Request represents the data obtained from the context with
	// detailed information about the http request made.
	Request struct {
		Url        string
		Method     string
		Headers    map[string][]string
		Query      map[string][]string
		Body       string
		Cookies    []*http.Cookie
		IP         string
		DataLength int
		UserAgent  string
		Referer    string
	}
)

// getData
//
// Retrieves all the necessary template data to show
// the recovery page. If a template executor has
// been set, the template file stack will be
// retrieved,
func (r *Recover) getData() *Data {

	// Retrieve the request body contents.
	body, err := ioutil.ReadAll(r.config.Context.Request.Body)
	if err != nil {
		body = nil
	}

	// Check if the template exec has been set, if it has
	// retrieve the file stack for the template.
	var tpl *FileStack = nil
	if r.config.TplExec != nil && r.config.TplFile != "" {
		path := r.config.TplExec.Config().GetRoot() + "/" + r.config.TplFile + r.config.TplExec.Config().GetExtension()
		tpl = &FileStack{
			File:     path,
			Line:     tplLineNumber(r.config.Error),
			Name:     r.config.TplFile,
			Contents: tplFileContents(path),
		}
	}

	return &Data{
		Error: Error{
			Code:      r.err.Code,
			Message:   r.err.Message,
			Operation: r.err.Operation,
			Err:       r.err.Error(),
		},
		Template: tpl,
		Request: Request{
			Url:        location.Get(r.config.Context).String() + r.config.Context.Request.URL.Path,
			Method:     r.config.Context.Request.Method,
			Headers:    r.config.Context.Request.Header,
			Query:      r.config.Context.Request.URL.Query(),
			Body:       string(body),
			Cookies:    r.config.Context.Request.Cookies(),
			IP:         r.config.Context.ClientIP(),
			DataLength: r.config.Context.Writer.Size(),
			UserAgent:  r.config.Context.Request.UserAgent(),
			Referer:    r.config.Context.Request.Referer(),
		},
		Post:     r.config.Post,
		Stack:    Stack(StackDepth, TraverseLength),
		Debug: environment.IsDebug(),
	}
}
