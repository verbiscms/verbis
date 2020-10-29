package controllers

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/minify"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// FrontendHandler defines methods for the frontend to interact with the server
type FrontendHandler interface {
	GetUploads(g *gin.Context)
	GetAssets(g *gin.Context)
	Serve(g *gin.Context)
}

// FrontendController defines the handler for all frontend routes
type FrontendController struct {
	server          *server.Server
	models 			*models.Store
	config 			config.Configuration
	minify 			minify.Minifier
}

// newFrontend - Construct
func newFrontend(m *models.Store, config config.Configuration) *FrontendController {
	min := minify.New(m.Options)
	return &FrontendController{
		models: m,
		config: config,
		minify: min,
	}
}

// GetUploads retrieves images in the uploads folder, returns webp if accepts
func (c *FrontendController) GetUploads(g *gin.Context) {
	const op = "FrontendHandler.GetUploads"

	acceptHeader := g.Request.Header.Get("Accept")
	acceptWebp := strings.Contains(acceptHeader, "image/webp")

	path := g.Request.URL.Path

	data, mime, err := c.models.Media.Serve(path, acceptWebp)
	if err != nil {
		c.NoPageFound(g)
		return
	}

	g.Data(200, mime, data)
}

func (c *FrontendController) GetAssets(g *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	// Get the site config for serving the assets
	theme, err := c.models.Site.GetThemeConfig()
	if err != nil {
		log.Fatal(err)
	}

	assetsPath := paths.Theme() + theme.AssetsPath
	fileName := strings.Replace(g.Request.URL.Path, "/assets", "", 1)

	file, err := os.Open(assetsPath + fileName)
	if err != nil {
		g.Writer.WriteHeader(http.StatusNotFound)
	}

	defer func() {
		file.Close()
	}()

	mime := mime.TypeByExtension(strings.Replace(filepath.Ext(fileName), ".", "", 1))

	// If the minified file is nil or the err is not empty, serve the original data
	minifiedFile, err := c.minify.Minify(file, mime)
	if err != nil || minifiedFile == nil {
		g.File(assetsPath + fileName)
	}

	g.Data(200, mime, minifiedFile)
}



// Serve the front end website
func (c *FrontendController) Serve(g *gin.Context) {
	const op = "FrontendHandler.Serve"

	path := g.Request.URL.Path
	post, err := c.models.Posts.GetBySlug(path)

	if err != nil {
		c.NoPageFound(g)
		return
	}

	_, err = g.Cookie("verbis-session")
	if err != nil && post.Status != "published" {
		c.NoPageFound(g)
		return
	}

	pt := "index"
	if post.PageTemplate != "default" {
		pt = c.config.Template.TemplateDir + "/" + post.PageTemplate
	}

	master := ""
	if post.Layout != "default" {
		master = c.config.Template.LayoutDir + "/" + post.Layout
	} else {
		pt = pt + c.config.Template.FileExtension
	}

	tf := templates.NewFunctions(g, c.models, &post)
	gvFrontend := goview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: c.models.Config.Template.FileExtension,
		Master:    master,
		Partials:  []string{},
		Funcs: tf.GetFunctions(),
		DisableCache: !environment.IsProduction(),
	})

	data, err := tf.GetData()
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
  	if err := gvFrontend.RenderWriter(&b, pt, data); err != nil {
		// TODO: Panic
  		fmt.Println(err)
  		panic(err)
	}

	minfied, err := c.minify.MinifyBytes(&b, "text/html")
	if err != nil || minfied == nil {
		g.Writer.Write(b.Bytes())
	}

	g.Writer.WriteHeader(200)
	g.Writer.Write(minfied)

}

func (c *FrontendController) NoPageFound(g *gin.Context) {
	gvError := goview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: c.config.Template.FileExtension,
		Partials:  []string{},
		DisableCache: true,
	})
	if err := gvError.Render(g.Writer, 404, "404", nil); err != nil {
		panic(err)
	}
	return
}

