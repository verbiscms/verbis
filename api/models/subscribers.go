package models

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

type SubscriberRepository interface {
	GetAll() ([]domain.Subscriber, error)
	GetById(id int) (domain.Subscriber, error)
	Create(u *domain.Subscriber) (int, error)
	Delete(id int) error
	Send(u *domain.Subscriber) ([]byte, error)
}

type SubscriberStore struct {
	db *sqlx.DB
}

//Construct
func newSubscriber(db *sqlx.DB) *SubscriberStore {
	return &SubscriberStore{
		db: db,
	}
}

// Get all subscribers
func (s *SubscriberStore) GetAll() ([]domain.Subscriber, error) {
	var u []domain.Subscriber
	if err := s.db.Select(&u, "SELECT * FROM subscribers"); err != nil {
		return nil, fmt.Errorf("Could not get subscribers - %w", err)
	}
	if len(u) == 0 {
		return []domain.Subscriber{}, nil
	}
	return u, nil
}

// Get the subscriber by ID
func (s *SubscriberStore) GetById(id int) (domain.Subscriber, error) {
	var u domain.Subscriber
	if err := s.db.Get(&u, "SELECT * FROM subscribers WHERE id = ?", id); err != nil {
		return domain.Subscriber{}, fmt.Errorf("Could not get user with ID %v - %w", id, err)
	}
	return u, nil
}

// Create subscriber
func (s *SubscriberStore) Create(u *domain.Subscriber) (int, error) {
	q := "INSERT INTO subscribers (email, created_at, updated_at) VALUES (?, NOW(), NOW())"
	c, err := s.db.Exec(q, u.Email)

	if err != nil {
		return 0, fmt.Errorf("Could not create the subscriber, %v - %w", u.Email, err)
	}

	id, err := c.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Could not get the newly subscriber user %v ID: %w", u.Email, err)
	}

	return int(id), nil
}


// Delete subscriber
func (s *SubscriberStore) Delete(id int) error {
	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM subscribers WHERE id = ?", id); err != nil {
		return fmt.Errorf("Could not delete user with the ID of %v - %w", id, err)
	}

	return nil
}

// ConvertKit add to subscribers
func (s *SubscriberStore) Send(u *domain.Subscriber) ([]byte, error) {

	postData, err := json.Marshal(map[string]string{
		"email": u.Email,
		"api_key": "HpKBNYGfqTx4M_cdfrwoFA",
	})

	if err != nil {
		return nil, err
	}

	formId := "1601485"
	url := "https://api.convertkit.com/v3/forms/" + formId + "/subscribe"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}


	return body, nil
}


