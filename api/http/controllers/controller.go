package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"encoding/json"
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
	Time 			string           `json:"request_time"`
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


	// Check if data is nil or an empty slice, if it is return empty object
	if data == nil {
		data = gin.H{}
	} else if reflect.TypeOf(data).Kind().String() == "slice" {
		s := reflect.ValueOf(data)
		ret := make([]interface{}, s.Len())
		if len(ret) == 0 {
			data = gin.H{}
		}
	}

	// Get the type of data
	dataType := reflect.TypeOf(data).String()

	// If data is of type validation errors, pass to validator
	if dataType == "validator.ValidationErrors" {
		validationErrors, _ := data.(validator.ValidationErrors)
		var v validation.Validator = validation.New()
		data = &ValidationErrJson{
			Errors: v.Process(validationErrors),
		}
	}

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
	}

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

	// Set up the response JSON
	respond := RespondJson{
		Status: status,
		Message: message,
		Error: hasError,
		Meta: Meta{
			Time: time.Now().UTC().String(),
			Pagination: returnPagination,
		},
		Data: data,
	}
	g.JSON(status, respond)

	return
}

// Abort with JSON
func AbortJSON(g *gin.Context, status int, message string, data interface{}) {

	// Check if data is nil, if it is return empty object
	if data == nil || reflect.ValueOf(data).IsNil() {
		data = gin.H{}
	}

	hasError := false
	if status != 200 {
		hasError = true
	}

	respond := RespondJson{
		Status: status,
		Error: hasError,
		Message: message,
		Meta: Meta{
			Time: time.Now().UTC().String(),
		},
		Data: data,
	}

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


