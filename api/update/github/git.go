// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gookit/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Github struct {
	RepoURL     string
	ArchiveName string
	tempDir     string
	archivePath string
	reader      *zip.ReadCloser
}

// tag struct used to unmarshal response from Github
// https://api.github.com/repos/ownerName/projectName/tags
type tag struct {
	Name string `json:"name"`
}

// A FileInfo describes a file given by a provider
type FileInfo struct {
	Path string
	Mode os.FileMode
}

// WalkFunc is the type of the function called for each file or directory
// visited by Walk.
// path is relative
type WalkFunc func(info *FileInfo) error

var (
	ErrNoCheckSum = errors.New("no checksum for archive found")
	tagsUrl       = "https://api.Github.com/repos/ainsleyclark/verbis/tags"
)

// getTags
func (g *Github) getTags() ([]tag, error) {
	//tagsUrl := "https://api.Github.com/repos/" + api.Repo + "/verbis/tags"

	resp, err := http.Get(tagsUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tags []tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// getLatestVersion
func (g *Github) getLatestVersion() (string, error) {
	tags, err := g.getTags()
	if err != nil {
		return "", err
	}

	if len(tags) < 1 {
		return "", errors.New("Verbis has no tags")
	}

	return tags[0].Name, nil
}

func (g *Github) Open() (err error) {
	version, err := g.getLatestVersion()
	if err != nil {
		return
	}

	archiveURL := g.getArchiveURL(version)
	resp, err := http.Get(archiveURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	g.tempDir, err = ioutil.TempDir("", "verbis-updater")
	if err != nil {
		return
	}

	g.archivePath = filepath.Join(g.tempDir, g.ArchiveName)
	archiveFile, err := os.Create(g.archivePath)
	if err != nil {
		return
	}

	_, err = io.Copy(archiveFile, resp.Body)
	archiveFile.Close()
	if err != nil {
		return
	}

	//one := g.checkSum(version)
	//two := ChecksumFileSHA256Bytes(g.archivePath)

	fmt.Println(string(g.checkSum(version)))

	g.reader, err = zip.OpenReader(g.archivePath)
	if err != nil {
		return
	}

	return
}

// Walk
func (g *Github) Walk(walkFn WalkFunc) error {
	if g.reader == nil {
		return errors.New("nil zip.ReadCloser")
	}

	for _, f := range g.reader.File {
		if f != nil {
			err := walkFn(&FileInfo{
				Path: f.Name,
				Mode: f.Mode(),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Github) Close() error {
	if len(g.tempDir) <= 0 {
		return nil
	}

	err := os.RemoveAll(g.tempDir)
	if err != nil {
		return err
	}
	g.tempDir = ""

	return g.reader.Close()
}

// GetArchiveURL
func (g *Github) getArchiveURL(tag string) string {
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", api.Repo, tag, g.ArchiveName)
}

// checkSum

func (g *Github) checkSum(tag string) []byte {
	const op = "Github.CheckSum"

	url := fmt.Sprintf("https://github.com/%s/releases/download/%s/checksums.txt", api.Repo, tag)

	resp, err := http.Get(url)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error obtaining checksum for file: " + g.ArchiveName, Operation: op, Err: ErrNoCheckSum})
		return nil
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "  ")
		if len(line) != 2 {
			continue
		}
		if line[1] == g.ArchiveName {
			color.Red.Println(line[0])
			return []byte(line[0])
		}
	}

	return nil
}
