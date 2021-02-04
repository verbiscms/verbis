package variables

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
)

type (
	// TemplateData represents the main datta to be returned
	// in templates.
	TemplateData struct {
		Site    domain.Site
		Theme   domain.ThemeConfig
		Post    domain.PostData
		Options Options
	}
	// Options represents Verbis options to be returned
	// in templates.
	Options struct {
		Social  Social
		Contact Contact
	}
	// Social represents social details to be returned
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

// Data returns the TemplateData for the front end which are
// bound to posts and the context.
func Data(d *deps.Deps, ctx *gin.Context, post *domain.PostData) interface{} {
	return TemplateData{
		Site:  d.Site,
		Theme: d.Theme,
		Post:  *post,
		Options: Options{
			Social: Social{
				Facebook:  d.Options.SocialFacebook,
				Twitter:   d.Options.SocialTwitter,
				Youtube:   d.Options.SocialYoutube,
				LinkedIn:  d.Options.SocialLinkedIn,
				Instagram: d.Options.SocialInstagram,
				Pintrest:  d.Options.SocialPinterest,
			},
			Contact: Contact{
				Email:     d.Options.ContactEmail,
				Telephone: d.Options.ContactTelephone,
				Address:   d.Options.ContactAddress,
			},
		},
	}
}



