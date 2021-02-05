package errors

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Request Request
	Post    *domain.PostData
}

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

func GetRequest(ctx *gin.Context) *Request {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		body = nil
	}
	return &Request{
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
