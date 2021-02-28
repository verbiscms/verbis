// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"os"
	"path/filepath"
	"strings"
)

// SiteRepository defines methods for Posts to interact with the database
type SiteRepository interface {
	GetGlobalConfig() domain.Site
	GetTemplates() (domain.Templates, error)
	GetLayouts() (domain.Layouts, error)
}

// SiteStore defines the data layer for Posts
type SiteStore struct {
	*StoreConfig
	optionsModel OptionsRepository
	cache        siteCache
}

// siteCache defines the options for caching
type siteCache struct {
	Site      bool
	Templates bool
	Resources bool
	Layout    bool
}

// newSite - Construct
func newSite(cfg *StoreConfig) *SiteStore {
	const op = "SiteRepository.newSite"

	s := &SiteStore{
		StoreConfig: cfg,
	}

	om := newOptions(cfg)
	s.optionsModel = om

	return s
}

// GetGlobalConfig gets the site global config
func (s *SiteStore) GetGlobalConfig() domain.Site {
	const op = "SiteRepository.GetGlobalConfig"

	opts := s.optionsModel.GetStruct()

	return domain.Site{
		Title:       opts.SiteTitle,
		Description: opts.SiteDescription,
		Logo:        opts.SiteLogo,
		Url:         opts.SiteUrl,
		Version:     api.App.Version,
	}
}

// GetTemplates
//
// Get all templates stored within the templates directory
// Returns errors.INTERNAL if the template path is invalid.
func (s *SiteStore) GetTemplates() (domain.Templates, error) {
	const op = "SiteRepository.GetTemplates"

	templateDir := s.Paths.Base + "/theme" + s.Options.Theme() + "/" + s.Config.TemplateDir

	files, err := s.walkMatch(templateDir, "*"+s.Config.FileExtension)
	if err != nil {
		return domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get templates from the path: %s, with the file extension: %s", templateDir, "*"+s.Config.FileExtension), Operation: op}
	}

	var templates []map[string]interface{}
	templates = append(templates, map[string]interface{}{
		"key":  "default",
		"name": "Default",
	})

	for _, file := range files {
		name := strings.Title(strings.ToLower(strings.Replace(file, "-", " ", -1)))
		t := map[string]interface{}{
			"key":  file,
			"name": name,
		}
		templates = append(templates, t)
	}

	t := domain.Templates{
		Template: templates,
	}

	if len(t.Template) == 0 {
		return domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "No page templates available", Err: fmt.Errorf("no page templates available"), Operation: op}
	}

	return t, nil
}

// GetLayouts
//
// Get all layouts stored within the layouts directory
// Returns errors.INTERNAL if the layout path is invalid.
func (s *SiteStore) GetLayouts() (domain.Layouts, error) {
	const op = "SiteRepository.GetLayouts"

	layoutDir := s.Paths.Base + "/theme" + s.Options.Theme() + "/" + s.Config.LayoutDir

	files, err := s.walkMatch(layoutDir, "*"+s.Config.FileExtension)
	if err != nil {
		return domain.Layouts{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get templates from the path: %s, with the file extension: %s", layoutDir, "*"+s.Config.FileExtension), Operation: op}
	}

	var layouts []map[string]interface{}
	layouts = append(layouts, map[string]interface{}{
		"key":  "default",
		"name": "Default",
	})

	for _, file := range files {
		name := strings.Title(strings.ToLower(strings.Replace(file, "-", " ", -1)))
		t := map[string]interface{}{
			"key":  file,
			"name": name,
		}
		layouts = append(layouts, t)
	}

	t := domain.Layouts{
		Layout: layouts,
	}

	if len(t.Layout) == 0 {
		return domain.Layouts{}, &errors.Error{Code: errors.NOTFOUND, Message: "No layouts available", Err: fmt.Errorf("no layouts available"), Operation: op}
	}

	return t, nil
}

// walkMatch
//
// Walk through root and return array of strings
// to the file path.
func (s *SiteStore) walkMatch(root, pattern string) ([]string, error) {
	const op = "SiteRepository.walkMatch"

	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			template := strings.Replace(path, root+"/", "", 1)
			template = strings.Replace(template, s.Config.FileExtension, "", -1)
			matches = append(matches, template)
		}
		return nil
	})

	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to find page templates", Err: err, Operation: op}
	}

	return matches, nil
}
