package variables

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
)

type TemplateData struct {
	Site domain.Site
	Theme domain.ThemeConfig
	Post domain.PostData
	Options tplOptions
}

type tplOptions struct {
	Social tplSocial
	Contact tplContact
}

type tplSocial struct {
	Facebook string
	Twitter  string
	Youtube string
	LinkedIn string
	Instagram string
	Pintrest string
}

type tplContact struct {
	Email string
	Telephone string
	Address string
}

// Nwq - Returns all the necessary data for template usage.
func New(d *deps.Deps, post domain.PostData) *TemplateData {
	return &TemplateData{
		Site:    d.Site,
		Theme:   d.Theme,
		Post:    post,
		Options: tplOptions{
			Social: tplSocial{
				Facebook:  d.Options.SocialFacebook,
				Twitter:   d.Options.SocialTwitter,
				Youtube:   d.Options.SocialYoutube,
				LinkedIn:  d.Options.SocialLinkedIn,
				Instagram: d.Options.SocialInstagram,
				Pintrest:  d.Options.SocialPinterest,
			},
			Contact: tplContact{
				Email:     d.Options.ContactEmail,
				Telephone: d.Options.ContactTelephone,
				Address:   d.Options.ContactAddress,
			},
		},
	}
}