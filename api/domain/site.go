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
	Resources map[string]Resource `json:"resources"`
}

// Resources
type Resources struct {
	Resource []Resource `json:"resources"`
}

type Resource struct {
	Name    string  `json:"name"`
	Options options `json:"options"`
}

type options struct {
	Slug string `json:"slug"`
	Icon string `json:"icon"`
}

// Templates
type Templates struct {
	Template  []map[string]interface{} `json:"templates"`
}