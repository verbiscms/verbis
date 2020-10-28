package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ghodss/yaml"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// Global Configuration, sets defaults to ensure that there are no
// empty values within the themes config to prevent any errors.
var (
	themeConfig = domain.ThemeConfig{
		Theme:      domain.Theme{},
		Resources:  nil,
		AssetsPath: "/assets",
		Editor:     domain.Editor{
			Modules: []string{
				"blockquote",
				"code_block",
				"code_block_highlight",
				"hardbreak",
				"h1",
				"h2",
				"h3",
				"h4",
				"h5",
				"h6",
				"paragraph",
				"hr",
				"ul",
				"ol",
				"bold",
				"code",
				"italic",
				"link",
				"strike",
				"underline",
				"history",
				"search",
				"trailing_node",
				"color",
			},
			Options: map[string]interface{}{
				"palette": []string{
					"#4D4D4D", "#999999", "#FFFFFF", "#F44E3B", "#FE9200", "#FCDC00",
					"#DBDF00", "#A4DD00", "#68CCCA", "#73D8FF", "#AEA1FF", "#FDA1FF",
					"#333333", "#808080", "#CCCCCC", "#D33115", "#E27300", "#FCC400",
					"#B0BC00", "#68BC00", "#16A5A5", "#009CE0", "#7B64FF", "#FA28FF",
					"#000000", "#666666", "#B3B3B3", "#9F0500", "#C45100", "#FB9E00",
					"#808900", "#194D33", "#0C797D", "#0062B1", "#653294", "#AB149E",
				},
			},
		},
	}
)

// SiteRepository defines methods for Posts to interact with the database
type SiteRepository interface {
	GetGlobalConfig() *domain.Site
	GetThemeConfig() (domain.ThemeConfig, error)
	GetTemplates() (*domain.Templates, error)
	GetLayouts() (*domain.Layouts, error)
}

// SiteStore defines the data layer for Posts
type SiteStore struct {
	db *sqlx.DB
	config config.Configuration
	optionsModel OptionsRepository
	cache siteCache
}

// siteCache defines the options for caching
type siteCache struct {
	Site bool
	Templates bool
	Resources bool
	Layout bool
}

// newSite - Construct
func newSite(db *sqlx.DB, config config.Configuration) *SiteStore {
	s := &SiteStore{
		db: db,
		config: config,
	}

	om := newOptions(db)
	s.optionsModel = om

	opts, err := s.optionsModel.GetStruct()
	if err != nil {
		log.Fatal(err)
	}

	// Cache the site config JSON file
	s.cache.Site = opts.CacheSite

	// Cache the templates, preventing reading from the
	// templates path when endpoint is hit.
	s.cache.Templates = opts.CacheTemplates

	// Cache the layouts, preventing reading from the
	// layouts path when endpoint is hit.
	s.cache.Layout = opts.CacheLayout

	// Cache the resources, preventing reading from the
	// config json file in the theme directory.
	s.cache.Resources = opts.CacheResources

	return s
}

// GetGlobalConfig gets the site global config
func (s *SiteStore) GetGlobalConfig() *domain.Site {

	// If the cache allows for caching of the site config &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Site && environment.IsProduction() {
		cached, found := cache.Store.Get("site_config")
		if found {
			return cached.(*domain.Site)
		}
	}

	opts, err := s.optionsModel.GetStruct()
	if err != nil {
		log.Fatal(err)
	}

	ds := domain.Site{
		Title:       opts.SiteTitle,
		Description: opts.SiteDescription,
		Logo:        opts.SiteLogo,
		Url:         opts.SiteUrl,
		Version:     api.App.Version,
	}

	// Set the cache for the site config if the cache was not found
	// and the options allow.
	if !found && s.cache.Site {
		cache.Store.Set("site_config", &ds, cache.RememberForever)
	}

	return &ds
}

// Get"s the themes configuration from the themes path
// Returns errors.INTERNAL if the unmarshalling was unsuccessful.
func (s *SiteStore) GetThemeConfig() (domain.ThemeConfig, error) {
	const op = "SiteRepository.GetThemeConfig"

	y, err := files.LoadFile(paths.Theme() + "/config.yml")
	if err != nil {
		return domain.ThemeConfig{}, err
	}
	if err := yaml.Unmarshal(y, &themeConfig); err != nil {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the config.yml file", Operation: op, Err: err}
	}

	return themeConfig, nil
}


// Get all templates stored within the templates directory
// Returns errors.INTERNAL if the template path is invalid.
func (s *SiteStore) GetTemplates() (*domain.Templates, error) {
	const op = "SiteRepository.GetTemplates"

	// If the cache allows for caching of the site templates &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Templates {
		cached, found := cache.Store.Get("site_templates")
		if found {
			return cached.(*domain.Templates), nil
		}
	}

	files, err := s.walkMatch(paths.Templates(), "*" + s.config.Template.FileExtension)
	if err != nil {
		return &domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get templates from the path & file extension: %s, %s", paths.Templates(), "*" + s.config.Template.FileExtension), Operation: op}
	}

	var templates []map[string]interface{}
	templates = append(templates, map[string]interface{}{
		"key": "default",
		"name": "Default",
	})

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

// Get all layouts stored within the layouts directory
// Returns errors.INTERNAL if the layout path is invalid.
func (s *SiteStore) GetLayouts() (*domain.Layouts, error) {
	const op = "SiteRepository.GetLayouts"

	// If the cache allows for caching of the site templates &
	// if the config has already been cached, return.
	var found bool
	if s.cache.Templates && environment.IsProduction() {
		cached, found := cache.Store.Get("site_layouts")
		if found {
			return cached.(*domain.Layouts), nil
		}
	}

	files, err := s.walkMatch(paths.Layouts(), "*" + s.config.Template.FileExtension)
	if err != nil {
		return &domain.Layouts{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get layouts from the path & file extension: %s, %s", paths.Templates(), "*" + s.config.Template.FileExtension), Operation: op}
	}

	var layouts []map[string]interface{}
	layouts = append(layouts, map[string]interface{}{
		"key": "default",
		"name": "Default",
	})

	for _, file := range files {
		name := strings.Title(strings.ToLower(strings.Replace(file, "-", " ", -1)))
		t := map[string]interface{}{
			"key": file,
			"name": name,
		}
		layouts = append(layouts, t)
	}

	t := domain.Layouts{
		Layout: layouts,
	}

	if len(t.Layout) == 0 {
		var m = make([]map[string]interface{}, 0)
		t := domain.Layouts{
			Layout: m,
		}
		return &t, nil
	}

	// Set the cache for the templates if the cache was not found
	// and the options allow.
	if !found && s.cache.Templates {
		cache.Store.Set("site_layouts", &t, cache.RememberForever)
	}

	return &t, nil
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
			str := strings.Split(path, "/")
			template := str[len(str)-1]
			template = strings.Replace(template, s.config.Template.FileExtension, "", -1)
			matches = append(matches, template)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}