// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

// Site
type Site struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
	Version     string `json:"version"`
}

// Theme
type ThemeConfig struct {
	Theme         Theme               `yaml:"theme" json:"theme"`
	Resources     map[string]Resource `yaml:"resources" json:"resources"`
	AssetsPath    string              `yaml:"assets_path" json:"assets_path"`
	FileExtension string              `yaml:"file_extension" json:"file_extension"`
	TemplateDir   string              `yaml:"template_dir" json:"template_dir"`
	LayoutDir     string              `yaml:"layout_dir" json:"layout_dir"`
	Editor        Editor              `yaml:"editor" json:"editor"`
}

type Theme struct {
	Title       string `yaml:"title" json:"title"`
	Description string `yaml:"description" json:"description"`
	Version     string `yaml:"version" json:"version"`
}

// Resources
type Resources struct {
	Resource []Resource `json:"resources"`
}

type Resource struct {
	Name             string `yaml:"name" json:"name"`
	FriendlyName     string `yaml:"friendly_name" json:"friendly_name"`
	SingularName     string `yaml:"singular_name" json:"singular_name"`
	Slug             string `yaml:"slug" json:"slug"`
	Icon             string `yaml:"icon" json:"icon"`
	Hidden           bool   `yaml:"hidden" json:"hidden"`
	HideCategorySlug bool   `yaml:"hide_category_slug" json:"hide_category_slug"`
}

// Templates
type Templates struct {
	Template []map[string]interface{} `json:"templates"`
}

// Layouts
type Layouts struct {
	Layout []map[string]interface{} `json:"layouts"`
}

// Editor
type Editor struct {
	Modules []string               `yaml:"modules" json:"modules"`
	Options map[string]interface{} `yaml:"options" json:"options"`
}
