// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

package api

// SuperAdmin defines if the app (Verbis) is being developed
// or is being packaged out for distribution.
var SuperAdminString = "true"
var SuperAdmin = true

// App defines default values before the the user has defined
// any custom properties by updating the database.
var App = struct {
	Title       string
	Description string
	Url         string
	Logo        string
	Version     string
}{
	Title:       "Verbis",
	Description: "A Verbis website. Publish online, build a business, work from home",
	Url:         "http://127.0.0.1:8080",
	Logo:        "/verbis/images/verbis-logo.svg",
	Version:     "0.0.1",
}

const (
	// The web route of the API.
	APIRoute = "/api/v1"

	DefaultTheme = ""
)

// 50
var AssetsChan = make(chan int, 10)
var UploadChan = make(chan int, 10)
var ServeChan = make(chan int, 50)
