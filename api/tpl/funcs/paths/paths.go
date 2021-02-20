// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

// Base
//
// Returns the base path of the project.
//
// Example: {{ basePath }}
func (ns *Namespace) Base() string {
	return ns.deps.Paths.Base
}

// Ad,om
//
// Returns the admin path of the project.
//
// Example: {{ adminPath }}
func (ns *Namespace) Admin() string {
	return ns.deps.Paths.Admin
}

// API
//
// Returns the API path of the project.
//
// Example: {{ apiPath }}
func (ns *Namespace) API() string {
	return ns.deps.Paths.API
}

// Theme
//
// Returns the Theme path of the project.
//
// Example: {{ themePath }}
func (ns *Namespace) Theme() string {
	return ns.deps.Paths.Theme
}

// Uploads
//
// Returns the Uploads path of the project.
//
// Example: {{ uploadsPath }}
func (ns *Namespace) Uploads() string {
	return ns.deps.Paths.Uploads
}

// Storage
//
// Returns the Storage path of the project.
//
// Example: {{ storagePath }}
func (ns *Namespace) Storage() string {
	return ns.deps.Paths.Storage
}

// Assets
//
// Returns the assets path of the theme.
//
// Example: {{ uploadsPath }}
func (ns *Namespace) Assets() string {
	return ns.deps.Theme.AssetsPath
}

// Templates
//
// Returns the directory where page templates
// are stored.
//
// Example: {{ templatesPath }}
func (ns *Namespace) Templates() string {
	return ns.deps.Paths.Theme + ns.deps.Theme.TemplateDir
}

// Layouts
//
// Returns the directory where page layouts
// are stored.
//
// Example: {{ layoutsPath }}
func (ns *Namespace) Layouts() string {
	return ns.deps.Paths.Theme + ns.deps.Theme.LayoutDir
}
