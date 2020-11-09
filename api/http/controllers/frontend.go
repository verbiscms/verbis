package controllers

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/frontend"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/minify"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/helpers/webp"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// FrontendHandler defines methods for the frontend to interact with the server
type FrontendHandler interface {
	GetUploads(g *gin.Context)
	GetAssets(g *gin.Context)
	GetCachedAsset(url string) (*[]byte, *string)
	Serve(g *gin.Context)
}

// FrontendController defines the handler for all frontend routes
type FrontendController struct {
	server          *server.Server
	models 			*models.Store
	config 			config.Configuration
	cacher 			frontend.Cacher
	minify 			minify.Minifier
}

// newFrontend - Construct
func newFrontend(m *models.Store, config config.Configuration) *FrontendController {
	return &FrontendController{
		models: m,
		config: config,
		cacher: frontend.NewCache(m.Options),
		minify: minify.New(m.Options),
	}
}

// GetUploads retrieves images in the uploads folder, returns webp if accepts
func (c *FrontendController) GetUploads(g *gin.Context) {
	const op = "FrontendHandler.GetUploads"

	// Get the base url e.g /uploads/2020/10/test.png
	url := g.Request.URL.Path

	// Check if the file has been cached
	cachedFile, cachedMime := c.GetCachedAsset(url)
	if cachedFile != nil && cachedMime != nil {
		fmt.Println("in cache")
		g.Data(200, *cachedMime, *cachedFile)
	}

	// Set cache headers
	c.cacher.Cache(g)

	// Get the data & mime type from the media store
	file, mimeType, err := c.models.Media.Serve(url, webp.Accepts(g))
	if err != nil {
		c.NoPageFound(g)
		return
	}

	// Set the cache if the app is in production
	defer func() {
		if environment.IsProduction() {
			cache.Store.Set(url, &file, cache.RememberForever)
			cache.Store.Set(url + "mimetype",  &mimeType, cache.RememberForever)
		}
	}()

	// If the minified file is nil or the err is not empty, serve the original data
	minifiedFile, err :=  c.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	if err != nil || minifiedFile != nil {
		file = minifiedFile
	}

	// Return the upload
	g.Data(200, mimeType, file)
}

// GetAssets
//
// It then obtains the assets path from the site model, and then checks
// if the file exists, by opening the file, if it doesnt it will return a
// 404.
// It then sets cache headers using the cacher interface & checks if a webp
// image is available with the path of .jpg.webp. The minify is the used
// to see if the file can be minfied.
func (c *FrontendController) GetAssets(g *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	// Get the base url e.g /assets/images/test.png
	url := g.Request.URL.Path

	// Check if the file has been cached
	cachedFile, cachedMime := c.GetCachedAsset(url)
	if cachedFile != nil && cachedMime != nil {
		g.Data(200, *cachedMime, *cachedFile)
	}

	// Get the options
	options, err := c.models.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	// Get the site config for serving the assets
	theme, err := c.models.Site.GetThemeConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get theme config", Operation: op, Err: err},
		}).Fatal()
	}

	// Get the relevant paths
	assetsPath := paths.Theme() + theme.AssetsPath
	fileName := strings.Replace(url, "/assets", "", 1)
	mimeType := mime.TypeByExtension(strings.Replace(filepath.Ext(fileName), ".", "", 1))

	// Open the file.
	file, err := ioutil.ReadFile(assetsPath + fileName)
	if err != nil {
		c.NoPageFound(g)
		return
	}

	// Set the cache if the app is in production
	defer func() {
		if environment.IsProduction() {
			cache.Store.Set(url, &file, cache.RememberForever)
			cache.Store.Set(url + "mimetype",  &mimeType, cache.RememberForever)
		}
	}()

	// Set cache headers
	c.cacher.Cache(g)

	// Check if the serving of webp's is allowed & get the
	// webp images and assign if not nil
	if options.MediaServeWebP && webp.Accepts(g) {
		webpFile := webp.GetData(g, assetsPath + fileName, mimeType)
		if webpFile != nil {
			mimeType = "image/webp"
			file = webpFile
		}
	}

	// If the minified file is nil or the err is not empty, serve the original data
	minifiedFile, err := c.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	if err != nil || minifiedFile != nil {
		file = minifiedFile
	}

	// Return the asset
	g.Data(200, mimeType, file)
}

// GetCachedAsset checks to see if there is a cached version of the file
// and mimetypes, returns nil for both if nothing was found.
func (c *FrontendController) GetCachedAsset(url string) (*[]byte, *string) {
	if environment.IsProduction() {
		return nil, nil
	}

	file, foundFile := cache.Store.Get(url)
	mime, foundMime := cache.Store.Get(url + "mimetype")

	if foundFile && foundMime {
		file := file.(*[]byte)
		mimeType := mime.(*string)
		return file, mimeType
	}

	return nil, nil
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
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get template data", Operation: op, Err: err},
		}).Fatal()
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

