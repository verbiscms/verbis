package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// controllerTest represents the suite of testing methods for controllers.
type controllerTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine   *gin.Engine
}

// newResponseRecorder - New recorder for testing
// controllers, initalises gin & sets gin mode.
func newResponseRecorder(t *testing.T) *controllerTest {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	gin, engine := gin.CreateTestContext(rr)

	return &controllerTest{
		testing:  t,
		recorder: rr,
		gin:      gin,
		engine:   engine,
	}
}

// runSuccess gets the response data from the recorder and marshalls
// the given data to a string it then asserts that the data matches
// if the status code is 200 & the content type is
// application/json
func (c *controllerTest) RunSuccess(data interface{}) {

	if data == nil {
		assert.Equal(c.testing, 200, c.recorder.Code)
		assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
	}

	got := c.Data()
	want, err := json.Marshal(data)
	if err != nil {
		c.testing.Error(fmt.Sprintf("error marshalling struct %v", err))
	}

	assert.JSONEq(c.testing, string(want), got)
	assert.Equal(c.testing, 200, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// RunInternalError gets the response data from the recorder and asserts
// that the data is empty, if the status code is 500 & the content
// type is application/json
func (c *controllerTest) RunInternalError() {
	got := c.Data()

	assert.JSONEq(c.testing, "{}", got)
	assert.Equal(c.testing, 500, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// RunValidationError gets the response data from the recorder and asserts
// that the data is is empty, if the status code is 400 & the content
// type is application/json
func (c *controllerTest) RunValidationError() {
	b, ok := c.Body()["data"]
	if !ok {
		c.testing.Fatal("no data within the response")
	}

	dict := b.(map[string]interface{})
	_, ok = dict["errors"]
	if !ok {
		c.testing.Error("no errors within the response")
	}

	dictE, ok := b.(map[string]interface{})
	if !ok {
		c.testing.Error("no errors within the response")
	}

	assert.Greater(c.testing, len(dictE), 0)
	assert.Equal(c.testing, 400, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// RunParamError gets the response data from the recorder and asserts
// that the data is is empty, if the status code is 400 & the content
// type is application/json
func (c *controllerTest) RunParamError() {
	got := c.Data()

	assert.JSONEq(c.testing, "{}", got)
	assert.Equal(c.testing, 400, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// Message gets the response message from the body
func (c *controllerTest) Message() string {
	b, ok := c.Body()["message"]
	if !ok {
		c.testing.Error("no message within the response")
	}
	return b.(string)
}

// Body gets the response body from the request
func (c *controllerTest) Body() map[string]interface{} {
	body := map[string]interface{}{}
	if err := json.Unmarshal(c.recorder.Body.Bytes(), &body); err != nil {
		c.testing.Fatal(fmt.Sprintf("error unmarshalling body %v", err))
	}
	return body
}

// getResponseData gets the response body & checks if the data key
// exists and then marshalls the data key to form a string.
//
// Returns a string of the marshalled data
func (c *controllerTest) Data() string {

	b, ok := c.Body()["data"]
	if !ok {
		c.testing.Fatal("no data within the response")
	}

	got, err := json.Marshal(b)
	if err != nil {
		c.testing.Fatal(fmt.Sprintf("error marshalling data %v", err))
	}

	return string(got)
}

// NewRequest makes a new http.Request and assigns the gin testing
// the request.
func (c *controllerTest) NewRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.testing.Fatal(err)
	}
	c.gin.Request = req
}


func (c *controllerTest) RequestAndServe(method string, url string, engineUrl string, body io.Reader,  handler func(ctx *gin.Context)) {
	switch method {
	case "GET": {
		c.engine.GET(engineUrl, handler)
	}
	case "POST": {
		c.engine.POST(engineUrl, handler)
	}
	case "PUT": {
		c.engine.PUT(engineUrl, handler)
	}
	case "DELETE": {
		c.engine.DELETE(engineUrl, handler)
	}
	}
	c.NewRequest(method, url, body)
	c.engine.ServeHTTP(c.recorder, c.gin.Request)
}