package models

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"

	//"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type OptionsRepository interface {
	GetAll() (domain.OptionsDB, error)
	GetByName(name string) (interface{}, error)
	GetStruct() (domain.Options, error)
	UpdateCreate(options domain.OptionsDB) error
	Exists(name string) bool
}

type OptionsStore struct {
	db *sqlx.DB
}

//Construct
func newOptions(db *sqlx.DB) *OptionsStore {
	return &OptionsStore{
		db: db,
	}
}

// Get all options
func (s *OptionsStore) GetAll() (domain.OptionsDB, error) {
	var o []domain.OptionDB
	if err := s.db.Select(&o, "SELECT * FROM options"); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not get options")
	}

	opts := make(domain.OptionsDB)
	for _, v := range o {
		unValue, err := s.unmarshalValue(v.Value)
		if err != nil {
			return domain.OptionsDB{}, err
		}
		opts[v.Name] = unValue
	}

	return opts, nil
}

// Get by name
func (s *OptionsStore) GetByName(name string) (interface{}, error) {
	opts, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	if val, ok := opts[name]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("Option value has not been found with the name: %v:", name)
}

// Get the options struct for use in the API
func (s *OptionsStore) GetStruct() (domain.Options, error) {
	var opts []domain.OptionDB
	if err := s.db.Select(&opts, "SELECT * FROM options"); err != nil {
		return domain.Options{}, err
	}

	unOpts := make(domain.OptionsDB)
	for _, v := range opts {
		unValue, err := s.unmarshalValue(v.Value)
		if err != nil {
			return domain.Options{}, err
		}
		unOpts[v.Name] = unValue
	}

	mOpts, err := json.Marshal(unOpts)
	if err != nil {
		return domain.Options{}, err
	}

	var options domain.Options
	if err := json.Unmarshal(mOpts, &options); err != nil {
		return domain.Options{}, err
	}

	return options, nil
}

// Update or create options
func (s *OptionsStore) UpdateCreate(options domain.OptionsDB) error {
	for name, value := range options {
		jsonValue, err := s.marshalValue(value)
		if err != nil {
			return err
		}
		if s.Exists(name) {
			if err := s.update(name, jsonValue); err != nil {
				return err
			}
		} else {
			if err := s.create(name, jsonValue); err != nil {
				return err
			}
		}
	}
	return nil
}

// Check if the option exists
func (s *OptionsStore) Exists(name string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT option_name FROM options WHERE option_name = ?)", name).Scan(&exists)
	return exists
}

// Create the option
func (s *OptionsStore) create(name string, value interface{}) error {
	q := "INSERT INTO options (option_name, option_value) VALUES (?, ?)"
	_, err := s.db.Exec(q, name, value)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not create the option record with the name: %v", name)
	}
	return nil
}

// Update the option
func (s *OptionsStore) update(name string, value interface{}) error {
	q := "UPDATE options SET option_name = ?, option_value = ? WHERE option_name = ?"
	_, err := s.db.Exec(q, name, value, name)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update the option with the name: %v", name)
	}
	return nil
}

// Unmarshal the value
func (s *OptionsStore) unmarshalValue(optValue json.RawMessage) (interface{}, error) {
	var value interface{}
	if err := json.Unmarshal(optValue, &value); err != nil {
		log.Error(err)
		return nil, err
	}
	return value, nil
}

// Marshal the value
func (s *OptionsStore) marshalValue(optValue interface{}) (json.RawMessage, error) {
	return json.Marshal(optValue)
}

