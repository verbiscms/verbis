package errors

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Error    *errors.Error
	Request  Request
	Post     domain.PostData
	Template *FileStack
	Stack    []*FileStack
}

func (r *Recover) GetData(ctx *gin.Context, err interface{}) *Data {
	e := getError(err)
	fmt.Println(e)
	return &Data{
		Error:   getError(err),
		Request: getRequest(ctx),
		Stack:   Stack(StackDepth, TraverseLength),
	}
}

// Request represents the data to passed back to the error page.
type Request struct {
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

// GetRequest
//
// Returns important information to the error page
// about the request made.
func getRequest(ctx *gin.Context) Request {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		body = nil
	}
	return Request{
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
	}
}
