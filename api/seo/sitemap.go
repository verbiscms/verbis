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
	viewData *SitemapViewData
	siteUrl  string
}

// SitemapPosts defines the array of posts for the sitemap.
type SitemapViewItem struct {
	Slug      string
	CreatedAt string
}

// SitemapViewData defines the data to executed on the sitemap.
type SitemapViewData struct {
	Home          string
	HomeCreatedAt string
	Items         []SitemapViewItem
}

var (
	// Template data for the pages sitemap.
	pageTmpl = `<urlset
			xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
					http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
			<url>
				<loc>{{ .Home }}</loc>
				<lastmod>{{ .HomeCreatedAt }}</lastmod>
				<priority>1.00</priority>
			</url>
			{{ range .Pages }}
			<url>
				<loc>{{ .Slug }}</loc>
				<lastmod>{{ .CreatedAt }}</lastmod>
				<priority>0.80</priority>
			</url>
			{{ end }}
			{{ range .Redirects }}
			<url>
				<loc>{{ .Slug }}</loc>
				<lastmod>{{ .CreatedAt }}</lastmod>
				<priority>0.60</priority>
			</url>
			{{ end }}
		</urlset>`
)

// NewSitemap - Construct
func NewSitemap(m *models.Store) *Sitemap {
	const op = "SiteMapper.NewSitemap"

	options, err := m.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: fmt.Errorf("could not get the options struct")},
		}).Fatal()
	}

	s := &Sitemap{
		models:  m,
		options: options,
		siteUrl: options.SiteUrl,
	}

	return s
}


func (s *Sitemap) GetIndex() ([]byte, error) {
	const op = "SiteMapper.GetIndex"

	theme, err := s.models.Site.GetThemeConfig()
	if err != nil {
		return nil, err
	}

	var data []SitemapViewItem
	for _, v := range theme.Resources {
		data = append(data, SitemapViewItem{
			Slug:      v.Slug,
			CreatedAt: time.Now().Format(time.RFC3339),
		})
	}

	fmt.Println(len(data))
	fmt.Println((paths.Api() + "/web/sitemaps/index.html"))

	t := template.Must(template.New("").Parse(paths.Api() + "/web/sitemaps/index.html"))
	var b bytes.Buffer


	err = t.Execute(&b, data)
	if err != nil {
		fmt.Println(err)
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to execute sitemap template.", Operation: op, Err: err}
	}

	fmt.Println("gogle er")

	return nil, nil
}

// GetPages first checks to see if the sitemap serving is enabled in the
// options, then goes on to retrieve the pages. Template data is then
// constructed and executed.
// Returns errors.CONFLICT if the sitemap serve options was not enabled.
// Returns errors.INTERNAL if the pages template was unable to be executed.
func (s *Sitemap) GetPages(resource string) ([]byte, error) {
	const op = "SiteMapper.GetPages"

	if !s.options.SeoSitemapServe {
		return nil, &errors.Error{Code: errors.CONFLICT, Message: "Sitemap should not be served due to user options preferences", Operation: op, Err: fmt.Errorf("sitemap could not be served due to preferences")}
	}

	s.viewData = &SitemapViewData{
		Home:          s.siteUrl,
		HomeCreatedAt: s.getHomeCreatedAt(),
		Items:         make([]SitemapViewItem, 0),
	}

	s.retrieveRedirects()
	err := s.retrievePages(resource)
	if err != nil {
		return nil, err
	}

	t := template.Must(template.New("sitemap").Parse(pageTmpl))
	var b bytes.Buffer
	err = t.Execute(&b, s.viewData)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to execute sitemap template.", Operation: op, Err: err}
	}

	return b.Bytes(), nil
}

// getPosts obtains all of the posts for the sitemap in created at
// descending order.
// Returns errors.INTERNAL if the posts could not be retrieved from the store.
func (s *Sitemap) retrievePages(resource string) error {
	const op = "SiteMapper.getPosts"

	posts, _, err := s.models.Posts.Get(http.Params{
		Page:           1,
		Limit:          http.PaginationAllLimit,
		OrderDirection: "desc",
		OrderBy:        "created_at",
	}, resource)

	if err != nil {
		return err
	}

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
			s.viewData.Items = append(s.viewData.Items, SitemapViewItem{
				Slug:      s.siteUrl + v.Slug,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
			})
		}
	}

	return nil
}

// retrieveRedirects first checks to see if the sitemap redirect serving
// is enabled in the options and the sets the view data to the range
// loop.
func (s *Sitemap) retrieveRedirects() {
	if s.options.SeoSitemapRedirects {
		//for _, v := range s.options.SeoRedirects {
		//	s.viewData.Redirects = append(s.viewData.Redirects, SitemapViewItem{
		//		Slug:      v.From,
		//		CreatedAt: time.Now().Format(time.RFC3339),
		//	})
		//}
	}
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
