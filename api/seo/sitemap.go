package seo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	"html/template"
	"time"
)

// TODO
//
// - Add images to the pages array by scanning the getField function as well as the <img>'s
// - Split the sitemaps up into 49,999 chunks
// - Add sitemap-index & move to web with XLS.
// - Viewdata & pages should be private, no need to expose.

// SiteMapper represents functions for executing the sitemap data.
type SiteMapper interface {
	GetPages() ([]byte, error)
	GetIndex() ([]byte, error)
}

// Sitemap represents the generation of sitemap.xml files for use
// with the frontend controller
type Sitemap struct {
	models   *models.Store
	options  domain.Options
	resources map[string]domain.Resource
	templatePath string
}

// SitemapPosts defines the array of posts for the sitemap.
type sitemapViewItem struct {
	Slug      string
	CreatedAt string
}

const (
	// MAPLIMIT defines how many items can be used within a
	// sitemap.xml before splitting into a new one.
	MAPLIMIT = 49999
)

// NewSitemap - Construct
func NewSitemap(m *models.Store) *Sitemap {
	const op = "SiteMapper.NewSitemap"

	options, err := m.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	theme, err := m.Site.GetThemeConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get resources", Operation: op, Err: err},
		}).Fatal()
	}

	s := &Sitemap{
		models:  m,
		options: options,
		resources: theme.Resources,
		templatePath: paths.Api() + "/web/sitemaps/",
	}

	return s
}

// GetIndex first checks to see if the sitemap serving is enabled in the
// options, then goes on to retrieve the pages. Template data is then
// constructed and executed.
//
// Returns errors.CONFLICT if the sitemap serve options was not enabled.
func (s *Sitemap) GetIndex() ([]byte, error) {
	const op = "SiteMapper.GetIndex"

	if !s.options.SeoSitemapServe {
		return nil, &errors.Error{Code: errors.CONFLICT, Message: "Sitemap should not be served due to user options preferences", Operation: op, Err: fmt.Errorf("sitemap could not be served due to preferences")}
	}

	var data []sitemapViewItem
	for _, v := range s.resources {
		posts, err := s.retrievePages(v.Name)
		if err != nil || len(posts) == 0 {
			continue
		}

		data = append(data, sitemapViewItem{
			Slug:      s.options.SiteUrl + "/sitemaps" + v.Slug + "/sitemap.xml",
			CreatedAt: time.Now().Format(time.RFC3339),
		})

		if len(posts) > MAPLIMIT {
			// do something
		}
	}
	
	if s.hasRedirects() {
		data = append(data, sitemapViewItem{
			Slug:     s.options.SiteUrl + "/sitemaps/redirects/sitemap.xml",
			CreatedAt: time.Now().Format(time.RFC3339),
		})
	}

	return s.executeTemplate("index.html", map[string]interface{}{
		"Items" : data,
	})
}


// GetPages first checks to see if the sitemap serving is enabled in the
// options, then goes on to retrieve the pages. Template data is then
// constructed and executed.
//
// Returns errors.CONFLICT if the sitemap serve options was not enabled.
// Returns errors.INTERNAL if the pages template was unable to be executed.
// Returns errors.NOTFOUND if the given resource was not found within the resource or redirects.
func (s *Sitemap) GetPages(resource string) ([]byte, error) {
	const op = "SiteMapper.GetPages"

	if !s.options.SeoSitemapServe {
		return nil, &errors.Error{Code: errors.CONFLICT, Message: "Sitemap should not be served due to user options preferences", Operation: op, Err: fmt.Errorf("sitemap could not be served due to preferences")}
	}

	found := false
	for _, v := range s.resources {
		if v.Name == resource {
			found = true
		}
	}

	if resource == "redirects" {
		found = true
	}

	if !found {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No resource available with the name: %s", resource), Operation: op, Err: fmt.Errorf("no resource found")}
	}

	var items []sitemapViewItem
	if resource == "redirects" {
		return s.executeTemplate("resource.html", map[string]interface{}{
			"Home": s.options.SiteUrl,
			"HomeCreatedAt": s.getHomeCreatedAt(),
			"Items": items,
		})
	}

	items, err := s.retrievePages(resource)
	if err != nil {
		return nil, err
	}

	return s.executeTemplate("resource.html", map[string]interface{}{
		"Home": s.options.SiteUrl,
		"HomeCreatedAt": s.getHomeCreatedAt(),
		"Items": items,
	})
}

// getPosts obtains all of the posts for the sitemap in created at
// descending order.
// Returns errors.INTERNAL if the posts could not be retrieved from the store.
func (s *Sitemap) retrievePages(resource string) ([]sitemapViewItem, error) {
	const op = "SiteMapper.retrievePages"

	posts, _, err := s.models.Posts.Get(http.Params{
		Page:           1,
		Limit:          http.PaginationAllLimit,
		OrderDirection: "desc",
		OrderBy:        "created_at",
	}, resource)

	if err != nil {
		return nil, err
	}

	var items []sitemapViewItem
	for _, v := range posts {
		resource := ""
		if v.Resource == nil {
			resource = "pages"
		} else {
			resource = *v.Resource
		}

		exclude := false
		if v.SeoMeta.Seo != nil {
			var seo *domain.PostSeo
			err := json.Unmarshal(*v.SeoMeta.Seo, &seo)
			if err == nil {
				exclude = seo.ExcludeSitemap
			}
		}

		if !helpers.StringInSlice(resource, s.options.SeoSitemapExcluded) && !exclude && v.Status == "published" {
			items = append(items, sitemapViewItem{
				Slug:      s.options.SiteUrl + v.Slug,
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
func (s *Sitemap) getRedirects() []sitemapViewItem {
	var data []sitemapViewItem
	if s.options.SeoSitemapRedirects {
		for _, v := range s.options.SeoRedirects {
			data = append(data, sitemapViewItem{
				Slug:      v.From,
				CreatedAt: time.Now().Format(time.RFC3339),
			})
		}
	}
	return data
}

// hasRedirects determines if there is any redirects set in the options.
//
// Returns true if found.
func (s *Sitemap) hasRedirects() bool {
	return len(s.getRedirects()) > 0
}

// getHomeCreatedAt - Get the homepage created at time or now if it
// is not set.
func (s *Sitemap) getHomeCreatedAt() string {
	home, err := s.models.Posts.GetBySlug("/")
	createdAt := time.Now().Format(time.RFC3339)
	if err == nil {
		createdAt = home.CreatedAt.Format(time.RFC3339)
	}
	return createdAt
}

// ExecuteTemplate - Execute the given file name along with any data passed
// to it.
//
// Returns errors.INTERNAL if the template failed to execute.
func (s *Sitemap) executeTemplate(file string, data interface{}) ([]byte, error) {
	const op = "SiteMapper.ExecuteTemplate"

	t := template.Must(template.New(file).ParseFiles(s.templatePath + file))
	var b bytes.Buffer
	err := t.Execute(&b, data)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to execute sitemap template.", Operation: op, Err: err}
	}

	return b.Bytes(), nil
}