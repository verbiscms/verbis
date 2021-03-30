// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/store/posts"
	"io/ioutil"
	"time"
)

// TODO
//
// - Add images to the pages array by scanning the getField function as well as the <img>'s
// - Split the sitemaps up into 49,999 chunks
// - Add sitemap-index & move to web with XSL.
// - Viewdata & pages should be private, no need to expose.

// SiteMapper represents functions for executing the sitemap data.
type SiteMapper interface {
	Index() ([]byte, error)
	Pages(resource string) ([]byte, error)
	XSL(index bool) ([]byte, error)
	ClearCache()
}

// Sitemap represents the generation of sitemap.xml files for use
// with the sitemap controller.
type Sitemap struct {
	deps         *deps.Deps
	options      *domain.Options
	resources    map[string]domain.Resource
	templatePath string
	indexXSL     string
	resourceXSL  string
}

// index defines the the XML data for rendering a the main (index) sitemap.
type index struct {
	XMLName xml.Name   `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 sitemapindex"`
	Items   []viewItem `xml:"sitemap"`
}

// resources defines the the XML data for rendering a resource sitemap.
type resources struct {
	XMLName           xml.Name   `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	XMLNSImage        string     `xml:"xmlns:image,attr"`
	XSI               string     `xml:"xmlns:xsi,attr"`
	XSISchemaLocation string     `xml:"xsi:schemaLocation,attr"`
	Items             []viewItem `xml:"url"`
}

// TODO! Add images
type image struct {
	Slug      string `xml:"loc"`
	CreatedAt string `xml:"lastmod"`
}

// sitemapViewItem defines the array of posts or items for both
// the index sitemap & resources sitemap.
type viewItem struct {
	Slug      string   `xml:"loc"`
	CreatedAt string   `xml:"lastmod"`
	Image     *[]image `xml:"image"`
}

const (
	// MAPLIMIT defines how many items can be used within a
	// sitemap.xml before splitting into a new one.
	MAPLIMIT = 49999
)

// NewSitemap - Construct
func NewSitemap(d *deps.Deps) *Sitemap {
	const op = "SiteMapper.NewSitemap"

	// TODO: Causing issue! Returning pages to frontend?
	resources := d.Config.Resources
	r := map[string]domain.Resource{}
	for k, v := range resources {
		r[k] = v
	}

	r["pages"] = domain.Resource{
		Name: "pages",
		Slug: "pages",
	}

	s := &Sitemap{
		deps:         d,
		options:      d.Options,
		resources:    r,
		templatePath: d.Paths.Web + "/sitemaps/",
		indexXSL:     "main-sitemap.xsl",
		resourceXSL:  "resource-sitemap.xsl",
	}

	return s
}

// Index
//
// Index first checks to see if the sitemap serving is enabled in the
// options, then goes on to retrieve the pages. Template data is then
// constructed and executed.
//
// Returns errors.CONFLICT if the sitemap serve options was not enabled.
func (s *Sitemap) Index() ([]byte, error) {
	const op = "SiteMapper.GetIndex"

	if !s.options.SeoSitemapServe {
		return nil, &errors.Error{Code: errors.CONFLICT, Message: "Sitemap should not be served due to user options preferences", Operation: op, Err: fmt.Errorf("sitemap could not be served due to preferences")}
	}

	if cached := s.getCachedFile("sitemap-index"); cached != nil {
		return cached, nil
	}

	viewData := index{}

	for _, v := range s.resources {
		posts, err := s.retrievePages(v.Name)
		if err != nil || len(posts) == 0 {
			continue
		}

		viewData.Items = append(viewData.Items, viewItem{
			Slug:      s.options.SiteUrl + "/sitemaps/" + v.Slug + "/sitemap.xml",
			CreatedAt: time.Now().Format(time.RFC3339),
		})
	}

	if s.hasRedirects() {
		viewData.Items = append(viewData.Items, viewItem{
			Slug:      s.options.SiteUrl + "/sitemaps/redirects/sitemap.xml",
			CreatedAt: time.Now().Format(time.RFC3339),
		})
	}

	xmlData, err := s.formatXML(viewData, true)
	if err != nil {
		return nil, err
	}

	go cache.Store.Set("sitemap-index", &xmlData, cache.RememberForever)

	return xmlData, nil
}

// GetXSL reads the main index XSL file from the sitemaps template
// path for use with the sitemap-xml file.
//
// Returns errors.INTERNAL if the ioutil function failed to read the path.
func (s *Sitemap) XSL(index bool) ([]byte, error) {
	const op = "SiteMapper.GeXLS"

	fileName := s.indexXSL
	if !index {
		fileName = s.resourceXSL
	}

	if cached := s.getCachedFile(fileName); cached != nil {
		return cached, nil
	}

	path := s.templatePath + fileName
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read the xsl file with the path: %s", path), Operation: op, Err: err}
	}

	go cache.Store.Set(fileName, &data, cache.RememberForever)

	return data, nil
}

// GetPages first checks to see if the sitemap serving is enabled in the
// options, then goes on to retrieve the pages. Template data is then
// constructed and executed.
//
// Returns errors.CONFLICT if the sitemap serve options was not enabled.
// Returns errors.INTERNAL if the pages template was unable to be executed.
// Returns errors.NOTFOUND if the given resource was not found within the resource or redirects.
func (s *Sitemap) Pages(resource string) ([]byte, error) {
	const op = "SiteMapper.GetPages"

	if !s.options.SeoSitemapServe {
		return nil, &errors.Error{Code: errors.CONFLICT, Message: "Sitemap should not be served due to user options preferences", Operation: op, Err: fmt.Errorf("sitemap could not be served due to preferences")}
	}

	if err := s.canServeResource(resource); err != nil {
		return nil, err
	}

	viewData := resources{}
	viewData.XSI = "http://www.w3.org/2001/XMLSchema-instance"
	viewData.XMLNSImage = "http://www.google.com/schemas/sitemap-image/1.1"
	viewData.XSISchemaLocation = "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd http://www.google.com/schemas/sitemap-image/1.1 http://www.google.com/schemas/sitemap-image/1.1/sitemap-image.xsd"

	if resource == "redirects" {
		viewData.Items = s.getRedirects()
	} else {
		r, err := s.findResourceBySlug(resource)
		if err != nil {
			return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No resource items available with the name: %s", resource), Operation: op, Err: fmt.Errorf("no resource items found")}
		}

		posts, err := s.retrievePages(r.Name)
		if err != nil {
			return nil, err
		}

		if len(posts) == 0 {
			return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No resource items available with the name: %s", resource), Operation: op, Err: fmt.Errorf("no resource items found")}
		}

		for _, v := range posts {
			viewData.Items = append(viewData.Items, viewItem{
				Slug:      v.Slug,
				CreatedAt: time.Now().Format(time.RFC3339),
			})
		}
	}

	xmlData, err := s.formatXML(viewData, true)
	if err != nil {
		return nil, err
	}

	return xmlData, nil
}

func (s *Sitemap) findResourceBySlug(slug string) (domain.Resource, error) {
	for _, v := range s.resources {
		if v.Slug == slug {
			return v, nil
		}
	}
	return domain.Resource{}, fmt.Errorf("not found")
}

// ClearCache - Clears all of the cached data from the index.xml file
// as well as the resources xml files.
//
// Returns no error.
func (s *Sitemap) ClearCache() {
	cache.Store.Delete("sitemap-index")
	for _, v := range s.resources {
		cache.Store.Delete("sitemap-" + v.Name)
	}
}

// getPosts obtains all of the posts for the sitemap in created at
// descending order.
// Returns errors.INTERNAL if the posts could not be retrieved from the store.
func (s *Sitemap) retrievePages(resource string) ([]viewItem, error) {
	const op = "SiteMapper.retrievePages"

	cfg := posts.ListConfig{
		Resource: resource,
		Status:   "published",
	}

	posts, _, err := s.deps.Store.Posts.List(params.Params{
		Page:           1,
		Limit:          0,
		LimitAll:       true,
		OrderDirection: "desc",
		OrderBy:        "created_at",
	}, false, cfg)

	if err != nil {
		return nil, err
	}

	var items []viewItem
	for _, v := range posts {
		if !v.HasResource() {
			resource = "pages"
		} else {
			resource = v.Resource
		}

		exclude := false
		if v.SeoMeta.Seo != nil {
			exclude = v.SeoMeta.Seo.ExcludeSitemap
		}

		if !helpers.StringInSlice(resource, s.options.SeoSitemapExcluded) && !exclude {
			items = append(items, viewItem{
				Slug:      s.options.SiteUrl + v.Permalink,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
			})
		}
	}

	return items, nil
}

// getRedirects first checks to see if the sitemap redirect serving
// is enabled in the options and the sets the view data to the range
// loop.
//
// Returns []sitemapViewItem containing slug & created at date.
func (s *Sitemap) getRedirects() []viewItem {
	const op = "SiteMapper.getRedirects"

	if !s.options.SeoSitemapRedirects {
		return nil
	}

	var data []viewItem
	redirects, _, err := s.deps.Store.Redirects.List(params.Params{LimitAll: true, OrderBy: "created_at", OrderDirection: "desc"})

	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error obtaining site redirects", Operation: op, Err: err}).Error()
		return nil
	}

	for _, v := range redirects {
		data = append(data, viewItem{
			Slug:      v.From,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
		})
	}

	return data
}

// hasRedirects determines if there is any redirects set in the options.
//
// Returns true if found.
func (s *Sitemap) hasRedirects() bool {
	return len(s.getRedirects()) > 0
}

// canServeResource - Determines if the resource passed exists in the
// sitemap struct or if the resource is a redirect & there are
// redirect items to serve.
//
// Returns errors.NOTFOUND if there was no matching resource found.
func (s *Sitemap) canServeResource(resource string) error {
	const op = "SiteMapper.canServeResource"

	found := false
	for _, v := range s.resources {
		if v.Slug == resource {
			found = true
		}
	}

	if resource == "redirects" && s.hasRedirects() {
		found = true
	}

	if !found {
		return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No resource available with the name: %s", resource), Operation: op, Err: fmt.Errorf("no resource found")}
	}

	return nil
}

// getHomeCreatedAt - Get the homepage created at time or now if it
// is not set.
//
//nolint
func (s *Sitemap) getHomeCreatedAt() string {
	home, err := s.deps.Store.Posts.FindBySlug("/")
	createdAt := time.Now().Format(time.RFC3339)
	if err == nil {
		createdAt = home.CreatedAt.Format(time.RFC3339)
	}
	return createdAt
}

// formatXML - Formats the XML []byte passed and adds headers to the
// XML file, if the sitemap is index, a different xsl file header
// will be appended.
//
// Returns back a []byte once formatted.
func (s *Sitemap) formatXML(data interface{}, index bool) ([]byte, error) {
	const op = "SiteMapper.formatXML"

	xmlString, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to execute the sitemap XML", Operation: op, Err: err}
	}

	var b bytes.Buffer
	b.WriteString(xml.Header)

	if index {
		b.WriteString(fmt.Sprintf(`<?xml-stylesheet type="text/xsl" href="%s/main-sitemap.xsl"?>`+"\n", s.options.SiteUrl))
	} else {
		b.WriteString(fmt.Sprintf(`<?xml-stylesheet type="text/xsl" href="%s/resources-sitemap.xsl"?>`+"\n", s.options.SiteUrl))
	}

	b.Write(xmlString)
	b.WriteString("\n")
	b.WriteString("<!-- XML Sitemap generated by Verbis -->")

	return b.Bytes(), nil
}

// getCachedFile -Obtains the cached sitemap xml file by key
//
// Returns [[byte if found or nil.
func (s *Sitemap) getCachedFile(key string) []byte {
	cachedIndex, found := cache.Store.Get(key)
	if found {
		cachedBytes := cachedIndex.(*[]byte)
		return *cachedBytes
	}
	return nil
}
