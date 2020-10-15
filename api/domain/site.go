package domain

// Site
type Site struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Url string `json:"url"`
	Version string `json:"version"`
}

// Theme
type ThemeConfig struct {
	Theme Theme `yaml:"theme" json:"theme"`
	Resources map[string]Resource `yaml:"resources" json:"resources"`
	AssetsPath string `yaml:"assets_path" json:"assets_path"`
	Editor Editor `yaml:"editor" json:"editor"`
}

type Theme struct {
	Title string `yaml:"title" json:"title"`
	Description string `yaml:"description" json:"description"`
	Version string `yaml:"version" json:"version"`
}

// Resources
type Resources struct {
	Resource []Resource `json:"resources"`
}

type Resource struct {
	Name string `yaml:"name" json:"name"`
	FriendlyName string `yaml:"friendly_name" json:"friendly_name"`
	SingularName string `yaml:"singular_name" json:"singular_name"`
	Slug string `yaml:"slug" json:"slug"`
	Icon string `yaml:"icon" json:"icon"`
}

// Templates
type Templates struct {
	Template  []map[string]interface{} `json:"templates"`
}

// Layouts
type Layouts struct {
	Layout  []map[string]interface{} `json:"layouts"`
}

// Editor
type Editor struct {
	Modules []string `yaml:"modules" json:"modules"`
	Options map[string]interface{} `yaml:"options" json:"options"`
}
