package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// FrontendHandler defines methods for the frontend to interact with the server
type FrontendHandler interface {
	Home(g *gin.Context)
	Test(g *gin.Context)
	GetUploads(g *gin.Context)
	Serve(g *gin.Context)
	Recovery(g *gin.Context, err interface{})
}

// FrontendController defines the handler for all frontend routes
type FrontendController struct {
	server          *server.Server
	models 			*models.Store
}


// newFrontend - Construct
func newFrontend(m *models.Store) *FrontendController {
	return &FrontendController{
		models: m,
	}
}

// TEMP - Home
func (c *FrontendController) Home(g *gin.Context) {
	g.HTML(200, "templates/home", nil)
}

// TEMP - Test
func (c *FrontendController) Test(g *gin.Context) {
	log.WithFields(log.Fields{
		"error": errors.Error{Code: errors.INTERNAL, Message: "Test", Operation: "op", Err: fmt.Errorf("dfdasfdasf")},
	}).Panic()
	return
}

// GetUploads retrieves images in the uploads folder, returns webp if accepts
func (c *FrontendController) GetUploads(g *gin.Context) {
	acceptHeader := g.Request.Header.Get("Accept")
	acceptWebp := strings.Contains(acceptHeader, "image/webp")

	path := g.Request.URL.Path

	data, mime, err := c.models.Media.Serve(path, acceptWebp)
	if err != nil {
		NoPageFound(g)
		return
	}

	g.Data(200, mime, data)
}

// Serve the front end website
func (c *FrontendController) Serve(g *gin.Context) {

	path := g.Request.URL.Path
	post, err := c.models.Posts.GetBySlug(path)

	if err != nil {
		NoPageFound(g)
		return
	}

	type ResourceData struct {
		Post		domain.Post
	}

	r := ResourceData{
		Post: post,
	}

	tf := templates.NewFunctions(g, c.models, &post)
	gvFrontend := goview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: config.Template.FileExtension,
		Master:    config.Template.LayoutDir + "/" + post.Layout,
		Partials:  []string{},
		Funcs: tf.GetFunctions(),
		DisableCache: !environment.IsProduction(),
	})

	if err := gvFrontend.Render(g.Writer, http.StatusOK, config.Template.TemplateDir + "/" + post.PageTemplate, r); err != nil {
		log.Error(err)
	}
}

// Errors
func (c *FrontendController) Recovery(g *gin.Context, err interface{}) {
	const op = "FrontendController.Recovery"

	// Load up the Verbis error pages
	gvRecovery := goview.New(goview.Config{
		Root:      paths.Web(),
		Extension: ".html",
		Master: "layouts/main",
		DisableCache: true,
	})

	// Obtain the error
	var errData error
	errData, ok := err.(error); if !ok {
		if entry, ok := err.(*log.Entry); ok {

			// TODO: Check the type of error here

			fmt.Println(entry.Data["error"])
		}
		return
	}


	// TemplateStack defines
	type TemplateStack struct {
		File string
		Line int
		Name string
		Message string
	}

	// Get the stack
	var stack []TemplateStack
	const stackDepth = 16
	for c := 0; c < stackDepth; c++ {
		t, file, line, ok := runtime.Caller(c)
		if ok {
			stack = append(stack, TemplateStack{
				File: file,
				Line: line,
				Name: runtime.FuncForPC(t).Name(),
			})
		}
	}


	// Get the file contents of the broken file
	var errFile string
	var language string
	tmpl := helpers.StringsBetween(errData.Error(), "name:", ",")
	lineStr := regexp.MustCompile("[0-9]+").FindAllString(errData.Error(), -1)
	var lineNum int
	// Template
	if len(lineStr) > 0 {
		line, err := strconv.Atoi(lineStr[0])
		if err != nil {
			log.Panic(errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not convert %s to int", line), Operation: op, Err: err})
		}
		lineNum = line
		errFile = paths.Theme() + "/" + tmpl + config.Template.FileExtension
		language = "handlebars"
	//Go files
	}  else {
		errFile = stack[0].File
		if len(stack) > 3 {
			lineNum = stack[3].Line
		}
		language = "go"
	}


	// Get the template file contents
	var fileContents string
	if ok := files.Exists(errFile); ok {
		var err error
		if fileContents, err = files.GetFileContents(errFile); err != nil {
			log.Panic(errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not convert get file contents with path %s", errFile), Operation: op, Err: err})
		}
	}

	// Get the lines near where the error happened
	fileLines := files.Lines(fileContents, lineNum, 10)

	// Set the error for the logger & middleware
	vErr := errors.Error{
		Code:      errors.TEMPLATE,
		Message:   fmt.Sprintf("Could not render the template %s", tmpl),
		Operation: "RenderTemplate",
		Err:       fmt.Errorf("%s on line %d", helpers.StringsSplitRight(errData.Error(), "function "), lineNum),
	}
	g.Set("verbis_error", &vErr)

	// Return the error page
	if err := gvRecovery.Render(g.Writer, http.StatusOK, "/templates/error", gin.H{
		"Stack": stack,
		"Message": errData.Error(),
		"RequestMethod": g.Request.Method,
		"File": fileLines,
		"LineNumber": lineNum,
		"FileLanguage": language,
		"Url": g.Request.URL.Path,
		"Ip": g.ClientIP(),
		"DataLength": g.Writer.Size(),
	}); err != nil {
		log.Panic(err)
	}
}







