package location

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Finder defines the method for obtaining field layouts.
type Finder interface {
	GetLayout(p domain.Post, a domain.User, c *domain.Category, cacheable bool) []domain.FieldGroup
}

// Location defines
type Location struct {
	// Groups defines the current field groups that
	// are saved to disk in storage.
	Groups []domain.FieldGroup
	// jsonPath defines where JSON files containing
	// domain.FieldGroups are kept
	JsonPath string
}

var (
	// Storage path of the application.
	storagePath = paths.Storage()
)

// NewLocation - Construct
func NewLocation() *Location {
	return &Location{
		JsonPath: storagePath + "/fields",
	}
}

// GetLayout
//
//
// Obtains layouts specific for the arguments passed. If
// caching allows and the domain.FieldGroups have
// been cached, it will return the cached
// version
func (l *Location) GetLayout(p domain.Post, a domain.User, c *domain.Category, cacheable bool) []domain.FieldGroup {

	// If the cache allows for caching of layouts & if the
	// layout has already been cached, return.
	var found bool
	if cacheable {
		cached, found := cache.Store.Get("field_layout_" + p.UUID.String())
		if found {
			return cached.([]domain.FieldGroup)
		}
	}

	fg, err := l.fieldGroupWalker()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
	}
	l.Groups = fg

	// Get the groups from the resolver
	groups := l.groupResolver(p, a, c)

	// Set the cache field layout if the cache was not found
	if !found && cacheable {
		cache.Store.Set("field_layout_"+p.UUID.String(), groups, cache.RememberForever)
	}

	return groups
}

// groupResolver
//
// Loops over all of the locations within the config json file
// that is defined. Compares the location sets with with
// properties of the post, user and category passed.
// Produces an array of field groups that can be
// returned for the post.
func (l *Location) groupResolver(p domain.Post, a domain.User, c *domain.Category) []domain.FieldGroup {
	var fg []domain.FieldGroup

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
							locationSet = append(locationSet, checkLocation(p.Status, rule))
						case "post":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(p.Id), rule))
						case "page_template":
							locationSet = append(locationSet, checkLocation(p.PageTemplate, rule))
						case "page_layout":
							locationSet = append(locationSet, checkLocation(p.PageLayout, rule))
						case "resource":
							if p.Resource != nil {
								locationSet = append(locationSet, checkLocation(*p.Resource, rule))
							} else {
								locationSet = append(locationSet, false)
							}
						case "category":
							if c != nil {
								locationSet = append(locationSet, checkLocation(strconv.Itoa(c.Id), rule))
							} else {
								locationSet = append(locationSet, false)
							}
						case "author":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(p.UserId), rule))
						case "role":
							locationSet = append(locationSet, checkLocation(strconv.Itoa(a.Role.Id), rule))
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
		fg = []domain.FieldGroup{}
	}

	return fg
}

// FieldGroupWalker
//
// This function will loop over all of the json files that have been
// stored in /storage/fields and append them to the array to be
// returned.
//
// Returns errors.INTERNAL if the path file not be read or be unmarshalled
func (l *Location) fieldGroupWalker() ([]domain.FieldGroup, error) {
	const op = "Fields.GetFieldGroups"

	var fg []domain.FieldGroup
	err := filepath.Walk(l.JsonPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No file or directory with the path: %s", path), Operation: op, Err: err}
		}

		if strings.Contains(info.Name(), "json") {

			file, err := ioutil.ReadFile(path)
			if err != nil {
				return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read field file with the path: %s", path), Operation: op, Err: err}
			}

			var fields domain.FieldGroup
			err = json.Unmarshal(file, &fields)
			if err != nil {
				return &errors.Error{Code: errors.INTERNAL, Message: "Unable to unmarshal the field struct", Operation: op, Err: err}
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

// checkLocation
//
// Checks to see if there has been a match within the location
// json passed, the string could be resource. page
// template or anything defined within fields.
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

// checkMatch
//
// Checks to see if the there has been a match within
// the location json block by using an array of booleans, if there
// has already been a match, it will return false. Useful
// for and location json blocks not or.
func checkMatch(matches []bool) bool {
	for _, a := range matches {
		if !a {
			return false
		}
	}
	return true
}

// hasBeenAdded
//
// Checks to see if there already been a match within
// the array of field groups by comparing key and
// the UUID.
func hasBeenAdded(key string, fg []domain.FieldGroup) bool {
	for _, v := range fg {
		if v.UUID.String() == key {
			return true
		}
	}
	return false
}
