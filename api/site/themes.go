// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// The default screenshot name within the theme's
	// directory.
	ScreenshotName = "screenshot"

	// TODO: TEMPORARY
	path = "/Users/ainsley/Desktop/Reddico/apis/verbis/themes"
)

var (
	ScrenshotExtensions = []string{
		".png",
		".svg",
		".jpg",
	}
)

// Themes
//
// TODO
//
func (s *Site) Themes(themePath string) (domain.Themes, error) {
	const op = "SiteRepository.Themes"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	var themes domain.Themes
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		cfg, err := config.Config(path + string(os.PathSeparator) + f.Name())
		if err != nil {
			logger.WithError(err).Error()
		}

		p, err := s.findScreenshot(path, f.Name())
		fmt.Println(p)
		fmt.Println(err)
		themes = append(themes, cfg.Theme)
	}

	return themes, nil
}

func (s *Site) findScreenshot(path, theme string) (string, error) {
	for _, v := range ScrenshotExtensions {
		name := path + string(os.PathSeparator) + theme + string(os.PathSeparator) + ScreenshotName + v
		fmt.Println(name)
		info, err := os.Stat(name)
		if err != nil {
			continue
		}
		fmt.Println(info.Name())
		return "test", nil
	}
	return "", fmt.Errorf("no screenshot found")
}

// Screenshot
//
// TODO
//
func (s *Site) Screenshot(basePath, theme, file string) ([]byte, string, error) {
	const op = "SiteRepository.Screenshot"

	ps := string(os.PathSeparator)
	filePath := basePath + ps + "themes" + ps + theme + ps + file

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: "Error finding screenshot with the path " + file, Operation: op, Err: fmt.Errorf("err")}
	}

	return b, mime.TypeByExtension(filepath.Ext(file)), nil
}
