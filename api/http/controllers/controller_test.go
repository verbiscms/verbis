package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// controllerTest represents the suite of testing methods for controllers.
type controllerTest struct {
	testing *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine *gin.Engine
}

// newResponseRecorder - New recorder for testing
// controllers, initalises gin & sets gin mode.
func newResponseRecorder(t *testing.T) *controllerTest {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	gin, engine := gin.CreateTestContext(rr)

	return &controllerTest{
		testing: t,
		recorder: rr,
		gin:      gin,
		engine: engine,
	}
}

// runSuccess gets the response data from the recorder and marshalls
// the given data to a string it then asserts that the data matches
// if the status code is 200 & the content type is
// application/json
func (c *controllerTest) runSuccess(data interface{}) {

	got := c.getResponseData()

	want, err := json.Marshal(data)
	if err != nil {
		c.testing.Error(fmt.Sprintf("error marshalling struct %v", err))
	}

	assert.JSONEq(c.testing, string(want), got)
	assert.Equal(c.testing, 200, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// runInternalError gets the response data from the recorder and asserts
// that the data is empty, if the status code is 500 & the content
// type is application/json
func (c *controllerTest) runInternalError() {
	got := c.getResponseData()

	assert.JSONEq(c.testing, "{}", got)
	assert.Equal(c.testing, 500, c.recorder.Code)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

// getResponseData nmarshalls the response body, checks if the data key
// exists and then marshalls the data key to form a string.
//
// Returns a string of the marshalled data
func (c *controllerTest) getResponseData() string {

	responseBody := map[string]interface{}{}
	if err := json.Unmarshal(c.recorder.Body.Bytes(), &responseBody); err != nil {
		c.testing.Error(fmt.Sprintf("error unmarshalling body %v", err))
	}

	b, ok := responseBody["data"]
	if !ok {
		c.testing.Error("no data within the response")
	}

	got, err := json.Marshal(b)
	if err != nil {
		c.testing.Error(fmt.Sprintf("error marshalling data %v", err))
	}

	return string(got)
}
