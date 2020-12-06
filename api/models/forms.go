package models

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/events"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
)

// FormsRepository defines methods for Posts to interact with the database
type FormsRepository interface {
	GetByUUID(uuid string) (domain.Form, error)
	GetValidation(form *domain.Form) dynamicstruct.Builder
	Send(form *domain.Form, ip string, agent string) error
}

// FormsStore defines the data layer for Forms
type FormsStore struct {
	db *sqlx.DB
	config config.Configuration
}

// newSeoMeta - Construct
func newForms(db *sqlx.DB, config config.Configuration) *FormsStore {
	return &FormsStore{
		db: db,
		config: config,
	}
}

// GetByUUID returns a form by UUID.
//
// Returns errors.NOTFOUND if the form was not found by the given UUID.
func (s *FormsStore) GetByUUID(uuid string) (domain.Form, error) {
	const op = "FormsRepository.GetByUUID"

	var f domain.Form
	if err := s.db.Get(&f, "SELECT * FROM forms WHERE uuid = ? LIMIT 1", uuid); err != nil {
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form with the UUID: %s", uuid), Operation: op, Err: err}
	}

	fields, err := s.GetFields(f.Id)
	if err == nil {
		f.Fields = fields
	}

	f.Body = s.GetValidation(&f).Build().New()

	return f, nil
}

// GetFields returns form fields by form ID.
//
// Returns errors.NOTFOUND if there were no fields found by the given form ID.
func (s *FormsStore) GetFields(id int) ([]domain.FormField, error) {
	const op = "FormsRepository.GetFields"
	var f []domain.FormField
	if err := s.db.Select(&f, "SELECT * FROM form_fields WHERE form_id = ?", id); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form fields with the form ID: %v", id), Operation: op, Err: err}
	}
	return f, nil
}

// GetValidation returns the dynamic struct used for validation.
func (s *FormsStore) GetValidation(form *domain.Form) dynamicstruct.Builder {
	instance := dynamicstruct.NewStruct()
	for _, v := range form.Fields {
		tag := fmt.Sprintf(`json:"%s"`, v.Key)
		if v.Required {
			tag = fmt.Sprintf("%s, `binding:\"required\"`", tag)
		}
		instance.AddField(v.Label.Name(), "", tag)
	}
	return instance
}

func (s *FormsStore) Send(form *domain.Form, ip string, agent string) error {
	const op = "FormsRepository.GetFields"
	//form.Reader = dynamicstruct.NewReader(body)
	if form.StoreDB {
		if err := s.storeSubmission(form, ip, agent); err != nil {
			return err
		}
	}
	if form.EmailSend {
		if err := s.mailSubmission(form); err != nil {
			return err
		}
	}
	return nil
}

func (s *FormsStore) mailSubmission(form *domain.Form) error {
	const op = "FormsRepository.mailSubmission"
	fs, err := events.NewFormSend(s.config)
	if err != nil {
		return err
	}
	if err := fs.Send(form); err != nil {
		return err
	}
	return nil
}

func (s *FormsStore) storeSubmission(form *domain.Form, ip string, agent string) error {
	const op = "FormsRepository.storeSubmission"

	fields, err := json.Marshal(form.Body)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not process the form fields for storing"), Operation: op, Err: err}
	}

	_, err = s.db.Exec("INSERT INTO form_submissions (uuid, form_id, fields, ip_address, user_agent, sent_at) VALUES (?, ?, ?, ?, ?, NOW())", uuid.New().String(), form.Id, fields, ip, agent)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the form submission with the ID: %v", form.Id), Operation: op, Err: err}
	}

	return nil
}
