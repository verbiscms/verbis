// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package location

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// Finder defines the method for obtaining field layouts.
type Finder interface {
	Layout(themePath string, post domain.PostDatum, cacheable bool) domain.FieldGroups
}

// Location defines
type Location struct {
	// Groups defines the current field groups that
	// are saved to disk in storage.
	Groups domain.FieldGroups
}

// FieldPath is where the JSON files reside within the
// theme
var FieldPath = "fields"

// Layout
//
// Obtains layouts specific for the arguments passed. If
// caching allows and the domain.FieldGroups have
// been cached, it will return the cached
// version
func (l *Location) Layout(themePath string, post domain.PostDatum, cacheable bool) domain.FieldGroups {
	// If the cache allows for caching of layouts & if the
	// layout has already been cached, return.
	var found bool
	if cacheable {
		cached, found := cache.Store.Get("field_layout_" + post.UUID.String())
		if found {
			return cached.(domain.FieldGroups)
		}
	}

	fg, err := l.fieldGroupWalker(filepath.Join(themePath, FieldPath))
	if err != nil {
		logger.WithError(err).Error()
	}
	l.Groups = fg

	// Get the groups from the resolver
	groups := l.groupResolver(post)

	// Set the cache field layout if the cache was not found
	if !found && cacheable {
		cache.Store.Set("field_layout_"+post.UUID.String(), groups, cache.RememberForever)
	}

	return groups
}

// groupResolver oops over all of the locations within the config
// json file that is defined. Compares the location sets with
// properties of the post, user and category passed.
// Produces an array of field groups that can be
// returned for the post.
// TODO: cyclomatic complexity 19 of func `(*Location).groupResolver` is high (> 15)
func (l *Location) groupResolver(post domain.PostDatum) domain.FieldGroups { // nolint
	var fg domain.FieldGroups

	// Loop over the groups
	for _, group := range l.Groups {
		// Check for empty locations json
		if len(group.Locations) == 0 && !hasBeenAdded(group.UUID.String(), fg) {
			fg = append(fg, group)
		} else {
			// Check and Loop over locations
			for _, location := range group.Locations {
				if !hasBeenAdded(group.UUID.String(), fg) {
					// Loop over rule sets
					var locationSet []bool
					for _, rule := range location {
						switch rule.Param {
						case "status":
							locationSet = append(locationSet, checkLocation(post.Status, rule))
						case "post":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(post.Id), rule))
						case "page_template":
							locationSet = append(locationSet, checkLocation(post.PageTemplate, rule))
						case "page_layout":
							locationSet = append(locationSet, checkLocation(post.PageLayout, rule))
						case "resource":
							locationSet = append(locationSet, checkLocation(post.Resource, rule))
						case "category":
							if post.Category != nil {
								locationSet = append(locationSet, checkLocation(strconv.Itoa(post.Category.Id), rule))
							} else {
								locationSet = append(locationSet, checkLocation("", rule))
							}
						case "author":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(post.Author.Id), rule))
						case "role":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(post.Author.Role.Id), rule))
						}
					}

					// Remove from the array for the front end
					group.Locations = nil

					if checkMatch(locationSet) {
						fg = append(fg, group)
					}
				}
			}
		}
	}

	// Append empty if nil
	if fg == nil {
		fg = domain.FieldGroups{}
	}

	return fg
}

// fieldGroupWalker will loop over all of the json files that have been
// stored in /storage/fields and append them to the array to be
// returned.
// Returns errors.INTERNAL if the file file not be read or be unmarshalled
func (l *Location) fieldGroupWalker(path string) (domain.FieldGroups, error) {
	const op = "Fields.FieldGroupWalker"

	var fg domain.FieldGroups
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No file or directory with the file: %s", path), Operation: op, Err: err}
		}

		if filepath.Ext(info.Name()) == ".json" {
			file, err := ioutil.ReadFile(path)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read field file with the file: %s", path), Operation: op, Err: err}).Error()
				return nil
			}

			var fields domain.FieldGroup
			err = json.Unmarshal(file, &fields)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Unable to unmarshal the field struct", Operation: op, Err: fmt.Errorf("cannot parse file %s: %s", info.Name(), err.Error())}).Error()
				return nil
			}

			if fields.Fields == nil {
				logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No fields exist with the file: %s", path), Operation: op, Err: fmt.Errorf("layout does not contain any fields")}).Error()
				return nil
			}

			fg = append(fg, fields)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fg, nil
}

// checkLocation checks to see if there has been a match within
// the location json passed, the string could be resource,
// page template or anything defined within fields.
func checkLocation(check string, location domain.FieldLocation) bool {
	var match = false

	switch location.Operator {
	case "==":
		if check == location.Value {
			match = true
		}
	case "!=":
		if check != location.Value {
			match = true
		}
	}

	return match
}

// checkMatch Checks to see if the there has been a match within
// the location json block by using an array of booleans, if
// there has already been a match, it will return false.
func checkMatch(matches []bool) bool {
	for _, a := range matches {
		if !a {
			return false
		}
	}
	return true
}

// hasBeenAdded Checks to see if there already been a match
// within the array of field groups by comparing key and
// the UUID.
func hasBeenAdded(key string, fg domain.FieldGroups) bool {
	for _, v := range fg {
		if v.UUID.String() == key {
			return true
		}
	}
	return false
}
