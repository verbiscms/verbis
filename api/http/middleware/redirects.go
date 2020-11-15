package middleware

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Redirects(o models.OptionsRepository) gin.HandlerFunc {
	const op = "Middleware.FrontEndCache"
	return func(g *gin.Context) {

		options, err := o.GetStruct()
		if err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
			}).Fatal()
		}

		path := location.Get(g).String() + g.Request.URL.String()
		for _, v := range options.SeoRedirects {
			if path == v.From {
				g.Redirect(v.Code, v.To)
				return
			}
		}

		g.Next()
	}
}

