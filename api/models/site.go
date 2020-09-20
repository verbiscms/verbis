package models

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	gojson "encoding/json"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

type SiteRepository interface {
	GetGlobalConfig() *domain.Site
	GetAllResources() (*[]domain.Resource, error)
	GetAllTemplates() (*domain.Templates, error)
}

type SiteStore struct {
	db *sqlx.DB
	optionsModel OptionsRepository
	cache siteCache
}

// Defines the options for caching
type siteCache struct {
	Site bool
	Templates bool
	Resources bool
}

//Construct
func newSite(db *sqlx.DB, om OptionsRepository) *SiteStore {
	s := &SiteStore{
		db: db,
		optionsModel: om,
	}

	// Cache the site config JSON file
	site, err := om.GetByName("cache_site")
	if err != nil {
		s.cache.Site = true
	} else {
		s.cache.Site = site.(bool)
	}

	// Cache the templates, preventing reading from the
	// templates path when endpoint is hit.
	resources, err := om.GetByName("cache_templates")
	if err != nil {
		s.cache.Resources = true
	} else {
		s.cache.Resources = resources.(bool)
	}

	// Cache the resources, preventing reading from the
	// config json file in the theme directory.
	templates, err := om.GetByName("cache_resources")
	if err != nil {
		s.cache.Templates = true
	} else {
		s.cache.Templates = templates.(bool)
	}

	return s
}

// Get the site config
func (s *SiteStore) GetGlobalConfig() *domain.Site {

	// If the cache allows for caching of the site config &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Site {
		cached, found := cache.Store.Get("site_config")
		if found {
			return cached.(*domain.Site)
		}
	}

	ds := domain.Site{}

	title, err := s.optionsModel.GetByName("site_title")
	if err != nil {
		ds.Title = api.AppTitle
	} else {
		ds.Title = title.(string)
	}

	description, err := s.optionsModel.GetByName("site_description")
	if err != nil {
		ds.Description = api.AppDescription
	} else {
		ds.Description = description.(string)
	}

	logo, err := s.optionsModel.GetByName("site_logo")
	if err != nil {
		ds.Logo = api.AppLogo
	} else {
		ds.Logo = logo.(string)
	}

	url, err := s.optionsModel.GetByName("site_url")
	if err != nil {
		ds.Url = api.AppUrl
	} else {
		ds.Url = url.(string)
	}

	ds.Version = api.AppVersion

	// Set the cache for the site config if the cache was not found
	// and the options allow.
	if !found && s.cache.Site {
		cache.Store.Set("site_config", &ds, cache.RememberForever)
	}

	return &ds
}

// Get all templates stored within the template file
func (s *SiteStore) GetAllTemplates() (*domain.Templates, error) {

	// If the cache allows for caching of the site templates &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Templates {
		cached, found := cache.Store.Get("site_templates")
		if found {
			log.Debug(cached)
			return cached.(*domain.Templates), nil
		}
	}

	files, err := s.walkMatch(paths.Templates(), "*" + config.Template.FileExtension)
	if err != nil {
		return &domain.Templates{}, err
	}

	var templates []map[string]interface{}
	for _, file := range files {
		name := strings.Title(strings.ToLower(strings.Replace(file, "-", " ", -1)))
		t := map[string]interface{}{
			"key": file,
			"name": name,
		}
		templates = append(templates, t)
	}

	t := domain.Templates{
		Template: templates,
	}

	if len(t.Template) == 0 {
		var m = make([]map[string]interface{}, 0)
		t := domain.Templates{
			Template: m,
		}
		return &t, nil
	}

	// Set the cache for the templates if the cache was not found
	// and the options allow.
	if !found && s.cache.Templates {
		cache.Store.Set("site_templates", &t, cache.RememberForever)
	}

	return &t, nil
}

// Get all resources
func (s *SiteStore) GetAllResources() (*[]domain.Resource, error) {

	// If the cache allows for caching of the site resources &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Resources {
		cached, found := cache.Store.Get("site_resources")
		if found {
			log.Debug(cached)
			return cached.(*[]domain.Resource), nil
		}
	}

	jsonResources, err := s.getConfig()
	if err != nil {
		return &[]domain.Resource{}, err
	}

	var resources []domain.Resource
	for k, _ := range jsonResources.Resources {
		r := jsonResources.Resources[k]
		resources = append(resources, r)
	}

	// Set the cache for the templates if the cache was not found
	// and the options allow.
	if !found && s.cache.Resources {
		cache.Store.Set("site_resources", &resources, cache.RememberForever)
	}

	return &resources, err

}

// Get json config
func (s *SiteStore) getConfig() (domain.ThemeConfig, error) {

	// Retrieve the theme JSON file
	jsonFile, err := helpers.ReadJson(paths.Theme() + "/config.json")
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	// Unmarshal the config file into the Config struct
	var r domain.ThemeConfig
	err = gojson.Unmarshal(jsonFile, &r)
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	return r, nil
}

// Walk through root and return array of strings
func (s *SiteStore) walkMatch(root, pattern string) ([]string, error) {
	var matches []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			s := strings.Split(path, "/")
			template := s[len(s)-1]
			template = strings.Replace(template, config.Template.FileExtension, "", -1)
			matches = append(matches, template)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}