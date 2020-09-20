package models

import (
	"cms/api/cache"
	"cms/api/domain"
	"cms/api/helpers"
	"cms/api/helpers/paths"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FieldsRepository interface {
	GetFieldGroups() (*[]domain.FieldGroup, error)
	GetLayout(p domain.Post, a domain.User, c []domain.Category) *[]domain.FieldGroup
}

type FieldsStore struct {
	db *sqlx.DB
	optionsModel OptionsRepository
	cache fieldCache
	jsonPath string
}

// Defines the options for caching
type fieldCache struct {
	All bool
	Layout bool
	Fields bool
}

// Construct
func newFields(db *sqlx.DB, om OptionsRepository) *FieldsStore {
	fs := FieldsStore{
		db: db,
		optionsModel: om,
		jsonPath: paths.Storage() + "/fields",
	}
	fs.cache = fs.initCache()
	return &fs
}

// Init the field cache and obtain options for caching the fields
func (s *FieldsStore) initCache() fieldCache {
	fc := fieldCache{}

	// Cache all global
	all, err := s.optionsModel.GetByName("cache")
	if err != nil {
		fc.All = true
	} else {
		fc.All = all.(bool)
	}

	// Cache the layout, this is used for caching the json
	// files when reading the fields for the post
	layout, err := s.optionsModel.GetByName("cache_layout")
	if err != nil {
		layout = true
	} else {
		fc.Layout = layout.(bool)
	}

	// TODO: Implement
	fields, err := s.optionsModel.GetByName("cache_fields")
	if err != nil {
		fields = true
	} else {
		fc.Fields = fields.(bool)
	}

	return fc
}

// GetLayout loops over all of the locations within the config json
// file that is defined. Produces an array of field groups that
// can be returned for the post
func (s *FieldsStore) GetLayout(p domain.Post, a domain.User, c []domain.Category) *[]domain.FieldGroup {
	var fg []domain.FieldGroup

	// If the cache allows for caching of layouts & if the
	// layout has already been cached, return.
	var found bool
	if s.cache.Layout {
		cached, found := cache.Store.Get("field_layout_" + p.UUID.String())
		if found {
			return cached.(*[]domain.FieldGroup)
		}
	}

	// Obtain json files
	fieldGroups, err := s.GetFieldGroups()
	if err != nil {
		log.Error(err)
	}

	// Loop over the groups
	for _, group := range *fieldGroups {

		// Loop over locations
		for _, location := range group.Locations {

			if !s.hasBeenAdded(group.UUID.String(), fg) {

				// Loop over rule sets
				var locationSet []bool
				for _, rule := range location {

					switch rule.Param {
						// Status
						case "status": {
							locationSet = append(locationSet, s.checkLocation(p.Status, rule))
							break
						}
						// Status
						case "post": {
							locationSet = append(locationSet, s.checkLocation(strconv.Itoa(p.Id), rule))
							break
						}
						// Page Templates
						case "page_template": {
							locationSet = append(locationSet, s.checkLocation(p.PageTemplate, rule))
							break
						}
						// Resources
						case "resource": {
							locationSet = append(locationSet, s.checkLocation(*p.Resource, rule))
							break
						}
						// Categories
						case "categories": {
							for _, category := range c {
								locationSet = append(locationSet, s.checkLocation(strconv.Itoa(category.Id), rule))
							}
							break
						}
						// Author
						case "author": {
							locationSet = append(locationSet, s.checkLocation(strconv.Itoa(p.UserId), rule))
							break
						}
						// Role
						case "role": {
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

	// Append empty if nil
	if fg == nil {
		fg = []domain.FieldGroup{}
	}

	// Set the cache field layout if the cache was not found
	if !found && s.cache.Layout {
		cache.Store.Set("field_layout_" + p.UUID.String(), &fg, cache.RememberForever)
	}

	return &fg
}

// GetFieldGroups will loop over all of the json files that have been
// stored in /storage/fields and append them to the array to be
// returned.
func (s *FieldsStore) GetFieldGroups() (*[]domain.FieldGroup, error) {
	var fg []domain.FieldGroup

	err := filepath.Walk(s.jsonPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), "json") {

			file, err := helpers.ReadJson(path)
			if err != nil {
				return err
			}

			var fields domain.FieldGroup
			err = json.Unmarshal(file, &fields)
			if err != nil {
				return err
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

	switch location.Operator  {
		case "==": {
			if check == location.Value {
				match = true
			}
		}
		case "!=": {
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

