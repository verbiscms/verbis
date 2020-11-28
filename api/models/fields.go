package models

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	//"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/jmoiron/sqlx"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FieldsRepository defines methods for Posts to interact with the database
type FieldsRepository interface {
	GetFieldGroups() (*[]domain.FieldGroup, error)
	GetLayout(p domain.Post, a domain.User, c *domain.Category) (*[]domain.FieldGroup, error)
}

// FieldsStore defines the data layer for Posts
type FieldsStore struct {
	db       *sqlx.DB
	options  domain.Options
	jsonPath string
}

// newFields - Construct
func newFields(db *sqlx.DB) *FieldsStore {
	const op = "FieldsRepository.newFields"

	fs := FieldsStore{
		db:       db,
		options:  newOptions(db).GetStruct(),
		jsonPath: paths.Storage() + "/fields",
	}

	return &fs
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(p domain.Post, a domain.User, c *domain.Category) (*[]domain.FieldGroup, error) {
	var fg []domain.FieldGroup

	// If the cache allows for caching of layouts & if the
	// layout has already been cached, return.
	var found bool
	if s.options.CacheServerFields {
		cached, found := cache.Store.Get("field_layout_" + p.UUID.String())
		if found {
			return cached.(*[]domain.FieldGroup), nil
		}
	}

	// Obtain json files
	fieldGroups, err := s.GetFieldGroups()
	if err != nil {
		return nil, err
	}

	// Loop over the groups
	for _, group := range *fieldGroups {

		// Check for empty location
		if len(group.Locations) == 0 && !s.hasBeenAdded(group.UUID.String(), fg) {
			fg = append(fg, group)

		} else {

			// Check and Loop over locations
			for _, location := range group.Locations {

				if !s.hasBeenAdded(group.UUID.String(), fg) {

					// Loop over rule sets
					var locationSet []bool
					for _, rule := range location {

						switch rule.Param {
						// Status
						case "status":
							{
								locationSet = append(locationSet, s.checkLocation(p.Status, rule))
								break
							}
						// Status
						case "post":
							{
								locationSet = append(locationSet, s.checkLocation(strconv.Itoa(p.Id), rule))
								break
							}
						// Page Templates
						case "page_template":
							{
								locationSet = append(locationSet, s.checkLocation(p.PageTemplate, rule))
								break
							}
						// Layout
						case "layout":
							{
								locationSet = append(locationSet, s.checkLocation(p.Layout, rule))
								break
							}
						// Resources
						case "resource":
							{
								if p.Resource != nil {
									locationSet = append(locationSet, s.checkLocation(*p.Resource, rule))
								} else {
									locationSet = append(locationSet, false)
								}
								break
							}
						// Categories
						case "categories":
							{
								if c != nil {
									locationSet = append(locationSet, s.checkLocation(strconv.Itoa(c.Id), rule))
								} else {
									locationSet = append(locationSet, false)
								}
								break
							}
						// Author
						case "author":
							{
								locationSet = append(locationSet, s.checkLocation(strconv.Itoa(p.UserId), rule))
								break
							}
						// Role
						case "role":
							{
								locationSet = append(locationSet, s.checkLocation(strconv.Itoa(a.Role.Id), rule))
								break
							}
						}
					}

					// Remove from the array for the front end
					group.Locations = nil

					if s.checkMatch(locationSet) {
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

	// Set the cache field layout if the cache was not found
	if !found && s.options.CacheServerFields {
		cache.Store.Set("field_layout_"+p.UUID.String(), &fg, cache.RememberForever)
	}

	return &fg, nil
}

// GetFieldGroups will loop over all of the json files that have been
// stored in /storage/fields and append them to the array to be
// returned.
// Returns errors.INTERNAL if the path file not be read or be unmarshalled
func (s *FieldsStore) GetFieldGroups() (*[]domain.FieldGroup, error) {
	const op = "FieldsRepository.GetFieldGroups"

	var fg []domain.FieldGroup
	err := filepath.Walk(s.jsonPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), "json") {

			file, err := files.ReadJson(path)
			if err != nil {
				return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read field path with the path: %s", path), Operation: op, Err: err}
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
		return &[]domain.FieldGroup{}, err
	}

	return &fg, nil
}

// checkLocation checks to see if there has been a match within
// the location passed, the string could be resource. page
// template or anything defined within fields.
func (s *FieldsStore) checkLocation(check string, location domain.FieldLocation) bool {
	var match = false

	switch location.Operator {
	case "==":
		{
			if check == location.Value {
				match = true
			}
		}
	case "!=":
		{
			if check != location.Value {
				match = true
			}
		}
	}

	return match
}

// checkMatch checks to see if the there has been a match within
// the location block by using an array of booleans, if there
// has already been a match, it will return false. Useful
// for and location blocks not or.
func (s *FieldsStore) checkMatch(matches []bool) bool {
	for _, a := range matches {
		if !a {
			return false
		}
	}
	return true
}

// hasBeenAdded checks to see if there already been a match within
// the array of field groups.
func (s *FieldsStore) hasBeenAdded(key string, fg []domain.FieldGroup) bool {
	for _, v := range fg {
		if v.UUID.String() == key {
			return true
		}
	}
	return false
}
