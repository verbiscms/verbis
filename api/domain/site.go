// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type (
	// Site defines the global Verbis object that is used in
	// the public facing API (without credentials). The
	// version is the version of Verbis the application
	// is currently running.
	Site struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Url         string `json:"url"` //nolint
		Version     string `json:"version"`
	}
	// ThemeConfig defines the data used for unmarshalling the
	// config.yml file found in the theme's directory.
	ThemeConfig struct {
		Theme         Theme               `yaml:"theme" json:"theme"`
		Resources     map[string]Resource `yaml:"resources" json:"resources"`
		AssetsPath    string              `yaml:"assets_path" json:"assets_path"`
		FileExtension string              `yaml:"file_extension" json:"file_extension"`
		TemplateDir   string              `yaml:"template_dir" json:"template_dir"`
		LayoutDir     string              `yaml:"layout_dir" json:"layout_dir"`
		Admin         AdminConfig         `yaml:"admin" json:"admin"`
		Media         MediaConfig         `yaml:"media" json:"media"`
		Editor        Editor              `yaml:"editor" json:"editor"`
	}

	// TODO
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

	AdminConfig struct {
		Path                string `yaml:"admin_path,omitempty" json:"admin_path,omitempty"`
		InactiveSessionTime int    `yaml:"inactive_session_time,omitempty" json:"inactive_session_time,omitempty"`
	}

	MediaConfig struct {
		UploadPath       string   `yaml:"upload_path" json:"upload_path"`
		AllowedFileTypes []string `yaml:"allowed_file_types" json:"allowed_file_types"`
	}

	Themes []Theme

	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

	// Theme defines the information for the currently active
	// theme.
	Theme struct {
		Title       string `yaml:"title" json:"title"`
		Description string `yaml:"description" json:"description"`
		Version     string `yaml:"version" json:"version"`
		Screenshot  string `yaml:"-" json:"screenshot"`
		FileName    string `yaml:"-" json:"-"`
	}
	// Resources defines the slice of resources declared in
	// the theme config.
	Resources struct {
		Resource []Resource `json:"resources"`
	}
	// Resource defines an individual resource or custom post
	// type declared in the theme config.
	Resource struct {
		Name               string   `yaml:"name" json:"name"`
		FriendlyName       string   `yaml:"friendly_name" json:"friendly_name"`
		SingularName       string   `yaml:"singular_name" json:"singular_name"`
		Slug               string   `yaml:"slug" json:"slug"`
		Icon               string   `yaml:"icon" json:"icon"`
		Hidden             bool     `yaml:"hidden" json:"hidden"`
		HideCategorySlug   bool     `yaml:"hide_category_slug" json:"hide_category_slug"`
		AvailableTemplates []string `yaml:"available_templates" json:"available_templates"`
	}
	// Template defines a page template that are available
	// from the theme's template directory.
	Template struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	// Layouts represents the slice of Layout's.
	Templates []Template
	// Layout defines a page layout that are available
	// from the theme's layouts directory.
	Layout struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	// Layouts represents the slice of Layout's.
	Layouts []Layout
	// Editor defines editor options for the admin interface.
	Editor struct {
		Modules []string               `yaml:"modules" json:"modules"`
		Options map[string]interface{} `yaml:"options" json:"options"`
	}
)
