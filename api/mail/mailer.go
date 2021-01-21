package mail

import (
	"fmt"
	sp "github.com/SparkPost/gosparkpost"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/html"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	log "github.com/sirupsen/logrus"
)

type Mailer struct {
	client       sp.Client
	Config       config.Configuration
	Transmission Sender
	FromAddress  string
	FromName     string
}

type Sender struct {
	To      []string
	Subject string
	HTML    string
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
// Returns errors.INTERNAL if the new mailer instance could not be created
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
// Returns errors.INVALID if the mail failed to send via sparkpost.
func (m *Mailer) Send(t *Sender) {
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
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Mail sending failed: %s", id), Operation: op, Err: err},
		})
		return
	}

	// TODO: Nil pointer dereference here for logging?
	fmt.Println("Email successfully sent")
}

// Execute the mail HTML files
// Returns errors.INTERNAL if the render failed
func (m *Mailer) ExecuteHTML(file string, data interface{}) (string, error) {
	const op = "mail.ExecuteHTML"
	path := paths.Web() + "/mail/" + file
	tmpl, err := html.RenderTemplate("main", data, paths.Web()+"/mail/main-layout.html", path)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to render the template: %s", path), Operation: op, Err: err}
	}
	return tmpl, nil
}
