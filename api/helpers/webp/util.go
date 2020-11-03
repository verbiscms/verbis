package webp

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

// Accepts checks to see if the browser accepts WebP images
func Accepts(g *gin.Context) bool {
	acceptHeader := g.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, "image/webp")
}

// GetData first checks to see if the browser accepts WebP images
// and if the mime type is jpg or a png.
// Returns a data was found, nil if it hasn't.
func GetData(g *gin.Context, path string, mime string) []byte {
	if Accepts(g) && mime == "image/jpeg" || mime == "image/png" {
		data, found := ioutil.ReadFile(path + ".webp")
		if found != nil {
			return nil
		}
		return data
	}
	return nil
}

// Converts an image to webp based on compression and decoded image.
// Compression level is also set.
func Convert(path string, compression int) {
	const op = "Webp.Convert"

	err := NewCWebP().
		Quality(uint(compression)).
		InputFile(path).
		OutputFile(path + ".webp").
		Run()

	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Could not convert the image to webp", Operation: op, Err: err},
		}).Error()
	}
}