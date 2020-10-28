package models

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"

	//"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// OptionsRepository defines methods for Options to interact with the database
type OptionsRepository interface {
	Get() (domain.OptionsDB, error)
	GetByName(name string) (interface{}, error)
	GetStruct() (domain.Options, error)
	UpdateCreate(options domain.OptionsDB) error
	Create(name string, value interface{}) error
	Update(name string, value interface{}) error
	Exists(name string) bool
}

// OptionsStore defines the data layer for Posts
type OptionsStore struct {
	db *sqlx.DB
}

// newOptions - Construct
func newOptions(db *sqlx.DB) *OptionsStore {
	return &OptionsStore{
		db: db,
	}
}

// Get all options
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no options available.
func (s *OptionsStore) Get() (domain.OptionsDB, error) {
	const op = "OptionsRepository.Get"

	var o []domain.OptionDB
	if err := s.db.Select(&o, "SELECT * FROM options"); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get options", Operation: op, Err: err}
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
// Returns errors.NOTFOUND if the post was not found by the given name.
func (s *OptionsStore) GetByName(name string) (interface{}, error) {
	const op = "OptionsRepository.GetByName"

	opts, err := s.Get()
	if err != nil {
		return nil, err
	}

	if val, ok := opts[name]; ok {
		return val, nil
	}

	return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get option with the name %s", name), Operation: op, Err: err}
}

// GetStruct gets the options struct for use in the API
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *OptionsStore) GetStruct() (domain.Options, error) {
	const op = "OptionsRepository.GetStruct"

	var opts []domain.OptionDB
	if err := s.db.Select(&opts, "SELECT * FROM options"); err != nil {
		return domain.Options{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not get options", Operation: op, Err: err,}
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

// UpdateCreate update's or create options depending on Exists check
func (s *OptionsStore) UpdateCreate(options domain.OptionsDB) error {
	const op = "OptionsRepository.UpdateCreate"
	for name, value := range options {
		jsonValue, err := s.marshalValue(value)
		if err != nil {
			return err
		}
		if s.Exists(name) {
			if err := s.Update(name, jsonValue); err != nil {
				return err
			}
		} else {
			if err := s.Create(name, jsonValue); err != nil {
				return err
			}
		}
	}
	return nil
}

// Exists checks if the option exists
func (s *OptionsStore) Exists(name string) bool {
	const op = "OptionsRepository.Exists"
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT option_name FROM options WHERE option_name = ?)", name).Scan(&exists)
	return exists
}

// Create the option
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *OptionsStore) Create(name string, value interface{}) error {
	const op = "OptionsRepository.create"
	q := "INSERT INTO options (option_name, option_value) VALUES (?, ?)"
	_, err := s.db.Exec(q, name, value)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the option with the name: %s", name), Operation: op, Err: err}
	}
	return nil
}

// Update the option
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *OptionsStore) Update(name string, value interface{}) error {
	const op = "OptionsRepository.update"
	q := "UPDATE options SET option_name = ?, option_value = ? WHERE option_name = ?"
	_, err := s.db.Exec(q, name, value, name)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the option with the name: %s", name), Operation: op, Err: err}
	}
	return nil
}

// Unmarshal the value
// Returns errors.INTERNAL if the unmarshalling failed
func (s *OptionsStore) unmarshalValue(optValue json.RawMessage) (interface{}, error) {
	const op = "OptionsRepository.unmarshalValue"
	var value interface{}
	if err := json.Unmarshal(optValue, &value); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the option", Operation: op, Err: err}
	}
	return value, nil
}

// Marshal the value
// Returns errors.INTERNAL if the mmarshalling failed
func (s *OptionsStore) marshalValue(optValue interface{}) (json.RawMessage, error) {
	const op = "OptionsRepository.marshalValue"
	m, err := json.Marshal(optValue)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not marshal the option", Operation: op}
	}
	return m, nil
}

