package mail

import (
	"fmt"
	sp "github.com/SparkPost/gosparkpost"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/html"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
)

type Mailer struct {
	client sp.Client
	Transmission Sender
	FromAddress string
	FromName string
}

type Sender struct {
	To		[]string
	Subject	string
	HTML	string
}

type Data map[string]interface{}

// New, Create a new mailable instance using environment variables.
func New() (*Mailer, error) {
	const op = "mail.New"
	m := &Mailer{}
	if err := m.load(); err != nil {
		return &Mailer{}, err
	}
	return m, nil
}

// Load the mailer and connect to sparkpost
func (m *Mailer) load() error {
	const op = "mail.Load"

	mailConf := environment.GetMailConfiguration()

	config := &sp.Config{
		BaseUrl:    mailConf.SparkpostUrl,
		ApiKey:     mailConf.SparkpostApiKey,
		ApiVersion: 1,
	}

	var client sp.Client
	err := client.Init(config)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not create a new mailer instance", Operation: op, Err: err}
	}
	m.client = client

	m.FromAddress = mailConf.FromAddress
	m.FromName = mailConf.FromName

	return nil
}


// Create a Transmission using an inline Recipient List
// and inline email Content.
func (m *Mailer) Send(t *Sender) (string, error) {
	const op = "mail.Send"

	tx := &sp.Transmission{
		Recipients: t.To,
		Content: sp.Content{
			HTML:    t.HTML,
			From:    m.FromAddress,
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