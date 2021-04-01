package res

import (
	gojson "encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/jmoiron/sqlx"
	"time"
)

type Resources struct {
	db     *sqlx.DB
	config *Config
}

type Resource struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	FriendlyName string    `db:"friendly_name" json:"friendly_name"`
	Slug         string    `db:"slug" json:"slug"`
	Icon         string    `db:"icon" json:"icon"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Config struct {
	Theme theme `json:"theme"`
	resources
}

type theme struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
}

type resources struct {
	Resource map[string]resource `json:"resources"`
}

type resource struct {
	Name    string  `json:"name"`
	Options options `json:"options"`
}

type options struct {
	Slug string `icon:"name"`
	Icon string `icon:"name"`
}

func NewResources(db *database.DB) (*Resources, error) {

	// Create new resources instance
	r := &Resources{
		db: db.Sqlx,
	}

	// Load config
	c, err := r.Load()
	if err != nil {
		return &Resources{}, err
	}
	r.config = c

	//Purge to database

	return r, nil
}

func (s *Resources) Load() (*Config, error) {

	// Retrieve the theme JSON file
	jsonFile, err := helpers.Read(paths.Theme() + "/config.json")
	if err != nil {
		return &Config{}, err
	}

	// Unmarshal the config file into the Theme struct
	var c Config
	err = gojson.Unmarshal(jsonFile, &c)
	if err != nil {
		fmt.Print("here")
		return &Config{}, err
	}

	// TEST: Pretty print
	pretty, _ := gojson.MarshalIndent(c, "", "\t")
	fmt.Print(string(pretty))

	return &c, nil
}

// Get all resources
func (s *Resources) GetAll() ([]Resource, error) {
	var r []Resource
	if err := s.db.Select(&r, "SELECT * FROM resources"); err != nil {
		return nil, fmt.Errorf("Could not get resources - %w", err)
	}
	return r, nil
}

// Get resource by ID
func (s *Resources) GetById(id int) (Resource, error) {
	var r Resource
	if err := s.db.Get(&r, "SELECT * FROM resources WHERE id = ?", id); err != nil {
		return Resource{}, fmt.Errorf("Could not get resource with the ID %v - %w", id, err)
	}
	return r, nil
}

// Get resource by name
func (s *Resources) GetByName(name string) (Resource, error) {
	var r Resource
	if err := s.db.Get(&r, "SELECT * FROM resources WHERE name = ?", name); err != nil {
		return Resource{}, fmt.Errorf("Could not get resource with the name %v - %w", name, err)
	}
	return r, nil
}

// Update resource
func (s *Resources) Update(r *Resource) error {
	_, err := s.GetById(r.ID)
	if err != nil {
		return err
	}

	// TODO: Update all posts that share a resource

	q := "UPDATE resources SET name = ?, friendly_name = ?, slug = ?, icon = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, r.Name, r.FriendlyName, r.Slug, r.Icon)
	if err != nil {
		return fmt.Errorf("Could not update the resource %v - %w", r.Name, err)
	}

	return nil
}

// Create a resource
func (s *Resources) Create(r *Resource) error {
	q := "INSERT INTO resources (name, friendly_name, slug, icon, updated_at, created_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	_, err := s.db.Exec(q, r.Name, r.FriendlyName, r.Slug, r.Icon)

	if err != nil {
		return fmt.Errorf("Could not create the resource, %v - %w", r.Name, err)
	}

	return nil
}

// Insert all resources into database
func (s *Resources) insertToDB() error {
	for k, v := range s.config.Resource {

		// Check to see if the resource exists
		_, err := s.GetByName(k)
		if err != nil {
			continue
		}

		r := Resource{
			Name:         k,
			FriendlyName: v.Name,
			Slug:         v.Options.Slug,
			Icon:         v.Options.Icon,
		}

		err = s.Create(&r)
		if err != nil {
			return err
		}
	}

	return nil
}

// Dump json file
