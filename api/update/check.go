// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package update

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/version"
	hashver "github.com/hashicorp/go-version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type Release struct {
	URL             string    `json:"url"`
	AssetsURL       string    `json:"assets_url"`
	UploadURL       string    `json:"upload_url"`
	HTMLURL         string    `json:"html_url"`
	ID              int       `json:"id"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []Asset   `json:"assets"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Body            string    `json:"body"`
}

type Asset struct {
	URL      string `json:"url"`
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	Uploader struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

// Tag represents a tag object.
// http get to github
// check version number and compare using version.version
// make backup

type Updater struct {
	CurrentVersion *hashver.Version
	GithubVersion  *hashver.Version
	Release        *Release
}

var ErrHashMismatch = errors.New("new file hash mismatch after patch")

func Init() (bool, error) {
	//r, err := GithubVersion()
	//if err != nil {
	//	return false, err
	//}
	//
	//gitv, err := hashver.NewSemver(r.Name)
	//if err != nil {
	//	fmt.Println(err)
	//	return false, err
	//}
	//
	//needsUpdate := gitv.GreaterThan(version.SemVer)
	//if needsUpdate {
	//
	//}


	//verbis_0.0.1_darwin_amd64.zip
	zip := fmt.Sprintf("verbis_%s_%s_%s.zip", version.Version, runtime.GOOS, runtime.GOARCH)
	fmt.Println(zip)
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/ainsleyclark/verbis",
			ArchiveName:   zip,
		},
		Version:       	"v0.0.0",
	}

	u.Provider.Open()

	err := u.Provider.Walk(func(info *provider.FileInfo) error {
		fmt.Println(info.Path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}



	status, err := u.Update()
	if err != nil {
		return false, err
	}

	fmt.Println()


	fmt.Println(status)

	return false, nil
}

const (
	ReleasesURL = "https://api.github.com/repos/ainsleyclark/verbis/releases/latest"
)

func GithubVersion() (*Release, error) {
	resp, err := http.Get(ReleasesURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}

func verifySha(bin []byte, sha []byte) bool {
	h := sha256.New()
	h.Write(bin)
	return bytes.Equal(h.Sum(nil), sha)
}
