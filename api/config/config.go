package config

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
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
	Logs = logs{
		AccessLog: "default",
		ErrorLog:  "default",
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

// Logs 
type logs struct {
	AccessLog string `yaml:"access_log"`
	ErrorLog string `yaml:"error_log"`
}


// Init the configuration, obtain all of the yaml files
// within the config directory and set variables.
// Returns errors.INTERNAL if the unmarshal was unsuccessful.
func Init() error {
	const op = "config.Init"

	// Admin
	a, err := loadConfig("/admin.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(a, &Admin); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the admin.yml file", Operation: op, Err: err}
	}

	// Media
	m, err := loadConfig("/media.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(m, &Media); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the media.yml file", Operation: op, Err: err}
	}

	// Resources
	r, err := loadConfig("/resources.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(r, &Resources); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the resources.yml file", Operation: op, Err: err}
	}

	// Resources
	t, err := loadConfig("/template.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(t, &Template); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the template.yml file", Operation: op, Err: err}
	}

	// Theme
	th, err := loadConfig("/theme.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(th, &Theme); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the template.yml file", Operation: op, Err: err}
	}

	// Logs
	l, err := loadConfig("/logs.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(l, &Logs); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the logs.yml file", Operation: op, Err: err}
	}

	return nil
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

// loadConfig load's the configuration file based on path
// and returns a []byte for yaml conversion
// Returns errors.INTERNAL if the configuration file failed to load.
func loadConfig (path string) ([]byte, error) {
	const op = "config.loadConfig"
	data, err := ioutil.ReadFile(getConfigPath() + path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not load the configuration file with the path: %s", getConfigPath() + path), Operation: op, Err: err}
	}
	return data, nil
}

// getConfigPath obtains the configuration path of the yaml files
func getConfigPath() string {
	const op = "config.getConfigPath"
	path := ""
	if environment.IsProduction() {
		path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		_, b, _, _ := runtime.Caller(0)
		path = filepath.Join(filepath.Dir(b), "../..")
	}
	return path + "/config"
}
