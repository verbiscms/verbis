package models

import (
	gojson "encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// SiteRepository defines methods for Posts to interact with the database
type SiteRepository interface {
	GetGlobalConfig() *domain.Site
	GetAllResources() (*[]domain.Resource, error)
	GetAllTemplates() (*domain.Templates, error)
}

// SiteStore defines the data layer for Posts
type SiteStore struct {
	db *sqlx.DB
	optionsRepo domain.Options
	cache siteCache
}

// siteCache defines the options for caching
type siteCache struct {
	Site bool
	Templates bool
	Resources bool
}

// newSite - Construct
func newSite(db *sqlx.DB) *SiteStore {
	s := &SiteStore{
		db: db,
	}

	om := newOptions(db)
	opts, err := om.GetStruct()
	if err != nil {
		log.Fatal(err)
	}
	s.optionsRepo = opts

	// Cache the site config JSON file
	s.cache.Site = s.optionsRepo.CacheSite

	// Cache the templates, preventing reading from the
	// templates path when endpoint is hit.
	s.cache.Templates = s.optionsRepo.CacheTemplates

	// Cache the resources, preventing reading from the
	// config json file in the theme directory.
	s.cache.Resources = s.optionsRepo.CacheResources

	return s
}

// GetGlobalConfig gets the site global config
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

	ds := domain.Site{
		Title:       s.optionsRepo.SiteTitle,
		Description: s.optionsRepo.SiteDescription,
		Logo:        s.optionsRepo.SiteLogo,
		Url:         s.optionsRepo.SiteUrl,
		Version:     api.App.Version,
	}

	// Set the cache for the site config if the cache was not found
	// and the options allow.
	if !found && s.cache.Site {
		cache.Store.Set("site_config", &ds, cache.RememberForever)
	}

	return &ds
}

// Get all templates stored within the template file
// Returns errors.INTERNAL if the template path is invalid.
func (s *SiteStore) GetAllTemplates() (*domain.Templates, error) {
	const op = "SiteRepository.Get"

	// If the cache allows for caching of the site templates &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Templates {
		cached, found := cache.Store.Get("site_templates")
		if found {
			return cached.(*domain.Templates), nil
		}
	}

	files, err := s.walkMatch(paths.Templates(), "*" + config.Template.FileExtension)
	if err != nil {
		return &domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get templates from the path & filextension: %s, %s", paths.Templates(), "*" + config.Template.FileExtension), Operation: op}
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

// GetAllResources gets all resources from the config file
func (s *SiteStore) GetAllResources() (*[]domain.Resource, error) {
	const op = "SiteRepository.GetAllResources"

	// If the cache allows for caching of the site resources &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Resources {
		cached, found := cache.Store.Get("site_resources")
		if found {
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

// getConfig gets the json config
// Returns errors.INVALID if the file could not be located.
// Returns errors.INTERNAL if the file could not be unmarshalled.
func (s *SiteStore) getConfig() (domain.ThemeConfig, error) {
	const op = "SiteRepository.getConfig"

	jsonFile, err := files.ReadJson(paths.Theme() + "/config.json")
	if err != nil {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not get the theme's config file with the path: %s", paths.Theme() + "/config.json"), Operation: op}
	}

	var r domain.ThemeConfig
	err = gojson.Unmarshal(jsonFile, &r)
	if err != nil {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the theme config file", Operation: op}
	}

	return r, nil
}

// walkMatch Walk through root and return array of strings
func (s *SiteStore) walkMatch(root, pattern string) ([]string, error) {
	const op = "SiteRepository.walkMatch"

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