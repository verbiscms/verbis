package config

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/environment"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// Global Configuration, sets defaults to ensure that there are no
// empty values within the configuration to prevent any errors.
var (
	Admin = admin{
		Path:                "admin",
		InactiveSessionTime: 60,
	}
	Media = media{
		UploadPath:       "",
		AllowedFileTypes: nil,
	}
	Resources = resources{
		Configuration: map[string]map[string]string{
			"posts": {
				"name": "posts",
				"friendly_name": "Posts",
				"slug": "/posts",
				"icon": "fal fa-newspaper",
			},
		},
	}
	Template = template{
		FileExtension: ".cms",
		TemplateDir:   "templates",
		LayoutDir:     "layouts",
	}
	Theme = theme{
		AssetsPath:        "/assets",
		ErrorPageNotFound: "404",
	}
)

// Admin
type admin struct {
	Path string `yaml:"admin_path,omitempty"`
	InactiveSessionTime int `yaml:"inactive_session_time,omitempty"`
}

// Media
type media struct {
	UploadPath string `yaml:"upload_path"`
	AllowedFileTypes []string `yaml:"allowed_file_types"`
}

// Resources
type resources struct {
	Configuration map[string]map[string]string `yaml:"resources"`
}

// Template
type template struct {
	FileExtension string `yaml:"file_extension"`
	TemplateDir string `yaml:"template_dir"`
	LayoutDir string `yaml:"layout_dir"`
}

// Theme
type theme struct {
	AssetsPath string `yaml:"assets_path"`
	ErrorPageNotFound string `yaml:"404_page"`
}

// Init the configuration, obtain all of the yaml files
// within the config directory and set variables.
func Init() {

	// Admin
	a := loadConfig("/admin.yml")
	if err := yaml.Unmarshal(a, &Admin); err != nil {
		log.Error(err)
	}

	// Media
	m := loadConfig("/media.yml")
	if err := yaml.Unmarshal(m, &Media); err != nil {
		log.Error(err)
	}

	// Resources
	r := loadConfig("/resources.yml")
	if err := yaml.Unmarshal(r, &Resources); err != nil {
		log.Error(err)
	}

	// Resources
	t := loadConfig("/template.yml")
	if err := yaml.Unmarshal(t, &Template); err != nil {
		log.Error(err)
	}

	// Theme
	th := loadConfig("/theme.yml")
	if err := yaml.Unmarshal(th, &Theme); err != nil {
		log.Error(err)
	}
}

// Cache the configuration
func Cache() {
	cache.Store.Set("config_admin", Admin, cache.RememberForever)
	cache.Store.Set("config_media", Media, cache.RememberForever)
	cache.Store.Set("config_resources", Resources, cache.RememberForever)
	cache.Store.Set("config_template", Template, cache.RememberForever)
	cache.Store.Set("config_theme", Theme, cache.RememberForever)
}

// Clear the configuration
func CacheClear() {
	cache.Store.Delete("config_admin")
	cache.Store.Delete("config_media")
	cache.Store.Delete("config_resources")
	cache.Store.Delete("config_template")
	cache.Store.Delete("config_theme")
}

// Load the configuration file based on path and return
// []byte for  yaml conversion
func loadConfig (path string) []byte {
	data, err := ioutil.ReadFile(getConfigPath() + path)
	if err != nil {
		log.Panic(err)
	}
	return data
}

// Get the configuration path of the yaml files
func getConfigPath() string {
	path := ""
	if environment.IsProduction() {
		path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		_, b, _, _ := runtime.Caller(0)
		path = filepath.Join(filepath.Dir(b), "../..")
	}
	return path + "/config"
}
