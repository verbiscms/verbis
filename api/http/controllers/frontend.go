package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strings"
)

type FrontendController struct {
	server          *server.Server
	models 			*models.Store
}

type FrontendHandler interface {
	Home(g *gin.Context)
	StyleGuide(g *gin.Context)
	Subscribe(g *gin.Context)
	GetUploads(g *gin.Context)
	Serve(g *gin.Context)
	Recovery(g *gin.Context, err interface{})
}

type ResourceData struct {
	Post		domain.Post
}

// Construct
func newFrontend(m *models.Store) *FrontendController {
	return &FrontendController{
		models: m,
	}
}

// Home
func (c *FrontendController) Home(g *gin.Context) {
	g.HTML(200, "templates/home", gin.H{})
}

// Style Guide
func (c *FrontendController) StyleGuide(g *gin.Context) {
	g.HTML(200, "templates/style-guide", gin.H{})
}

// Subscribe to Newsletter
func (c *FrontendController) Subscribe(g *gin.Context) {
	var subscriber domain.Subscriber
	if err := g.ShouldBindJSON(&subscriber); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	_, err := c.models.Subscriber.Create(&subscriber)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	_, err = c.models.Subscriber.Send(&subscriber)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully inserted subscriber", subscriber)
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

	gvRecovery := goview.New(goview.Config{
		Root:      paths.Web(),
		Extension: ".html",
		Master: "",
		DisableCache: true,
	})

	errData := err.(error)

	type stackError struct {
		File string
		Line int
		Name string
		Message string
	}

	// Get the stack
	var stack []stackError
	const stackDepth = 32
	for c := 2; c < stackDepth; c++ {
		t, file, line, ok := runtime.Caller(c)
		if ok {
			stack = append(stack, stackError{
				File: file,
				Line: line,
				Name: runtime.FuncForPC(t).Name(),
			})
		}
	}

	tmpl := helpers.StringsBetween(errData.Error(), "name:", ",")
	line := regexp.MustCompile("[0-9]+").FindAllString(errData.Error(), -1)
	path := paths.Theme() + "/" + tmpl + config.Template.FileExtension


	if ok := files.Exists(path); ok {
		fmt.Println("Opening a file ")
		var file, err = ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(file))
	}

	fmt.Println(line)

	// Return the error page
	if err := gvRecovery.Render(g.Writer, http.StatusOK, "/templates/error", gin.H{
		"Stack": stack,
		"Message": errData.Error(),
		"RequestMethod": g.Request.Method,
		"Ip": g.ClientIP(),
		"DataLength": g.Writer.Size(),
	}); err != nil {
		log.Panic(err)
	}
}



