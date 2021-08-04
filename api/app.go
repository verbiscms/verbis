// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

var (
	// ProductionString defines if the app (Verbis) is being developed
	// or is being packaged out for distribution.
	ProductionString = "false"
	// Production is the shortcut for determining if Verbis is in
	// dev/production.
	Production = false
)

// App defines default values before the the user has defined
// any custom properties by updating the database.
var App = struct {
	Title       string
	Description string
	URL         string
	Logo        string
	Version     string
}{
	Title:       "Verbis",
	Description: "A Verbis website. Publish online, build a business, work from home",
	URL:         "http://127.0.0.1:8080",
	Logo:        "/verbis/images/verbis-logo.svg",
}

const (
	// HTTPAPIRoute is the URI route for incoming and
	// outgoing requests via the API.
	HTTPAPIRoute = "/api/v1"
	// Repo is the URI for the hosted github Verbis
	// repository.
	Repo = "https://github.com/verbiscms/verbis"
	// AdminPath is the Url for serving the Vue SPA
	// backend.
	AdminPath = "/admin"
	// AssetsChannel is the maximum amount of concurrent
	// requests for serving assets on the frontend.
	AssetsChannel = 10
	// UploadChannel is the maximum amount of concurrent
	// requests for serving uploads on the frontend.
	UploadChannel = 10
	// ServerChannel is the maximum amount of concurrent
	// requests for serving posts on the frontend.
	ServerChannel = 50
)
