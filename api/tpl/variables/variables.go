package variables

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
)

type data struct {
	deps *deps.Deps
	ctx  *gin.Context
	post *domain.PostData
}

// Creates a new Funcs
func New(d *deps.Deps) *data {
	return &data{deps: d}
}

func (d *data) Data(ctx *gin.Context, post *domain.PostData) interface{} {
	d.ctx = ctx
	d.post = post
	return d.getData()
}

type (
	// TemplateData represents the main datta to be returned
	// in templates.
	Data struct {
		Site    domain.Site
		Theme   domain.ThemeConfig
		Post    domain.PostData
		Options Options
	}
	// tplOptions represents Verbis options to be returned
	// in templates.
	Options struct {
		Social  Social
		Contact Contact
	}
	// tplSocial represents social details to be returned
	// in templates.
	Social struct {
		Facebook  string
		Twitter   string
		Youtube   string
		LinkedIn  string
		Instagram string
		Pintrest  string
	}
	// Contact represents contact details of the site
	// to be returned in templates.
	Contact struct {
		Email     string
		Telephone string
		Address   string
	}
)

// Nwq - Returns all the necessary data for template usage.
func (d *data) getData() interface{} {
	return Data{
		Site:  d.deps.Site,
		Theme: d.deps.Theme,
		Post:  *d.post,
		Options: Options{
			Social: Social{
				Facebook:  d.deps.Options.SocialFacebook,
				Twitter:   d.deps.Options.SocialTwitter,
				Youtube:   d.deps.Options.SocialYoutube,
				LinkedIn:  d.deps.Options.SocialLinkedIn,
				Instagram: d.deps.Options.SocialInstagram,
				Pintrest:  d.deps.Options.SocialPinterest,
			},
			Contact: Contact{
				Email:     d.deps.Options.ContactEmail,
				Telephone: d.deps.Options.ContactTelephone,
				Address:   d.deps.Options.ContactAddress,
			},
		},
	}
}
