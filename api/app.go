// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

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
	Version:     "0.0.1",
}

const (
	APIRoute      = "/api/v1"
	AssetsChannel = 10
	UploadChannel = 10
	ServerChannel = 50
	Repo          = "ainsleyclark/verbis"
)

var AssetsChan = make(chan int, AssetsChannel)
var UploadChan = make(chan int, UploadChannel)
var ServeChan = make(chan int, ServerChannel)
