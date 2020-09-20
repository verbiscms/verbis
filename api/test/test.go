package test

import (
	"cms/api/helpers/paths"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Store Config
)

type Config struct {
	Admin Admin
	Media Media
	Resources Resources
	Template Template
	Theme Theme
}

type Admin struct {
	AdminPath string `yaml:"admin_path"`
	InactiveSessionTime int `yaml:"inactive_session_time"`
}

type Media struct {
	UploadPath string `yaml:"upload_path"`
	AllowedFileTypes []string `yaml:"allowed_file_types"`
}

type Resources struct {
	Configuration map[string]map[string]string `yaml:"resources"`
}

type Template struct {
	AdminPath string `yaml:"file_extension"`
	TemplateDir string `yaml:"template_dir"`
	LayoutDir string `yaml:"layout_dir"`
}

type Theme struct {
	AssetsPath string `yaml:"assets_path"`
	ErrorPageNotFound string `yaml:"404_page"`
}

func GetConfig() {

	// Admin
	a := loadConfig("/admin.yml")
	admin := Admin{}
	if err := yaml.Unmarshal(a, &admin); err != nil {
		log.Error(err)
	}

	// Media
	m := loadConfig("/media.yml")
	media := Media{}
	if err := yaml.Unmarshal(m, &media); err != nil {
		log.Error(err)
	}

	// Resources
	r := loadConfig("/resources.yml")
	resources := Resources{}
	if err := yaml.Unmarshal(r, &resources); err != nil {
		log.Error(err)
	}

	// Resources
	t := loadConfig("/template.yml")
	template := Template{}
	if err := yaml.Unmarshal(t, &template); err != nil {
		log.Error(err)
	}

	// Theme
	th := loadConfig("/theme.yml")
	theme := Theme{}
	if err := yaml.Unmarshal(th, &theme); err != nil {
		log.Error(err)
	}

	Store = Config{
		Admin:     admin,
		Media:     media,
		Resources: resources,
		Template:  template,
		Theme:     theme,
	}

	fmt.Print(Store)
}

// Load the configuration file based on path and return
// []byte for  yaml conversion
func loadConfig (path string) []byte {
	data, err := ioutil.ReadFile(paths.Config() + path)
	if err != nil {
		log.Panic(err)
	}
	return data
}



