package recovery

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type (
	// Data represents the main struct for sending back data to the
	// template for recovery.
	Data struct {
		Error    Error
		Request  Request
		Post     *domain.PostData
		Template *FileStack
		Stack    []*FileStack
	}
	// Error represents a errors.Error in friendly form (strings) to
	// for the recovery template.
	Error struct {
		Code string
		Message string
		Operation string
		Err string
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
func (r *Recover) getData(ctx *gin.Context) *Data {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		body = nil
	}

	return &Data{
		Error: Error{
			Code:      r.err.Code,
			Message:   r.err.Message,
			Operation: r.err.Operation,
			Err:       r.err.Error(),
		},
		Request: Request{
			Url:        location.Get(ctx).String() + ctx.Request.URL.Path,
			Method:     ctx.Request.Method,
			Headers:    ctx.Request.Header,
			Query:      ctx.Request.URL.Query(),
			Body:       string(body),
			Cookies:    ctx.Request.Cookies(),
			IP:         ctx.ClientIP(),
			DataLength: ctx.Writer.Size(),
			UserAgent:  ctx.Request.UserAgent(),
			Referer:    ctx.Request.Referer(),
		},
		Stack: Stack(StackDepth, TraverseLength),
	}
}