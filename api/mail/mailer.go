package mail

import (
	"cms/api/environment"
	"cms/api/helpers/html"
	"cms/api/helpers/paths"
	"fmt"
	sp "github.com/SparkPost/gosparkpost"
)

type Mailer struct {
	client sp.Client
	Transmission Sender
}

type Sender struct {
	To		[]string
	Subject	string
	HTML	string
}

type Data map[string]interface{}

// Create a new mailable instance using environment.
func New() (*Mailer, error) {
	m := &Mailer{}

	if err := m.load(); err != nil {
		return &Mailer{}, err
	}

	return m, nil
}

// Load the mailer and connect to sparkpost
func (m *Mailer) load() error {

	config := &sp.Config{
		BaseUrl:    environment.Env.SparkpostUrl,
		ApiKey:     environment.Env.SparkpostApiKey,
		ApiVersion: 1,
	}

	var client sp.Client
	err := client.Init(config)
	if err != nil {
		return fmt.Errorf("SparkPost client init failed: %s\n", err)
	}
	m.client = client

	return nil
}


// Create a Transmission using an inline Recipient List
// and inline email Content.
func (m *Mailer) Send(t *Sender) (string, error) {

	tx := &sp.Transmission{
		Recipients: t.To,
		Content: sp.Content{
			HTML:    t.HTML,
			From:    environment.Env.MailFromAddress,
			Subject: t.Subject,
		},
	}

	id, _, err := m.client.Send(tx)
	if err != nil {
		return id, fmt.Errorf("Mail sending failed: %s\n", err.Error())
	}

	return id, nil
}

// Execute the mail HTML files
func (m *Mailer) ExecuteHTML(file string, data interface{}) (string, error) {
	html, err := html.RenderTemplate("main", data, paths.Api() + "/mail/main-layout.html", paths.Api() + "/mail/" + file)
	if err != nil {
		return "", err
	}
	return html, nil
}