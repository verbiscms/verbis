// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"os"
	"path/filepath"
)

// Paths represent the struct of paths for use with the
/// application.
type Paths struct {
	Base    string
	Admin   string
	API     string
	Uploads string
	Storage string
	Themes  string
	Web     string
	Forms   string
	Bin     string
}

const (
	// Admin defines the file path for the Vue SPA.
	Admin = string(os.PathSeparator) + "admin"
	// API defines the file path backend code.
	API = string(os.PathSeparator) + "api"
	// Storage defines the file path for uploads, forms and
	// anything that needs to be stored within Verbis.
	Storage = string(os.PathSeparator) + "storage"
	// Themes defines the file path for all themes.
	Themes = string(os.PathSeparator) + "themes"
	// Web defines the file path for any web files that the API
	// needs to serve.
	Web = API + string(os.PathSeparator) + "www"
	// Uploads defines the file path for media uploads within
	// Verbis.
	Uploads = Storage + string(os.PathSeparator) + "uploads"
	// Forms defines the file path for form dumps within Verbis.
	Forms = Storage + string(os.PathSeparator) + "forms"
	// Bin defines the file path any independent executables.
	Bin = string(os.PathSeparator) + "bin"
)

// Get retrieves relevant paths for the application.
func Get() Paths {
	base := base()
	return Paths{
		Base:    base,
		Admin:   base + Admin,
		API:     base + API,
		Uploads: base + Uploads,
		Storage: base + Storage,
		Themes:  base + Themes,
		Web:     base + Web,
		Forms:   base + Forms,
		Bin:     base + Bin,
	}
}

// filepath.Abs is the stdlib Absolute function for
// obtaining the base path of the project
var abs = filepath.Abs

// base returns base path of project.
func base() string {
	dir, err := abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
