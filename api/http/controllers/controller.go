package controllers

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
)

type Controller struct {
	Auth 		AuthHandler
	Categories 	CategoryHandler
	Frontend 	FrontendHandler
	Media 		MediaHandler
	Options 	OptionsHandler
	Posts		PostHandler
	Spa 		SpaHandler
	Site 		SiteHandler
	User 		UserHandler
	server 		*server.Server
}

type RespondJson struct {
	Status 			int 			`json:"status"`
	Error 			bool 			`json:"error"`
	Message			string 			`json:"message"`
	Meta 			Meta			`json:"meta"`
	Data 			interface{} 	`json:"data"`
}

type Meta struct {
	RequestTime 	string           `json:"request_time"`
	ResponseTime 	string           `json:"response_time"`
	LatencyTime 	string           `json:"latency_time"`
	Pagination 		interface{}		 `json:"pagination,omitempty"`
}

type ValidationErrJson struct {
	Errors 		interface{} 		`json:"errors"`
}

// Construct
func New(m *models.Store) (*Controller, error) {

	c := Controller{
		Auth: newAuth(m.Auth, m.Session),
		Categories: newCategories(m.Categories),
		Frontend: newFrontend(m),
		Media: newMedia(m.Media, m.User),
		Options: newOptions(m.Options),
		Posts: newPosts(m.Posts, m.Fields, m.User, m.Categories),
		Spa: newSpa(),
		Site: newSite(m.Site),
		User: newUser(m.User),
	}

	return &c, nil
}

// Main JSON responder.
func Respond(g *gin.Context, status int, message string, data interface{}, pagination ...http.Pagination) {

	// Check the response data
	if d, changed := checkResponseData(g, data); changed {
		data = d
	}

	g.Set("verbis_message", message)

	// If there is no error set the status to 200
	hasError := false
	if status != 200 {
		hasError = true
	}

	// Check if the pagination is empty
	var returnPagination interface{}
	if len(pagination) == 0 {
		returnPagination = nil
	} else {
		returnPagination = pagination[0]
	}

	// Construct meta
	m := calculateRequestTime(g)
	m.Pagination = returnPagination

	// Set up the response JSON
	respond := RespondJson{
		Status: status,
		Message: message,
		Error: hasError,
		Meta: m,
		Data: data,
	}

	// Respond
	g.JSON(status, respond)

	return
}

// Abort with JSON
func AbortJSON(g *gin.Context, status int, message string, data interface{}) {

	// Check the response data
	if d, changed := checkResponseData(g, data); changed {
		data = d
	}

	// If there is no error set the status to 200
	hasError := false
	if status != 200 {
		hasError = true
	}

	// Set up the response JSON
	respond := RespondJson{
		Status: status,
		Error: hasError,
		Message: message,
		Meta: calculateRequestTime(g),
		Data: data,
	}

	// Respond
	g.AbortWithStatusJSON(status, respond)
}

// Handle 404s.
func NoPageFound(g *gin.Context) {
	g.HTML(404, config.Theme.ErrorPageNotFound, gin.H{})
}

// Abort message
func AbortMsg(g *gin.Context, code int, err error) {
	g.String(code, "Oops! Please retry.")
	// A custom error page with HTML templates can be shown by c.HTML()
	//g.HTML(code, "PUT VIEW HRTEW"", gin.H{})
	g.String(code, "Not working", err)
	g.Error(err)
	g.Abort()
}

// checkResponseData checks what type of data is passed and processes it
// accordingly. errors, empty slices & interfaces as well as validation.
// Returns true if the data has changed.
func checkResponseData(g *gin.Context, data interface{}) (interface{}, bool) {

	if data != nil {

		// Get the type of data
		dataType := reflect.TypeOf(data).String()

		// Report to the log if data is an error
		if dataType == "*errors.Error" {
			errData := data.(*errors.Error)
			g.Set("verbis_error", errData)

			if errData.Err != nil {
				errType := reflect.TypeOf(errData.Err).String()
				if errType == "validator.ValidationErrors" && errData.Code == errors.INVALID  {
					validationErrors, _ := errData.Err.(validator.ValidationErrors)
					var v validation.Validator = validation.New()
					data = &ValidationErrJson{
						Errors: v.Process(validationErrors),
					}
					return data, true
				} else {
					return errData, true
				}
			}

			return gin.H{}, true
		}

		// Check if data is nil or an empty slice, if it is return empty object
		if reflect.TypeOf(data).Kind().String() == "slice" {
			s := reflect.ValueOf(data)
			ret := make([]interface{}, s.Len())
			if len(ret) == 0 {
				return gin.H{}, true
			}
		}

		// If data is of type validation errors, pass to validator


		// If the data is type unmarshal error
		if dataType == "*json.UnmarshalTypeError" {
			e, _ := data.(*json.UnmarshalTypeError)
			data = &ValidationErrJson{
				Errors: validation.ValidationError{
					Key:     e.Field,
					Type:    "Unmarshal error",
					Message: "Invalid type passed to " + e.Struct + " struct.",
				},
			}
			return data, true
		}
	}

	return gin.H{}, false
}

// calculateRequestTime processes the request and response time and works out latency time.
// Returns Meta
func calculateRequestTime(g *gin.Context) Meta {
	var startTime time.Time

	if t, exists := g.Get("request_time"); exists {
		startTime = t.(time.Time)
	} else {
		startTime = time.Now()
	}
	latencyTime := time.Since(startTime)

	return Meta{
		RequestTime: startTime.UTC().String(),
		ResponseTime: time.Now().UTC().String(),
		LatencyTime: latencyTime.Round(time.Microsecond).String(),
	}
}
